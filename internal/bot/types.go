package bot

import (
	"fmt"
	"math/big"

	"github.com/bcnmy/hyphen-arbitrage/internal/pool"
)

type arbitrationProfit struct {
	route          *arbitrationRoute
	netTokenProfit *big.Int
}

type arbitrationRoute struct {
	from   pool.Pool
	to     pool.Pool
	reward *potentialReward
	fee    *potentialFee
}

func (r *arbitrationRoute) String() string {
	return fmt.Sprintf("%s:%s", r.from.NetworkName(), r.to.NetworkName())
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
