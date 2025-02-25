package util

import (
	"fmt"
	"math"
	"math/big"
)

var (
	// ether1 is 1 token in wei (18 decimals).
	Ether1 = dec(1, 18)

	// million Gwei.
	MilliEther = dec(1, 15)
)

func dec(amt float64, decimals int) *big.Int {
	unit := math.Pow10(decimals)

	p := amt * unit

	_, dec := math.Modf(p)
	if dec != 0 {
		panic(fmt.Sprintf("amt float64 must be an int multiple of 1e%d", decimals))
	}

	return new(big.Int).SetUint64(uint64(p))
}
