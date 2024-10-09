package omnitoken_test

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/contracts/omnitoken"

	"github.com/stretchr/testify/require"
)

// TestTotalSupply confirms total supply is 100M.
func TestTotalSupply(t *testing.T) {
	t.Parallel()

	totalSupply, ok := new(big.Int).SetString("100000000000000000000000000", 10)
	require.True(t, ok)
	require.Equal(t, omnitoken.TotalSupply, totalSupply)
}
