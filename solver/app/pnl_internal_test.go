package app

import (
	"testing"

	"github.com/omni-network/omni/lib/bi"

	"github.com/stretchr/testify/require"
)

func TestTknAmtToGweiF64(t *testing.T) {
	t.Parallel()

	require.InEpsilon(t, 1e9, tknAmtToGweiF64(bi.Ether(1), 18), 1e-9)
	require.InEpsilon(t, 1e9, tknAmtToGweiF64(bi.Dec6(1), 6), 1e-9)

	require.InEpsilon(t, 5e9, tknAmtToGweiF64(bi.Ether(5), 18), 1e-9)
	require.InEpsilon(t, 5e9, tknAmtToGweiF64(bi.Dec6(5), 6), 1e-9)

	require.InEpsilon(t, 1e10, tknAmtToGweiF64(bi.Ether(10), 18), 1e-9)
	require.InEpsilon(t, 1e10, tknAmtToGweiF64(bi.Dec6(10), 6), 1e-9)

	require.InEpsilon(t, 1e8, tknAmtToGweiF64(bi.Ether(0.1), 18), 1e-9)
	require.InEpsilon(t, 1e8, tknAmtToGweiF64(bi.Dec6(0.1), 6), 1e-9)

	require.InEpsilon(t, 1e7, tknAmtToGweiF64(bi.Ether(0.01), 18), 1e-9)
	require.InEpsilon(t, 1e7, tknAmtToGweiF64(bi.Dec6(0.01), 6), 1e-9)

	require.InEpsilon(t, 1e10, tknAmtToGweiF64(bi.Ether(1), 17), 1e-9)
	require.InEpsilon(t, 1e11, tknAmtToGweiF64(bi.Ether(1), 16), 1e-9)
}
