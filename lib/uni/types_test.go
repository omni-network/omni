package uni_test

import (
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/uni"

	"github.com/stretchr/testify/require"
)

func TestIsSolChain(t *testing.T) {
	t.Parallel()

	require.True(t, uni.IsSVMChain(evmchain.IDSolana))
	require.True(t, uni.IsSVMChain(evmchain.IDSolanaTest))
	require.True(t, uni.IsSVMChain(evmchain.IDSolanaLocal))

	require.False(t, uni.IsSVMChain(evmchain.IDBase))
	require.False(t, uni.IsSVMChain(evmchain.IDOptimism))
	require.False(t, uni.IsSVMChain(evmchain.IDArbitrumOne))
}
