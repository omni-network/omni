package xchain_test

import (
	"math/rand"
	"testing"

	"github.com/omni-network/omni/lib/xchain"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

// TestMsgTreeNoVerify is mostly for test coverage.
// Go test coverage doesn't count tests by other packages.
// See ../merkle/block_test.go for a more thorough test that actually verifies proof.
func TestMsgTreeNoVerify(t *testing.T) {
	t.Parallel()

	var msgs []xchain.Msg
	fuzz.New().NilChance(0).NumElements(1, 64).Fuzz(&msgs)

	tree, err := xchain.NewMsgTree(msgs)
	require.NoError(t, err)

	// Prove some random messages
	for end := 1; end < len(msgs); end++ {
		start := rand.Intn(end)

		_, err := tree.Proof(msgs[start:end])
		require.NoError(t, err)
	}
}
