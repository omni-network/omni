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

// MaxUint256 is the maximum value that can be represented by a uint256.
var MaxUint256 = func() *big.Int {
	// Copied from uint256 package.
	const twoPow256Sub1 = "115792089237316195423570985008687907853269984665640564039457584007913129639935"
	maxUint256, ok := new(big.Int).SetString(twoPow256Sub1, 10)
	if !ok {
		panic("invalid max uint256")
	}

	return maxUint256
}()

var MaxUint128 = func() *big.Int {
	const twoPow128Sub1 = "340282366920938463463374607431768211455"
	maxUint128, ok := new(big.Int).SetString(twoPow128Sub1, 10)
	if !ok {
		panic("invalid max uint128")
	}

	return maxUint128
}()

var MaxUint96 = func() *big.Int {
	const twoPow96Sub1 = "79228162514264337593543950335"
	maxUint96, ok := new(big.Int).SetString(twoPow96Sub1, 10)
	if !ok {
		panic("invalid max uint96")
	}

	return maxUint96
}()

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

func MustToUint32[N constraints.Integer](i N) uint32 {
	resp, err := ToUint32(i)
	if err != nil {
		panic(err)
	}

	return resp
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

// ToUint8 returns i as an uint8 or an error if it cannot be represented as such.
func ToUint8[N constraints.Integer](i N) (uint8, error) {
	if i < 0 {
		return 0, errors.New("underflow")
	} else if uint64(i) > math.MaxUint8 {
		return 0, errors.New("overflow")
	}

	return uint8(i), nil
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

func ToInt[N constraints.Integer](i N) (int, error) {
	if i < 0 {
		return 0, errors.New("underflow")
	} else if uint64(i) > math.MaxInt {
		return 0, errors.New("overflow")
	}

	return int(i), nil
}
