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

	require.Equal(t, bi.Ether(1), bi.ToWei(bi.Ether(1), 18))
	require.Equal(t, bi.Ether(1), bi.ToWei(bi.Dec6(1), 6))
	require.Equal(t, bi.Ether(1), bi.ToWei(bi.Gwei(1), 9))
	require.Equal(t, bi.Ether(1), bi.ToWei(bi.Wei(10), 1))
	require.Equal(t, bi.Ether(1), bi.ToWei(bi.Wei(1), 0))
}

func TestDec6(t *testing.T) {
	t.Parallel()

	require.Equal(t, bi.Wei(1_000_000), bi.Dec6(1))
	require.Equal(t, bi.Gwei(1), bi.Dec6(1_000))
	require.Equal(t, bi.Ether(1), bi.Dec6(1_000_000_000_000))

	require.Equal(t, bi.N(1), bi.Dec6(0.000_001))
	require.Equal(t, bi.Zero(), bi.Dec6(0.000_000_1))

	min1 := new(big.Int).Mul(bi.Dec6(1), big.NewInt(-1))
	min1K := new(big.Int).Mul(bi.Dec6(1), big.NewInt(-1_000))
	min1M := new(big.Int).Mul(bi.Dec6(1), big.NewInt(-1_000_000))
	min1G := new(big.Int).Mul(bi.Dec6(1), big.NewInt(-1_000_000_000))

	require.Equal(t, min1, bi.Dec6(-1))
	require.Equal(t, min1K, bi.Dec6(-1_000))
	require.Equal(t, min1M, bi.Dec6(-1_000_000))
	require.Equal(t, min1G, bi.Dec6(-1_000_000_000))
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

	require.Equal(t, 1.0, bi.ToF64(bi.Ether(1), 18))
	require.Equal(t, 1.0, bi.ToF64(bi.Dec6(1), 6))
	require.Equal(t, 5.0, bi.ToF64(bi.Ether(5), 18))
	require.Equal(t, 5.0, bi.ToF64(bi.Dec6(5), 6))
	require.Equal(t, 10.0, bi.ToF64(bi.Ether(10), 18))
	require.Equal(t, 10.0, bi.ToF64(bi.Dec6(10), 6))
	require.Equal(t, 0.1, bi.ToF64(bi.Ether(0.1), 18))
	require.Equal(t, 0.1, bi.ToF64(bi.Dec6(0.1), 6))
	require.Equal(t, 0.01, bi.ToF64(bi.Ether(0.01), 18))
	require.Equal(t, 0.01, bi.ToF64(bi.Dec6(0.01), 6))
	require.Equal(t, 10.0, bi.ToF64(bi.Ether(1), 17))
	require.Equal(t, 100.0, bi.ToF64(bi.Ether(1), 16))
}

func TestMulF64(t *testing.T) {
	t.Parallel()

	require.Equal(t, "0", bi.MulF64(bi.Wei(1), 0).String())
	require.Equal(t, "0", bi.MulF64(bi.Wei(1), 0.0000001).String())
	require.Equal(t, "0", bi.MulF64(bi.Wei(1), 0.11).String())
	require.Equal(t, "1", bi.MulF64(bi.Wei(1), 1).String())
	require.Equal(t, "1", bi.MulF64(bi.Wei(1), 1.5).String())
	require.Equal(t, "1", bi.MulF64(bi.Wei(1), 1.6).String())
	require.Equal(t, "2", bi.MulF64(bi.Wei(1), 2).String())
	require.Equal(t, "1000000", bi.MulF64(bi.Wei(1), 1_000_000.123).String())

	require.Equal(t, "0", bi.MulF64(bi.Gwei(1), 0).String())
	require.Equal(t, "99", bi.MulF64(bi.Gwei(1), 0.0000001).String())
	require.Equal(t, "110000000", bi.MulF64(bi.Gwei(1), 0.11).String())
	require.Equal(t, "1000000000", bi.MulF64(bi.Gwei(1), 1).String())
	require.Equal(t, "1500000000", bi.MulF64(bi.Gwei(1), 1.5).String())
	require.Equal(t, "1600000000", bi.MulF64(bi.Gwei(1), 1.6).String())
	require.Equal(t, "2000000000", bi.MulF64(bi.Gwei(1), 2).String())
	require.Equal(t, "1000000123000000", bi.MulF64(bi.Gwei(1), 1_000_000.123).String())

	require.Equal(t, "0", bi.MulF64(bi.Ether(1), 0).String())
	require.Equal(t, "99999999999", bi.MulF64(bi.Ether(1), 0.0000001).String())
	require.Equal(t, "110000000000000000", bi.MulF64(bi.Ether(1), 0.11).String())
	require.Equal(t, "1000000000000000000", bi.MulF64(bi.Ether(1), 1).String())
	require.Equal(t, "1500000000000000000", bi.MulF64(bi.Ether(1), 1.5).String())
	require.Equal(t, "1600000000000000088", bi.MulF64(bi.Ether(1), 1.6).String())
	require.Equal(t, "2000000000000000000", bi.MulF64(bi.Ether(1), 2).String())
	require.Equal(t, "1000000123000000021397504", bi.MulF64(bi.Ether(1), 1_000_000.123).String())

	big.NewInt(1e18)
	oneM := bi.Ether(1_000_000)
	require.Equal(t, "0", bi.MulF64(oneM, 0).String())
	require.Equal(t, "99999999999999995", bi.MulF64(oneM, 0.0000001).String())
	require.Equal(t, "110000000000000000555111", bi.MulF64(oneM, 0.11).String())
	require.Equal(t, "1000000000000000000000000", bi.MulF64(oneM, 1).String())
	require.Equal(t, "1500000000000000000000000", bi.MulF64(oneM, 1.5).String())
	require.Equal(t, "1600000000000000088817840", bi.MulF64(oneM, 1.6).String())
	require.Equal(t, "2000000000000000000000000", bi.MulF64(oneM, 2).String())
	require.Equal(t, "1000000123000000021420418531328", bi.MulF64(oneM, 1_000_000.123).String())
}
