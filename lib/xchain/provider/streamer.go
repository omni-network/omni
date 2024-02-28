package provider

import (
	"context"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
)

// streamBlocks triggers a continuously running routine that fetches and delivers xBlocks.
func (p *Provider) streamBlocks(ctx context.Context, chainName string, chainID uint64, height uint64,
	callback xchain.ProviderCallback,
) {
	go func() {
		backoff, reset := p.backoffFunc(ctx)
		currentHeight := height

		// Stream blocks until the context is canceled
		for ctx.Err() == nil {
			currCtx := log.WithCtx(ctx, "height", currentHeight)

			// Fetch the next xblock.
			xBlock, exists, err := p.GetBlock(currCtx, chainID, currentHeight)
			if currCtx.Err() != nil {
				return // Application context is killed
			} else if err != nil {
				log.Warn(currCtx, "Failed fetching xBlock (will retry)", err)
				fetchErrTotal.WithLabelValues(chainName).Inc()
				backoff() // Backoff and retry fetching the block

				continue
			} else if !exists {
				// We reached the head of the (finalized) chain, wait for new blocks.
				backoff()

				continue
			}
			reset() // Reset fetch backoff

			// deliver the fetched xBlock
			for currCtx.Err() == nil {
				err := callback(currCtx, xBlock)
				if currCtx.Err() != nil {
					return // Application context is killed
				} else if err != nil {
					log.Warn(currCtx, "Failure delivering xblock callback (will retry)", err)
					callbackErrTotal.WithLabelValues(chainName).Inc()
					backoff()

					continue
				}

				streamHeight.WithLabelValues(chainName).Set(float64(currentHeight))

				break // successfully delivered the xBlock
			}
			reset() // delivery backoff reset

			currentHeight++
		}
	}()
}
