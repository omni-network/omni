// Package umath provides some useful unsigned math functions to prevent underflows.
// It also provides some type conversion functions to convert between different
// integer types to prevent underflows and overflows or overflows.
package umath

import (
	"math"
	"math/big"

	"github.com/omni-network/omni/lib/errors"

	"golang.org/x/exp/constraints"
)

// Subtract returns a - b and true if a >= b, otherwise 0 and false.
func Subtract(a, b uint64) (uint64, bool) {
	if a < b {
		return 0, false
	}

	return a - b, true
}

// SubtractOrZero returns a - b if a >= b, otherwise 0.
// This is a convenience function for inline usage.
func SubtractOrZero(a, b uint64) uint64 {
	resp, _ := Subtract(a, b)
	return resp
}

// NewBigInt return a big.Int version of the provided uint64.
// This is a convenience function to avoid gosec complaining about big.NewInt(int64(i))).
func NewBigInt(i uint64) *big.Int {
	return new(big.Int).SetUint64(i)
}

// Len returns the length of the slice as a uint64.
// This convenience function to avoid gosec complaining about uint64(len(slice)).
func Len[T any](slice []T) uint64 {
	l := len(slice)
	if l < 0 {
		panic("impossible")
	}

	return uint64(l)
}

// ToUint64 returns i as an uint64 or an error if it cannot be represented as such.
func ToUint64[N constraints.Integer](i N) (uint64, error) {
	if i < 0 {
		return 0, errors.New("underflow")
	}

	return uint64(i), nil
}

// ToInt64 returns i as an int64 or an error if it cannot be represented as such.
func ToInt64[N constraints.Integer](n N) (int64, error) {
	// All negative int values are valid int64
	if n < 0 {
		return int64(n), nil
	}

	if uint64(n) > math.MaxInt64 {
		return 0, errors.New("overflow")
	}

	return int64(n), nil
}

// ToUint32 returns i as an uint32 or an error if it cannot be represented as such.
func ToUint32[N constraints.Integer](i N) (uint32, error) {
	if i < 0 {
		return 0, errors.New("underflow")
	} else if uint64(i) > math.MaxUint32 {
		return 0, errors.New("overflow")
	}

	return uint32(i), nil
}

// ToInt32 returns i as an int32 or an error if it cannot be represented as such.
func ToInt32[N constraints.Integer](i N) (int32, error) {
	// Using float64 for int32 is fine since rounding not a problem.
	if float64(i) > math.MaxInt32 {
		return 0, errors.New("overflow")
	} else if float64(i) < math.MinInt32 {
		return 0, errors.New("underflow")
	}

	return int32(i), nil
}
