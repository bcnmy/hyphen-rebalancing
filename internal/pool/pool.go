package pool

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bcnmy/hyphen-rebalancing/internal/config"
)

// Pool represents single liquidity pool on chain
type Pool interface {
	Dial(context.Context) (*ethclient.Client, error)
	ChainID() *big.Int
	Token() common.Address
	Address() common.Address

	NetworkName() string
	String() string
}

type pool struct {
	network config.Network
	token   common.Address
	address common.Address

	networkName string
}

func (p *pool) Dial(ctx context.Context) (*ethclient.Client, error) {
	return ethclient.DialContext(ctx, p.network.RPCURL.Value)
}

func (p *pool) ChainID() *big.Int {
	return p.network.ChainID
}

func (p *pool) Token() common.Address {
	return p.token
}

func (p *pool) Address() common.Address {
	return p.address
}

func (p *pool) NetworkName() string {
	return p.networkName
}

func (p *pool) String() string {
	return fmt.Sprintf("%s:%s", p.address.String(), p.networkName)
}
