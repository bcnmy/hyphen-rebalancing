package bot

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/bcnmy/hyphen-arbitrage/internal/erc20"
	"github.com/bcnmy/hyphen-arbitrage/internal/pool"
)

const connectionTimeout = time.Second * 15
const jobTimeout = time.Second * 15

type Bot interface {
	// Starts bot worker
	Run(ctx context.Context) error

	// Gracefully stops bot worker
	Stop(ctx context.Context) error
}

func New(mgr pool.Manager, logger log.Logger) (Bot, error) {
	return &bot{
		mgr: mgr,
		logger: log.With(
			logger,
			"mgr", fmt.Sprintf("%s:%s:%s", mgr.PoolName(), mgr.AccountName(), mgr.TokenName()),
		),

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
		go b.processPool(ctx, p)
	}

	<-ctx.Done()
	close(b.running)
	b.poolWg.Wait()

	return nil
}

func (b *bot) Stop(ctx context.Context) error {
	<-ctx.Done()
	return nil
}

func (b *bot) processPool(ctx context.Context, p pool.Pool) {
	level.Info(b.logger).Log("network", p.NetworkName(), "msg", "starting processing pool")

	b.poolWg.Add(1)
	defer b.poolWg.Done()

	dialCtx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()

	client, err := p.Dial(dialCtx)
	if err != nil {
		level.Error(b.logger).Log("msg", "unable to create EVM client for pool", "err", err)
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
		case <-jobs:
			jobCtx, cancel := context.WithTimeout(ctx, jobTimeout)
			defer cancel()

			err = b.processPoolJob(jobCtx, client, p)
			if err != nil {
				level.Error(b.logger).Log("msg", "unable to process pool job", "err", err)
				return
			}
		}
	}
}

func (b *bot) processPoolJob(ctx context.Context, client *ethclient.Client, p pool.Pool) error {
	level.Info(b.logger).Log("network", p.NetworkName(), "msg", "processing new pool job")
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

	level.Debug(b.logger).Log("network", p.NetworkName(), "msg", fmt.Sprintf("token balance: %s", balance.String()))

	_, err = b.calculateProfitability(client, p, *balance)
	if err != nil {
		return err
	}

	return nil
}

func (b *bot) calculateProfitability(client *ethclient.Client, p pool.Pool, maxDeposit big.Int) (*profitability, error) {
	// TODO: Implementation
	return &profitability{}, nil
}

type profitability struct {
	deposit  big.Int
	reward   big.Int
	gasPrice big.Int
}
