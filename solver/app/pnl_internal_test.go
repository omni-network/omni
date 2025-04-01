package app

import (
	"testing"

	"github.com/omni-network/omni/lib/bi"

	"github.com/stretchr/testify/require"
)

//nolint:testifylint // Epsilon comparison not required
func TestToGweiF64(t *testing.T) {
	t.Parallel()

	require.Equal(t, 1e9, toGweiF64(bi.Ether(1), 18))
	require.Equal(t, 1e9, toGweiF64(bi.Dec6(1), 6))

	require.Equal(t, 5e9, toGweiF64(bi.Ether(5), 18))
	require.Equal(t, 5e9, toGweiF64(bi.Dec6(5), 6))

	require.Equal(t, 1e10, toGweiF64(bi.Ether(10), 18))
	require.Equal(t, 1e10, toGweiF64(bi.Dec6(10), 6))

	require.Equal(t, 1e8, toGweiF64(bi.Ether(0.1), 18))
	require.Equal(t, 1e8, toGweiF64(bi.Dec6(0.1), 6))

	require.Equal(t, 1e7, toGweiF64(bi.Ether(0.01), 18))
	require.Equal(t, 1e7, toGweiF64(bi.Dec6(0.01), 6))

	require.Equal(t, 1e10, toGweiF64(bi.Ether(1), 17))
	require.Equal(t, 1e11, toGweiF64(bi.Ether(1), 16))
}
