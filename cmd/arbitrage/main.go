package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/bcnmy/hyphen-arbitrage/internal/bot"
	"github.com/bcnmy/hyphen-arbitrage/internal/config"
	"github.com/bcnmy/hyphen-arbitrage/internal/pool"
)

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))

	cfgFile := getConfigFile()
	level.Info(logger).Log("msg", "reading config", "file", cfgFile)
	conf, err := config.FromFile(cfgFile)
	if err != nil {
		panic(err)
	}

	level.Info(logger).Log(
		"msg", "read config",
		"networks", len(conf.Networks),
		"tokens", len(conf.Tokens),
		"accounts", len(conf.Accounts),
		"pools", len(conf.Pools),
	)

	poolManagers, err := pool.FromConfig(conf, logger)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	runBots(ctx, poolManagers, &wg, logger)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		level.Info(logger).Log("msg", fmt.Sprintf("got %v signal, stopping", sig))
		cancel()
	}()

	wg.Wait()
}

func runBots(ctx context.Context, poolManagers []pool.Manager, wg *sync.WaitGroup, logger log.Logger) {
	for _, mgr := range poolManagers {
		wg.Add(1)

		go func(mgr pool.Manager) {
			defer wg.Done()

			bot, err := bot.New(mgr, logger)
			if err != nil {
				level.Error(logger).Log("msg", "unable to create bot", "err", err)
				return
			}

			err = bot.Run(ctx)
			if err != nil {
				level.Error(logger).Log("msg", "bot running failure", "err", err)
				return
			}
		}(mgr)
	}
}

func getConfigFile() string {
	if filename, ok := os.LookupEnv("CONFIG_FILE"); ok {
		return filename
	}

	return "config.yaml"
}
