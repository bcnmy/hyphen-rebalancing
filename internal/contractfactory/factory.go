package contractfactory

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bcnmy/hyphen-rebalancing/internal/generated/erc20"
	"github.com/bcnmy/hyphen-rebalancing/internal/generated/hyphen"
	"github.com/bcnmy/hyphen-rebalancing/internal/pool"
	"github.com/bcnmy/hyphen-rebalancing/internal/utils"
)

const (
	dialTimeout = time.Second * 60
	callTimeout = time.Second * 30
)

// Factory allows to create contract instances from pool interface.
// Contract instances are cached per chain id.
type Factory interface {
	LiquidityPool(p pool.Pool) (*hyphen.LiquidityPool, error)
	LiquidityProviders(p pool.Pool) (*hyphen.LiquidityProviders, error)
	TokenManager(p pool.Pool) (*hyphen.TokenManager, error)
	ERC20(p pool.Pool) (*erc20.Erc20, error)
	Client(p pool.Pool) (*ethclient.Client, error)
}

func New() (Factory, error) {
	return &factory{
		liquidityPoolMap:      make(map[uint64]*hyphen.LiquidityPool),
		liquidityProvidersMap: make(map[uint64]*hyphen.LiquidityProviders),
		tokenManagerMap:       make(map[uint64]*hyphen.TokenManager),
		erc20Map:              make(map[uint64]*erc20.Erc20),

		clientMap: make(map[uint64]*ethclient.Client),
	}, nil
}

type factory struct {
	liquidityPoolMap   map[uint64]*hyphen.LiquidityPool
	liquidityPoolMutex sync.Mutex

	liquidityProvidersMap   map[uint64]*hyphen.LiquidityProviders
	liquidityProvidersMutex sync.Mutex

	tokenManagerMap   map[uint64]*hyphen.TokenManager
	tokenManagerMutex sync.Mutex

	erc20Map   map[uint64]*erc20.Erc20
	erc20Mutex sync.Mutex

	clientMap   map[uint64]*ethclient.Client
	clientMutex sync.Mutex
}

func (f *factory) LiquidityPool(p pool.Pool) (*hyphen.LiquidityPool, error) {
	f.liquidityPoolMutex.Lock()
	defer f.liquidityPoolMutex.Unlock()

	liquidityPool, ok := f.liquidityPoolMap[f.getChainID(p)]
	if ok {
		return liquidityPool, nil
	}

	client, err := f.Client(p)
	if err != nil {
		return nil, err
	}

	liquidityPool, err = hyphen.NewLiquidityPool(p.Address(), client)
	if err != nil {
		return nil, err
	}

	f.liquidityPoolMap[f.getChainID(p)] = liquidityPool

	return liquidityPool, nil
}

func (f *factory) LiquidityProviders(p pool.Pool) (*hyphen.LiquidityProviders, error) {
	f.liquidityProvidersMutex.Lock()
	defer f.liquidityProvidersMutex.Unlock()

	liquidityProviders, ok := f.liquidityProvidersMap[f.getChainID(p)]
	if ok {
		return liquidityProviders, nil
	}

	liquidityPool, err := f.LiquidityPool(p)
	if err != nil {
		return nil, err
	}

	client, err := f.Client(p)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), callTimeout)
	defer cancel()
	liquidityProvidersAddress, err := liquidityPool.LiquidityProviders(utils.NewCallOpts(ctx))
	if err != nil {
		return nil, err
	}

	liquidityProviders, err = hyphen.NewLiquidityProviders(liquidityProvidersAddress, client)
	if err != nil {
		return nil, err
	}

	f.liquidityProvidersMap[f.getChainID(p)] = liquidityProviders

	return liquidityProviders, nil
}

func (f *factory) TokenManager(p pool.Pool) (*hyphen.TokenManager, error) {
	f.tokenManagerMutex.Lock()
	defer f.tokenManagerMutex.Unlock()

	tokenManager, ok := f.tokenManagerMap[f.getChainID(p)]
	if ok {
		return tokenManager, nil
	}

	liquidityPool, err := f.LiquidityPool(p)
	if err != nil {
		return nil, err
	}

	client, err := f.Client(p)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), callTimeout)
	defer cancel()
	tokenManagerAddress, err := liquidityPool.TokenManager(utils.NewCallOpts(ctx))
	if err != nil {
		return nil, err
	}

	tokenManager, err = hyphen.NewTokenManager(tokenManagerAddress, client)
	if err != nil {
		return nil, err
	}

	f.tokenManagerMap[f.getChainID(p)] = tokenManager

	return tokenManager, nil
}

func (f *factory) ERC20(p pool.Pool) (*erc20.Erc20, error) {
	f.erc20Mutex.Lock()
	defer f.erc20Mutex.Unlock()

	token, ok := f.erc20Map[f.getChainID(p)]
	if ok {
		return token, nil
	}

	client, err := f.Client(p)
	if err != nil {
		return nil, err
	}

	token, err = erc20.NewErc20(p.Token(), client)
	if err != nil {
		return nil, err
	}

	f.erc20Map[f.getChainID(p)] = token

	return token, nil
}

func (f *factory) Client(p pool.Pool) (*ethclient.Client, error) {
	f.clientMutex.Lock()
	defer f.clientMutex.Unlock()

	client, ok := f.clientMap[f.getChainID(p)]
	if ok {
		return client, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()

	client, err := p.Dial(ctx)
	if err != nil {
		return nil, err
	}

	f.clientMap[f.getChainID(p)] = client

	return client, nil
}

func (f *factory) getChainID(p pool.Pool) uint64 {
	return p.ChainID().Uint64()
}
