// Package bi ("big int") provides an improved API for working with *big.Int values
// since the native API is clunky and error-prone (since it is mutable).
//
// Note that it doesn't do nil-checks, that should be done by the application
// when data enters the system as nil values aare invalid.
package bi

import (
	"math/big"

	"golang.org/x/exp/constraints"
)

// N returns a big.Int presentation of the provided uint64 or int64.
// N is an abbreviation of "New", it can also be thought of as "Number"; i=N.
func N[I constraints.Integer](i I) *big.Int {
	if i < 0 {
		return big.NewInt(int64(i))
	}

	return new(big.Int).SetUint64(uint64(i))
}

// Clone returns a new big.Int with the same value.
// This mitigates the mutable nature of big.Int.
func Clone(a *big.Int) *big.Int {
	return new(big.Int).Set(a)
}

// Sub returns a - b0 [ - b1 - b2 ...].
// Note that SubRaw isn't provided since subtracting literal weis is not common.
func Sub(a *big.Int, bs ...*big.Int) *big.Int {
	resp := Clone(a)
	for _, b := range bs {
		resp.Sub(resp, b)
	}

	return resp
}

// Add returns a + b0 [+ b1 + b2 ...].
// Note that AddRaw isn't provided since adding literal weis is not common.
func Add(a *big.Int, bs ...*big.Int) *big.Int {
	resp := Clone(a)
	for _, b := range bs {
		resp.Add(resp, b)
	}

	return resp
}

// Mul returns a * b0 [* b1 * b2 ...].
// It is a convenience function improving the big.Int API.
func Mul(a *big.Int, bs ...*big.Int) *big.Int {
	resp := Clone(a)
	for _, b := range bs {
		resp.Mul(resp, b)
	}

	return resp
}

// MulRaw returns a * b0 [* b1 * b2 ...].
func MulRaw[I constraints.Integer](a *big.Int, bs ...I) *big.Int {
	resp := Clone(a)
	for _, b := range bs {
		resp.Mul(resp, N(b))
	}

	return resp
}

// MulF64 returns a * b0 [* b1 * b2 ...].
// Note that floats are not accurate.
func MulF64(a *big.Int, bs ...float64) *big.Int {
	f := new(big.Float)
	f.SetMode(big.ToZero)
	f = f.SetInt(a)
	for _, b := range bs {
		f.Mul(f, big.NewFloat(b))
	}

	resp, _ := f.Int(new(big.Int))

	return resp
}

// Div returns a / b0 [/ b1 / b2 ...].
func Div(a *big.Int, bs ...*big.Int) *big.Int {
	resp := Clone(a)
	for _, b := range bs {
		resp.Div(resp, b)
	}

	return resp
}

// DivRaw returns a / b0 [/ b1 / b2 ...].
func DivRaw[I constraints.Integer](a *big.Int, bs ...I) *big.Int {
	resp := Clone(a)
	for _, b := range bs {
		resp = Div(resp, N(b))
	}

	return resp
}

// Mod returns a % b0 [% b1 % b2 ...].
func Mod(a *big.Int, bs ...*big.Int) *big.Int {
	resp := Clone(a)
	for _, b := range bs {
		resp.Mod(resp, b)
	}

	return resp
}

// ModRaw returns a % b0 [% b1 % b2 ...].
func ModRaw[I constraints.Integer](a *big.Int, bs ...I) *big.Int {
	resp := Clone(a)
	for _, b := range bs {
		resp.Mod(resp, N(b))
	}

	return resp
}

// IsZero returns true if i is zero.
func IsZero(i *big.Int) bool {
	return i.Sign() == 0
}

// EQ returns true if a == b.
func EQ(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}

// GT returns true if a > b.
func GT(a, b *big.Int) bool {
	return a.Cmp(b) == 1
}

// GTE returns true if a >= b.
func GTE(a, b *big.Int) bool {
	return a.Cmp(b) >= 0
}

// LT returns true if a < b.
func LT(a, b *big.Int) bool {
	return a.Cmp(b) == -1
}

// LTE returns true if a <= b.
func LTE(a, b *big.Int) bool {
	return a.Cmp(b) <= 0
}
