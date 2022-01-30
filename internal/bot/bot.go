package bot

import (
	"context"
	"fmt"

	"github.com/go-kit/log"

	"github.com/bcnmy/hyphen-arbitrage/internal/pool"
)

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
	}, nil
}

type bot struct {
	mgr    pool.Manager
	logger log.Logger
}

func (b *bot) Run(ctx context.Context) error {
	<-ctx.Done()
	return nil
}

func (b *bot) Stop(ctx context.Context) error {
	<-ctx.Done()
	return nil
}
