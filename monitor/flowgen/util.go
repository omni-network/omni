package flowgen

import (
	"math/big"

	"github.com/omni-network/omni/lib/tokens"
)

func spend(token tokens.Token, amt float64) Spend {
	return Spend{token: unit(amt, token.Decimals)}
}

func unit(amt float64, decimals uint) *big.Int {
	g := amt * float64(1^decimals)
	return new(big.Int).SetUint64(uint64(g))
}
