package bot

import (
	"math/big"

	"github.com/bcnmy/hyphen-arbitrage/internal/pool"
)

type arbitrationProfit struct {
	route          *arbitrationRoute
	netTokenProfix *big.Int
}

type arbitrationRoute struct {
	from   pool.Pool
	to     pool.Pool
	reward *potentialReward
	fee    *potentialFee
}

type potentialReward struct {
	balance   *big.Int
	allowance *big.Int
	deposit   *big.Int
	reward    *big.Int
}

type potentialFee struct {
	tokenFee *big.Int
}
