package config

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v3"
)

const nullAddressHex = "0x0000000000000000000000000000000000000000"

type Config struct {
	Networks map[string]Network `yaml:"networks"`
	Tokens   map[string]Token   `yaml:"tokens"`
	Accounts map[string]Account `yaml:"accounts"`
	Pools    map[string]Pool    `yaml:"pools"`
}

type Network struct {
	ChainID         uint64   `yaml:"chainID"`
	RPCURL          string   `yaml:"rpcURL"`
	DefaultGasPrice *big.Int `yaml:"defaultGasPrice"`
	MaxGasPrice     *big.Int `yaml:"maxGasPrice"`
}

type Token struct {
	Addresses map[string]Address `yaml:"addresses"`
}

type Account struct {
	PrivateKey PrivateKey `yaml:"privateKey"`
}

type Pool struct {
	Tokens                 []string           `yaml:"tokens"`
	Accounts               []string           `yaml:"accounts"`
	Addresses              map[string]Address `yaml:"addresses"`
	ProfitabilityThreshold float64            `yaml:"profitabilityThreshold"`
	InfiniteApprove        bool               `yaml:"infiniteApprove"`
	CancelOnFail           bool               `yaml:"cancelOnFail"`
}

type Address struct {
	Value common.Address
}

func (a *Address) UnmarshalYAML(value *yaml.Node) error {
	address := common.HexToAddress(value.Value)
	nullAddress := common.HexToAddress(nullAddressHex)
	if address.Hex() == nullAddress.Hex() {
		return errors.New("invalid address")
	}

	a.Value = address
	return nil
}

type PrivateKey struct {
	Value *ecdsa.PrivateKey
}

func (p *PrivateKey) UnmarshalYAML(value *yaml.Node) error {
	key, err := crypto.HexToECDSA(value.Value)
	if err != nil {
		return err
	}

	p.Value = key
	return nil
}
