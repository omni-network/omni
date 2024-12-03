package umath_test

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/umath"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func TestMaxUint256(t *testing.T) {
	t.Parallel()
	_, overflow := uint256.FromBig(umath.MaxUint256)
	require.False(t, overflow, "don't expect overflow")

	maxPlus1 := new(big.Int).Add(umath.MaxUint256, big.NewInt(1))
	_, overflow = uint256.FromBig(maxPlus1)
	require.True(t, overflow, "expect overflow")
}

func TestUintTo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		i        uint64
		okInt64  bool
		okUint64 bool
		okInt32  bool
		okUint32 bool
	}{
		{math.MaxUint64, false, true, false, false},
		{math.MaxInt64 + 1, false, true, false, false},
		{math.MaxInt64, true, true, false, false},
		{math.MaxUint32 + 1, true, true, false, false},
		{math.MaxUint32, true, true, false, true},
		{math.MaxInt32 + 1, true, true, false, true},
		{math.MaxInt32, true, true, true, true},
		{1, true, true, true, true},
		{0, true, true, true, true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.i), func(t *testing.T) {
			t.Parallel()
			if _, err := umath.ToInt64(test.i); (err == nil) != test.okInt64 {
				require.Fail(t, "ToInt64", "unexpected result: %v", err)
			}
			if _, err := umath.ToInt32(test.i); (err == nil) != test.okInt32 {
				require.Fail(t, "ToInt32", "unexpected result: %v", err)
			}
			if _, err := umath.ToUint64(test.i); (err == nil) != test.okUint64 {
				require.Fail(t, "ToUint64", "unexpected result: %v", err)
			}
			if _, err := umath.ToUint32(test.i); (err == nil) != test.okUint32 {
				require.Fail(t, "ToUint32", "unexpected result: %v", err)
			}
		})
	}
}

func TestIntTo(t *testing.T) {
	t.Parallel()
	tests := []struct {
		i        int64
		okInt64  bool
		okUint64 bool
		okInt32  bool
		okUint32 bool
	}{
		{math.MaxInt64, true, true, false, false},
		{math.MaxUint32 + 1, true, true, false, false},
		{math.MaxUint32, true, true, false, true},
		{math.MaxInt32 + 1, true, true, false, true},
		{math.MaxInt32, true, true, true, true},
		{1, true, true, true, true},
		{0, true, true, true, true},
		{-1, true, false, true, false},
		{math.MinInt32, true, false, true, false},
		{math.MinInt32 - 1, true, false, false, false},
		{math.MinInt64, true, false, false, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.i), func(t *testing.T) {
			t.Parallel()
			if _, err := umath.ToInt64(test.i); (err == nil) != test.okInt64 {
				require.Fail(t, "ToInt64", "unexpected result: %v", err)
			}
			if _, err := umath.ToInt32(test.i); (err == nil) != test.okInt32 {
				require.Fail(t, "ToInt32", "unexpected result: %v", err)
			}
			if _, err := umath.ToUint64(test.i); (err == nil) != test.okUint64 {
				require.Fail(t, "ToUint64", "unexpected result: %v", err)
			}
			if _, err := umath.ToUint32(test.i); (err == nil) != test.okUint32 {
				require.Fail(t, "ToUint32", "unexpected result: %v", err)
			}
		})
	}
}
