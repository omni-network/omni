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

	// Ensure msg.LogIndex is increasing
	for i := 1; i < len(msgs); i++ {
		msgs[i].LogIndex = msgs[i-1].LogIndex + 1 + uint64(rand.Intn(1000))
	}

	tree, err := xchain.NewMsgTree(msgs)
	require.NoError(t, err)

	// Prove some random messages
	for end := 1; end < len(msgs); end++ {
		start := rand.Intn(end)

		_, err := tree.Proof(msgs[start:end])
		require.NoError(t, err)
	}
}

func TestOrderedMsgs(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 64)

	isOrdered := func(msgs []xchain.Msg) bool {
		for i := 1; i < len(msgs); i++ {
			if msgs[i].LogIndex < msgs[i-1].LogIndex {
				return false
			}
		}

		return true
	}

	var msgs []xchain.Msg
	fuzzer.Fuzz(&msgs)
	for isOrdered(msgs) {
		fuzzer.Fuzz(&msgs)
	}

	require.False(t, isOrdered(msgs))
	_, err := xchain.NewMsgTree(msgs)
	require.ErrorContains(t, err, "not ordered")
}
