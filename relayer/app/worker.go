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

	cursors, initialOffsets, err := getSubmittedCursors(ctx, w.network, w.destChain.ID, w.xProvider)
	if err != nil {
		return err
	}

	sender, err := w.sendProvider()
	if err != nil {
		return err
	}

	buf := newActiveBuffer(w.destChain.Name, mempoolLimit, sender)

	var logAttrs []any //nolint:prealloc // Not worth it
	for srcChainID, fromHeight := range fromHeights(cursors, w.destChain, w.network.Chains, w.state) {
		if srcChainID == w.destChain.ID { // Sanity check
			return errors.New("unexpected cursor [BUG]")
		}

		callback := newCallback(w.xProvider, initialOffsets, w.creator, buf.AddInput, w.destChain.ID, newMsgStreamMapper(w.network), w.awaitValSet)
		wrapCb := wrapStatePersist(callback, w.state, w.destChain.ID)

		w.cProvider.Subscribe(ctx, srcChainID, fromHeight, w.destChain.Name, wrapCb)

		srcChain, f := w.network.Chain(srcChainID)
		if !f {
			continue
		}
		logAttrs = append(logAttrs, srcChain.Name, fromHeight)
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
				}
				resp[streamID] = append(resp[streamID], msg)
			}
		}

		return resp
	}
}

func newCallback(xProvider xchain.Provider, initialOffsets map[xchain.StreamID]uint64, creator CreateFunc,
	sender SendFunc, destChainID uint64, msgStreamMapper msgStreamMapper, awaitValSet awaitValSet) cchain.ProviderCallback {
	return func(ctx context.Context, att xchain.Attestation) error {
		// Get the xblock from the source chain.
		block, ok, err := xProvider.GetBlock(ctx, att.SourceChainID, att.BlockHeight)
		if err != nil {
			return err
		} else if !ok { // Sanity check, should never happen.
			return errors.New("attestation block not finalized [BUG!]")
		} else if block.BlockHash != att.BlockHash { // Sanity check, should never happen.
			return errors.New("attestation block hash mismatch [BUG!]",
				log.Hex7("attestation_hash", att.BlockHash[:]),
				log.Hex7("block_hash", block.BlockHash[:]),
			)
		} else if len(block.Msgs) == 0 {
			return nil
		}

		tree, err := xchain.NewBlockTree(block)
		if err != nil {
			return err
		}
		// Split into streams
		for streamID, msgs := range msgStreamMapper(block.Msgs) {
			if streamID.DestChainID != destChainID {
				continue
			}

			if err := awaitValSet(ctx, att.ValidatorSetID); err != nil {
				return errors.Wrap(err, "await validator set")
			}

			msgs = filterMsgs(msgs, initialOffsets, streamID) // Filter out any partially submitted stream updates.
			if len(msgs) == 0 {
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

func wrapStatePersist(cb cchain.ProviderCallback, state *State, destChainID uint64) cchain.ProviderCallback {
	return func(ctx context.Context, att xchain.Attestation) error {
		if err := cb(ctx, att); err != nil {
			return err
		}

		if err := state.Persist(destChainID, att.SourceChainID, att.BlockHeight); err != nil {
			return errors.Wrap(err, "persist state")
		}

		return nil
	}
}
