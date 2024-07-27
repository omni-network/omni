package voter

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
	"sync"
	"time"

	"github.com/omni-network/omni/halo/attest/types"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/cometbft/cometbft/libs/tempfile"

	"github.com/ethereum/go-ethereum/common"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	prodBackoff         = time.Second
	maxAvailable        = 1_000
	initialAttestOffset = uint64(1)
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
	cChainID    uint64
	privKey     crypto.PrivKey
	network     netconf.Network
	address     common.Address
	provider    xchain.Provider
	deps        types.VoterDeps
	backoffFunc func(context.Context) func()
	wg          sync.WaitGroup

	mu          sync.Mutex
	latest      map[xchain.ChainVersion]*types.Vote // Latest vote per chain
	available   []*types.Vote
	proposed    []*types.Vote
	committed   []*types.Vote
	minsByChain map[xchain.ChainVersion]uint64 // map[chainID]offset
	isVal       bool
	valSetID    uint64
}

// GenEmptyStateFile generates an empty attester state file at the given path.
// This must be called before LoadVoter.
func GenEmptyStateFile(path string) error {
	return (&Voter{path: path}).saveUnsafe()
}

// LoadVoter returns a new attester with state loaded from disk.
func LoadVoter(privKey crypto.PrivKey, path string, provider xchain.Provider, deps types.VoterDeps,
	network netconf.Network,
) (*Voter, error) {
	if len(privKey.PubKey().Bytes()) != k1.PubKeySize {
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
		cChainID: network.ID.Static().OmniConsensusChainIDUint64(),
		address:  addr,
		path:     path,
		network:  network,
		provider: provider,
		deps:     deps,
		backoffFunc: func(ctx context.Context) func() {
			return expbackoff.New(ctx, expbackoff.WithPeriodicConfig(prodBackoff))
		},

		available: s.Available,
		proposed:  s.Proposed,
		committed: s.Committed,
		latest:    latestFromJSON(s.Latest),
	}, nil
}

// Start starts runners that attest to each source chain. It does not block, it returns immediately.
func (v *Voter) Start(ctx context.Context) {
	for _, chain := range v.network.Chains {
		for _, chainVer := range chain.ChainVersions() {
			go v.runForever(ctx, chainVer)
		}
	}
}

// WaitDone waits for all runners to exit. Note the original Start context must be canceled to exit.
func (v *Voter) WaitDone() {
	v.wg.Wait()
}

// minWindow returns the minimum vote window (attestation offset).
func (v *Voter) minWindow(chainVer xchain.ChainVersion) (uint64, bool) {
	v.mu.Lock()
	defer v.mu.Unlock()

	resp, ok := v.minsByChain[chainVer]

	return resp, ok
}

// isValidator returns true if the local address is a validator in the active set.
func (v *Voter) isValidator() bool {
	v.mu.Lock()
	defer v.mu.Unlock()

	return v.isVal
}

// runForever blocks, repeatedly calling runOnce (with backoff) for the provided chain until the context is canceled.
func (v *Voter) runForever(ctx context.Context, chainVer xchain.ChainVersion) {
	v.wg.Add(1)
	defer v.wg.Done()

	backoff := v.backoffFunc(ctx)
	for ctx.Err() == nil {
		if !v.isValidator() {
			backoff()
			continue
		}

		err := v.runOnce(ctx, chainVer)
		if ctx.Err() != nil {
			return // Don't log or sleep on context cancel.
		}

		log.Warn(ctx, "Vote runner failed (will retry)", err, "chain", v.network.ChainVersionName(chainVer))
		backoff()
	}
}

// runOnce blocks, streaming xblocks from the provided chain until an error is encountered.
// It always returns a non-nil error.
func (v *Voter) runOnce(ctx context.Context, chainVer xchain.ChainVersion) error {
	chain, ok := v.network.Chain(chainVer.ID)
	if !ok {
		return errors.New("unknown chain ID")
	}

	maybeDebugLog := newDebugLogFilterer(time.Minute) // Log empty blocks once per minute.
	first := true                                     // Allow skipping on first attestation.

	// Use actual chain version to calculate offset to start voting from (to prevent double signing).
	_, skipBeforeOffset, err := v.getFromHeightAndOffset(ctx, chainVer)
	if err != nil {
		return errors.Wrap(err, "get from height and offset")
	}

	// Use finalized chain version to calculate height and offset to start streaming from (for correct offset calcs).
	finalVer := xchain.ChainVersion{ID: chainVer.ID, ConfLevel: xchain.ConfFinalized}
	fromBlockHeight, fromAttestOffset, err := v.getFromHeightAndOffset(ctx, finalVer)
	if err != nil {
		return errors.Wrap(err, "get from height and offset")
	}

	log.Info(ctx, "Voting started for chain",
		"from_block_height", fromBlockHeight,
		"from_attest_offset", fromAttestOffset,
		"skip_before_offset", skipBeforeOffset,
		"chain", v.network.ChainVersionName(chainVer),
	)

	req := xchain.ProviderRequest{
		ChainID:   chainVer.ID,
		Height:    fromBlockHeight,
		ConfLevel: chainVer.ConfLevel,
	}

	tracker := newOffsetTracker(fromAttestOffset)
	streamOffsets := make(map[xchain.StreamID]uint64)
	var prevBlock *xchain.Block

	return v.provider.StreamBlocks(ctx, req,
		func(ctx context.Context, block xchain.Block) error {
			if !v.isValidator() {
				return errors.New("not a validator anymore")
			}

			if err := detectReorg(chainVer, prevBlock, block, streamOffsets); err != nil {
				reorgTotal.WithLabelValues(v.network.ChainVersionName(chainVer)).Inc()
				// Restart stream, recalculating block offset from finalized version.

				return err
			}
			prevBlock = &block

			if !block.ShouldAttest(chain.AttestInterval) {
				maybeDebugLog(ctx, "Not creating vote for empty cross chain block")

				return nil // Do not vote for empty blocks.
			}

			attestOffset, err := tracker.NextAttestOffset(block.BlockHeight)
			if err != nil {
				return errors.Wrap(err, "next attestation offset")
			}

			// Create a vote for the block.
			attHeader := xchain.AttestHeader{
				ConsensusChainID: v.cChainID,
				ChainVersion:     chainVer,
				AttestOffset:     attestOffset,
			}

			if attestOffset < skipBeforeOffset {
				maybeDebugLog(ctx, "Skipping previously voted block on startup", "attest_offset", attestOffset, "skip_before_offset", skipBeforeOffset)

				return nil // Do not vote for offsets already approved or that we voted for previously (this risks double signing).
			}

			minimum, ok := v.minWindow(chainVer)
			if ok && attestOffset < minimum {
				// Restart stream, jumping ahead to middle of vote window.
				return errors.New("behind vote window (too slow)", "attest_offset", attestOffset, "window_minimum", minimum)
			}

			backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second*5))
			for v.AvailableCount() > maxAvailable {
				log.Warn(ctx, "Voting paused, latest approved attestation is too far behind (stuck?)", nil, "attest_offset", attestOffset, "block_height", block.BlockHeight)
				backoff()
			}

			if err := v.Vote(attHeader, block, first); err != nil {
				return errors.Wrap(err, "vote")
			}
			first = false

			// TODO(corver): Remove if this becomes too noisy.
			logVoteCreated(ctx, v.network, attHeader, block)

			return nil
		},
	)
}

// getFromHeightAndOffset returns the height and offset to start streaming from for the given chain version.
//
//nolint:nonamedreturns // Ambiguous return values.
func (v *Voter) getFromHeightAndOffset(ctx context.Context, chainVer xchain.ChainVersion) (fromBlockHeight uint64, fromAttestOffset uint64, err error) {
	fromAttestOffset = initialAttestOffset // Default to initial offset.
	// Note that initialisation of fromBlockHeight is handled in xprovider.

	// Get latest state from disk.
	if latest, ok := v.latestByChain(chainVer); ok {
		fromBlockHeight = latest.BlockHeader.BlockHeight + 1
		fromAttestOffset = latest.AttestHeader.AttestOffset + 1
	}

	// Get latest approved attestation from the chain.
	if latest, ok, err := v.deps.LatestAttestation(ctx, chainVer); err != nil {
		return 0, 0, errors.Wrap(err, "latest attestation")
	} else if ok && fromBlockHeight < latest.BlockHeight+1 {
		// Allows skipping ahead of we were behind for some reason.
		fromBlockHeight = latest.BlockHeight + 1
		fromAttestOffset = latest.AttestHeader.AttestOffset + 1
	}

	return fromBlockHeight, fromAttestOffset, nil
}

// Vote creates a vote for the given block and adds it to the internal state.
func (v *Voter) Vote(attHeader xchain.AttestHeader, block xchain.Block, allowSkip bool) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	vote, err := CreateVote(v.privKey, attHeader, block)
	if err != nil {
		return err
	} else if err := vote.Verify(); err != nil {
		return errors.Wrap(err, "verify vote")
	}

	chainVer := vote.AttestHeader.XChainVersion()

	// Ensure attestation is sequential and not a duplicate.
	latest, ok := v.latest[chainVer]
	if ok && latest.AttestHeader.AttestOffset >= vote.AttestHeader.AttestOffset {
		return errors.New("attestation height already exists",
			"latest", latest.AttestHeader.AttestOffset, "new", vote.AttestHeader.AttestOffset)
	} else if ok && !allowSkip && latest.AttestHeader.AttestOffset+1 != vote.AttestHeader.AttestOffset {
		return errors.New("attestation is not sequential",
			"existing", latest.AttestHeader.AttestOffset, "new", vote.AttestHeader.AttestOffset)
	}

	v.latest[chainVer] = vote
	v.available = append(v.available, vote)

	lag := time.Since(block.Timestamp).Seconds()
	name := v.network.ChainVersionName(chainVer)
	createLag.WithLabelValues(name).Set(lag)
	createHeight.WithLabelValues(name).Set(float64(vote.BlockHeader.BlockHeight))
	createAttestOffset.WithLabelValues(name).Set(float64(vote.AttestHeader.AttestOffset))
	for stream, msgOffset := range latestMsgOffsets(block.Msgs) {
		createMsgOffset.WithLabelValues(v.network.StreamName(stream)).Set(float64(msgOffset))
	}

	return v.saveUnsafe()
}

// UpdateValidatorSet caches whether this voter is a validator in the provided set.
func (v *Voter) UpdateValidatorSet(valset *vtypes.ValidatorSetResponse) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	isVal, err := valset.IsValidator(v.address)
	if err != nil {
		return err
	}

	if v.valSetID > valset.Id {
		return errors.New("unexpected validator set ID [BUG]", "old", v.valSetID, "new", valset.Id)
	} else if v.valSetID == valset.Id && v.isVal != isVal {
		return errors.New("unexpected isVal [BUG]", "old", v.valSetID, "new", valset.Id)
	}

	v.isVal = isVal
	v.valSetID = valset.Id

	// Note this results in outgoing validators not voting in their last block.
	// This is ok, the added complexity of tracking valset.ActivatedHeight vs currentHeight isn't worth it.

	return nil
}

// TrimBehind trims all available and proposed votes that are behind the vote window thresholds (map[chainID]offset)
// and returns the number that was deleted.
func (v *Voter) TrimBehind(minsByChain map[xchain.ChainVersion]uint64) int {
	v.mu.Lock()
	defer v.mu.Unlock()

	trim := func(votes []*types.Vote) []*types.Vote {
		var remaining []*types.Vote
		for _, vote := range votes {
			chainVer := vote.AttestHeader.XChainVersion()
			minimum, ok := minsByChain[chainVer]
			if ok && vote.AttestHeader.AttestOffset < minimum {
				trimTotal.WithLabelValues(v.network.ChainVersionName(chainVer)).Inc()
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
func (v *Voter) SetProposed(headers []*types.AttestHeader) error {
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
		if proposed[vote.AttestHeader.ToXChain()] {
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
func (v *Voter) SetCommitted(headers []*types.AttestHeader) error {
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
		if committed[vote.AttestHeader.ToXChain()] {
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
		commitHeight.WithLabelValues(v.network.ChainVersionName(vote.AttestHeader.XChainVersion())).Set(float64(vote.BlockHeader.BlockHeight))
	}

	return v.saveUnsafe()
}

func (v *Voter) LocalAddress() common.Address {
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

func (v *Voter) latestByChain(chainVer xchain.ChainVersion) (*types.Vote, bool) {
	v.mu.Lock()
	defer v.mu.Unlock()

	vote, ok := v.latest[chainVer]

	return vote, ok
}

// saveUnsafe saves the state to disk. It is unsafe since it assumes the lock is held.
func (v *Voter) saveUnsafe() error {
	sortVotes := func(atts []*types.Vote) {
		sort.Slice(atts, func(i, j int) bool {
			if atts[i].BlockHeader.ChainId != atts[j].BlockHeader.ChainId {
				return atts[i].BlockHeader.ChainId < atts[j].BlockHeader.ChainId
			}

			return atts[i].AttestHeader.AttestOffset < atts[j].AttestHeader.AttestOffset
		})
	}
	sortVotes(v.available)
	sortVotes(v.proposed)
	sortVotes(v.committed)

	s := stateJSON{
		Available: v.available,
		Proposed:  v.proposed,
		Committed: v.committed,
		Latest:    latestToJSON(v.latest),
	}
	bz, err := json.MarshalIndent(s, "", " ")
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
		counts := make(map[xchain.ChainVersion]int)
		for _, vote := range atts {
			counts[vote.AttestHeader.XChainVersion()]++
		}

		for chain, count := range counts {
			gaugeVec.WithLabelValues(v.network.ChainVersionName(chain)).Set(float64(count))
		}
	}

	count(v.available, availableCount)
	count(v.proposed, proposedCount)
}

// detectReorg returns an error if a reorg is detected based on the following conditions:
// - Previous block hash doesn't match the next block's parent hash.
// - Stream offsets are not consecutive.
func detectReorg(chainVer xchain.ChainVersion, prevBlock *xchain.Block, block xchain.Block, streamOffsets map[xchain.StreamID]uint64) error {
	if prevBlock == nil {
		return nil // Skip first block (without previous).
	}

	if prevBlock.BlockHeight+1 != block.BlockHeight {
		return errors.New("consecutive block height mismatch [BUG]", "prev_height", prevBlock.BlockHeight, "new_height", block.BlockHeight)
	}

	for _, xmsg := range block.Msgs {
		offset, ok := streamOffsets[xmsg.StreamID]
		if ok && xmsg.StreamOffset != offset+1 {
			return errors.New("non-consecutive message offsets", "streamID", xmsg.StreamID, "prev_offset", offset, "new_offset", xmsg.StreamOffset)
		}

		// Update the cursor
		streamOffsets[xmsg.StreamID] = xmsg.StreamOffset
	}

	if block.BlockHash == (common.Hash{}) {
		return nil // Skip consensus chain blocks without block hashes.
	}
	if prevBlock.BlockHash == block.ParentHash {
		return nil // No reorg detected.
	}

	if chainVer.ConfLevel.IsFuzzy() {
		return errors.New("fuzzy chain reorg detected", "height", block.BlockHeight, "parent_hash", prevBlock.BlockHash, "new_parent_hash", block.ParentHash)
	}

	return errors.New("finalized chain reorg detected [BUG]", "height", block.BlockHeight, "parent_hash", prevBlock.BlockHash, "new_parent_hash", block.ParentHash)
}

// stateJSON is the JSON representation of the attester state.
type stateJSON struct {
	Available []*types.Vote `json:"available"`
	Proposed  []*types.Vote `json:"proposed"`
	Committed []*types.Vote `json:"committed"`
	Latest    []*types.Vote `json:"latest"`
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

	verify := func(voteSets ...[]*types.Vote) error {
		for _, votes := range voteSets {
			for _, vote := range votes {
				if err := vote.Verify(); err != nil {
					return errors.Wrap(err, "verify vote")
				}
			}
		}

		return nil
	}

	if err := verify(s.Latest, s.Proposed, s.Committed, s.Available); err != nil {
		return stateJSON{}, err
	}

	return s, nil
}

// headerMap converts a list of headers to a bool map (set).
func headerMap(headers []*types.AttestHeader) map[xchain.AttestHeader]bool {
	resp := make(map[xchain.AttestHeader]bool) // Can't use protos as map keys.
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
		if ok && latestAtt.AttestHeader.AttestOffset >= vote.AttestHeader.AttestOffset {
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

// newDebugLogFilterer returns a debug log function that only logs once per period per message.
func newDebugLogFilterer(period time.Duration) func(context.Context, string, ...any) {
	lastByMsg := make(map[string]time.Time)
	return func(ctx context.Context, msg string, args ...any) {
		if last, ok := lastByMsg[msg]; ok && time.Since(last) < period {
			return
		}
		lastByMsg[msg] = time.Now()

		log.Debug(ctx, msg, args...)
	}
}

func latestMsgOffsets(msgs []xchain.Msg) map[xchain.StreamID]uint64 {
	resp := make(map[xchain.StreamID]uint64)
	for _, msg := range msgs {
		if resp[msg.StreamID] < msg.StreamOffset {
			resp[msg.StreamID] = msg.StreamOffset
		}
	}

	return resp
}

func latestToJSON(latest map[xchain.ChainVersion]*types.Vote) []*types.Vote {
	resp := make([]*types.Vote, 0, len(latest))
	for _, v := range latest {
		resp = append(resp, v)
	}

	return resp
}

func latestFromJSON(latest []*types.Vote) map[xchain.ChainVersion]*types.Vote {
	resp := make(map[xchain.ChainVersion]*types.Vote, len(latest))
	for _, v := range latest {
		resp[v.AttestHeader.XChainVersion()] = v
	}

	return resp
}

func logVoteCreated(ctx context.Context, network netconf.Network, attHeader xchain.AttestHeader, block xchain.Block) {
	// Collect start offsets per shard.
	startOffsets := make(map[string]uint64)
	for _, msg := range block.Msgs {
		emitShard := fmt.Sprintf("%s|%s", msg.ShardID.Label(), network.ChainName(msg.DestChainID))
		if _, ok := startOffsets[emitShard]; ok {
			continue
		}
		startOffsets[emitShard] = msg.StreamOffset
	}

	attrs := []any{
		"offset", attHeader.AttestOffset,
		"msgs", len(block.Msgs),
	}
	for shard, offset := range startOffsets {
		attrs = append(attrs, shard, offset)
	}

	log.Debug(ctx, "📬 Created vote for cross chain block", attrs...)
}
