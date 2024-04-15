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
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/libs/tempfile"

	"github.com/ethereum/go-ethereum/common"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	prodBackoff  = time.Second
	maxAvailable = 1_000
)

var _ types.Voter = (*Voter)(nil)

// Voter implements the types.Voter interface.
// It is responsible for creating and persisting votes.
// The goal is to ensure "all blocks are votes for".
//
// Note Start must be called only once on startup.
// GetAvailable, SetProposed, and SetCommitted are thread safe, but must be called after Start.
type Voter struct {
	path        string
	privKey     crypto.PrivKey
	chains      map[uint64]string
	address     common.Address
	provider    xchain.Provider
	deps        types.VoterDeps
	backoffFunc func(context.Context) func()
	wg          sync.WaitGroup

	mu          sync.Mutex
	latest      map[uint64]*types.Vote // Latest vote per chain
	available   []*types.Vote
	proposed    []*types.Vote
	committed   []*types.Vote
	minsByChain map[uint64]uint64
	isVal       bool
}

// GenEmptyStateFile generates an empty attester state file at the given path.
// This must be called before LoadVoter.
func GenEmptyStateFile(path string) error {
	return (&Voter{path: path}).saveUnsafe()
}

// LoadVoter returns a new attester with state loaded from disk.
func LoadVoter(privKey crypto.PrivKey, path string, provider xchain.Provider, deps types.VoterDeps,
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

	return &Voter{
		privKey:  privKey,
		address:  addr,
		path:     path,
		chains:   chains,
		provider: provider,
		deps:     deps,
		backoffFunc: func(ctx context.Context) func() {
			return expbackoff.New(ctx, expbackoff.WithPeriodicConfig(prodBackoff))
		},

		available: s.Available,
		proposed:  s.Proposed,
		committed: s.Committed,
		latest:    s.Latest,
	}, nil
}

// Start starts runners that attest to each source chain. It does not block, it returns immediately.
func (v *Voter) Start(ctx context.Context) {
	for chainID := range v.chains {
		go v.runForever(ctx, chainID)
	}
}

// WaitDone waits for all runners to exit. Note the original Start context must be canceled to exit.
func (v *Voter) WaitDone() {
	v.wg.Wait()
}

// minWindow returns the minimum.
func (v *Voter) minWindow(chainID uint64) (uint64, bool) {
	v.mu.Lock()
	defer v.mu.Unlock()

	resp, ok := v.minsByChain[chainID]

	return resp, ok
}

// minWindow returns the minimum.
func (v *Voter) isValidator() bool {
	v.mu.Lock()
	defer v.mu.Unlock()

	return v.isVal
}

// runForever blocks, repeatedly calling runOnce (with backoff) for the provided chain until the context is canceled.
func (v *Voter) runForever(ctx context.Context, chainID uint64) {
	v.wg.Add(1)
	defer v.wg.Done()

	ctx = log.WithCtx(ctx, "chain", v.chains[chainID])

	backoff := v.backoffFunc(ctx)
	for ctx.Err() == nil {
		if !v.isValidator() {
			backoff()
			continue
		}

		err := v.runOnce(ctx, chainID)
		if ctx.Err() != nil {
			return // Don't log or sleep on context cancel.
		}

		log.Warn(ctx, "Vote runner failed (will retry)", err)
		backoff()
	}
}

// runOnce blocks, streaming xblocks from the provided chain until an error is encountered.
// It always returns a non-nil error.
func (v *Voter) runOnce(ctx context.Context, chainID uint64) error {
	// Determine what height to stream from.
	var fromHeight uint64

	// Get latest state from disk.
	if latest, ok := v.latestByChain(chainID); ok {
		fromHeight = latest.BlockHeader.Height + 1
	}

	// Get latest approved attestation from the chain.
	if latest, ok, err := v.deps.LatestAttestationHeight(ctx, chainID); err != nil {
		return errors.Wrap(err, "latest attestation")
	} else if ok && fromHeight < latest+1 {
		// Allows skipping ahead of we were behind for some reason.
		fromHeight = latest + 1
	}

	first := true // Allow skipping on first attestation.

	return v.provider.StreamBlocks(ctx, chainID, fromHeight,
		func(ctx context.Context, block xchain.Block) error {
			if !v.isValidator() {
				return errors.New("not a validator anymore")
			}

			minimum, ok := v.minWindow(block.SourceChainID)
			if ok && block.BlockHeight < minimum {
				return errors.New("behind vote window (too slow)", "vote_height", block.BlockHeight, "window_minimum", minimum)
			}

			backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second*5))
			for v.AvailableCount() > maxAvailable {
				log.Warn(ctx, "Voting paused, latest approved attestation is too far behind (stuck?)", nil, "vote_height", block.BlockHeight)
				backoff()
			}

			if err := v.Vote(block, first); err != nil {
				return errors.Wrap(err, "vote")
			}
			first = false

			return nil
		},
	)
}

// Vote creates a vote for the given block and adds it to the internal state.
func (v *Voter) Vote(block xchain.Block, allowSkip bool) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	vote, err := CreateVote(v.privKey, block)
	if err != nil {
		return err
	}

	// Ensure attestation is sequential and not a duplicate.
	latest, ok := v.latest[vote.BlockHeader.ChainId]
	if ok && latest.BlockHeader.Height >= vote.BlockHeader.Height {
		return errors.New("attestation height already exists",
			"latest", latest.BlockHeader.Height, "new", vote.BlockHeader.Height)
	} else if ok && !allowSkip && latest.BlockHeader.Height+1 != vote.BlockHeader.Height {
		return errors.New("attestation is not sequential",
			"existing", latest.BlockHeader.Height, "new", vote.BlockHeader.Height)
	}

	v.latest[vote.BlockHeader.ChainId] = vote
	v.available = append(v.available, vote)

	lag := time.Since(block.Timestamp).Seconds()
	createLag.WithLabelValues(v.chains[vote.BlockHeader.ChainId]).Set(lag)
	createHeight.WithLabelValues(v.chains[vote.BlockHeader.ChainId]).Set(float64(vote.BlockHeader.Height))

	return v.saveUnsafe()
}

// UpdateValidators caches whether this voter is a validator in the provided set.
func (v *Voter) UpdateValidators(valset []abci.ValidatorUpdate) {
	v.mu.Lock()
	defer v.mu.Unlock()

	for _, val := range valset {
		addr, err := k1util.PubKeyPBToAddress(val.PubKey)
		if err != nil {
			continue
		}
		if v.address != addr {
			continue
		}

		v.isVal = val.Power > 0

		return
	}
	// If validator not updated, then isVal didn't change.
}

// TrimBehind trims all available and proposed votes that are behind the vote window thresholds (map[chainID]height)
// and returns the number that was deleted.
func (v *Voter) TrimBehind(minsByChain map[uint64]uint64) int {
	v.mu.Lock()
	defer v.mu.Unlock()

	trim := func(votes []*types.Vote) []*types.Vote {
		var remaining []*types.Vote
		for _, vote := range votes {
			minimum, ok := minsByChain[vote.BlockHeader.ChainId]
			if ok && vote.BlockHeader.Height < minimum {
				trimTotal.WithLabelValues(v.chains[vote.BlockHeader.ChainId]).Inc()
				continue // Skip/Trim
			}
			remaining = append(remaining, vote) // Retain all others
		}

		return remaining
	}

	remainingAvailable := trim(v.available)
	remainingProposed := trim(v.proposed)

	trimmed := (len(v.available) - len(remainingAvailable)) + (len(v.proposed) - len(remainingProposed))

	v.available = remainingAvailable
	v.proposed = remainingProposed

	v.minsByChain = minsByChain

	return trimmed
}

// AvailableCount returns the number of available votes.
func (v *Voter) AvailableCount() int {
	v.mu.Lock()
	defer v.mu.Unlock()

	return len(v.available)
}

// GetAvailable returns a copy of all the available votes.
func (v *Voter) GetAvailable() []*types.Vote {
	v.mu.Lock()
	defer v.mu.Unlock()

	return slices.Clone(v.available)
}

// SetProposed sets the votes as proposed.
func (v *Voter) SetProposed(headers []*types.BlockHeader) error {
	proposedPerBlock.Observe(float64(len(headers)))

	if len(headers) == 0 {
		return nil
	}

	v.mu.Lock()
	defer v.mu.Unlock()

	proposed := headerMap(headers)

	var newAvailable, newProposed []*types.Vote
	for _, vote := range v.availableAndProposedUnsafe() {
		// If proposed, move it to the proposed list.
		if proposed[vote.BlockHeader.ToXChain()] {
			newProposed = append(newProposed, vote)
		} else { // Otherwise, keep or move it to in the available list.
			newAvailable = append(newAvailable, vote)
		}
	}

	v.available = newAvailable
	v.proposed = newProposed

	return v.saveUnsafe()
}

// SetCommitted sets the votes as committed. Persisting the result to disk.
func (v *Voter) SetCommitted(headers []*types.BlockHeader) error {
	committedPerBlock.Observe(float64(len(headers)))

	if len(headers) == 0 {
		return nil
	}

	v.mu.Lock()
	defer v.mu.Unlock()

	committed := headerMap(headers)

	newCommitted := v.committed

	var newAvailable []*types.Vote
	for _, vote := range v.availableAndProposedUnsafe() {
		// If newly committed, add to committed.
		if committed[vote.BlockHeader.ToXChain()] {
			newCommitted = append(newCommitted, vote)
		} else { // Otherwise, keep/move it back to available.
			newAvailable = append(newAvailable, vote)
		}
	}

	v.available = newAvailable
	v.proposed = nil
	v.committed = pruneLatestPerChain(newCommitted)

	// Update committed height metrics.
	for _, vote := range v.committed {
		commitHeight.WithLabelValues(v.chains[vote.BlockHeader.ChainId]).Set(float64(vote.BlockHeader.Height))
	}

	return v.saveUnsafe()
}

func (v *Voter) LocalAddress() common.Address {
	v.mu.Lock()
	defer v.mu.Unlock()

	return v.address
}

// availableAndProposed returns all the available and proposed votes.
// It is unsafe since it assumes the lock is held.
func (v *Voter) availableAndProposedUnsafe() []*types.Vote {
	var resp []*types.Vote
	resp = append(resp, v.available...)
	resp = append(resp, v.proposed...)

	return resp
}

func (v *Voter) latestByChain(chainID uint64) (*types.Vote, bool) {
	v.mu.Lock()
	defer v.mu.Unlock()

	vote, ok := v.latest[chainID]

	return vote, ok
}

// saveUnsafe saves the state to disk. It is unsafe since it assumes the lock is held.
func (v *Voter) saveUnsafe() error {
	sortVotes := func(atts []*types.Vote) {
		sort.Slice(atts, func(i, j int) bool {
			if atts[i].BlockHeader.ChainId != atts[j].BlockHeader.ChainId {
				return atts[i].BlockHeader.ChainId < atts[j].BlockHeader.ChainId
			}

			return atts[i].BlockHeader.Height < atts[j].BlockHeader.Height
		})
	}
	sortVotes(v.available)
	sortVotes(v.proposed)
	sortVotes(v.committed)

	s := stateJSON{
		Available: v.available,
		Proposed:  v.proposed,
		Committed: v.committed,
		Latest:    v.latest,
	}
	bz, err := json.Marshal(s)
	if err != nil {
		return errors.Wrap(err, "marshal state path")
	}

	if err := tempfile.WriteFileAtomic(v.path, bz, 0o600); err != nil {
		return errors.Wrap(err, "write state path")
	}

	v.instrumentUnsafe()

	return nil
}

// instrumentUnsafe updates metrics. It is unsafe since it assumes the lock is held.
func (v *Voter) instrumentUnsafe() {
	count := func(atts []*types.Vote, gaugeVec *prometheus.GaugeVec) {
		counts := make(map[uint64]int)
		for _, vote := range atts {
			counts[vote.BlockHeader.ChainId]++
		}

		for chain, count := range counts {
			gaugeVec.WithLabelValues(v.chains[chain]).Set(float64(count))
		}
	}

	count(v.available, availableCount)
	count(v.proposed, proposedCount)
}

// stateJSON is the JSON representation of the attester state.
type stateJSON struct {
	Available []*types.Vote          `json:"available"`
	Proposed  []*types.Vote          `json:"proposed"`
	Committed []*types.Vote          `json:"committed"`
	Latest    map[uint64]*types.Vote `json:"latest"`
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
	if s.Latest == nil {
		s.Latest = bootstrapLatest(s)
	}

	return s, nil
}

// bootstrapLatest returns the latest attestations for the given state.
// It can be used to bootstrap the new field if not present on disk.
func bootstrapLatest(s stateJSON) map[uint64]*types.Vote {
	resp := make(map[uint64]*types.Vote)
	iter := func(votes []*types.Vote) {
		for _, vote := range votes {
			existing, ok := resp[vote.BlockHeader.ChainId]
			if ok && vote.BlockHeader.Height <= existing.BlockHeader.Height {
				continue
			}
			resp[vote.BlockHeader.ChainId] = vote
		}
	}

	iter(s.Available)
	iter(s.Proposed)
	iter(s.Committed)

	return resp
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
