package attest_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/lib/xchain"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestAttester(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 64)

	path := filepath.Join(t.TempDir(), "state.json")
	err := attest.GenEmptyStateFile(path)
	require.NoError(t, err)

	pk := k1.GenPrivKey()

	const (
		chain1 = 1
		chain2 = 2
		chain3 = 3
	)

	// reloadAttester reloads the attester from disk.
	reloadAttester := func(t *testing.T, from1, from2 uint64) *attest.Attester {
		t.Helper()
		p := make(stubProvider)
		a, err := attest.LoadAttester(ctx, pk, path, p, map[uint64]string{chain1: "", chain2: "", chain3: ""})
		require.NoError(t, err)

		require.EqualValues(t, from1, p[chain1])
		require.EqualValues(t, from2, p[chain2])
		require.Empty(t, p[chain3])

		return a
	}

	a := reloadAttester(t, 0, 0)

	add := func(t *testing.T, chainID, height uint64) {
		t.Helper()
		var block xchain.Block
		fuzzer.Fuzz(&block)
		block.BlockHeader = xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
		}

		err := a.Attest(ctx, block)
		require.NoError(t, err)
	}

	propose := func(chainID, height uint64) {
		header := xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
		}

		err := a.SetProposed([]xchain.BlockHeader{header})
		require.NoError(t, err)
	}

	commit := func(chainID, height uint64) {
		header := xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
		}

		err := a.SetCommitted([]xchain.BlockHeader{header})
		require.NoError(t, err)
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

	// Reload
	a = reloadAttester(t, 4, 0)

	// Add noise
	add(t, 2, 1)

	// Assert all are available
	available(t, 1, 1, true)
	available(t, 1, 2, true)
	available(t, 1, 3, true)
	available(t, 2, 1, true)

	// Reload
	a = reloadAttester(t, 4, 2)

	// Propose and commit 3 only
	propose(1, 3)
	commit(1, 3)

	// Assert 1, 2 are available
	available(t, 1, 1, true)
	available(t, 1, 2, true)
	available(t, 1, 3, false)

	// Reload
	a = reloadAttester(t, 4, 2)

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

	// Reload
	a = reloadAttester(t, 4, 2)

	// Commit 2 and noise
	commit(1, 2)
	commit(2, 1)

	// All committed
	bz, err := os.ReadFile(path)
	require.NoError(t, err)
	var stateJSON map[string]any
	require.NoError(t, json.Unmarshal(bz, &stateJSON))

	require.Len(t, stateJSON, 3)
	require.Empty(t, stateJSON["available"])
	require.Empty(t, stateJSON["proposed"])
	require.Len(t, stateJSON["committed"], 2) // One per chain

	// Ensure non-sequential asserts fail
	addErr := func(t *testing.T, chainID, height uint64) {
		t.Helper()
		var block xchain.Block
		fuzzer.Fuzz(&block)
		block.BlockHeader = xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
		}

		err := a.Attest(ctx, block)
		require.Error(t, err)
	}

	addErr(t, 1, 3)
	addErr(t, 1, 2)
	addErr(t, 1, 5)
}

var _ xchain.Provider = stubProvider{}

type stubProvider map[uint64]uint64

func (p stubProvider) Subscribe(_ context.Context, chainID uint64, fromHeight uint64, _ xchain.ProviderCallback) error {
	p[chainID] = fromHeight
	return nil
}

func (stubProvider) GetBlock(context.Context, uint64, uint64) (xchain.Block, bool, error) {
	panic("unexpected")
}

func (stubProvider) GetSubmittedCursor(context.Context, uint64, uint64) (xchain.StreamCursor, bool, error) {
	panic("unexpected")
}

func (stubProvider) GetEmittedCursor(context.Context, uint64, uint64) (xchain.StreamCursor, bool, error) {
	panic("unexpected")
}
