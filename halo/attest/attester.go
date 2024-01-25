package attest

import (
	"context"
	"encoding/json"
	"os"
	"slices"
	"sort"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/libs/tempfile"
)

var _ Service = (*Attester)(nil)

// Attester implements the Service interface.
// It is responsible for creating and persisting attestations.
// The goal is to ensure "all blocks are attested to".
type Attester struct {
	path    string
	privKey crypto.PrivKey

	mu        sync.Mutex
	available []xchain.Attestation
	proposed  []xchain.Attestation
	committed []xchain.Attestation
}

// GenEmptyStateFile generates an empty attester state file at the given path.
// This must be called before LoadAttester.
func GenEmptyStateFile(path string) error {
	return (&Attester{path: path}).saveUnsafe()
}

// LoadAttester returns a new attester with state loaded from disk.
func LoadAttester(ctx context.Context, privKey crypto.PrivKey, path string, provider xchain.Provider, chains []uint64,
) (*Attester, error) {
	if len(privKey.PubKey().Bytes()) != 33 {
		return nil, errors.New("invalid private key")
	}

	s, err := loadState(path)
	if err != nil {
		return nil, err
	}

	a := Attester{
		privKey: privKey,
		path:    path,

		available: s.Available,
		proposed:  s.Proposed,
		committed: s.Committed,
	}

	// Subscribe to latest block provider
	for _, chainID := range chains {
		var fromHeight uint64
		if latest, ok := a.latestByChainUnsafe(chainID); ok {
			fromHeight = latest.BlockHeight + 1
		}
		err := provider.Subscribe(ctx, chainID, fromHeight, a.Attest)
		if err != nil {
			return nil, err
		}
	}

	return &a, nil
}

// Attest creates an attestation for the given block and adds it to the internal state.
func (a *Attester) Attest(_ context.Context, block xchain.Block) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	att, err := CreateAttestation(a.privKey, block)
	if err != nil {
		return err
	}

	// Ensure attestation is sequential and not a duplicate.
	latest, ok := a.latestByChainUnsafe(att.SourceChainID)
	if ok && latest.BlockHeight >= att.BlockHeight {
		return errors.New("attestation height already exists",
			"latest", latest.BlockHeight, "new", att.BlockHeight)
	} else if ok && latest.BlockHeight+1 != att.BlockHeight {
		return errors.New("attestation is not sequential",
			"existing", latest.BlockHeight, "new", att.BlockHeight)
	}

	a.available = append(a.available, att)

	return a.saveUnsafe()
}

// GetAvailable returns all the available attestations.
func (a *Attester) GetAvailable() []xchain.Attestation {
	a.mu.Lock()
	defer a.mu.Unlock()

	return slices.Clone(a.available)
}

// SetProposed sets the attestations as proposed.
func (a *Attester) SetProposed(headers []xchain.BlockHeader) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	proposed := headerMap(headers)

	var newAvailable, newProposed []xchain.Attestation
	for _, att := range a.availableAndProposedUnsafe() {
		// If proposed, move it to the proposed list.
		if proposed[att.BlockHeader] {
			newProposed = append(newProposed, att)
		} else { // Otherwise, keep or move it to in the available list.
			newAvailable = append(newAvailable, att)
		}
	}

	a.available = newAvailable
	a.proposed = newProposed

	return a.saveUnsafe()
}

// SetCommitted sets the attestations as committed. Persisting the result to disk,.
func (a *Attester) SetCommitted(headers []xchain.BlockHeader) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	committed := headerMap(headers)

	newCommitted := a.committed

	var newAvailable []xchain.Attestation
	for _, att := range a.availableAndProposedUnsafe() {
		// If newly committed, add to committed.
		if committed[att.BlockHeader] {
			newCommitted = append(newCommitted, att)
		} else { // Otherwise, keep/move it back to available.
			newAvailable = append(newAvailable, att)
		}
	}

	a.available = newAvailable
	a.proposed = nil
	a.committed = pruneLatestPerChain(newCommitted)

	return a.saveUnsafe()
}

func (a *Attester) LocalPubKey() [33]byte {
	a.mu.Lock()
	defer a.mu.Unlock()

	bz := a.privKey.PubKey().Bytes()

	return [33]byte(bz)
}

// availableAndProposed returns all the available and proposed attestations.
// It is unsafe since it assumes the lock is held.
func (a *Attester) availableAndProposedUnsafe() []xchain.Attestation {
	var resp []xchain.Attestation
	resp = append(resp, a.available...)
	resp = append(resp, a.proposed...)

	return resp
}

// LatestByChainUnsafe returns the latest attestation for the given chain.
// It is unsafe since it assumes the lock is held.
func (a *Attester) latestByChainUnsafe(chainID uint64) (xchain.Attestation, bool) {
	var resp xchain.Attestation
	var found bool

	iter := func(atts []xchain.Attestation) {
		for _, att := range atts {
			if att.SourceChainID != chainID || att.BlockHeight <= resp.BlockHeight {
				continue
			}
			resp = att
			found = true
		}
	}

	iter(a.available)
	iter(a.proposed)
	iter(a.committed)

	return resp, found
}

// saveUnsafe saves the state to disk. It is unsafe since it assumes the lock is held.
func (a *Attester) saveUnsafe() error {
	sortAtts := func(atts []xchain.Attestation) {
		sort.Slice(atts, func(i, j int) bool {
			if atts[i].SourceChainID != atts[j].SourceChainID {
				return atts[i].SourceChainID < atts[j].SourceChainID
			}

			return atts[i].BlockHeight < atts[j].BlockHeight
		})
	}
	sortAtts(a.available)
	sortAtts(a.proposed)
	sortAtts(a.committed)

	s := stateJSON{
		Available: a.available,
		Proposed:  a.proposed,
		Committed: a.committed,
	}
	bz, err := json.Marshal(s)
	if err != nil {
		return errors.Wrap(err, "marshal state path")
	}

	if err := tempfile.WriteFileAtomic(a.path, bz, 0o600); err != nil {
		return errors.Wrap(err, "write state path")
	}

	return nil
}

// stateJSON is the JSON representation of the attester state.
type stateJSON struct {
	Available []xchain.Attestation `json:"available"`
	Proposed  []xchain.Attestation `json:"proposed"`
	Committed []xchain.Attestation `json:"committed"`
}

// loadState loads a path state from the given path.
func loadState(path string) (stateJSON, error) {
	bz, err := os.ReadFile(path)
	if err != nil {
		return stateJSON{}, errors.Wrap(err, "read state path")
	}

	var s stateJSON
	if err := json.Unmarshal(bz, &s); err != nil {
		return stateJSON{}, errors.Wrap(err, "unmarshal state path")
	}

	return s, nil
}

// headerMap converts a list of headers to a bool map (set).
func headerMap(headers []xchain.BlockHeader) map[xchain.BlockHeader]bool {
	resp := make(map[xchain.BlockHeader]bool)
	for _, header := range headers {
		resp[header] = true
	}

	return resp
}

// pruneLatestPerChain returns only the latest attestation per chain.
func pruneLatestPerChain(atts []xchain.Attestation) []xchain.Attestation {
	latest := make(map[uint64]xchain.Attestation)
	for _, att := range atts {
		if latest[att.SourceChainID].BlockHeight >= att.BlockHeight {
			continue
		}
		latest[att.SourceChainID] = att
	}

	// Flatten
	resp := make([]xchain.Attestation, 0, len(latest))
	for _, att := range latest {
		resp = append(resp, att)
	}

	return resp
}
