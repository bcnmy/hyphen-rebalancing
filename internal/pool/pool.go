package pool

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bcnmy/hyphen-arbitrage/internal/config"
)

// Pool represents single liquidity pool on chain
type Pool interface {
	NetworkName() string
	Dial(context.Context) (*ethclient.Client, error)
	ChainID() uint64
	Token() common.Address
	Address() common.Address
	Providers() common.Address
}

type pool struct {
	networkName string
	network     config.Network
	token       common.Address
	address     common.Address
	providers   common.Address
}

func (p *pool) NetworkName() string {
	return p.networkName
}

func (p *pool) Dial(ctx context.Context) (*ethclient.Client, error) {
	return ethclient.DialContext(ctx, p.network.RPCURL)
}

func (p *pool) ChainID() uint64 {
	return p.network.ChainID
}

func (p *pool) Token() common.Address {
	return p.token
}

func (p *pool) Address() common.Address {
	return p.address
}

func (p *pool) Providers() common.Address {
	return p.providers
}
