package provider

import (
	"context"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
)

// Streamer maintains the config and the destination for each chain.
type Streamer struct {
	chainConfig *ChainConfig            // the chain config which also has the subscription information
	callback    xchain.ProviderCallback // the callback to call on receiving a xblock
	backoffFunc func(context.Context) (func(), func())
}

// NewStreamer manages the rpc client to collect xblocks and delivers it to the
// subscriber through callback.
func NewStreamer(config *ChainConfig,
	callback xchain.ProviderCallback,
	backoffFunc func(context.Context) (func(), func()),
) *Streamer {
	// initialize the streamer structure with the received configuration
	stream := &Streamer{
		chainConfig: config,
		callback:    callback,
		backoffFunc: backoffFunc,
	}

	return stream
}

// streamBlocks triggers a continuously running routine that fetches and delivers xBlocks.
func (s *Streamer) streamBlocks(ctx context.Context, height uint64) {
	go func() {
		backoff, reset := s.backoffFunc(ctx)
		currentHeight := height

		// stream blocks until the context is canceled
		for ctx.Err() == nil {
			// fetch xBlock
			log.Debug(ctx, "Fetching block", "height", currentHeight)
			xBlock, exists := s.fetchXBlock(ctx, currentHeight, backoff, reset)

			if !exists {
				// no cross chain logs in this height, so go to the next height
				log.Debug(ctx, "No cross chain block", "height", currentHeight)
				currentHeight++

				continue
			}

			// deliver the fetched xBlock
			s.deliverXBlock(ctx, currentHeight, xBlock, backoff, reset)
			log.Debug(ctx, "Delivered xBlock", "height", currentHeight)

			currentHeight++
		}
	}()
}

func (s *Streamer) fetchXBlock(ctx context.Context,
	currentHeight uint64,
	backoff func(),
	reset func(),
) (xchain.Block, bool) {
	// fetch xBlock
	for ctx.Err() == nil {
		// get the message and receipts from the chain for this block if any
		xBlock, exists, err := s.chainConfig.rpcClient.GetBlock(ctx, currentHeight)
		if ctx.Err() != nil {
			return xBlock, false
		}
		if err != nil {
			log.Warn(ctx, "Could not fetch xBlock, will retry again after sometime", err,
				"height", currentHeight)
			backoff() // backoff and retry fetching the block

			continue
		}

		// err == nil and exists == false means the height is not finalized yet
		// so backoff
		if !exists {
			backoff()

			continue
		}

		// err == nil and exists = true means we have a xBlock
		// so reset the backoff and return
		reset()

		return xBlock, exists
	}

	return xchain.Block{}, false
}

func (s *Streamer) deliverXBlock(ctx context.Context,
	currentHeight uint64,
	xBlock xchain.Block,
	backoff func(),
	reset func(),
) {
	// deliver the fetched xBlock
	for ctx.Err() == nil {
		err := s.callback(ctx, &xBlock)
		if ctx.Err() != nil {
			return
		}
		if err != nil {
			log.Warn(ctx, "Could not deliver xBlock, will retry again after sometime", err,
				"height", currentHeight)
			backoff() // try delivering after sometime

			continue
		}
		reset() // delivery backoff reset

		break // successfully delivered the xBlock
	}
}
