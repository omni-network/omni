package relayer

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

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
	state        *State
	sendProvider func() (SendFunc, error)
	awaitValSet  awaitValSet
}

// NewWorker creates a new worker for a single destination chain.
func NewWorker(destChain netconf.Chain, network netconf.Network, cProvider cchain.Provider,
	xProvider xchain.Provider, creator CreateFunc, sendProvider func() (SendFunc, error), state *State,
	awaitValSet awaitValSet,
) *Worker {
	return &Worker{
		destChain:    destChain,
		network:      network,
		cProvider:    cProvider,
		xProvider:    xProvider,
		creator:      creator,
		sendProvider: sendProvider,
		state:        state,
		awaitValSet:  awaitValSet,
	}
}

func (w *Worker) Run(ctx context.Context) {
	ctx = log.WithCtx(ctx, "dst_chain", w.destChain.Name)
	backoff := expbackoff.NewWithAutoReset(ctx)
	for ctx.Err() == nil {
		err := w.runOnce(ctx)
		if ctx.Err() != nil {
			return
		}

		log.Error(ctx, "Worker failed, resetting", err)

		if err := w.state.Clear(w.destChain.ID); err != nil {
			log.Error(ctx, "Failed to clear worker state", err)
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

	for _, cursor := range cursors {
		log.Info(ctx, "Worker fetched submitted cursor",
			"stream", w.network.StreamName(cursor.StreamID),
			"block_offset", cursor.BlockOffset,
			"msg_offset", cursor.MsgOffset,
			"valset_id", cursor.ValidatorSetID,
		)
	}

	sender, err := w.sendProvider()
	if err != nil {
		return err
	}

	buf := newActiveBuffer(w.destChain.Name, mempoolLimit, sender)

	blockOffsets, err := fromChainVersionOffsets(w.destChain.ID, cursors, w.network.ChainVersionsTo(w.destChain.ID), w.state)
	if err != nil {
		return err
	}

	msgFilter, err := newMsgOffsetFilter(cursors)
	if err != nil {
		return err
	}

	var logAttrs []any //nolint:prealloc // Not worth it
	for chainVer, fromOffset := range blockOffsets {
		if chainVer.ID == w.destChain.ID { // Sanity check
			return errors.New("unexpected chain version [BUG]")
		}

		callback := newCallback(w.xProvider, msgFilter, w.creator, buf.AddInput, w.destChain.ID, newMsgStreamMapper(w.network), w.awaitValSet)
		wrapCb := wrapStatePersist(callback, w.state, w.destChain.ID)

		w.cProvider.Subscribe(ctx, chainVer, fromOffset, w.destChain.Name, wrapCb)

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
			power, err := portal.ValidatorSetTotalPower(&bind.CallOpts{Context: ctx}, valsetID)
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

func newCallback(
	xProvider xchain.Provider,
	msgFilter *msgCursorFilter,
	creator CreateFunc,
	sender SendFunc,
	destChainID uint64,
	msgStreamMapper msgStreamMapper,
	awaitValSet awaitValSet,
) cchain.ProviderCallback {
	return func(ctx context.Context, att xchain.Attestation) error {
		block, ok, err := fetchXBlock(ctx, xProvider, att)
		if err != nil {
			return err
		} else if !ok {
			return nil // Mismatching fuzzy attestation, skip.
		}

		tree, err := xchain.NewBlockTree(block)
		if err != nil {
			return err
		}

		// Split into streams
		for streamID, msgs := range msgStreamMapper(block.Msgs) {
			if streamID.DestChainID != destChainID {
				continue // Skip streams not destined for this worker.
			} else if !attestationForShard(att, streamID.ShardID) {
				continue // Skip streams not applicable to this attestation.
			}

			if err := awaitValSet(ctx, att.ValidatorSetID); err != nil {
				return errors.Wrap(err, "await validator set")
			}

			// Filter out any previously submitted message offsets
			msgs, err = filterMsgs(ctx, streamID, att.ValidatorSetID, msgs, msgFilter)
			if err != nil {
				return err
			} else if len(msgs) == 0 {
				continue
			}

			update := StreamUpdate{
				StreamID:    streamID,
				Attestation: att,
				Msgs:        msgs,
				Tree:        tree,
			}

			submissions, err := creator(update)
			if err != nil {
				return err
			}

			for _, subs := range submissions {
				if err := sender(ctx, subs); err != nil {
					return err
				}
			}
		}

		return nil
	}
}

// wrapStatePersist wraps a provider callback, persisting successful processed block offsets per chain version to local state.
func wrapStatePersist(cb cchain.ProviderCallback, state *State, destChainID uint64) cchain.ProviderCallback {
	return func(ctx context.Context, att xchain.Attestation) error {
		if err := cb(ctx, att); err != nil {
			return err
		}

		if err := state.Persist(destChainID, att.ChainVersion(), att.BlockOffset); err != nil {
			return errors.Wrap(err, "persist state")
		}

		return nil
	}
}

// fetchXBlock gets the xblock from the source chain (retry up to 10s if block-not-finalized).
func fetchXBlock(rootCtx context.Context, xProvider xchain.Provider, att xchain.Attestation) (xchain.Block, bool, error) {
	ctx, cancel := context.WithTimeout(rootCtx, 10*time.Second)
	defer cancel()

	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second))
	for {
		req := xchain.ProviderRequest{
			ChainID:   att.SourceChainID,
			Height:    att.BlockHeight,
			Offset:    att.BlockOffset,
			ConfLevel: att.ConfLevel,
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
			if att.ConfLevel.IsFuzzy() {
				log.Warn(ctx, "Skipping fuzzy attestation mismatching block", err)
				return block, false, nil
			}

			return xchain.Block{}, false, errors.Wrap(err, "mismatching block vs finalized attestation [BUG]")
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
	} else if len(block.Msgs) == 0 {
		// All attestations must map to non-empty xblocks
		return errors.New("unexpected empty xblock")
	} else if block.BlockOffset != att.BlockOffset {
		// All attestations must map to non-empty xblocks with XBlockOffset populated.
		return errors.New("unexpected XBlockOffset")
	}

	tree, _ := xchain.NewBlockTree(block)
	if root := tree.Root(); att.AttestationRoot != root {
		return errors.New("attestation root mismatch",
			log.Hex7("attestation_root", att.AttestationRoot[:]),
			log.Hex7("block_root", root[:]),
		)
	}

	return nil
}

// attestationForShard returns true if the attestation proof contains messages for the shard.
// Fuzzy attestations cannot be used to prove finalized shards. But finalized attestations can prove all shards.
func attestationForShard(att xchain.Attestation, shard xchain.ShardID) bool {
	if att.ConfLevel == xchain.ConfFinalized {
		return true // Finalized attestation, matches all streams.
	}

	return att.ConfLevel == shard.ConfLevel()
}
