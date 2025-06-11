package evmchain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerify(t *testing.T) {
	t.Parallel()

	uniqNames := make(map[string]bool)
	uniqChainIDs := make(map[uint64]bool)

	for chainID, metadata := range static {
		require.Equal(t, chainID, metadata.ChainID)
		require.NotEmpty(t, metadata.BlockPeriod)

		if metadata.Name != omniEVMName {
			require.False(t, uniqNames[metadata.Name])
		}
		require.False(t, uniqChainIDs[metadata.ChainID])
		uniqNames[metadata.Name] = true
		uniqChainIDs[metadata.ChainID] = true
	}
}

func TestIsSolChain(t *testing.T) {
	t.Parallel()

	require.True(t, IsSVM(IDSolana))
	require.True(t, IsSVM(IDSolanaDevnet))
	require.True(t, IsSVM(IDSolanaLocal))

	require.False(t, IsSVM(IDBase))
	require.False(t, IsSVM(IDOptimism))
	require.False(t, IsSVM(IDArbitrumOne))
}
