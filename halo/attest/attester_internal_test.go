package attest

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto/secp256k1"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestAttester(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 64)

	a := NewAttesterForT(t, secp256k1.GenPrivKey())

	add := func(t *testing.T, chainID, height uint64) {
		t.Helper()
		var block xchain.Block
		fuzzer.Fuzz(&block)
		block.BlockHeader = xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
		}

		err := a.Attest(ctx, &block)
		require.NoError(t, err)
	}

	propose := func(chainID, height uint64) {
		header := xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
		}

		a.SetProposed([]xchain.BlockHeader{header})
	}

	commit := func(chainID, height uint64) {
		header := xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
		}

		a.SetCommitted([]xchain.BlockHeader{header})
	}

	available := func(t *testing.T, chainID, height uint64, ok bool) {
		t.Helper()

		var found bool
		for _, att := range a.GetAvailable() {
			if att.SourceChainID == chainID && att.BlockHeight == height {
				found = true
				break
			}
		}

		require.Equal(t, ok, found)
	}

	// Add 1, 2, 3
	add(t, 1, 1)
	add(t, 1, 2)
	add(t, 1, 3)

	// Add noise
	add(t, 2, 1)

	// Assert all are available
	available(t, 1, 1, true)
	available(t, 1, 2, true)
	available(t, 1, 3, true)
	available(t, 2, 1, true)

	// Propose and commit 3 only
	propose(1, 3)
	commit(1, 3)

	// Assert 1, 2 are available
	available(t, 1, 1, true)
	available(t, 1, 2, true)
	available(t, 1, 3, false)

	// Propose 1
	propose(1, 1)
	available(t, 1, 1, false)

	// Propose 2 (resets 1)
	propose(1, 2)
	available(t, 1, 1, true)
	available(t, 1, 2, false)

	// Commit 1 (resets 2)
	commit(1, 1)
	available(t, 1, 1, false)
	available(t, 1, 2, true)

	// Commit 2 and noise
	commit(1, 2)
	commit(2, 1)

	// All committed
	a.mu.Lock()
	defer a.mu.Unlock()
	require.Empty(t, a.available)
	require.Empty(t, a.proposed)
}
