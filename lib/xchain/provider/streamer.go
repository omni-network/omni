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
			currCtx := log.WithCtx(ctx, "height", currentHeight)

			xBlock, ok := p.fetchXBlock(currCtx, chainID, currentHeight, backoff, reset)
			if !ok { // this will happen only if the context is killed
				return
			}

			// deliver the fetched xBlock
			for currCtx.Err() == nil {
				err := callback(currCtx, xBlock)
				if currCtx.Err() != nil {
					return // Application context is killed
				} else if err != nil {
					log.Warn(currCtx, "Failure delivering xblock callback (will retry)", err)
					backoff()

					continue
				}

				break // successfully delivered the xBlock
			}
			reset() // delivery backoff reset

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
