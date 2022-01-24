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

	"github.com/bcnmy/hyphen-arbitrage/internal/config"
	"github.com/bcnmy/hyphen-arbitrage/internal/pool"
)

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))

	level.Info(logger).Log("msg", "reading config", "file", configFile())
	conf, err := config.FromFile(configFile())
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

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	for _, mgr := range poolManagers {
		wg.Add(1)

		go func(mgr pool.Manager) {
			mgr.Run(ctx)
			wg.Done()
		}(mgr)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		level.Info(logger).Log("msg", fmt.Sprintf("got %v signal, stopping", sig))
		cancel()
	}()

	wg.Wait()
}

func configFile() string {
	if filename, ok := os.LookupEnv("CONFIG_FILE"); ok {
		return filename
	}

	return "config.yaml"
}
