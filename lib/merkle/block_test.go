package merkle_test

import (
	"math/rand"
	"testing"

	"github.com/omni-network/omni/lib/merkle"
	"github.com/omni-network/omni/lib/xchain"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestBlockTree(t *testing.T) {
	t.Parallel()

	var block xchain.Block
	fuzz.New().NilChance(0).NumElements(1, 64).Fuzz(&block)

	tree, err := xchain.NewBlockTree(block)
	require.NoError(t, err)

	// Prove some random messages
	for end := 1; end < len(block.Msgs); end++ {
		start := rand.Intn(end)

		multi, err := tree.Proof(block.BlockHeader, block.Msgs[start:end])
		require.NoError(t, err)

		// Verify the proof
		root, err := merkle.ProcessMultiProof(multi)
		require.NoError(t, err)

		require.Equal(t, tree.Root(), root)
	}
}
