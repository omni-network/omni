package comet

import (
	"crypto/sha256"
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/tempfile"
)

const (
	pubKeyLen         = 33
	stateFileName     = "halo_state.json"
	prevStateFileName = "halo_app_state.json"
)

type validator struct {
	PubKey [pubKeyLen]byte
	Power  int64
}

// State is the halo application state.
type State struct {
	// Immutable config fields (not persisted to disk)
	filePath        string
	prevFilePath    string
	persistInterval uint64 // Persist state every internal

	// Mutable state fields (persisted to disk)
	mu            sync.RWMutex
	initialHeight uint64
	height        uint64
	valSetID      uint64
	validators    []validator
	pendingAggs   []xchain.AggAttestation
	approvedAggs  []xchain.AggAttestation

	hash [32]byte // Calculated from above fields
}

// LoadOrGenState loads the state from the dir or creates a new state. It will persist to that dir.
func LoadOrGenState(dir string, persistInterval uint64) (*State, error) {
	filePath := filepath.Join(dir, stateFileName)
	prevFilePath := filepath.Join(dir, prevStateFileName)

	stateJSON, err := loadStateJSON(filePath, prevFilePath)
	if errors.Is(err, os.ErrNotExist) { //nolint:revive // Catch and swallow this error.
		// No file on disk, use empty state
	} else if err != nil {
		return nil, err
	}

	state := &State{
		filePath:        filePath,
		prevFilePath:    prevFilePath,
		persistInterval: persistInterval,
	}
	populateStateFromJSON(state, stateJSON)

	state.hash, err = state.calcHashUnsafe() // No need to lock since just created above.
	if err != nil {
		return nil, errors.Wrap(err, "calculate state hash")
	}

	return state, nil
}

// InitValidators sets consensus validators.
// It returns an error if the validators are already set.
// This implies validators can only be set once and never change.
func (s *State) InitValidators(vals []abci.ValidatorUpdate) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.valSetID++ // Update from 0 to 1.

	for _, v := range vals {
		if len(v.PubKey.GetEd25519()) > 0 {
			return errors.New("ed25519 keys not supported")
		}

		pk := v.PubKey.GetSecp256K1()
		if len(pk) != pubKeyLen {
			return errors.New("invalid pubkey length")
		}

		s.validators = append(s.validators, validator{
			PubKey: [pubKeyLen]byte(pk),
			Power:  v.Power,
		})
	}

	var err error
	s.hash, err = s.calcHashUnsafe()
	if err != nil {
		return errors.Wrap(err, "calculate state hash")
	}

	return nil
}

// AddAttestations adds the provided aggregates to the state, moving pending to approved if applicable.
func (s *State) AddAttestations(aggregates []xchain.AggAttestation) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Merge the new aggregates with existing, either pending or approved, add non-matching to remaining.
	var remaining []xchain.AggAttestation
	s.pendingAggs, remaining = mergeAggregates(s.pendingAggs, aggregates)
	s.approvedAggs, remaining = mergeAggregates(s.approvedAggs, remaining)

	// Add remaining non-matching to pending.
	s.pendingAggs = sortAggregates(append(s.pendingAggs, remaining...))

	// Check which pending are newly approved, and which are still pending.
	var stillPending, newApproved []xchain.AggAttestation
	for _, agg := range s.pendingAggs {
		if isApproved(agg, s.validators) {
			newApproved = append(newApproved, agg)
		} else {
			stillPending = append(stillPending, agg)
		}
	}

	// Update pending and approved.
	s.pendingAggs = sortAggregates(stillPending)
	s.approvedAggs = sortAggregates(append(s.approvedAggs, newApproved...))
}

// ApprovedAggregates returns a copy of the approved aggregates.
// For testing purposes only.
func (s *State) ApprovedAggregates() []xchain.AggAttestation {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return slices.Clone(s.approvedAggs)
}

// ApprovedFrom returns a sequential range of approved aggregates from the provided chain ID and height.
// It returns at most 100 aggregates. Their block heights are sequentially increasing.
func (s *State) ApprovedFrom(chainID uint64, height uint64) []xchain.AggAttestation {
	s.mu.Lock()
	defer s.mu.Unlock()

	next := height

	const limit = 100

	resp := make([]xchain.AggAttestation, 0, limit)
	for _, agg := range s.approvedAggs { // approvedAggs is sorted by block height.
		if agg.SourceChainID != chainID || agg.BlockHeight != next {
			continue
		}

		resp = append(resp, agg)
		next++

		if uint64(len(resp)) >= limit {
			break
		}
	}

	return resp
}

// saveUnsafe saves the state to disk. It is unsafe since it assumes the lock is held.
func (s *State) saveUnsafe() error {
	bz, err := s.marshalJSONUnsafe()
	if err != nil {
		return err
	}

	// Replace previous file with current, if it exists.
	if _, err := os.Stat(s.filePath); err == nil {
		if err := os.Rename(s.filePath, s.prevFilePath); err != nil {
			return errors.Wrap(err, "replace previous state")
		}
	}

	if err := tempfile.WriteFileAtomic(s.filePath, bz, 0o644); err != nil {
		return errors.Wrap(err, "write state file")
	}

	return nil
}

// Hash returns the existing state hash.
func (s *State) Hash() []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.hash[:]
}

// Height returns the current height of the state.
func (s *State) Height() uint64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.height
}

// Export returns the state json, used for state sync snapshots.
// Additionally, returns the current height and hash.
func (s *State) Export() ([]byte, uint64, [32]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	bz, err := s.marshalJSONUnsafe()
	if err != nil {
		return nil, 0, [32]byte{}, err
	}

	return bz, s.height, s.hash, nil
}

// Import imports state JSON bytes (overwriting existing state). It used for InitChain.AppStateBytes and
// state sync snapshots. It also saves the state once imported.
func (s *State) Import(height uint64, jsonBytes []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var stateJSON stateJSON
	err := json.Unmarshal(jsonBytes, &stateJSON)
	if err != nil {
		return errors.Wrap(err, "unmarshal state")
	} else if stateJSON.Height != height {
		return errors.New("mismatching height")
	}

	*s = State{
		filePath:        s.filePath,
		prevFilePath:    s.prevFilePath,
		persistInterval: s.persistInterval,
	}
	populateStateFromJSON(s, stateJSON)

	s.hash, err = s.calcHashUnsafe()
	if err != nil {
		return errors.Wrap(err, "calculate state hash")
	}

	return s.saveUnsafe()
}

// Finalize is called after applying a block, it updates the height and returns the new state hash.
func (s *State) Finalize() ([32]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.height > 0 {
		s.height++
	} else if s.initialHeight > 0 {
		s.height = s.initialHeight
	} else {
		s.height = 1
	}

	var err error
	s.hash, err = s.calcHashUnsafe()
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "calculate state hash")
	}

	return s.hash, nil
}

// Commit commits the current state, returning the current height.
func (s *State) Commit() (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.persistInterval > 0 && s.height%s.persistInterval == 0 {
		err := s.saveUnsafe()
		if err != nil {
			return 0, err
		}
	}

	return s.height, nil
}

// Rollback rolls back the state to the previous state file.
func (s *State) Rollback() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	stateJSON, err := loadStateJSON(s.prevFilePath, "")
	if err != nil {
		return errors.Wrap(err, "read prev state file")
	}

	*s = State{
		filePath:        s.filePath,
		prevFilePath:    s.prevFilePath,
		persistInterval: s.persistInterval,
	}
	populateStateFromJSON(s, stateJSON)

	s.hash, err = s.calcHashUnsafe()
	if err != nil {
		return errors.Wrap(err, "calculate state hash")
	}

	return nil
}

// marshalJSONUnsafe marshals the state to JSON.
// Note this is unsafe, it assumes the lock is held.
func (s *State) marshalJSONUnsafe() ([]byte, error) {
	stateJSON := &stateJSON{
		InitialHeight: s.initialHeight,
		Height:        s.height,
		ValSetID:      s.valSetID,
		Validators:    s.validators,
		PendingAggs:   s.pendingAggs,
		ApprovedAggs:  s.approvedAggs,
	}

	bz, err := json.Marshal(stateJSON)
	if err != nil {
		return nil, errors.Wrap(err, "marshal state")
	}

	return bz, nil
}

// calcHash calculate the hash of the state.
// It is unsafe since it assumes the lock is held.
func (s *State) calcHashUnsafe() ([32]byte, error) {
	// TODO(corver): Improve this, make it more robust.
	bz, err := s.marshalJSONUnsafe()
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "marshal state")
	}

	return sha256.Sum256(bz), nil
}

// stateJSON represents the JSON state file.
type stateJSON struct {
	InitialHeight uint64                  `json:"initial_height"`
	Height        uint64                  `json:"height"`
	ValSetID      uint64                  `json:"val_set_id"`
	Validators    []validator             `json:"validators"`
	PendingAggs   []xchain.AggAttestation `json:"pending_aggs"`
	ApprovedAggs  []xchain.AggAttestation `json:"approved_aggs"`
}

// loadState loads stateJSON from filePath, trying fallback if filePath does not exist.
func loadStateJSON(filePath, fallback string) (stateJSON, error) {
	bz, err := os.ReadFile(filePath)
	if errors.Is(err, os.ErrNotExist) && fallback != "" {
		bz, err = os.ReadFile(fallback)
		if err != nil {
			return stateJSON{}, errors.Wrap(err, "read both state files")
		}
	} else if err != nil {
		return stateJSON{}, errors.Wrap(err, "read state file")
	}

	var state stateJSON
	if err := json.Unmarshal(bz, &state); err != nil {
		return stateJSON{}, errors.Wrap(err, "unmarshal state file")
	}

	return state, nil
}

// populateStateFromJSON populates the state fields from the provided stateJSON.
func populateStateFromJSON(state *State, stateJSON stateJSON) {
	state.initialHeight = stateJSON.InitialHeight
	state.height = stateJSON.Height
	state.valSetID = stateJSON.ValSetID
	state.validators = stateJSON.Validators
	state.pendingAggs = stateJSON.PendingAggs
	state.approvedAggs = stateJSON.ApprovedAggs
}
