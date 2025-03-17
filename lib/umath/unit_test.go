package umath_test

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/umath"

	"github.com/stretchr/testify/require"
)

func TestToWei(t *testing.T) {
	t.Parallel()

	require.Equal(t, umath.Wei(1), umath.Gwei(0.000_000_001))
	require.Equal(t, umath.Wei(1_000_000_000), umath.Gwei(1))
	require.Equal(t, umath.Ether(1), umath.Gwei(1_000_000_000))

	require.Equal(t, umath.Wei(1), umath.Ether(0.000_000_000_000_000_001))
	require.Equal(t, umath.Zero(), umath.Gwei(0.000_000_000_1))
	require.Equal(t, umath.Gwei(100_000_000), umath.Ether(0.1))

	ether1K := new(big.Int).Mul(umath.Ether(1), big.NewInt(1_000))
	ether1M := new(big.Int).Mul(umath.Ether(1), big.NewInt(1_000_000))
	ether1G := new(big.Int).Mul(umath.Ether(1), big.NewInt(1_000_000_000))

	require.Equal(t, ether1K, umath.Ether(1_000))
	require.Equal(t, ether1M, umath.Ether(1_000_000))
	require.Equal(t, ether1G, umath.Ether(1_000_000_000))

	min1 := new(big.Int).Mul(umath.Ether(1), big.NewInt(-1))
	min1K := new(big.Int).Mul(umath.Ether(1), big.NewInt(-1_000))
	min1M := new(big.Int).Mul(umath.Ether(1), big.NewInt(-1_000_000))
	min1G := new(big.Int).Mul(umath.Ether(1), big.NewInt(-1_000_000_000))

	require.Equal(t, min1, umath.Ether(-1))
	require.Equal(t, min1K, umath.Ether(-1_000))
	require.Equal(t, min1M, umath.Ether(-1_000_000))
	require.Equal(t, min1G, umath.Ether(-1_000_000_000))
}

//nolint:testifylint // Epsilon comparison not required
func TestWeiTo(t *testing.T) {
	t.Parallel()

	require.Equal(t, 0.0, umath.ToGweiF64(umath.Zero()))
	require.Equal(t, 0.000_000_001, umath.ToGweiF64(umath.Wei(1)))
	require.Equal(t, 1.0, umath.ToGweiF64(umath.Gwei(1)))
	require.Equal(t, 1_000_000_000.0, umath.ToGweiF64(umath.Ether(1)))

	require.Equal(t, 0.0, umath.ToEtherF64(umath.Zero()))
	require.Equal(t, 0.000_000_000_000_000_001, umath.ToEtherF64(umath.Wei(1)))
	require.Equal(t, 0.000_000_001, umath.ToEtherF64(umath.Gwei(1)))
	require.Equal(t, 1.0, umath.ToEtherF64(umath.Ether(1)))
	require.Equal(t, 1_000_000.0, umath.ToEtherF64(umath.Ether(1_000_000)))
	require.Equal(t, 0.1, umath.ToEtherF64(umath.Ether(0.1)))
}
