package provider

import (
	"context"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
)

// streamBlocks triggers a continuously running routine that fetches and delivers xBlocks.
func (p *Provider) streamBlocks(ctx context.Context, chainID uint64, height uint64, callback xchain.ProviderCallback) {
	go func() {
		backoff, reset := p.backoffFunc(ctx)
		currentHeight := height

		// stream blocks until the context is canceled
		for ctx.Err() == nil {
			// fetch xBlock
			log.Debug(ctx, "Fetching block", "height", currentHeight)
			xBlock, ok := p.fetchXBlock(ctx, chainID, currentHeight, backoff, reset)
			if !ok {
				// this will happen only if the context is killed
				return
			}

			// deliver the fetched xBlock
			deliverXBlock(ctx, currentHeight, xBlock, callback, backoff, reset)
			log.Debug(ctx, "Delivered xBlock", "height", currentHeight)

			currentHeight++
		}
	}()
}

func (p *Provider) fetchXBlock(ctx context.Context,
	chainID uint64,
	currentHeight uint64,
	backoff func(),
	reset func(),
) (xchain.Block, bool) {
	// fetch xBlock
	for ctx.Err() == nil {
		// get the message and receipts from the chain for this block if any
		xBlock, exists, err := p.GetBlock(ctx, chainID, currentHeight)
		if ctx.Err() != nil {
			return xchain.Block{}, false
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

func deliverXBlock(ctx context.Context,
	currentHeight uint64,
	xBlock xchain.Block,
	callback xchain.ProviderCallback,
	backoff func(),
	reset func(),
) {
	// deliver the fetched xBlock
	for ctx.Err() == nil {
		err := callback(ctx, &xBlock)
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
