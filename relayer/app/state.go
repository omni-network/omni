package relayer

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/omni-network/omni/lib/errors"

	"github.com/cometbft/cometbft/libs/tempfile"
)

type PersistentState struct {
	mu       sync.Mutex
	filePath string
	cursors  map[uint64]map[uint64]uint64 // destChainID -> srcChainID -> height
}

func NewPersistentState(filePath string) *PersistentState {
	return &PersistentState{
		filePath: filePath,
		cursors:  make(map[uint64]map[uint64]uint64),
	}
}

// Get returns the current state.
func (p *PersistentState) Get() map[uint64]map[uint64]uint64 {
	p.mu.Lock()
	defer p.mu.Unlock()
	// Create a copy of the map to avoid race conditions
	copyMap := make(map[uint64]map[uint64]uint64)
	for k, v := range p.cursors {
		innerMap := make(map[uint64]uint64)
		for k2, v2 := range v {
			innerMap[k2] = v2
		}
		copyMap[k] = innerMap
	}

	return copyMap
}

// Persist saves the given height for the given chainID.
func (p *PersistentState) Persist(srcID, dstID, height uint64) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	srcMap, ok := p.cursors[dstID]
	if !ok {
		srcMap = make(map[uint64]uint64)
	}
	srcMap[srcID] = height

	p.cursors[dstID] = srcMap

	return p.saveUnsafe()
}

// saveUnsafe saves the state to disk. It is labeled as "unsafe" because it assumes the caller holds the necessary lock to ensure
// concurrent access safety. This function serializes the state to JSON format and atomically writes it to the specified file path.
func (p *PersistentState) saveUnsafe() error {
	bytes, err := json.Marshal(p.cursors)
	if err != nil {
		return errors.Wrap(err, "marshal file")
	}

	if err := tempfile.WriteFileAtomic(p.filePath, bytes, 0o600); err != nil {
		return errors.Wrap(err, "write persistent file")
	}

	return nil
}

// LoadCursors loads a file state from the given path.
func LoadCursors(path string) (*PersistentState, bool, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false, nil
		}

		return nil, false, errors.Wrap(err, "read state file")
	}

	cursors := make(map[uint64]map[uint64]uint64)
	if err := json.Unmarshal(bytes, &cursors); err != nil {
		return nil, false, errors.Wrap(err, "unmarshal state file")
	}

	return &PersistentState{
		cursors:  cursors,
		filePath: path,
	}, true, nil
}
