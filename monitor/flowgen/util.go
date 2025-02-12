package flowgen

import (
	"math"
	"math/big"

	"github.com/omni-network/omni/lib/tokens"
)

func spend(token tokens.Token, amt *big.Int) Spend { return Spend{token: amt} }

var (
	// ether1 is 1 token in wei (18 decimals).
	ether1 = dec18(1)
)

func dec18(amt float64) *big.Int {
	const unit = 1e18

	p := amt * unit

	_, dec := math.Modf(p)
	if dec != 0 {
		panic("amt float64 must be an int multiple of 1e18")
	}

	return new(big.Int).SetUint64(uint64(p))
}
