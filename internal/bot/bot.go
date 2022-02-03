package bot

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/bcnmy/hyphen-arbitrage/internal/contractfactory"
	"github.com/bcnmy/hyphen-arbitrage/internal/pool"
	"github.com/bcnmy/hyphen-arbitrage/internal/utils"
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
	cf, err := contractfactory.New()
	if err != nil {
		return nil, err
	}

	return &bot{
		mgr: mgr,
		cf:  cf,

		logger: logger,
	}, nil
}

type bot struct {
	mgr pool.Manager
	cf  contractfactory.Factory

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

	err := b.executeProfitablePath(ctx, profit)
	if err != nil {
		level.Error(b.logger).Log("msg", "unable to execute most profitable path", "route", profit.route, "err", err)
		return
	}
}

func (b *bot) executeProfitablePath(ctx context.Context, profit *arbitrationProfit) error {
	reward := profit.route.reward

	level.Info(b.logger).Log(
		"msg", "found most profitable path",
		"route", profit.route,
		"netprof", profit.netTokenProfit,
		"deposit", reward.deposit,
		"reward", reward.reward,
	)

	if b.mgr.MinimalProfit().Cmp(profit.netTokenProfit) > 0 {
		level.Info(b.logger).Log("msg", "net profit is not enough", "route", profit.route, "netprof", profit.netTokenProfit, "minprof", b.mgr.MinimalProfit())
		return nil
	}

	err := b.updateAllowance(ctx, profit.route.from, reward.allowance, reward.deposit)
	if err != nil {
		return err
	}

	err = b.executeDeposit(ctx, profit.route.from, profit.route.to, reward.deposit)
	if err != nil {
		return err
	}

	return nil
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
	balance, allowance, err := b.getBalanceAndAllowance(ctx, p)
	if err != nil {
		return nil, err
	}

	deposit, reward, err := b.estimateDepositAndReward(ctx, p, balance)
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
	liquidityPool, err := b.cf.LiquidityPool(p)
	if err != nil {
		return nil, err
	}

	callOpts := utils.NewCallOpts(ctx)
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

func (b *bot) estimateDepositAndReward(ctx context.Context, p pool.Pool, maxDeposit *big.Int) (*big.Int, *big.Int, error) {
	liquidityPool, err := b.cf.LiquidityPool(p)
	if err != nil {
		return nil, nil, err
	}

	callOpts := utils.NewCallOpts(ctx)
	liquidityProviders, err := b.cf.LiquidityProviders(p)
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

	tokenManager, err := b.cf.TokenManager(p)
	if err != nil {
		return nil, nil, err
	}

	tokenInfo, err := tokenManager.TokensInfo(callOpts, p.Token())
	if err != nil {
		return nil, nil, err
	}

	liquidityDeficit := big.NewInt(0).Sub(suppliedLiquidity, currentLiquidity)
	deposit := utils.Min(maxDeposit, liquidityDeficit, tokenInfo.MaxCap)
	deposit = utils.Max(deposit, tokenInfo.MinCap)

	reward, err := liquidityPool.GetRewardAmount(callOpts, deposit, p.Token())
	if err != nil {
		return nil, nil, err
	}

	return deposit, reward, nil
}

func (b *bot) getBalanceAndAllowance(ctx context.Context, p pool.Pool) (*big.Int, *big.Int, error) {
	erc20Contract, err := b.cf.ERC20(p)
	if err != nil {
		return nil, nil, err
	}

	callOpts := utils.NewCallOpts(ctx)
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

func (b *bot) updateAllowance(ctx context.Context, p pool.Pool, currentAllowance *big.Int, requiredAllowance *big.Int) error {
	if currentAllowance.Cmp(requiredAllowance) > 0 {
		return nil
	}

	toApprove := requiredAllowance
	if b.mgr.InfiniteApprove() {
		toApprove = abi.MaxUint256
	}

	erc20Contract, err := b.cf.ERC20(p)
	if err != nil {
		return err
	}

	opts, err := utils.NewTransactor(ctx, b.mgr.PrivateKey(), p.ChainID())
	if err != nil {
		return err
	}

	approveTx, err := erc20Contract.Approve(opts, p.Address(), toApprove)
	if err != nil {
		return err
	}

	level.Info(b.logger).Log("msg", "updating allowance", "tx", approveTx.Hash())

	client, err := b.cf.Client(p)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, client, approveTx)
	if err != nil {
		return err
	}

	return nil
}

func (b *bot) executeDeposit(ctx context.Context, from pool.Pool, to pool.Pool, amount *big.Int) error {
	liquidityPool, err := b.cf.LiquidityPool(from)
	if err != nil {
		return err
	}

	txOpts, err := utils.NewTransactor(ctx, b.mgr.PrivateKey(), from.ChainID())
	if err != nil {
		return err
	}

	depositTx, err := liquidityPool.DepositErc20(txOpts, to.ChainID(), from.Token(), b.mgr.AccountAddress(), amount, "")
	if err != nil {
		return err
	}

	level.Info(b.logger).Log("msg", "executing arbitrage transaction", "tx", depositTx.Hash())

	client, err := b.cf.Client(from)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, client, depositTx)
	if err != nil {
		return err
	}

	return nil
}

func (b *bot) poolID(p pool.Pool) string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", b.mgr.PoolName(), b.mgr.AccountName(), b.mgr.TokenName(), p.NetworkName(), p.ChainID().String())
}
