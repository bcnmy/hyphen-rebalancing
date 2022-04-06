package pool

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kit/log"

	"github.com/bcnmy/hyphen-rebalancing/internal/config"
)

// Manager contains several interconnected liquidity pools on different chains
type Manager interface {
	Pools() []Pool
	PrivateKey() *ecdsa.PrivateKey
	PublicKey() *ecdsa.PublicKey
	AccountAddress() common.Address
	MinimalProfit() *big.Int
	InfiniteApprove() bool

	PoolName() string
	TokenName() string
	AccountName() string
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

				privateKey, err := accountConf.GetPrivateKey()
				if err != nil {
					return nil, fmt.Errorf("account '%s' error: %s", accountName, err.Error())
				}

				mgr := new(manager)
				mgr.pools = []Pool{}
				mgr.privateKey = privateKey
				mgr.minimalProfit = tokenConf.MinimalProfit
				mgr.infiniteApprove = tokenConf.InfiniteApprove

				mgr.poolName = poolName
				mgr.tokenName = tokenName
				mgr.accountName = accountName

				for networkName, poolAddresses := range poolConf.Addresses {
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
						address:     poolAddresses.Value,
					})
				}

				result = append(result, mgr)
			}
		}
	}

	return result, nil
}

type manager struct {
	pools           []Pool
	privateKey      *ecdsa.PrivateKey
	minimalProfit   *big.Int
	infiniteApprove bool

	poolName    string
	tokenName   string
	accountName string
}

func (m *manager) Pools() []Pool {
	return m.pools
}

func (m *manager) PrivateKey() *ecdsa.PrivateKey {
	return m.privateKey
}

func (m *manager) PublicKey() *ecdsa.PublicKey {
	return &m.privateKey.PublicKey
}

func (m *manager) AccountAddress() common.Address {
	return crypto.PubkeyToAddress(m.privateKey.PublicKey)
}

func (m *manager) MinimalProfit() *big.Int {
	return m.minimalProfit
}

func (m *manager) InfiniteApprove() bool {
	return m.infiniteApprove
}

func (m *manager) AccountName() string {
	return m.accountName
}

func (m *manager) PoolName() string {
	return m.poolName
}

func (m *manager) TokenName() string {
	return m.tokenName
}
