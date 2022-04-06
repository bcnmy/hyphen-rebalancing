package pricing

import (
	"context"
	"errors"

	"github.com/bcnmy/hyphen-rebalancing/internal/config"
	"github.com/go-kit/log"
)

const defaultUpdateInterval uint64 = 60

// Provider allows to get token pricing data
type Provider interface {
	Run(ctx context.Context) error
	PriceBySymbol(symbol string) (float64, error)
}

func New(conf config.Pricing, tokens map[string]config.Token, logger log.Logger) (Provider, error) {
	updateInterval := defaultUpdateInterval
	if conf.UpdateInterval > 0 {
		updateInterval = conf.UpdateInterval
	}

	if conf.CoinMarketCapAPIKey.Value != "" {
		return NewCoinMarketCapProvider(updateInterval, conf.CoinMarketCapAPIKey.Value, tokens, logger)
	}

	return nil, errors.New("no pricing providers specified")
}
