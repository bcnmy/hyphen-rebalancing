package bot

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/bcnmy/hyphen-arbitrage/internal/generated/erc20"
	"github.com/bcnmy/hyphen-arbitrage/internal/generated/hyphen"
	"github.com/bcnmy/hyphen-arbitrage/internal/pool"
)

const connectionTimeout = time.Second * 15
const jobTimeout = time.Second * 30

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

	poolWg sync.WaitGroup
}

func (b *bot) Run(ctx context.Context) error {
	for _, p := range b.mgr.Pools() {
		b.poolWg.Add(1)
		go b.processPool(ctx, p)
	}

	b.poolWg.Wait()

	return nil
}

func (b *bot) processPool(ctx context.Context, p pool.Pool) {
	level.Info(b.logger).Log("msg", "starting processing pool", "network", p.NetworkName())
	defer b.poolWg.Done()

	dialCtx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	client, err := p.Dial(dialCtx)
	if err != nil {
		level.Error(b.logger).Log("msg", "unable to create EVM client for pool", "err", err, "network", p.NetworkName())
		return
	}

	jobs := make(chan bool, 4)
	jobs <- true

	for {
		select {
		case <-ctx.Done():
			return
		case <-jobs:
			jobCtx, cancel := context.WithTimeout(ctx, jobTimeout)
			defer cancel()

			err = b.processPoolJob(jobCtx, client, p)
			if err != nil {
				level.Error(b.logger).Log("pool", b.poolID(p), "msg", "unable to process pool job", "err", err)
			}
		}
	}
}

func (b *bot) processPoolJob(ctx context.Context, client *ethclient.Client, p pool.Pool) error {
	balance, allowance, err := b.getBalanceAndAllowance(ctx, client, p)
	if err != nil {
		return err
	}

	level.Debug(b.logger).Log("pool", b.poolID(p), "msg", "fetched token balance", "balance", balance, "allowance", allowance)

	deposit, reward, err := b.estimateDepositAndReward(ctx, client, p, balance)
	if err != nil {
		return err
	}

	level.Debug(b.logger).Log("pool", b.poolID(p), "msg", "calculated potential rewards", "deposit", deposit, "reward", reward)
	if reward.Cmp(big.NewInt(0)) <= 0 {
		return nil
	}

	if allowance.Cmp(deposit) < 0 {
		amount := deposit
		if b.mgr.InfiniteApprove() {
			amount = abi.MaxUint256
		}

		err = b.setAllowance(ctx, client, p, amount)
		if err != nil {
			return err
		}
	}

	gasLimit, err := b.estimateDepositGas(ctx, client, p, deposit)
	if err != nil {
		return err
	}

	level.Debug(b.logger).Log("pool", b.poolID(p), "msg", fmt.Sprintf("gas limit: %d", gasLimit))

	return nil
}

func (b *bot) setAllowance(ctx context.Context, client *ethclient.Client, p pool.Pool, amount *big.Int) error {
	erc20Contract, err := erc20.NewErc20(p.Token(), client)
	if err != nil {
		return err
	}

	opts, err := b.newTransactor(ctx, p.ChainID())
	if err != nil {
		return err
	}

	tx, err := erc20Contract.Approve(opts, p.Address(), amount)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		return err
	}

	return nil
}

func (b *bot) estimateDepositAndReward(ctx context.Context, client *ethclient.Client, p pool.Pool, maxDeposit *big.Int) (*big.Int, *big.Int, error) {
	liquidityPool, err := hyphen.NewLiquidityPool(p.Address(), client)
	if err != nil {
		return nil, nil, err
	}

	liquidityProviders, err := hyphen.NewLiquidityProviders(p.Providers(), client)
	if err != nil {
		return nil, nil, err
	}

	currentLiquidity, err := liquidityPool.GetCurrentLiquidity(b.newCallOpts(ctx), p.Token())
	if err != nil {
		return nil, nil, err
	}

	suppliedLiquidity, err := liquidityProviders.GetSuppliedLiquidityByToken(b.newCallOpts(ctx), p.Token())
	if err != nil {
		return nil, nil, err
	}

	liquidityDeficit := big.NewInt(0).Sub(suppliedLiquidity, currentLiquidity)

	deposit := liquidityDeficit
	if liquidityDeficit.Cmp(maxDeposit) > 0 {
		deposit = maxDeposit
	}

	reward, err := liquidityPool.GetRewardAmount(b.newCallOpts(ctx), deposit, p.Token())
	if err != nil {
		return nil, nil, err
	}

	return deposit, reward, nil
}

func (b *bot) estimateDepositGas(ctx context.Context, client *ethclient.Client, p pool.Pool, deposit *big.Int) (uint64, error) {
	liquidityPoolAbi := &abi.ABI{}
	err := liquidityPoolAbi.UnmarshalJSON([]byte(hyphen.LiquidityPoolABI))
	if err != nil {
		return 0, err
	}

	var args []interface{}
	args = append(args, big.NewInt(0), p.Token(), b.mgr.AccountAddress(), deposit, "tag")
	input, err := liquidityPoolAbi.Pack("depositErc20", args...)
	if err != nil {
		return 0, err
	}

	sendTo := p.Address()
	callMsg := ethereum.CallMsg{
		From:     b.mgr.AccountAddress(),
		To:       &sendTo,
		Gas:      0,
		GasPrice: big.NewInt(0),
		Value:    big.NewInt(0),
		Data:     input,
	}

	return client.EstimateGas(ctx, callMsg)
}

func (b *bot) getBalanceAndAllowance(ctx context.Context, client *ethclient.Client, p pool.Pool) (*big.Int, *big.Int, error) {
	erc20Contract, err := erc20.NewErc20(p.Token(), client)
	if err != nil {
		return nil, nil, err
	}

	balance, err := erc20Contract.BalanceOf(b.newCallOpts(ctx), b.mgr.AccountAddress())
	if err != nil {
		return nil, nil, err
	}

	allowance, err := erc20Contract.Allowance(b.newCallOpts(ctx), b.mgr.AccountAddress(), p.Address())

	if err != nil {
		return nil, nil, err
	}

	return balance, allowance, nil
}

func (b *bot) newTransactor(ctx context.Context, chainId *big.Int) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(b.mgr.PrivateKey(), chainId)
	if err != nil {
		return nil, err
	}

	opts.Context = ctx

	return opts, nil
}

func (b *bot) newCallOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx}
}

func (b *bot) poolID(p pool.Pool) string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", b.mgr.PoolName(), b.mgr.AccountName(), b.mgr.TokenName(), p.NetworkName(), p.ChainID().String())
}
