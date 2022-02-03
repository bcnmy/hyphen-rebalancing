package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func NewCallOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx}
}

func NewTransactor(ctx context.Context, key *ecdsa.PrivateKey, chainId *big.Int) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(key, chainId)
	if err != nil {
		return nil, err
	}

	opts.Context = ctx

	return opts, nil
}

func Min(num *big.Int, v ...*big.Int) *big.Int {
	min := num
	for _, num := range v {
		if num.Cmp(min) < 0 {
			min = num
		}
	}

	return min
}

func Max(num *big.Int, v ...*big.Int) *big.Int {
	max := num
	for _, num := range v {
		if num.Cmp(max) > 0 {
			max = num
		}
	}

	return max
}
