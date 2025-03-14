package umath_test

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/umath"

	"github.com/stretchr/testify/require"
)

func TestToWei(t *testing.T) {
	t.Parallel()

	require.Equal(t, umath.Wei, umath.GweiToWei(0.000_000_001))
	require.Equal(t, umath.Gwei, umath.GweiToWei(1))
	require.Equal(t, umath.Ether, umath.GweiToWei(1_000_000_000))

	require.Equal(t, umath.Wei, umath.EtherToWei(0.000_000_000_000_000_001))
	require.Equal(t, umath.Zero, umath.GweiToWei(0.000_000_000_1))
	require.Equal(t, umath.Ether, umath.EtherToWei(1))

	ether1K := new(big.Int).Mul(umath.Ether, big.NewInt(1_000))
	ether1M := new(big.Int).Mul(umath.Ether, big.NewInt(1_000_000))
	ether1G := new(big.Int).Mul(umath.Ether, big.NewInt(1_000_000_000))

	require.Equal(t, ether1K, umath.EtherToWei(1_000))
	require.Equal(t, ether1M, umath.EtherToWei(1_000_000))
	require.Equal(t, ether1G, umath.EtherToWei(1_000_000_000))

	min1 := new(big.Int).Mul(umath.Ether, big.NewInt(-1))
	min1K := new(big.Int).Mul(umath.Ether, big.NewInt(-1_000))
	min1M := new(big.Int).Mul(umath.Ether, big.NewInt(-1_000_000))
	min1G := new(big.Int).Mul(umath.Ether, big.NewInt(-1_000_000_000))

	require.Equal(t, min1, umath.EtherToWei(-1))
	require.Equal(t, min1K, umath.EtherToWei(-1_000))
	require.Equal(t, min1M, umath.EtherToWei(-1_000_000))
	require.Equal(t, min1G, umath.EtherToWei(-1_000_000_000))
}

//nolint:testifylint // Epsilon comparison not required
func TestWeiTo(t *testing.T) {
	t.Parallel()

	require.Equal(t, 0.0, umath.WeiToGweiF64(umath.Zero))
	require.Equal(t, 0.000_000_001, umath.WeiToGweiF64(umath.Wei))
	require.Equal(t, 1.0, umath.WeiToGweiF64(umath.Gwei))
	require.Equal(t, 1_000_000_000.0, umath.WeiToGweiF64(umath.Ether))

	require.Equal(t, 0.0, umath.WeiToEtherF64(umath.Zero))
	require.Equal(t, 0.000_000_000_000_000_001, umath.WeiToEtherF64(umath.Wei))
	require.Equal(t, 0.000_000_001, umath.WeiToEtherF64(umath.Gwei))
	require.Equal(t, 1.0, umath.WeiToEtherF64(umath.Ether))
}
