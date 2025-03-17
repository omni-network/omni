package bi

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/params"

	"golang.org/x/exp/constraints"
)

// Zero returns a new big.Int value representing 0.
func Zero() *big.Int {
	return big.NewInt(0)
}

// One returns a new big.Int value representing 1.
func One() *big.Int {
	return big.NewInt(1)
}

// Wei is an alias for New.
// It doesn't support floats like Ether or Gwei, since wei is the smallest unit.
func Wei[I constraints.Integer](i I) *big.Int {
	return N(i)
}

// number is a generic number type for int*/uint*/float*.
type number interface {
	constraints.Float | constraints.Integer
}

// ToEtherF64 converts big.Int wei to float64 ether (wei/1e18).
// Note that this is not accurate, only use for logging/metrics/display, not math.
func ToEtherF64(wei *big.Int) float64 {
	if GTE(wei, Ether(1)) {
		// Avoid float division of large numbers, rather trim to gwei first.
		wgei := Div(wei, Gwei(1))
		f, _ := wgei.Float64()

		return f / params.GWei
	}

	f, _ := wei.Float64()

	return f / params.Ether
}

// ToGweiF64 converts big.Int wei to float64 gwei (wei/1e9).
// Note that this is not accurate, only use for logging/metrics/display, not math.
func ToGweiF64(wei *big.Int) float64 {
	f, _ := wei.Float64()
	return f / params.GWei
}

// Gwei converts gwei float/int/uint in to wei big.Int; i * 1e9.
// Note this can be lossy for large floats.
func Gwei[N number](i N) *big.Int {
	if iU64, ok := numToU64(i); ok {
		return MulRaw(big.NewInt(params.GWei), iU64)
	} else if iI64, ok := numToI64(i); ok {
		return MulRaw(big.NewInt(params.GWei), iI64)
	}

	wei, _ := new(big.Float).Mul(
		big.NewFloat(float64(i)),
		big.NewFloat(params.GWei)).
		Int(nil)

	return wei
}

// Ether converts ether float/int/uint in to wei big.Int; i * 1e18.
// Note this can be lossy for large floats.
func Ether[N number](i N) *big.Int {
	if iU64, ok := numToU64(i); ok {
		return MulRaw(big.NewInt(params.Ether), iU64)
	} else if iI64, ok := numToI64(i); ok {
		return MulRaw(big.NewInt(params.Ether), iI64)
	}

	wei, _ := new(big.Float).Mul(
		big.NewFloat(float64(i)),
		big.NewFloat(params.Ether)).
		Int(nil)

	return wei
}

// numToU64 converts a number to uint64 if lossless.
func numToU64[N number](i N) (uint64, bool) {
	if i < 0 || float64(i) > math.MaxUint64 {
		return 0, false
	}

	// Test whether converting to uint64 and back to float64 is lossless.
	iU64 := uint64(i)

	return iU64, float64(i) == float64(iU64)
}

// numToI64 converts a number to int64 if lossless.
func numToI64[N number](i N) (int64, bool) {
	if float64(i) < math.MinInt64 || float64(i) > math.MaxInt64 {
		return 0, false
	}

	// Test whether converting to int64 and back to float64 is lossless.
	iI64 := int64(i)

	return iI64, float64(i) == float64(iI64)
}
