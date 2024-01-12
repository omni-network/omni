package attest

import (
	"encoding/json"
	"os"
	"sort"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/libs/tempfile"
)

// State contains the latest attestations for each chain.
// It is used to prevent double signing and as a cursor store.
type State interface {
	// Get returns all the attestations in the state.
	Get() []xchain.Attestation
	// Add adds the attestation to the state, persisting it if necessary.
	Add(attestations xchain.Attestation) error
}

var _ State = (*FileState)(nil)

// FileState is a simple file-backed implementation of the State interface.
type FileState struct {
	atts map[uint64]xchain.Attestation
	path string
}

// Add adds the attestation to the state, persisting it.
func (s *FileState) Add(att xchain.Attestation) error {
	if existing, ok := s.atts[att.SourceChainID]; ok {
		if existing.BlockHeight >= att.BlockHeight {
			return errors.New("attestation height already exists",
				"existing", existing.BlockHeight, "new", att.BlockHeight)
		}

		if existing.BlockHeight+1 != att.BlockHeight {
			return errors.New("attestation is not sequential",
				"existing", existing.BlockHeight, "new", att.BlockHeight)
		}
	}

	s.atts[att.SourceChainID] = att

	return s.save()
}

// Get returns a copy of all the attestations in the state in a deterministic order.
func (s *FileState) Get() []xchain.Attestation {
	atts := make([]xchain.Attestation, 0, len(s.atts))
	for _, att := range s.atts {
		atts = append(atts, att)
	}

	sort.Slice(atts, func(i, j int) bool {
		if atts[i].SourceChainID != atts[j].SourceChainID {
			return atts[i].SourceChainID < atts[j].SourceChainID
		}

		return atts[i].BlockHeight < atts[j].BlockHeight
	})

	return atts
}

func (s *FileState) save() error {
	bz, err := json.Marshal(s.Get())
	if err != nil {
		return errors.Wrap(err, "marshal state file")
	}

	if err := tempfile.WriteFileAtomic(s.path, bz, 0o600); err != nil {
		return errors.Wrap(err, "write state file")
	}

	return nil
}

// LoadState loads a file state from the given path.
func LoadState(path string) (*FileState, error) {
	bz, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "read state file")
	}

	var atts []xchain.Attestation
	if err := json.Unmarshal(bz, &atts); err != nil {
		return nil, errors.Wrap(err, "unmarshal state file")
	}

	attMap := make(map[uint64]xchain.Attestation)
	for _, att := range atts {
		attMap[att.SourceChainID] = att
	}

	if len(attMap) != len(atts) {
		return nil, errors.New("invalid state file, duplicate chains")
	}

	return &FileState{
		atts: attMap,
		path: path,
	}, nil
}
