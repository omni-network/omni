package rebalance_test

import (
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/solver/rebalance"

	"github.com/stretchr/testify/require"
)

func TestNoCircularDependencies(t *testing.T) {
	t.Parallel()

	chains := []uint64{
		evmchain.IDEthereum,
		evmchain.IDBase,
		evmchain.IDArbitrumOne,
		evmchain.IDOptimism,
		evmchain.IDMantle,
	}

	for _, chainID := range chains {
		require.False(t, hasCycle(chainID, make(map[uint64]bool)),
			"Found circular dependency starting from chain %d (%s)",
			chainID, evmchain.Name(chainID))
	}
}

func hasCycle(chainID uint64, seen map[uint64]bool) bool {
	if seen[chainID] {
		return true
	}

	seen[chainID] = true

	for _, dep := range rebalance.GetDependents(chainID) {
		if hasCycle(dep, seen) {
			return true
		}
	}

	return false
}
