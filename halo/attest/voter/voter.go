package voter

import (
	"context"
	"encoding/json"
	"os"
	"slices"
	"sort"
	"sync"
	"time"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/libs/tempfile"

	"github.com/ethereum/go-ethereum/common"

	"github.com/prometheus/client_golang/prometheus"
)

var _ types.Voter = (*Voter)(nil)

// Voter implements the types.Voter interface.
// It is responsible for creating and persisting votes.
// The goal is to ensure "all blocks are votes for".
type Voter struct {
	path    string
	privKey crypto.PrivKey
	chains  map[uint64]string
	address common.Address

	mu        sync.Mutex
	available []*types.Vote
	proposed  []*types.Vote
	committed []*types.Vote
}

// GenEmptyStateFile generates an empty attester state file at the given path.
// This must be called before LoadVoter.
func GenEmptyStateFile(path string) error {
	return (&Voter{path: path}).saveUnsafe()
}

// LoadVoter returns a new attester with state loaded from disk.
func LoadVoter(ctx context.Context, privKey crypto.PrivKey, path string, provider xchain.Provider,
	chains map[uint64]string,
) (*Voter, error) {
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

	a := Voter{
		privKey: privKey,
		address: addr,
		path:    path,
		chains:  chains,

		available: s.Available,
		proposed:  s.Proposed,
		committed: s.Committed,
	}

	// This shouldn't be required, but -race otherwise complains...
	a.mu.Lock()
	defer a.mu.Unlock()

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
func (a *Voter) Attest(_ context.Context, block xchain.Block) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	vote, err := CreateVote(a.privKey, block)
	if err != nil {
		return err
	}

	// Ensure attestation is sequential and not a duplicate.
	latest, ok := a.latestByChainUnsafe(vote.BlockHeader.ChainId)
	if ok && latest.BlockHeader.Height >= vote.BlockHeader.Height {
		return errors.New("attestation height already exists",
			"latest", latest.BlockHeader.Height, "new", vote.BlockHeader.Height)
	} else if ok && latest.BlockHeader.Height+1 != vote.BlockHeader.Height {
		return errors.New("attestation is not sequential",
			"existing", latest.BlockHeader.Height, "new", vote.BlockHeader.Height)
	}

	a.available = append(a.available, vote)

	lag := time.Since(block.Timestamp).Seconds()
	createLag.WithLabelValues(a.chains[vote.BlockHeader.ChainId]).Set(lag)
	createHeight.WithLabelValues(a.chains[vote.BlockHeader.ChainId]).Set(float64(vote.BlockHeader.Height))

	return a.saveUnsafe()
}

// GetAvailable returns all the available attestations.
func (a *Voter) GetAvailable() []*types.Vote {
	a.mu.Lock()
	defer a.mu.Unlock()

	return slices.Clone(a.available)
}

// SetProposed sets the attestations as proposed.
func (a *Voter) SetProposed(headers []*types.BlockHeader) error {
	if len(headers) == 0 {
		return nil
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	proposed := headerMap(headers)

	var newAvailable, newProposed []*types.Vote
	for _, vote := range a.availableAndProposedUnsafe() {
		// If proposed, move it to the proposed list.
		if proposed[vote.BlockHeader.ToXChain()] {
			newProposed = append(newProposed, vote)
		} else { // Otherwise, keep or move it to in the available list.
			newAvailable = append(newAvailable, vote)
		}
	}

	a.available = newAvailable
	a.proposed = newProposed

	return a.saveUnsafe()
}

// SetCommitted sets the attestations as committed. Persisting the result to disk.
func (a *Voter) SetCommitted(headers []*types.BlockHeader) error {
	if len(headers) == 0 {
		return nil
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	committed := headerMap(headers)

	newCommitted := a.committed

	var newAvailable []*types.Vote
	for _, vote := range a.availableAndProposedUnsafe() {
		// If newly committed, add to committed.
		if committed[vote.BlockHeader.ToXChain()] {
			newCommitted = append(newCommitted, vote)
		} else { // Otherwise, keep/move it back to available.
			newAvailable = append(newAvailable, vote)
		}
	}

	a.available = newAvailable
	a.proposed = nil
	a.committed = pruneLatestPerChain(newCommitted)

	// Update committed height metrics.
	for _, vote := range a.committed {
		commitHeight.WithLabelValues(a.chains[vote.BlockHeader.ChainId]).Set(float64(vote.BlockHeader.Height))
	}

	return a.saveUnsafe()
}

func (a *Voter) LocalAddress() common.Address {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.address
}

// availableAndProposed returns all the available and proposed attestations.
// It is unsafe since it assumes the lock is held.
func (a *Voter) availableAndProposedUnsafe() []*types.Vote {
	var resp []*types.Vote
	resp = append(resp, a.available...)
	resp = append(resp, a.proposed...)

	return resp
}

// LatestByChainUnsafe returns the latest attestation for the given chain.
// It is unsafe since it assumes the lock is held.
func (a *Voter) latestByChainUnsafe(chainID uint64) (*types.Vote, bool) {
	var found bool
	resp := &types.Vote{BlockHeader: &types.BlockHeader{}} // Zero value is safe to use.
	iter := func(votes []*types.Vote) {
		for _, vote := range votes {
			if vote.BlockHeader.ChainId != chainID || vote.BlockHeader.Height <= resp.BlockHeader.Height {
				continue
			}
			resp = vote
			found = true
		}
	}

	iter(a.available)
	iter(a.proposed)
	iter(a.committed)

	return resp, found
}

// saveUnsafe saves the state to disk. It is unsafe since it assumes the lock is held.
func (a *Voter) saveUnsafe() error {
	sortVotes := func(atts []*types.Vote) {
		sort.Slice(atts, func(i, j int) bool {
			if atts[i].BlockHeader.ChainId != atts[j].BlockHeader.ChainId {
				return atts[i].BlockHeader.ChainId < atts[j].BlockHeader.ChainId
			}

			return atts[i].BlockHeader.Height < atts[j].BlockHeader.Height
		})
	}
	sortVotes(a.available)
	sortVotes(a.proposed)
	sortVotes(a.committed)

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
func (a *Voter) instrumentUnsafe() {
	count := func(atts []*types.Vote, gaugeVec *prometheus.GaugeVec) {
		counts := make(map[uint64]int)
		for _, vote := range atts {
			counts[vote.BlockHeader.ChainId]++
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
	Available []*types.Vote `json:"available"`
	Proposed  []*types.Vote `json:"proposed"`
	Committed []*types.Vote `json:"committed"`
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
func headerMap(headers []*types.BlockHeader) map[xchain.BlockHeader]bool {
	resp := make(map[xchain.BlockHeader]bool) // Can't use protos as map keys.
	for _, header := range headers {
		resp[header.ToXChain()] = true
	}

	return resp
}

// pruneLatestPerChain returns only the latest attestation per chain.
func pruneLatestPerChain(atts []*types.Vote) []*types.Vote {
	latest := make(map[uint64]*types.Vote)
	for _, vote := range atts {
		latestAtt, ok := latest[vote.BlockHeader.ChainId]
		if ok && latestAtt.BlockHeader.Height >= vote.BlockHeader.Height {
			continue
		}

		latest[vote.BlockHeader.ChainId] = vote
	}

	// Flatten
	resp := make([]*types.Vote, 0, len(latest))
	for _, vote := range latest {
		resp = append(resp, vote)
	}

	return resp
}
