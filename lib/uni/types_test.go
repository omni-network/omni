package uni_test

import (
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/uni"

	"github.com/stretchr/testify/require"
)

func TestIsSolChain(t *testing.T) {
	t.Parallel()

	require.True(t, uni.IsSolChain(evmchain.IDSolana))
	require.True(t, uni.IsSolChain(evmchain.IDSolanaTest))
	require.True(t, uni.IsSolChain(evmchain.IDSolanaLocal))

	require.False(t, uni.IsSolChain(evmchain.IDBase))
	require.False(t, uni.IsSolChain(evmchain.IDOptimism))
	require.False(t, uni.IsSolChain(evmchain.IDArbitrumOne))
}
