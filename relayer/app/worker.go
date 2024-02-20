package relayer

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
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
	sendProvider func() (SendFunc, error)
}

// NewWorker creates a new worker for a single destination chain.
func NewWorker(destChain netconf.Chain, network netconf.Network, cProvider cchain.Provider,
	xProvider xchain.Provider, creator CreateFunc, sendProvider func() (SendFunc, error),
) *Worker {
	return &Worker{
		destChain:    destChain,
		network:      network,
		cProvider:    cProvider,
		xProvider:    xProvider,
		creator:      creator,
		sendProvider: sendProvider,
	}
}

func (w *Worker) Run(ctx context.Context) {
	ctx = log.WithCtx(ctx, "dst_chain", w.destChain.Name)
	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second)) // TODO(corver): Improve backoff.
	for ctx.Err() == nil {
		err := w.runOnce(ctx)
		if ctx.Err() != nil {
			return
		}

		log.Error(ctx, "Worker failed, resetting", err)
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
	for srcChainID, fromHeight := range FromHeights(cursors, w.destChain, w.network.Chains) {
		if srcChainID == w.destChain.ID { // Sanity check
			return errors.New("unexpected cursor [BUG]")
		}

		callback := newCallback(w.xProvider, initialOffsets, w.creator, buf.AddInput, w.destChain.ID)

		w.cProvider.Subscribe(ctx, srcChainID, fromHeight, w.destChain.Name, callback)

		srcChain, f := w.network.Chain(srcChainID)
		if !f {
			continue
		}
		logAttrs = append(logAttrs, srcChain.Name, fromHeight)
	}

	log.Info(ctx, "Worker subscribed to chains", logAttrs...)

	return buf.Run(ctx)
}

func newCallback(xProvider xchain.Provider, initialOffsets map[xchain.StreamID]uint64, creator CreateFunc,
	sender SendFunc, destChainID uint64) cchain.ProviderCallback {
	return func(ctx context.Context, att xchain.AggAttestation) error {
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
		for streamID, msgs := range mapByStreamID(block.Msgs) {
			if streamID.DestChainID != destChainID {
				continue
			}

			msgs = filterMsgs(msgs, initialOffsets, streamID) // Filter out any partially submitted stream updates.
			if len(msgs) == 0 {
				continue
			}

			update := StreamUpdate{
				StreamID:       streamID,
				AggAttestation: att,
				Msgs:           msgs,
				Tree:           tree,
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
