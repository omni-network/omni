package relayer

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/cometbft/cometbft/libs/tempfile"
	"github.com/omni-network/omni/lib/errors"
)

type State interface {
	Get() map[uint64]uint64
	Persist(chainID uint64, height uint64) error
}

type PersistentState struct {
	mu       sync.Mutex
	filePath string
	cursors  map[uint64]uint64
}

// Get returns the current state.
func (p *PersistentState) Get() map[uint64]uint64 {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.cursors
}

// Persist saves the given height for the given chainID. It returns an error if the height is already saved.
func (p *PersistentState) Persist(chainID uint64, height uint64) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if existing, ok := p.cursors[chainID]; ok {
		return errors.New("height already saved", "chainID", chainID, "existing", existing, "new", height)
	}

	p.cursors[chainID] = height

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

// Load loads a file state from the given path.
func Load(path string) (*PersistentState, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "read state file")
	}

	cursors := make(map[uint64]uint64)
	if err := json.Unmarshal(bytes, &cursors); err != nil {
		return nil, errors.Wrap(err, "unmarshal state file")
	}

	return &PersistentState{
		cursors:  cursors,
		filePath: path,
	}, nil
}
