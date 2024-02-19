package attester

import (
	"context"
	"encoding/json"
	"os"
	"slices"
	"sort"
	"sync"
	"time"

	"github.com/omni-network/omni/halo2/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/libs/tempfile"

	"github.com/ethereum/go-ethereum/common"

	"github.com/prometheus/client_golang/prometheus"
)

var _ types.Attester = (*Attester)(nil)

// Attester implements the types.Attester interface.
// It is responsible for creating and persisting attestations.
// The goal is to ensure "all blocks are attested to".
type Attester struct {
	path    string
	privKey crypto.PrivKey
	chains  map[uint64]string
	address common.Address

	mu        sync.Mutex
	available []*types.Attestation
	proposed  []*types.Attestation
	committed []*types.Attestation
}

// GenEmptyStateFile generates an empty attester state file at the given path.
// This must be called before LoadAttester.
func GenEmptyStateFile(path string) error {
	return (&Attester{path: path}).saveUnsafe()
}

// LoadAttester returns a new attester with state loaded from disk.
func LoadAttester(ctx context.Context, privKey crypto.PrivKey, path string, provider xchain.Provider,
	chains map[uint64]string,
) (*Attester, error) {
	if len(privKey.PubKey().Bytes()) != 33 {
		return nil, errors.New("invalid private key")
	}

	s, err := loadState(path)
	if err != nil {
		return nil, err
	}

	addr, err := k1util.PubKeyToAddress(privKey.PubKey())
	if err != nil {
		return nil, err
	}

	a := Attester{
		privKey: privKey,
		address: addr,
		path:    path,
		chains:  chains,

		available: s.Available,
		proposed:  s.Proposed,
		committed: s.Committed,
	}

	// Subscribe to latest block provider
	for chainID := range chains {
		var fromHeight uint64
		if latest, ok := a.latestByChainUnsafe(chainID); ok {
			fromHeight = latest.BlockHeader.Height + 1
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
	latest, ok := a.latestByChainUnsafe(att.BlockHeader.ChainId)
	if ok && latest.BlockHeader.Height >= att.BlockHeader.Height {
		return errors.New("attestation height already exists",
			"latest", latest.BlockHeader.Height, "new", att.BlockHeader.Height)
	} else if ok && latest.BlockHeader.Height+1 != att.BlockHeader.Height {
		return errors.New("attestation is not sequential",
			"existing", latest.BlockHeader.Height, "new", att.BlockHeader.Height)
	}

	a.available = append(a.available, att)

	lag := time.Since(block.Timestamp).Seconds()
	createLag.WithLabelValues(a.chains[att.BlockHeader.ChainId]).Set(lag)
	createHeight.WithLabelValues(a.chains[att.BlockHeader.ChainId]).Set(float64(att.BlockHeader.Height))

	return a.saveUnsafe()
}

// GetAvailable returns all the available attestations.
func (a *Attester) GetAvailable() []*types.Attestation {
	a.mu.Lock()
	defer a.mu.Unlock()

	return slices.Clone(a.available)
}

// SetProposed sets the attestations as proposed.
func (a *Attester) SetProposed(headers []*types.BlockHeader) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	proposed := headerMap(headers)

	var newAvailable, newProposed []*types.Attestation
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

// SetCommitted sets the attestations as committed. Persisting the result to disk.
func (a *Attester) SetCommitted(headers []*types.BlockHeader) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	committed := headerMap(headers)

	newCommitted := a.committed

	var newAvailable []*types.Attestation
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

	// Update committed height metrics.
	for _, att := range a.committed {
		commitHeight.WithLabelValues(a.chains[att.BlockHeader.ChainId]).Set(float64(att.BlockHeader.Height))
	}

	return a.saveUnsafe()
}

func (a *Attester) LocalAddress() common.Address {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.address
}

// availableAndProposed returns all the available and proposed attestations.
// It is unsafe since it assumes the lock is held.
func (a *Attester) availableAndProposedUnsafe() []*types.Attestation {
	var resp []*types.Attestation
	resp = append(resp, a.available...)
	resp = append(resp, a.proposed...)

	return resp
}

// LatestByChainUnsafe returns the latest attestation for the given chain.
// It is unsafe since it assumes the lock is held.
func (a *Attester) latestByChainUnsafe(chainID uint64) (*types.Attestation, bool) {
	var found bool
	resp := &types.Attestation{BlockHeader: &types.BlockHeader{}} // Zero value is safe to use.
	iter := func(atts []*types.Attestation) {
		for _, att := range atts {
			if att.BlockHeader.ChainId != chainID || att.BlockHeader.Height <= resp.BlockHeader.Height {
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
	sortAtts := func(atts []*types.Attestation) {
		sort.Slice(atts, func(i, j int) bool {
			if atts[i].BlockHeader.ChainId != atts[j].BlockHeader.ChainId {
				return atts[i].BlockHeader.ChainId < atts[j].BlockHeader.ChainId
			}

			return atts[i].BlockHeader.Height < atts[j].BlockHeader.Height
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

	a.instrumentUnsafe()

	return nil
}

// instrumentUnsafe updates metrics. It is unsafe since it assumes the lock is held.
func (a *Attester) instrumentUnsafe() {
	count := func(atts []*types.Attestation, gaugeVec *prometheus.GaugeVec) {
		counts := make(map[uint64]int)
		for _, att := range atts {
			counts[att.BlockHeader.ChainId]++
		}

		for chain, count := range counts {
			gaugeVec.WithLabelValues(a.chains[chain]).Set(float64(count))
		}
	}

	count(a.available, availableCount)
	count(a.proposed, proposedCount)
}

// stateJSON is the JSON representation of the attester state.
type stateJSON struct {
	Available []*types.Attestation `json:"available"`
	Proposed  []*types.Attestation `json:"proposed"`
	Committed []*types.Attestation `json:"committed"`
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
func headerMap(headers []*types.BlockHeader) map[*types.BlockHeader]bool {
	resp := make(map[*types.BlockHeader]bool)
	for _, header := range headers {
		resp[header] = true
	}

	return resp
}

// pruneLatestPerChain returns only the latest attestation per chain.
func pruneLatestPerChain(atts []*types.Attestation) []*types.Attestation {
	latest := make(map[uint64]*types.Attestation)
	for _, att := range atts {
		if latest[att.BlockHeader.ChainId].BlockHeader.Height >= att.BlockHeader.Height {
			continue
		}
		latest[att.BlockHeader.ChainId] = att
	}

	// Flatten
	resp := make([]*types.Attestation, 0, len(latest))
	for _, att := range latest {
		resp = append(resp, att)
	}

	return resp
}
