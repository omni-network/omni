package umath

import (
	"math/big"

	"golang.org/x/exp/constraints"
)

// New return a big.Int presentation of the provided uint64 or int64.
// It is a convenience function improving the big.Int API.
func New[N constraints.Integer](i N) *big.Int {
	if i < 0 {
		return big.NewInt(int64(i))
	}

	return new(big.Int).SetUint64(uint64(i))
}

// Clone returns a new big.Int with the same value as a.
func Clone(a *big.Int) *big.Int {
	return new(big.Int).Set(a)
}

// Sub returns a - b0 [ - b1 - b2 ...].
// It is a convenience function improving the big.Int API.
func Sub(a *big.Int, bs ...*big.Int) *big.Int {
	resp := Clone(a)
	for _, b := range bs {
		resp.Sub(resp, b)
	}

	return resp
}

// Add returns a + b0 [+ b1 + b2 ...].
// It is a convenience function improving the big.Int API.
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
// It is a convenience function improving the big.Int API.
func MulRaw[N constraints.Integer](a *big.Int, bs ...N) *big.Int {
	resp := Clone(a)
	for _, b := range bs {
		resp.Mul(resp, New(b))
	}

	return resp
}

// Div returns a / b0 [/ b1 / b2 ...].
// It is a convenience function improving the big.Int API.
func Div(a *big.Int, bs ...*big.Int) *big.Int {
	resp := Clone(a)
	for _, b := range bs {
		resp.Div(resp, b)
	}

	return resp
}

// DivRaw returns a / b0 [/ b1 / b2 ...].
// It is a convenience function improving the big.Int API.
func DivRaw[N constraints.Integer](a *big.Int, bs ...N) *big.Int {
	resp := Clone(a)
	for _, b := range bs {
		resp.Div(resp, New(b))
	}

	return resp
}

func IsZero(a *big.Int) bool {
	return a.Sign() == 0
}

func EQ(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}

// GT returns true if a > b.
// It is a convenience function improving the big.Int API.
func GT(a, b *big.Int) bool {
	return a.Cmp(b) == 1
}

// GTE returns true if a >= b.
// It is a convenience function improving the big.Int API.
func GTE(a, b *big.Int) bool {
	return a.Cmp(b) >= 0
}

// LT returns true if a < b.
// It is a convenience function improving the big.Int API.
func LT(a, b *big.Int) bool {
	return a.Cmp(b) == -1
}

// LTE returns true if a <= b.
// It is a convenience function improving the big.Int API.
func LTE(a, b *big.Int) bool {
	return a.Cmp(b) <= 0
}
