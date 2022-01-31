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
	// Starts bot worker
	Run(ctx context.Context) error

	// Gracefully stops bot worker
	Stop(ctx context.Context) error
}

func New(mgr pool.Manager, logger log.Logger) (Bot, error) {
	return &bot{
		mgr:    mgr,
		logger: logger,

		running: make(chan bool),
	}, nil
}

type bot struct {
	mgr    pool.Manager
	logger log.Logger

	running chan bool
	poolWg  sync.WaitGroup
}

func (b *bot) Run(ctx context.Context) error {
	for _, p := range b.mgr.Pools() {
		b.poolWg.Add(1)
		go b.processPool(ctx, p)
	}

	b.poolWg.Wait()

	return nil
}

func (b *bot) Stop(ctx context.Context) error {
	close(b.running)
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
		case _, closed := <-b.running:
			if closed {
				return
			}
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
	level.Info(b.logger).Log("pool", b.poolID(p), "msg", "processing new pool job")
	erc20Contract, err := erc20.NewErc20(p.Token(), client)
	if err != nil {
		return err
	}

	balance, err := erc20Contract.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, b.mgr.AccountAddress())

	if err != nil {
		return err
	}

	allowance, err := erc20Contract.Allowance(&bind.CallOpts{
		Context: ctx,
	}, b.mgr.AccountAddress(), p.Address())

	if err != nil {
		return err
	}

	level.Debug(b.logger).Log("pool", b.poolID(p), "msg", "fetched token balance", "balance", balance, "allowance", allowance)

	// Implementation draft:
	// 1. Check if pool is in deficit state
	// 2. Calculate liquidityDeficit = lProv.suppliedLiquidity - lPool.currentLiquidity ✅
	// 3. Calculate deposit = min(maxDeposit, liquidityDeficit) ✅
	// 4. Calculate rewards = lPool.getRewardAmount(deposit) ✅
	// 5. Calculate deposit gas fees
	// 6. Calculate send gas fees
	// 7. Calculate gas transter fees on destination chain

	potentialReward, err := b.calculatePotentialRewards(ctx, client, p, balance)
	if err != nil {
		return err
	}

	level.Debug(b.logger).Log("pool", b.poolID(p), "msg", "calculated potential rewards", "deposit", potentialReward.deposit, "reward", potentialReward.reward, "percentage", potentialReward.rewardPercentage)

	if potentialReward.reward.Cmp(big.NewInt(0)) <= 0 || potentialReward.deposit.Cmp(new(big.Int)) <= 0 {
		level.Debug(b.logger).Log("pool", b.poolID(p), "msg", "not enough rewards")
		return nil
	}

	opts, err := bind.NewKeyedTransactorWithChainID(b.mgr.PrivateKey(), p.ChainID())
	if err != nil {
		return err
	}

	if allowance.Cmp(potentialReward.deposit) < 0 {
		amount := potentialReward.deposit
		if b.mgr.InfiniteApprove() {
			amount = abi.MaxUint256
		}

		tx, err := erc20Contract.Approve(opts, p.Address(), amount)
		if err != nil {
			return err
		}

		level.Debug(b.logger).Log("pool", b.poolID(p), "msg", "approved liquidity pool spending", "amount", amount, "tx", tx.Hash())
	}

	// TODO: erc20Contract.Approve(&bind.TransactOpts{})

	liquidityPoolAbi := &abi.ABI{}
	err = liquidityPoolAbi.UnmarshalJSON([]byte(hyphen.LiquidityPoolABI))
	if err != nil {
		return err
	}

	var args []interface{}
	args = append(args, big.NewInt(0), p.Token(), b.mgr.AccountAddress(), potentialReward.deposit, "tag")
	input, err := liquidityPoolAbi.Pack("depositErc20", args...)
	if err != nil {
		return err
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

	gasLimit, err := client.EstimateGas(ctx, callMsg)
	if err != nil {
		return err
	}

	level.Debug(b.logger).Log("pool", b.poolID(p), "msg", fmt.Sprintf("gas limit: %d", gasLimit))

	return nil
}

func (b *bot) calculatePotentialRewards(ctx context.Context, client *ethclient.Client, p pool.Pool, maxDeposit *big.Int) (*rewardsEstimate, error) {
	liquidityPool, err := hyphen.NewLiquidityPool(p.Address(), client)
	if err != nil {
		return nil, err
	}

	liquidityProviders, err := hyphen.NewLiquidityProviders(p.Providers(), client)
	if err != nil {
		return nil, err
	}

	currentLiquidity, err := liquidityPool.GetCurrentLiquidity(&bind.CallOpts{}, p.Token())
	if err != nil {
		return nil, err
	}

	suppliedLiquidity, err := liquidityProviders.GetSuppliedLiquidityByToken(&bind.CallOpts{}, p.Token())
	if err != nil {
		return nil, err
	}

	liquidityDeficit := big.NewInt(0).Sub(suppliedLiquidity, currentLiquidity)

	deposit := liquidityDeficit
	if liquidityDeficit.Cmp(maxDeposit) > 0 {
		deposit = maxDeposit
	}

	reward, err := liquidityPool.GetRewardAmount(&bind.CallOpts{}, deposit, p.Token())
	if err != nil {
		return nil, err
	}

	rewardPercentage := new(big.Float)
	if deposit.Cmp(new(big.Int)) > 0 && reward.Cmp(new(big.Int)) > 0 {
		rewardPercentage.Quo(new(big.Float).SetInt(reward), new(big.Float).SetInt(deposit))
		rewardPercentage.Mul(rewardPercentage, big.NewFloat(100))
	}

	return &rewardsEstimate{
		deposit:          deposit,
		reward:           reward,
		rewardPercentage: rewardPercentage,
	}, nil
}

func (b *bot) poolID(p pool.Pool) string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", b.mgr.PoolName(), b.mgr.AccountName(), b.mgr.TokenName(), p.NetworkName(), p.ChainID().String())
}

type rewardsEstimate struct {
	deposit          *big.Int
	reward           *big.Int
	rewardPercentage *big.Float
}
