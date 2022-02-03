package bot

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/bcnmy/hyphen-arbitrage/internal/generated/erc20"
	"github.com/bcnmy/hyphen-arbitrage/internal/generated/hyphen"
	"github.com/bcnmy/hyphen-arbitrage/internal/pool"
)

const (
	connectionTimeout = time.Second * 30
	jobTimeout        = time.Second * 60
)

const (
	baseDivisor = 10000000000
)

type Bot interface {
	Run(ctx context.Context) error
}

func New(mgr pool.Manager, logger log.Logger) (Bot, error) {
	return &bot{
		mgr:    mgr,
		logger: logger,
	}, nil
}

type bot struct {
	mgr    pool.Manager
	logger log.Logger
}

func (b *bot) Run(ctx context.Context) error {
	jobs := make(chan bool, 4)
	jobs <- true

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-jobs:
			b.processJob(ctx)
		}
	}
}

func (b *bot) processJob(ctx context.Context) {
	jobCtx, cancel := context.WithTimeout(ctx, jobTimeout)
	defer cancel()

	jobWg := sync.WaitGroup{}
	profitWg := sync.WaitGroup{}
	profitChan := make(chan *arbitrationProfit)
	mostProfitChan := make(chan *arbitrationProfit)
	for _, p := range b.mgr.Pools() {
		jobWg.Add(1)
		go b.processPool(jobCtx, p, profitChan, &jobWg)
	}

	profitWg.Add(1)
	go b.processProfits(ctx, profitChan, mostProfitChan, &profitWg)

	jobWg.Wait()
	close(profitChan)

	profit, ok := <-mostProfitChan
	if !ok {
		return
	}

	route := fmt.Sprintf("%s -> %s", profit.route.from.NetworkName(), profit.route.to.NetworkName())
	level.Debug(b.logger).Log("msg", "most profitable path", "route", route, "netprof", profit.netTokenProfit)
}

func (b *bot) processProfits(ctx context.Context, profitChan <-chan *arbitrationProfit, mostProfitChan chan<- *arbitrationProfit, profitWg *sync.WaitGroup) {
	defer profitWg.Done()
	defer close(mostProfitChan)

	var mostProfitable *arbitrationProfit
	for profit := range profitChan {
		if mostProfitable == nil || profit.netTokenProfit.Cmp(mostProfitable.netTokenProfit) > 0 {
			mostProfitable = profit
		}
	}

	if mostProfitable != nil {
		mostProfitChan <- mostProfitable
	}
}

func (b *bot) processPool(ctx context.Context, p pool.Pool, profitChan chan<- *arbitrationProfit, jobWg *sync.WaitGroup) {
	defer jobWg.Done()

	reward, err := b.estimatePotentialReward(ctx, p)
	if err != nil {
		level.Error(b.logger).Log("pool", b.poolID(p), "msg", "unable to estimate potential reward", "err", err)
		return
	}

	for _, poolTo := range b.mgr.Pools() {
		if p.ChainID() == poolTo.ChainID() && p.Address() == poolTo.Address() {
			continue
		}

		fee, err := b.estimatePotentialFee(ctx, poolTo, reward)
		if err != nil {
			level.Error(b.logger).Log("pool", b.poolID(p), "msg", "unable to estimate potential fee", "err", err)
			return
		}

		route := &arbitrationRoute{
			from:   p,
			to:     poolTo,
			reward: reward,
			fee:    fee,
		}

		profit, err := b.calculateArbitrationProfit(ctx, route)
		if err != nil {
			level.Error(b.logger).Log("pool", b.poolID(p), "msg", "unable to calculate arbitration profit", "err", err)
			return
		}

		profitChan <- profit
	}
}

func (b *bot) estimatePotentialReward(ctx context.Context, p pool.Pool) (*potentialReward, error) {
	dialCtx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	client, err := p.Dial(dialCtx)
	if err != nil {
		return nil, err
	}

	balance, allowance, err := b.getBalanceAndAllowance(ctx, client, p)
	if err != nil {
		return nil, err
	}

	deposit, reward, err := b.estimateDepositAndReward(ctx, client, p, balance)
	if err != nil {
		return nil, err
	}

	return &potentialReward{
		balance:   balance,
		allowance: allowance,
		deposit:   deposit,
		reward:    reward,
	}, nil
}

func (b *bot) estimatePotentialFee(ctx context.Context, p pool.Pool, reward *potentialReward) (*potentialFee, error) {
	dialCtx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	client, err := p.Dial(dialCtx)
	if err != nil {
		return nil, err
	}

	liquidityPool, err := hyphen.NewLiquidityPool(p.Address(), client)
	if err != nil {
		return nil, err
	}

	callOpts := newCallOpts(ctx)
	amount := new(big.Int).Add(reward.deposit, reward.reward)
	transferFeePerc, err := liquidityPool.GetTransferFee(callOpts, p.Token(), amount)
	if err != nil {
		return nil, err
	}

	transferFee := new(big.Int).Div(new(big.Int).Mul(amount, transferFeePerc), big.NewInt(baseDivisor))

	return &potentialFee{
		tokenFee: transferFee,
	}, nil
}

func (b *bot) calculateArbitrationProfit(ctx context.Context, route *arbitrationRoute) (*arbitrationProfit, error) {
	netTokenProfit := new(big.Int).Sub(route.reward.reward, route.fee.tokenFee)

	return &arbitrationProfit{
		route:          route,
		netTokenProfit: netTokenProfit,
	}, nil
}

func (b *bot) estimateDepositAndReward(ctx context.Context, client *ethclient.Client, p pool.Pool, maxDeposit *big.Int) (*big.Int, *big.Int, error) {
	liquidityPool, err := hyphen.NewLiquidityPool(p.Address(), client)
	if err != nil {
		return nil, nil, err
	}

	callOpts := newCallOpts(ctx)
	liquidityProvidersAddress, err := liquidityPool.LiquidityProviders(callOpts)
	if err != nil {
		return nil, nil, err
	}

	liquidityProviders, err := hyphen.NewLiquidityProviders(liquidityProvidersAddress, client)
	if err != nil {
		return nil, nil, err
	}

	currentLiquidity, err := liquidityPool.GetCurrentLiquidity(callOpts, p.Token())
	if err != nil {
		return nil, nil, err
	}

	suppliedLiquidity, err := liquidityProviders.GetSuppliedLiquidityByToken(callOpts, p.Token())
	if err != nil {
		return nil, nil, err
	}

	tokenManagerAddress, err := liquidityPool.TokenManager(callOpts)
	if err != nil {
		return nil, nil, err
	}

	tokenManager, err := hyphen.NewTokenManager(tokenManagerAddress, client)
	if err != nil {
		return nil, nil, err
	}

	tokenInfo, err := tokenManager.TokensInfo(callOpts, p.Token())
	if err != nil {
		return nil, nil, err
	}

	liquidityDeficit := big.NewInt(0).Sub(suppliedLiquidity, currentLiquidity)
	deposit := min(maxDeposit, liquidityDeficit, tokenInfo.MaxCap)
	deposit = max(deposit, tokenInfo.MinCap)

	reward, err := liquidityPool.GetRewardAmount(callOpts, deposit, p.Token())
	if err != nil {
		return nil, nil, err
	}

	return deposit, reward, nil
}

func (b *bot) getBalanceAndAllowance(ctx context.Context, client *ethclient.Client, p pool.Pool) (*big.Int, *big.Int, error) {
	erc20Contract, err := erc20.NewErc20(p.Token(), client)
	if err != nil {
		return nil, nil, err
	}

	callOpts := newCallOpts(ctx)
	balance, err := erc20Contract.BalanceOf(callOpts, b.mgr.AccountAddress())
	if err != nil {
		return nil, nil, err
	}

	allowance, err := erc20Contract.Allowance(callOpts, b.mgr.AccountAddress(), p.Address())

	if err != nil {
		return nil, nil, err
	}

	return balance, allowance, nil
}

func (b *bot) updateAllowance(ctx context.Context, client *ethclient.Client, p pool.Pool, currentAllowance *big.Int, requiredAllowance *big.Int) error {
	if currentAllowance.Cmp(requiredAllowance) > 0 {
		return nil
	}

	toApprove := requiredAllowance
	if b.mgr.InfiniteApprove() {
		toApprove = abi.MaxUint256
	}

	erc20Contract, err := erc20.NewErc20(p.Token(), client)
	if err != nil {
		return err
	}

	opts, err := newTransactor(ctx, b.mgr.PrivateKey(), p.ChainID())
	if err != nil {
		return err
	}

	tx, err := erc20Contract.Approve(opts, p.Address(), toApprove)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		return err
	}

	return nil
}

func (b *bot) poolID(p pool.Pool) string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", b.mgr.PoolName(), b.mgr.AccountName(), b.mgr.TokenName(), p.NetworkName(), p.ChainID().String())
}
