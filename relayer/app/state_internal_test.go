package relayer

import (
	"path/filepath"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestPersistState(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 64)
	path := filepath.Join(t.TempDir(), "relayer.json")

	ps := &PersistentState{
		cursors:  make(map[uint64]uint64),
		filePath: path,
	}

	expected := make(map[uint64]uint64)
	fuzzer.Fuzz(&expected)

	for chainID, height := range expected {
		err := ps.Persist(chainID, height)
		require.NoError(t, err)
	}

	loadedState, err := Load(path)
	require.NoError(t, err)
	require.NotNil(t, loadedState)
	require.True(t, mapsEqual(expected, loadedState.Get()))
}

// mapsEqual compares two maps to check if they are equal
func mapsEqual(expected, actual map[uint64]uint64) bool {
	if len(expected) != len(actual) {
		return false
	}

	for key, val := range expected {
		if loadedVal, ok := actual[key]; !ok || loadedVal != val {
			return false
		}
	}

	return true
}
