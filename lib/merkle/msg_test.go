package merkle_test

import (
	"math/rand"
	"testing"

	"github.com/omni-network/omni/lib/merkle"
	"github.com/omni-network/omni/lib/xchain"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestMsgTree(t *testing.T) {
	t.Parallel()

	var msgs []xchain.Msg
	fuzz.New().NilChance(0).NumElements(1, 64).Fuzz(&msgs)

	tree, err := xchain.NewMsgTree(msgs)
	require.NoError(t, err)

	// Prove some random messages
	for end := 1; end < len(msgs); end++ {
		start := rand.Intn(end)

		multi, err := tree.Proof(msgs[start:end])
		require.NoError(t, err)
		// Verify the proof

		root, err := merkle.ProcessMultiProof(multi)
		require.NoError(t, err)

		require.Equal(t, tree.MsgRoot(), root)
	}
}
