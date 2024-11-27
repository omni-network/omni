package relayer

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/chaos"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/relayer/app/cursor"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const (
	// mempoolLimit is the maximum number of transactions we want to submit to the mempool at once.
	mempoolLimit = 16
)

type Worker struct {
	destChain    netconf.Chain // Destination chain
	network      netconf.Network
	cProvider    cchain.Provider
	xProvider    xchain.Provider
	creator      CreateFunc
	sendProvider func() (SendAsync, error)
	awaitValSet  awaitValSet
	cursors      *cursor.Store
}

// NewWorker creates a new worker for a single destination chain.
func NewWorker(
	destChain netconf.Chain,
	network netconf.Network,
	cProvider cchain.Provider,
	xProvider xchain.Provider,
	creator CreateFunc,
	sendProvider func() (SendAsync, error),
	awaitValSet awaitValSet,
	cursors *cursor.Store,
) *Worker {
	return &Worker{
		destChain:    destChain,
		network:      network,
		cProvider:    cProvider,
		xProvider:    xProvider,
		creator:      creator,
		sendProvider: sendProvider,
		awaitValSet:  awaitValSet,
		cursors:      cursors,
	}
}

func (w *Worker) Run(ctx context.Context) {
	ctx = log.WithCtx(ctx, "dst_chain", w.destChain.Name)
	backoff := expbackoff.NewWithAutoReset(ctx)
	for ctx.Err() == nil {
		err := w.runOnce(ctx)
		if ctx.Err() != nil {
			return
		} else if errors.Is(err, chaos.ErrChaos) {
			log.InfoErr(ctx, "Worker failed due to chaos testing, resetting", err)
		} else {
			log.Warn(ctx, "Worker failed, resetting", err)
		}

		workerResets.WithLabelValues(w.destChain.Name).Inc()
		backoff()
	}
}

func (w *Worker) runOnce(ctx context.Context) error {
	log.Info(ctx, "Worker starting")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cursors, err := getSubmittedCursors(ctx, w.network, w.destChain.ID, w.xProvider)
	if err != nil {
		return err
	}

	for _, c := range cursors {
		log.Info(ctx, "Worker fetched submitted cursor",
			"stream", w.network.StreamName(c.StreamID),
			"attest_offset", c.AttestOffset,
			"msg_offset", c.MsgOffset,
		)
	}

	sender, err := w.sendProvider()
	if err != nil {
		return err
	}

	buf := newActiveBuffer(w.destChain.Name, mempoolLimit, sender)

	attestOffsets, err := fromChainVersionOffsets(cursors, w.network.ChainVersionsTo(w.destChain.ID))
	if err != nil {
		return err
	}

	// TODO(corver): Remove this bootstrap workaround.
	for chainVer, offset := range attestOffsets {
		bootstrapOffset, ok := getBootstrapOffset(w.network, w.destChain, chainVer)
		if !ok {
			continue
		} else if bootstrapOffset > offset {
			log.Info(ctx, "Worker using bootstrap attest offset",
				"chain_version", w.network.ChainVersionName(chainVer),
				"prev", offset,
				"bootstrap", bootstrapOffset,
			)
			attestOffsets[chainVer] = bootstrapOffset
		}
	}

	// Update offsets if any stored cursor is higher
	stored, err := w.cursors.WorkerOffsets(ctx, w.destChain.ID)
	if err != nil {
		return err
	}
	for chainVer, offset := range attestOffsets {
		if stored[chainVer] > offset {
			log.Info(ctx, "Worker using stored attest offset",
				"chain_version", w.network.ChainVersionName(chainVer),
				"prev", offset,
				"bootstrap", stored[chainVer],
			)
			attestOffsets[chainVer] = stored[chainVer]
		}
	}

	msgFilter, err := newMsgOffsetFilter(cursors)
	if err != nil {
		return err
	}

	var logAttrs []any //nolint:prealloc // Not worth it
	for chainVer, fromOffset := range attestOffsets {
		if chainVer.ID == w.destChain.ID { // Sanity check
			return errors.New("unexpected chain version [BUG]")
		}

		callback := w.newCallback(
			msgFilter,
			buf.AddInput,
			newMsgStreamMapper(w.network),
			chainVer,
		)

		w.cProvider.StreamAsync(ctx, chainVer, fromOffset, w.destChain.Name, callback)

		logAttrs = append(logAttrs, w.network.ChainVersionName(chainVer), fromOffset)
	}

	log.Info(ctx, "Worker subscribed to chains", logAttrs...)

	return buf.Run(ctx)
}

// awaitValSet blocks until the portal is aware of this validator set ID.
type awaitValSet func(ctx context.Context, valsetID uint64) error

// newValSetAwaiter creates a new awaitValSet function for the given portal.
func newValSetAwaiter(portal *bindings.OmniPortal, blockPeriod time.Duration) awaitValSet {
	var prev atomic.Uint64 // Cache previous to reduce network lookups.
	return func(ctx context.Context, valsetID uint64) error {
		if prev.Load() == valsetID {
			return nil
		}
		backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(blockPeriod))
		var attempt int
		for ctx.Err() == nil {
			power, err := portal.ValSetTotalPower(&bind.CallOpts{Context: ctx}, valsetID)
			if err != nil {
				return errors.Wrap(err, "get validator set power")
			}
			if power == 0 {
				attempt++
				if attempt%10 == 0 {
					log.Warn(ctx, "Validator set not known by portal (will retry)", nil, "valset_id", valsetID, "attempt", attempt)
				}
				backoff()

				continue
			}

			prev.Store(valsetID)

			return nil
		}

		return errors.Wrap(ctx.Err(), "context done")
	}
}

// msgStreamMapper maps messages by stream ID.
type msgStreamMapper func([]xchain.Msg) map[xchain.StreamID][]xchain.Msg

// newMsgStreamMapper creates a new message stream mapper for the given network.
// It maps consensus chain messages to all EVM chains (broadcast), and normal messages to their stream ID.
func newMsgStreamMapper(network netconf.Network) msgStreamMapper {
	consensusChain, _ := network.OmniConsensusChain()

	return func(msgs []xchain.Msg) map[xchain.StreamID][]xchain.Msg {
		resp := make(map[xchain.StreamID][]xchain.Msg)
		for _, msg := range msgs {
			// Normal messages are mapped to their stream ID.
			if msg.SourceChainID != consensusChain.ID {
				resp[msg.StreamID] = append(resp[msg.StreamID], msg)
				continue
			}

			// Consensus chain messages are broadcasted to all EVM chains.
			for _, evmChain := range network.EVMChains() {
				streamID := xchain.StreamID{
					SourceChainID: consensusChain.ID,
					DestChainID:   evmChain.ID,
					ShardID:       msg.ShardID,
				}
				resp[streamID] = append(resp[streamID], msg)
			}
		}

		return resp
	}
}

func (w *Worker) newCallback(
	msgFilter *msgCursorFilter,
	sendBuffer func(context.Context, xchain.Submission) error,
	msgStreamMapper msgStreamMapper,
	streamerChainVer xchain.ChainVersion, // Use streamer chain version for cursors since fuzzy overrides otherwise store latest streamed offsets to finalized streamer.
) cchain.ProviderCallback {
	var cachedValSetID uint64
	var cachedValSet []cchain.PortalValidator

	return func(ctx context.Context, att xchain.Attestation) error {
		saveCursors := func(streamMsgs map[xchain.StreamID][]xchain.Msg) error {
			return w.cursors.Save(ctx, streamerChainVer, w.destChain.ID, att.AttestOffset, streamMsgs)
		}

		block, ok, err := fetchXBlock(ctx, w.xProvider, att)
		if err != nil {
			return err
		} else if !ok {
			return nil // Mismatching fuzzy attestation, skip.
		} else if len(block.Msgs) == 0 {
			// No messages, nothing to do, just update cursors
			return saveCursors(nil)
		}

		msgTree, err := xchain.NewMsgTree(block.Msgs)
		if err != nil {
			return err
		}

		if att.ValidatorSetID != cachedValSetID {
			cachedValSetID = att.ValidatorSetID
			cachedValSet, ok, err = w.cProvider.PortalValidatorSet(ctx, att.ValidatorSetID)
			if err != nil {
				return errors.Wrap(err, "fetch validator set")
			} else if !ok {
				return errors.New("validator set not found [BUG]", "valset", att.ValidatorSetID)
			}
		}

		// Add all conf-level applicable messages to cursor store
		applicable := make(map[xchain.StreamID][]xchain.Msg)

		// Split into streams
		for streamID, msgs := range msgStreamMapper(block.Msgs) {
			if streamID.DestChainID != w.destChain.ID {
				continue // Skip streams not destined for this worker.
			} else if !attestationForShard(att, streamID.ShardID) {
				continue // Skip streams not applicable to this attestation conf level.
			}

			applicable[streamID] = msgs

			if err := w.awaitValSet(ctx, att.ValidatorSetID); err != nil {
				return errors.Wrap(err, "await validator set")
			}

			// Filter out any previously submitted message offsets
			msgs, err = filterMsgs(ctx, streamID, w.network.StreamName, msgs, msgFilter)
			if err != nil {
				return err
			} else if len(msgs) == 0 {
				continue
			}

			update := StreamUpdate{
				StreamID:    streamID,
				Attestation: att,
				Msgs:        msgs,
				MsgTree:     msgTree,
				ValSet:      cachedValSet,
			}

			submissions, err := w.creator(update)
			if err != nil {
				return err
			}

			for _, subs := range submissions {
				if err := sendBuffer(ctx, subs); err != nil {
					return err
				}
			}
		}

		return saveCursors(applicable)
	}
}

// fetchXBlock gets the xblock from the source chain (retry up to 10s if block-not-finalized).
func fetchXBlock(rootCtx context.Context, xProvider xchain.Provider, att xchain.Attestation) (xchain.Block, bool, error) {
	ctx, cancel := context.WithTimeout(rootCtx, 10*time.Second)
	defer cancel()

	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second))
	for {
		req := xchain.ProviderRequest{
			ChainID:   att.ChainID,
			Height:    att.BlockHeight,
			ConfLevel: att.ChainVersion.ConfLevel,
		}
		block, ok, err := xProvider.GetBlock(ctx, req)
		if rootCtx.Err() != nil {
			return xchain.Block{}, false, errors.Wrap(rootCtx.Err(), "canceled") // Root context closed, shutting down
		} else if ctx.Err() != nil {
			return xchain.Block{}, false, errors.New("attestation block still not finalized (node lagging?)")
		} else if err != nil {
			return xchain.Block{}, false, err
		} else if !ok {
			// This happens sometimes if the evm node relayer is querying is lagging behind
			// the chain itself. Especially for omni_evm with instant finality, this does happen sometimes.
			// Just backoff and retry a few times.
			backoff()
			continue
		}

		if err := verifyAttBlock(att, block); err != nil {
			if att.ChainVersion.ConfLevel.IsFuzzy() {
				log.Warn(ctx, "Skipping fuzzy attestation mismatching block", err)
				return block, false, nil
			}

			return xchain.Block{}, false, errors.Wrap(err, "mismatching block vs finalized attestation [BUG]",
				"block_height", block.BlockHeight)
		}

		// We got the xblock, it is finalized and its hash matches the attestation block hash.
		return block, true, nil
	}
}

// verifyAttBlock verifies the attestation matches the xblock.
func verifyAttBlock(att xchain.Attestation, block xchain.Block) error {
	if block.BlockHash != att.BlockHash {
		return errors.New("attestation block hash mismatch",
			log.Hex7("attestation_hash", att.BlockHash[:]),
			log.Hex7("block_hash", block.BlockHash[:]),
		)
	} else if block.BlockHeader != att.BlockHeader {
		return errors.New("attestation block header mismatch")
	}

	var msgRoot [32]byte
	if len(block.Msgs) > 0 {
		msgTree, err := xchain.NewMsgTree(block.Msgs)
		if err != nil {
			return err
		}
		msgRoot = msgTree.MsgRoot()
	}

	if att.MsgRoot != msgRoot {
		return errors.New("attestation message root mismatch",
			log.Hex7("att_msg_root", att.MsgRoot[:]),
			log.Hex7("block_msg_root", msgRoot[:]),
		)
	}

	return nil
}

// attestationForShard returns true if the attestation proof contains messages for the shard.
// Fuzzy attestations cannot be used to prove finalized shards. But finalized attestations can prove all shards.
func attestationForShard(att xchain.Attestation, shard xchain.ShardID) bool {
	if att.ChainVersion.ConfLevel == xchain.ConfFinalized {
		return true // Finalized attestation, matches all streams.
	}

	return att.ChainVersion.ConfLevel == shard.ConfLevel()
}
