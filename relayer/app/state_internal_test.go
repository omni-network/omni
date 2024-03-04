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

	ps := &PersistentState{
		cursors:  make(map[uint64]map[uint64]uint64),
		filePath: path,
	}

	expected := make(map[uint64]map[uint64]uint64)
	fuzzer.Fuzz(&expected)

	for dstChainID, sourceMap := range expected {
		for srcChainID, height := range sourceMap {
			err := ps.Persist(srcChainID, dstChainID, height)
			require.NoError(t, err)
		}
	}

	loadedState, err := Load(path)
	require.NoError(t, err)
	require.NotNil(t, loadedState)
	require.True(t, mapsEqual(expected, loadedState.Get()))
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
