package umath

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/params"

	"golang.org/x/exp/constraints"
)

var (
	// Zero is a big.Int zero value.
	Zero = new(big.Int)

	// One is a big.Int value representing 1.
	One = big.NewInt(1)

	// Wei is a big.Int value representing 1 wei.
	Wei = big.NewInt(1)

	// Gwei is a big.Int value representing 1 gwei (1e9 wei).
	Gwei = big.NewInt(params.GWei)

	// Ether is a big.Int value representing 1 ether (1e18 wei).
	Ether = big.NewInt(params.Ether)
)

// number is a generic number type for int*/uint*/float*.
type number interface {
	constraints.Float | constraints.Integer
}

// WeiToEtherF64 converts big.Int wei to float64 ether (wei/1e18).
// Note that this is not accurate, only use for logging/metrics/display, not math.
func WeiToEtherF64(wei *big.Int) float64 {
	f, _ := wei.Float64()
	return f / params.Ether
}

// WeiToGweiF64 converts big.Int wei to float64 gwei (wei/1e9).
// Note that this is not accurate, only use for logging/metrics/display, not math.
func WeiToGweiF64(wei *big.Int) float64 {
	f, _ := wei.Float64()
	return f / params.GWei
}

// GweiToWei convert a gwei float/int/uint to wei big.Int.
// Note this can be lossy for large floats.
func GweiToWei[N number](i N) *big.Int {
	if iU64, ok := numToU64(i); ok {
		return new(big.Int).Mul(Gwei, New(iU64))
	} else if iI64, ok := numToI64(i); ok {
		return new(big.Int).Mul(Gwei, big.NewInt(iI64))
	}

	wei, _ := new(big.Float).Mul(
		big.NewFloat(float64(i)),
		big.NewFloat(params.GWei)).
		Int(nil)

	return wei
}

// EtherToWei convert an ether float/int/uint to wei big.Int.
// Note this can be lossy for large floats.
func EtherToWei[N number](i N) *big.Int {
	if iU64, ok := numToU64(i); ok {
		return new(big.Int).Mul(Ether, New(iU64))
	} else if iI64, ok := numToI64(i); ok {
		return new(big.Int).Mul(Ether, big.NewInt(iI64))
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
