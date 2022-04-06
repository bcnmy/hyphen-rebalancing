package pricing

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/bcnmy/hyphen-rebalancing/internal/config"
)

type CoinMarketCapProvider struct {
	updateInterval uint64
	apiKey         string
	tokens         map[string]config.Token
	latestPrices   map[string]float64
	logger         log.Logger

	priceMutex sync.Mutex
}

func NewCoinMarketCapProvider(updateInterval uint64, apiKey string, tokens map[string]config.Token, logger log.Logger) (Provider, error) {
	return &CoinMarketCapProvider{
		updateInterval: updateInterval,
		apiKey:         apiKey,
		tokens:         tokens,
		latestPrices:   make(map[string]float64),
		logger:         log.With(logger, "service", "pricing:coinmarketcap"),
	}, nil
}

func (p *CoinMarketCapProvider) Run(ctx context.Context) error {
	p.updatePrices(ctx)

	for {
		select {
		case <-time.After(time.Second * time.Duration(p.updateInterval)):
			p.updatePrices(ctx)
		case <-ctx.Done():
			return nil
		}
	}
}

func (p *CoinMarketCapProvider) PriceBySymbol(symbol string) (float64, error) {
	p.priceMutex.Lock()
	defer p.priceMutex.Unlock()
	price, ok := p.latestPrices[symbol]
	if !ok {
		return 0, fmt.Errorf("price is not available for symbol: %s", symbol)
	}

	return price, nil
}

func (p *CoinMarketCapProvider) updatePrices(ctx context.Context) {
	level.Info(p.logger).Log("msg", "updating prices...")

	client := &http.Client{
		Timeout: time.Second * time.Duration(p.updateInterval),
	}

	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest", nil)
	if err != nil {
		level.Error(p.logger).Log("msg", "unable to create request", "err", err)
		return
	}

	req.Header.Add("X-CMC_PRO_API_KEY", p.apiKey)

	var symbols []string
	for symbol, _ := range p.tokens {
		symbols = append(symbols, strings.ToUpper(symbol))
	}

	q := req.URL.Query()
	q.Add("symbol", strings.Join(symbols, ","))
	q.Add("convert", "USD")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		level.Error(p.logger).Log("msg", "unable to send request", "err", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		level.Error(p.logger).Log("msg", "unable to read response", "err", err)
		return
	}

	result := QuotesLatestResponse{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		level.Error(p.logger).Log("msg", "unable to decode response", "err", err)
		return
	}

	newPrices := make(map[string]float64)
	for symbol, data := range result.Data {
		symbol = strings.ToUpper(symbol)
		if len(data) == 0 {
			level.Warn(p.logger).Log("msg", fmt.Sprintf("unable to get USD quote for %s", symbol))
			continue
		}

		usdQuote, ok := data[0].Quote["USD"]
		if !ok {
			level.Warn(p.logger).Log("msg", fmt.Sprintf("unable to get USD quote for %s", symbol))
			continue
		}

		newPrices[symbol] = usdQuote.Price
	}

	p.priceMutex.Lock()
	defer p.priceMutex.Unlock()

	for symbol, price := range newPrices {
		p.latestPrices[symbol] = price
		level.Info(p.logger).Log("msg", fmt.Sprintf("updating price for %s to %f", symbol, price))
	}
}

type QuotesLatestResponse struct {
	Data map[string][]QuotesData `json:"data"`
}

type QuotesData struct {
	ID     int              `json:"id"`
	Name   string           `json:"name"`
	Symbol string           `json:"symbol"`
	Slug   string           `json:"slug"`
	Quote  map[string]Quote `json:"quote"`
}

type Quote struct {
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"last_updated"`
}
