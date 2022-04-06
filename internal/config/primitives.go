package config

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v3"
)

const nullAddressHex = "0x0000000000000000000000000000000000000000"

type Address struct {
	Value common.Address
}

func (a *Address) UnmarshalYAML(value *yaml.Node) error {
	stringValue, err := prepareValue(value.Value)
	if err != nil {
		return err
	}

	address := common.HexToAddress(stringValue)
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
	stringValue, err := prepareValue(value.Value)
	if err != nil {
		return err
	}

	key, err := crypto.HexToECDSA(stringValue)
	if err != nil {
		return err
	}

	p.Value = key
	return nil
}

type Secret struct {
	Value string
}

func (s *Secret) UnmarshalYAML(value *yaml.Node) error {
	stringValue, err := prepareValue(value.Value)
	if err != nil {
		return err
	}

	s.Value = stringValue
	return nil
}

func prepareValue(value string) (string, error) {
	if !strings.HasPrefix(value, "$") {
		return value, nil
	}

	envKey := strings.TrimLeft(value, "$")
	envValue, ok := os.LookupEnv(envKey)
	if !ok {
		return "", fmt.Errorf("environment variable not found: %s", envKey)
	}

	return envValue, nil
}
