package config

import (
	"math/big"
)

type Config struct {
	General  General            `yaml:"general"`
	Networks map[string]Network `yaml:"networks"`
	Tokens   map[string]Token   `yaml:"tokens"`
	Accounts map[string]Account `yaml:"accounts"`
	Pools    map[string]Pool    `yaml:"pools"`
	Pricing  Pricing            `yaml:"pricing"`
}

type General struct {
	UpdateInterval  uint64 `yaml:"updateInterval"`
	JobTimeout      uint64 `yaml:"jobTimeout"`
	HalfRebalancing bool   `yaml:"halfRebalancing"`
}

type Network struct {
	ChainID *big.Int `yaml:"chainID"`
	RPCURL  Secret   `yaml:"rpcURL"`
}

type Token struct {
	Addresses       map[string]Address `yaml:"addresses"`
	MinimalProfit   *big.Int           `yaml:"minimalProfit"`
	InfiniteApprove bool               `yaml:"infiniteApprove"`
}

type Account struct {
	PrivateKey PrivateKey `yaml:"privateKey"`
}

type Pool struct {
	Tokens    []string           `yaml:"tokens"`
	Accounts  []string           `yaml:"accounts"`
	Addresses map[string]Address `yaml:"addresses"`
}

type Pricing struct {
	UpdateInterval      uint64 `yaml:"updateInterval"`
	CoinMarketCapAPIKey Secret `yaml:"coinMarketCapAPIKey"`
}
