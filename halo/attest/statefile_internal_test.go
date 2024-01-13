package attest

import (
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/xchain"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestFileStateFromHeights(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 64)
	path := filepath.Join(t.TempDir(), "state.json")

	state := &FileState{
		atts: make(map[uint64]xchain.Attestation),
		path: path,
	}

	chainsInState := []uint64{1, 2, 3}
	for _, id := range chainsInState {
		// Create a random attestation for each chain with known height.
		var att xchain.Attestation
		fuzzer.Fuzz(&att)
		att.SourceChainID = id
		att.BlockHeight = id

		require.NoError(t, state.Add(att))
	}

	var moreChains []uint64
	moreChains = append(moreChains, chainsInState...)
	moreChains = append(moreChains, 4, 5, 6)

	state2, err := LoadState(path)
	require.NoError(t, err)

	expectHeights := map[uint64]uint64{
		1: 2,
		2: 3,
		3: 4,
		4: 0,
		5: 0,
		6: 0,
	}
	actualHeights := fromHeights(state2, moreChains)
	require.Equal(t, expectHeights, actualHeights)
}

func TestStateFileAdd(t *testing.T) {
	t.Parallel()
	path := filepath.Join(t.TempDir(), "state.json")

	state := &FileState{
		atts: make(map[uint64]xchain.Attestation),
		path: path,
	}

	assert := func(t *testing.T, chainID, height uint64, ok bool) {
		t.Helper()
		att := xchain.Attestation{
			BlockHeader: xchain.BlockHeader{
				SourceChainID: chainID,
				BlockHeight:   height,
			},
		}
		err := state.Add(att)
		if ok {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}

	assert(t, 1, 1, true)
	assert(t, 1, 2, true)
	assert(t, 1, 3, true)
	assert(t, 1, 5, false)
	assert(t, 1, 3, false)
	assert(t, 1, 2, false)
	assert(t, 1, 4, true)
	assert(t, 1, 5, true)
	assert(t, 2, 1, true)
}
