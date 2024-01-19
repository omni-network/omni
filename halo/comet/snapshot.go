package comet

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"sync"

	"github.com/omni-network/omni/lib/errors"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/tempfile"
)

const (
	metadataFileName  = "snapshot_metadata.json"
	snapshotChunkSize = 100 * (1 << 10) // 100KB

	// Keep only the most recent 10 snapshots. Older snapshots are pruned.
	maxSnapshotCount = 10
	snapshotFormat   = 1
)

// SnapshotStore stores state sync snapshots. Snapshots are stored simply as
// JSON files, and chunks are generated on-the-fly by splitting the JSON data
// into fixed-size chunks.
type SnapshotStore struct {
	// Immutable config fields (not persisted to disk)
	dir          string
	metadataPath string

	// Mutable state fields (persisted to disk)
	mu       sync.RWMutex
	metadata []abci.Snapshot
}

// NewSnapshotStore creates a new snapshot store.
// It loads existing metadata if it exists.
// It ensures the snapshot store directory exists.
func NewSnapshotStore(dir string) (*SnapshotStore, error) {
	// Ensure the snapshot store directory exists
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, errors.Wrap(err, "create snapshot store dir")
	}

	metadataPath := filepath.Join(dir, metadataFileName)

	// Load metadata if it exists
	metadata, err := maybeLoadMetadata(metadataPath)
	if err != nil {
		return nil, err
	}

	store := SnapshotStore{
		dir:          dir,
		metadataPath: metadataPath,
		metadata:     metadata,
	}

	return &store, nil
}

// Create creates a snapshot of the given application state.
func (s *SnapshotStore) Create(state *State) (abci.Snapshot, error) {
	bz, height, stateHash, err := state.Export()
	if err != nil {
		return abci.Snapshot{}, err
	}

	snapshot := abci.Snapshot{
		Height: height,
		Format: snapshotFormat,
		Hash:   stateHash[:],
		Chunks: countChunks(bz),
	}

	err = os.WriteFile(s.snapshotPath(height), bz, 0o644) //nolint:gosec // 0o644 is the default umask
	if err != nil {
		return abci.Snapshot{}, errors.Wrap(err, "write snapshot file")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Add and persist snapshot metadata
	s.metadata = append(s.metadata, snapshot)
	if err = s.saveMetadataUnsafe(); err != nil {
		return abci.Snapshot{}, err
	}

	return snapshot, nil
}

// Prune removes old snapshots ensuring only the maxSnapshotCount recent n snapshots remain.
func (s *SnapshotStore) Prune() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Snapshots are appended to the metadata struct, hence pruning removes from
	// the front of the slice
	keepFromIdx := len(s.metadata) - maxSnapshotCount

	if keepFromIdx <= 0 {
		return nil // Keep everything, nothing to prune
	}

	for i := 0; i < keepFromIdx; i++ {
		height := s.metadata[i].Height
		if err := os.Remove(s.snapshotPath(height)); err != nil {
			return errors.Wrap(err, "remove snapshot file")
		}
	}

	// Update metadata by removing the pruned snapshots
	remaining := make([]abci.Snapshot, len(s.metadata[keepFromIdx:]))
	copy(remaining, s.metadata[keepFromIdx:])

	s.metadata = remaining

	return s.saveMetadataUnsafe()
}

// List returns a copy of all the snapshot metadata.
func (s *SnapshotStore) List() []abci.Snapshot {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return slices.Clone(s.metadata)
}

// LoadChunk loads a snapshot chunk from disk.
func (s *SnapshotStore) LoadChunk(height uint64, format uint32, chunk uint32) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, snapshot := range s.metadata {
		if snapshot.Height != height || snapshot.Format != format {
			continue
		}

		bz, err := os.ReadFile(s.snapshotPath(height))
		if err != nil {
			return nil, errors.Wrap(err, "read snapshot file")
		}

		return newChunk(bz, chunk), nil
	}

	return nil, errors.New("snapshot not found")
}

// snapshotPath returns the path to the snapshot file for a given height.
func (s *SnapshotStore) snapshotPath(height uint64) string {
	return filepath.Join(s.dir, fmt.Sprintf("%v.json", height))
}

// saveMetadataUnsafe saves snapshot metadata to disk.
// It is unsafe since it assumes the lock is held.
func (s *SnapshotStore) saveMetadataUnsafe() error {
	bz, err := json.Marshal(s.metadata)
	if err != nil {
		return errors.Wrap(err, "marshal snapshot metadata")
	}

	if err := tempfile.WriteFileAtomic(s.metadataPath, bz, 0o644); err != nil {
		return errors.Wrap(err, "write snapshot metadata")
	}

	return nil
}

// maybeLoadMetadata loads snapshot metadata from disk if it exists.
// It returns an empty slice if the file does not exist.
func maybeLoadMetadata(filePath string) ([]abci.Snapshot, error) {
	bz, err := os.ReadFile(filePath)
	if errors.Is(err, os.ErrNotExist) {
		// No file on disk, return empty metadata
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "read snapshot metadata")
	}

	var resp []abci.Snapshot
	if err = json.Unmarshal(bz, &resp); err != nil {
		return nil, errors.Wrap(err, "unmarshal snapshot metadata")
	}

	return resp, nil
}

// newChunk returns the chunk at a given index from the full byte slice.
func newChunk(bz []byte, index uint32) []byte {
	start := int(index * snapshotChunkSize)
	end := int((index + 1) * snapshotChunkSize)

	if start >= len(bz) {
		return nil
	} else if end >= len(bz) {
		return bz[start:]
	}

	return bz[start:end]
}

// countChunks calculates the number of chunks in the byte slice.
func countChunks(bz []byte) uint32 {
	return uint32(math.Ceil(float64(len(bz)) / snapshotChunkSize))
}
