package relayer

import (
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/xchain"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestPersistState(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 5)
	path := filepath.Join(t.TempDir(), "state.json")

	ps := &State{
		cursors:  make(map[uint64]map[uint64]map[xchain.ConfLevel]uint64),
		filePath: path,
	}

	expected := make(map[uint64]map[uint64]map[xchain.ConfLevel]uint64)
	fuzzer.Fuzz(&expected)

	for dstChainID, inner1 := range expected {
		for srcChainID, inner2 := range inner1 {
			for confLevel, offset := range inner2 {
				err := ps.Persist(dstChainID, xchain.ChainVersion{ID: srcChainID, ConfLevel: confLevel}, offset)
				require.NoError(t, err)
			}
		}
	}

	load := func(t *testing.T) *State {
		t.Helper()
		loadedState, ok, err := LoadCursors(path)
		require.NoError(t, err)
		require.True(t, ok)
		require.NotNil(t, loadedState)

		return loadedState
	}

	require.EqualValues(t, expected, load(t).cursors)

	// Clear each destination
	for dstChainID := range expected {
		require.NotEmpty(t, ps.cursors[dstChainID])
		require.NotEmpty(t, load(t).cursors[dstChainID])

		require.NoError(t, ps.Clear(dstChainID))

		require.Empty(t, ps.cursors[dstChainID])
		require.Empty(t, load(t).cursors[dstChainID])
	}

	require.Empty(t, ps.cursors)
	require.Empty(t, load(t).cursors)
}
