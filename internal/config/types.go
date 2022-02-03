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
}

type General struct {
	UpdateInterval uint64 `yaml:"updateInterval"`
	JobTimeout     uint64 `yaml:"jobTimeout"`
}

type Network struct {
	ChainID *big.Int `yaml:"chainID"`
	RPCURL  string   `yaml:"rpcURL"`
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
