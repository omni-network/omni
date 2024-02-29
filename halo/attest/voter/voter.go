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

	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/libs/tempfile"

	"github.com/ethereum/go-ethereum/common"

	"github.com/prometheus/client_golang/prometheus"
)

const prodBackoff = time.Second

var _ types.Voter = (*Voter)(nil)

// Voter implements the types.Voter interface.
// It is responsible for creating and persisting votes.
// The goal is to ensure "all blocks are votes for".
//
// Note Start must be called only once on startup.
// GetAvailable, SetProposed, and SetCommitted are thread safe, but must be called after Start.
type Voter struct {
	path          string
	privKey       crypto.PrivKey
	chains        map[uint64]string
	address       common.Address
	provider      xchain.Provider
	deps          types.VoterDeps
	backoffPeriod time.Duration
	wg            sync.WaitGroup

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
		privKey:       privKey,
		address:       addr,
		path:          path,
		chains:        chains,
		provider:      provider,
		deps:          deps,
		backoffPeriod: prodBackoff,

		available: s.Available,
		proposed:  s.Proposed,
		committed: s.Committed,
	}, nil
}

// Start starts runners that attest to each source chain. It does not block, it returns immediately.
func (a *Voter) Start(ctx context.Context) {
	for chainID := range a.chains {
		go a.runForever(ctx, chainID)
	}
}

// WaitDone waits for all runners to exit. Note the original Start context must be canceled to exit.
func (a *Voter) WaitDone() {
	a.wg.Wait()
}

// runForever blocks, repeatedly calling runOnce (with backoff) for the provided chain until the context is canceled.
func (a *Voter) runForever(ctx context.Context, chainID uint64) {
	a.wg.Add(1)
	defer a.wg.Done()

	ctx = log.WithCtx(ctx, "chain", a.chains[chainID])

	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(a.backoffPeriod))
	for ctx.Err() == nil {
		if !a.deps.IsValidator(ctx, a.address) {
			backoff()
			continue
		}

		err := a.runOnce(ctx, chainID)
		if ctx.Err() != nil {
			return // Don't log or sleep on context cancel.
		}

		log.Warn(ctx, "Vote runner failed (will retry)", err)
		backoff()
	}
}

// runOnce blocks, streaming xblocks from the provided chain until an error is encountered.
// It always returns a non-nil error.
func (a *Voter) runOnce(ctx context.Context, chainID uint64) error {
	// Determine what height to stream from.
	var fromHeight uint64

	// Get latest state from disk.
	if latest, ok := a.latestByChain(chainID); ok {
		fromHeight = latest.BlockHeader.Height + 1
	}

	// Get latest approved attestation from the chain.
	if latest, ok, err := a.deps.LatestAttestationHeight(ctx, chainID); err != nil {
		return errors.Wrap(err, "latest attestation")
	} else if ok && fromHeight < latest+1 {
		// Allows skipping ahead of we were behind for some reason.
		fromHeight = latest + 1
	}

	// Channel shenanigans to wait for async subscription.
	errChan := make(chan error)
	ctx, cancel := context.WithCancel(ctx)
	handleErr := func(err error) error {
		select {
		case errChan <- err:
		case <-ctx.Done(): // Just in case of race.
		}
		cancel()

		return errors.Wrap(err, "voter runner")
	}

	// Start async subscription.
	first := true // Allow skipping on first attestation.
	err := a.provider.Subscribe(ctx, chainID, fromHeight,
		func(ctx context.Context, block xchain.Block) error {
			if !a.deps.IsValidator(ctx, a.address) {
				return handleErr(errors.New("not a validator anymore"))
			}

			cmp, err := a.deps.WindowCompare(ctx, block.SourceChainID, block.BlockHeight)
			if err != nil {
				return handleErr(errors.Wrap(err, "window compare"))
			} else if cmp < 0 {
				return handleErr(errors.New("behind vote window (too slow)"))
			} // Being ahead is not a problem, since we buffer on disk.

			if err := a.Vote(block, first); err != nil {
				return handleErr(errors.Wrap(err, "attest"))
			}
			first = false

			return nil
		},
	)
	if err != nil {
		return err
	}

	// Wait for error or context cancel.
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Vote creates a vote for the given block and adds it to the internal state.
func (a *Voter) Vote(block xchain.Block, allowSkip bool) error {
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
	} else if ok && !allowSkip && latest.BlockHeader.Height+1 != vote.BlockHeader.Height {
		return errors.New("attestation is not sequential",
			"existing", latest.BlockHeader.Height, "new", vote.BlockHeader.Height)
	}

	a.available = append(a.available, vote)

	lag := time.Since(block.Timestamp).Seconds()
	createLag.WithLabelValues(a.chains[vote.BlockHeader.ChainId]).Set(lag)
	createHeight.WithLabelValues(a.chains[vote.BlockHeader.ChainId]).Set(float64(vote.BlockHeader.Height))

	return a.saveUnsafe()
}

// TrimBehind trims all votes that are behind the vote window thresholds (map[chainID]height) and returns the number that was deleted.
func (a *Voter) TrimBehind(thresholds map[uint64]uint64) int {
	a.mu.Lock()
	defer a.mu.Unlock()

	var stillAvailable []*types.Vote
	for _, vote := range a.available {
		threshold, ok := thresholds[vote.BlockHeader.ChainId]
		if ok && vote.BlockHeader.Height < threshold {
			trimTotal.WithLabelValues(a.chains[vote.BlockHeader.ChainId]).Inc()
			continue // Skip/Trim
		}
		stillAvailable = append(stillAvailable, vote) // Retain all others
	}

	trimmed := len(a.available) - len(stillAvailable)

	a.available = stillAvailable

	return trimmed
}

// GetAvailable returns a copy of all the available votes.
func (a *Voter) GetAvailable() []*types.Vote {
	a.mu.Lock()
	defer a.mu.Unlock()

	return slices.Clone(a.available)
}

// SetProposed sets the votes as proposed.
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

// SetCommitted sets the votes as committed. Persisting the result to disk.
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

// availableAndProposed returns all the available and proposed votes.
// It is unsafe since it assumes the lock is held.
func (a *Voter) availableAndProposedUnsafe() []*types.Vote {
	var resp []*types.Vote
	resp = append(resp, a.available...)
	resp = append(resp, a.proposed...)

	return resp
}

func (a *Voter) latestByChain(chainID uint64) (*types.Vote, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.latestByChainUnsafe(chainID)
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
