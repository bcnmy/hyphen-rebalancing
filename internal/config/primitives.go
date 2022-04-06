package config

import (
	"crypto/ecdsa"
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v3"
)

const nullAddressHex = "0x0000000000000000000000000000000000000000"

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

type PrivateKeyEnv struct {
	Value *ecdsa.PrivateKey
}

func (p *PrivateKeyEnv) UnmarshalYAML(value *yaml.Node) error {
	privateKeyValue, ok := os.LookupEnv(value.Value)
	if !ok {
		return errors.New("private key env variable is not set")
	}

	key, err := crypto.HexToECDSA(privateKeyValue)
	if err != nil {
		return err
	}

	p.Value = key
	return nil
}
