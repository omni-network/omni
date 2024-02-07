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
	// bufferSize is the size of the activeBuffer for async send.
	bufferSize = 1024
	// mempoolLimit is the maximum number of transactions we want to submit to the mempool at once.
	mempoolLimit = 16
)

type Worker struct {
	chain        netconf.Chain // Destination chain
	network      netconf.Network
	cProvider    cchain.Provider
	xProvider    xchain.Provider
	creator      CreateFunc
	sendProvider func() (SendFunc, error)
}

// NewWorker creates a new worker for a single destination chain.
func NewWorker(chain netconf.Chain, network netconf.Network, cProvider cchain.Provider,
	xProvider xchain.Provider, creator CreateFunc, sendProvider func() (SendFunc, error),
) *Worker {
	return &Worker{
		chain:        chain,
		network:      network,
		cProvider:    cProvider,
		xProvider:    xProvider,
		creator:      creator,
		sendProvider: sendProvider,
	}
}

func (w *Worker) Run(ctx context.Context) {
	ctx = log.WithCtx(ctx, "chain", w.chain.Name)

	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second)) // TODO(corver): Improve backoff.
	for ctx.Err() == nil {
		err := w.runOnce(ctx)
		if ctx.Err() != nil {
			return
		}

		log.Error(ctx, "Worker failed, resetting", err)
		workerResets.WithLabelValues(w.chain.Name).Inc()
		backoff()
	}
}

func (w *Worker) runOnce(ctx context.Context) error {
	log.Info(ctx, "Worker starting")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cursors, initialOffsets, err := getSubmittedCursors(ctx, w.network.ChainIDs(), []uint64{w.chain.ID}, w.xProvider)
	if err != nil {
		return err
	}

	sender, err := w.sendProvider()
	if err != nil {
		return err
	}

	buf := newActiveBuffer(w.chain.Name, mempoolLimit, bufferSize, sender)

	callback := newCallback(w.xProvider, initialOffsets, w.creator, buf.AddInput)

	var logAttrs []any //nolint:prealloc // Not worth it
	for chainID, fromHeight := range FromHeights(cursors, w.network.Chains) {
		w.cProvider.Subscribe(ctx, chainID, fromHeight, callback)

		srcChain, _ := w.network.Chain(chainID)
		logAttrs = append(logAttrs, srcChain.Name, fromHeight)
	}

	log.Info(ctx, "Worker subscribed to chains", logAttrs...)

	return buf.Run(ctx)
}

func newCallback(xProvider xchain.Provider, initialOffsets map[xchain.StreamID]uint64, creator CreateFunc,
	sender SendFunc) cchain.ProviderCallback {
	return func(ctx context.Context, att xchain.AggAttestation) error {
		// Get the xblock from the source chain.
		block, ok, err := xProvider.GetBlock(ctx, att.SourceChainID, att.BlockHeight)
		if err != nil {
			return err
		} else if !ok { // Sanity check, should never happen.
			return errors.New("attestation block not finalized [BUG!]",
				"chain", att.SourceChainID,
				"height", att.BlockHeight,
			)
		} else if block.BlockHash != att.BlockHash { // Sanity check, should never happen.
			return errors.New("attestation block hash mismatch [BUG!]",
				"chain", att.SourceChainID,
				"height", att.BlockHeight,
				log.Hex7("attestation_hash", att.BlockHash[:]),
				log.Hex7("block_hash", block.BlockHash[:]),
			)
		} else if len(block.Msgs) == 0 {
			log.Debug(ctx, "Skipping empty attested block",
				"height", att.BlockHeight, "source_chain_id", att.SourceChainID)

			return nil
		}

		// Split into streams
		for streamID, msgs := range mapByStreamID(block.Msgs) {
			msgs = filterMsgs(msgs, initialOffsets, streamID) // Filter out any partially submitted stream updates.
			if len(msgs) == 0 {
				continue
			}

			update := StreamUpdate{
				StreamID:       streamID,
				AggAttestation: att,
				Msgs:           msgs,
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
