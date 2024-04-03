package relayer

import (
	"path/filepath"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestPersistState(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 5)
	path := filepath.Join(t.TempDir(), "state.json")

	ps := &State{
		cursors:  make(map[uint64]map[uint64]uint64),
		filePath: path,
	}

	expected := make(map[uint64]map[uint64]uint64)
	fuzzer.Fuzz(&expected)

	for dstChainID, sourceMap := range expected {
		for srcChainID, height := range sourceMap {
			err := ps.Persist(dstChainID, srcChainID, height)
			require.NoError(t, err)
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

	require.True(t, mapsEqual(expected, load(t).cursors))

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

func mapsEqual(map1, map2 map[uint64]map[uint64]uint64) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value := range map1 {
		if val2, ok := map2[key]; !ok || !subMapsEqual(value, val2) {
			return false
		}
	}

	return true
}

func subMapsEqual(subMap1, subMap2 map[uint64]uint64) bool {
	if len(subMap1) != len(subMap2) {
		return false
	}

	for k, v := range subMap1 {
		if v2, ok := subMap2[k]; !ok || v != v2 {
			return false
		}
	}

	return true
}
