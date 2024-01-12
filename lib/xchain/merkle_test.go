package xchain_test

import (
	"math/rand"
	"testing"

	"github.com/omni-network/omni/lib/xchain"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

// TestBlockTreeNoVerify is mostly for test coverage.
// Go test coverage doesn't count tests by other packages.
// See ../merkle/block_test.go for a more thorough test that actually verifies proof.
func TestBlockTreeNoVerify(t *testing.T) {
	t.Parallel()

	var block xchain.Block
	fuzz.New().NilChance(0).NumElements(1, 64).Fuzz(&block)

	tree, err := xchain.NewBlockTree(block)
	require.NoError(t, err)

	// Prove some random messages
	for end := 1; end < len(block.Msgs); end++ {
		start := rand.Intn(end)

		_, err := tree.Proof(block.BlockHeader, block.Msgs[start:end])
		require.NoError(t, err)
	}
}
