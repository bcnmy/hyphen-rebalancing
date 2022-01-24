package pool

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/bcnmy/hyphen-arbitrage/internal/config"
)

type Manager interface {
	PoolName() string
	TokenName() string
	AccountName() string
	AccountAddress() common.Address
	PrivateKey() *ecdsa.PrivateKey
	PublicKey() *ecdsa.PublicKey
	Pools() []Pool
	Run(context.Context)
}

type Pool interface {
	NetworkName() string
	ChainID() uint64
	Token() common.Address
	Address() common.Address
}

func FromConfig(conf *config.Config, logger log.Logger) ([]Manager, error) {
	result := []Manager{}

	for poolName, poolConf := range conf.Pools {
		if len(poolConf.Addresses) < 2 {
			return nil, fmt.Errorf("pool config '%s' contains %d addresses, should be >= 2", poolName, len(poolConf.Addresses))
		}

		if len(poolConf.Accounts) == 0 {
			return nil, fmt.Errorf("pool config '%s' doesn't contains any accounts", poolName)
		}

		if len(poolConf.Tokens) == 0 {
			return nil, fmt.Errorf("pool config '%s' doesn't contains any tokens", poolName)
		}

		for _, tokenName := range poolConf.Tokens {
			tokenConf, ok := conf.Tokens[tokenName]
			if !ok {
				return nil, fmt.Errorf("pool config '%s' contains unknown token '%s'", poolName, tokenName)
			}

			for _, accountName := range poolConf.Accounts {
				accountConf, ok := conf.Accounts[accountName]
				if !ok {
					return nil, fmt.Errorf("pool config '%s' contains unknown account '%s'", poolName, accountName)
				}

				mgr := new(manager)
				mgr.poolName = poolName
				mgr.tokenName = tokenName
				mgr.accountName = accountName
				mgr.privateKey = accountConf.PrivateKey.Value
				mgr.pools = []Pool{}

				for networkName, poolAddress := range poolConf.Addresses {
					networkConf, ok := conf.Networks[networkName]
					if !ok {
						return nil, fmt.Errorf("pool config '%s' contains address for unknown network '%s'", poolName, networkName)
					}

					tokenAddress, ok := tokenConf.Addresses[networkName]
					if !ok {
						return nil, fmt.Errorf("pool config '%s' contains token '%s' which doesn't support network '%s'", poolName, tokenName, networkName)
					}

					mgr.pools = append(mgr.pools, &pool{
						networkName: networkName,
						network:     networkConf,
						token:       tokenAddress.Value,
						address:     poolAddress.Value,
					})
				}

				mgr.logger = log.With(
					logger,
					"pool", mgr.PoolName(),
					"account", mgr.AccountName(),
					"token", mgr.TokenName(),
				)

				result = append(result, mgr)
			}
		}
	}

	return result, nil
}

type manager struct {
	logger      log.Logger
	poolName    string
	tokenName   string
	accountName string
	privateKey  *ecdsa.PrivateKey
	pools       []Pool
}

func (m *manager) PoolName() string {
	return m.poolName
}

func (m *manager) TokenName() string {
	return m.tokenName
}

func (m *manager) AccountName() string {
	return m.accountName
}

func (m *manager) AccountAddress() common.Address {
	return crypto.PubkeyToAddress(m.privateKey.PublicKey)
}

func (m *manager) PrivateKey() *ecdsa.PrivateKey {
	return m.privateKey
}

func (m *manager) PublicKey() *ecdsa.PublicKey {
	return &m.privateKey.PublicKey
}

func (m *manager) Pools() []Pool {
	return m.pools
}

func (m *manager) Run(ctx context.Context) {
	level.Info(m.logger).Log("msg", "starting pool manager")
	<-ctx.Done()
	level.Info(m.logger).Log("msg", "stopping pool manager")
}

type pool struct {
	networkName string
	network     config.Network
	token       common.Address
	address     common.Address
}

func (p *pool) NetworkName() string {
	return p.networkName
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
