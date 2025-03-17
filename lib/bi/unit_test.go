package bi_test

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/bi"

	"github.com/stretchr/testify/require"
)

func TestToWei(t *testing.T) {
	t.Parallel()

	require.Equal(t, bi.Wei(1), bi.Gwei(0.000_000_001))
	require.Equal(t, bi.Wei(1_000_000_000), bi.Gwei(1))
	require.Equal(t, bi.Ether(1), bi.Gwei(1_000_000_000))

	require.Equal(t, bi.Wei(1), bi.Ether(0.000_000_000_000_000_001))
	require.Equal(t, bi.Zero(), bi.Gwei(0.000_000_000_1))
	require.Equal(t, bi.Gwei(100_000_000), bi.Ether(0.1))

	ether1K := new(big.Int).Mul(bi.Ether(1), big.NewInt(1_000))
	ether1M := new(big.Int).Mul(bi.Ether(1), big.NewInt(1_000_000))
	ether1G := new(big.Int).Mul(bi.Ether(1), big.NewInt(1_000_000_000))

	require.Equal(t, ether1K, bi.Ether(1_000))
	require.Equal(t, ether1M, bi.Ether(1_000_000))
	require.Equal(t, ether1G, bi.Ether(1_000_000_000))

	min1 := new(big.Int).Mul(bi.Ether(1), big.NewInt(-1))
	min1K := new(big.Int).Mul(bi.Ether(1), big.NewInt(-1_000))
	min1M := new(big.Int).Mul(bi.Ether(1), big.NewInt(-1_000_000))
	min1G := new(big.Int).Mul(bi.Ether(1), big.NewInt(-1_000_000_000))

	require.Equal(t, min1, bi.Ether(-1))
	require.Equal(t, min1K, bi.Ether(-1_000))
	require.Equal(t, min1M, bi.Ether(-1_000_000))
	require.Equal(t, min1G, bi.Ether(-1_000_000_000))
}

//nolint:testifylint // Epsilon comparison not required
func TestWeiTo(t *testing.T) {
	t.Parallel()

	require.Equal(t, 0.0, bi.ToGweiF64(bi.Zero()))
	require.Equal(t, 0.000_000_001, bi.ToGweiF64(bi.Wei(1)))
	require.Equal(t, 1.0, bi.ToGweiF64(bi.Gwei(1)))
	require.Equal(t, 1_000_000_000.0, bi.ToGweiF64(bi.Ether(1)))

	require.Equal(t, 0.0, bi.ToEtherF64(bi.Zero()))
	require.Equal(t, 0.000_000_000_000_000_001, bi.ToEtherF64(bi.Wei(1)))
	require.Equal(t, 0.000_000_001, bi.ToEtherF64(bi.Gwei(1)))
	require.Equal(t, 1.0, bi.ToEtherF64(bi.Ether(1)))
	require.Equal(t, 1_000_000.0, bi.ToEtherF64(bi.Ether(1_000_000)))
	require.Equal(t, 0.1, bi.ToEtherF64(bi.Ether(0.1)))
}
