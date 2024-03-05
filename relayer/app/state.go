package relayer

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/omni-network/omni/lib/errors"

	"github.com/cometbft/cometbft/libs/tempfile"
)

// State represents the state of the relayer. It keeps track of the last submitted height for each source chain on each destination chain.
type State struct {
	mu       sync.Mutex
	filePath string
	cursors  map[uint64]map[uint64]uint64 // destChainID -> srcChainID -> height
}

func NewEmptyState(filePath string) *State {
	return &State{
		filePath: filePath,
		cursors:  make(map[uint64]map[uint64]uint64),
	}
}

// Get returns the current state.
func (s *State) Get() map[uint64]map[uint64]uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Create a copy of the map to avoid race conditions
	copyMap := make(map[uint64]map[uint64]uint64)
	for k, v := range s.cursors {
		innerMap := make(map[uint64]uint64)
		for k2, v2 := range v {
			innerMap[k2] = v2
		}
		copyMap[k] = innerMap
	}

	return copyMap
}

// GetHeight returns the last submitted height for the given destChainID and srcChainID.
func (s *State) GetHeight(dstID, srcID uint64) uint64 {
	srcMap, ok := s.Get()[dstID]
	if !ok {
		return 0
	}

	height, ok := srcMap[srcID]
	if !ok {
		return 0
	}

	return height
}

// Persist saves the given height for the given chainID.
func (s *State) Persist(dstID, srcID, height uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	srcMap, ok := s.cursors[dstID]
	if !ok {
		srcMap = make(map[uint64]uint64)
	}
	srcMap[srcID] = height

	s.cursors[dstID] = srcMap

	return s.saveUnsafe()
}

// saveUnsafe saves the state to disk. It is labeled as "unsafe" because it assumes the caller holds the necessary lock to ensure
// concurrent access safety. This function serializes the state to JSON format and atomically writes it to the specified file path.
func (s *State) saveUnsafe() error {
	bytes, err := json.Marshal(s.cursors)
	if err != nil {
		return errors.Wrap(err, "marshal file")
	}

	if err := tempfile.WriteFileAtomic(s.filePath, bytes, 0o600); err != nil {
		return errors.Wrap(err, "write persistent file")
	}

	return nil
}

// LoadCursors loads a file state from the given path.
func LoadCursors(path string) (*State, bool, error) {
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

	return &State{
		cursors:  cursors,
		filePath: path,
	}, true, nil
}
