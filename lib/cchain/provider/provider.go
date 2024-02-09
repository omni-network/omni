// Package provider implements the cchain.Provider interface.
package provider

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
)

var _ cchain.Provider = Provider{}

// FetchFunc abstracts fetching attestation from the consensus chain.
type FetchFunc func(ctx context.Context, chainID uint64, fromHeight uint64,
) ([]xchain.AggAttestation, error)

// Provider implements cchain.Provider.
type Provider struct {
	fetch       FetchFunc
	backoffFunc func(context.Context) (func(), func())
}

// TODO(corver): Add prod constructor once halo has an API.

// NewProviderForT creates a new provider for testing.
func NewProviderForT(_ *testing.T, fetch FetchFunc,
	backoffFunc func(context.Context) (func(), func()),
) Provider {
	return Provider{
		fetch:       fetch,
		backoffFunc: backoffFunc,
	}
}

func (p Provider) ApprovedFrom(ctx context.Context, sourceChainID uint64, sourceHeight uint64,
) ([]xchain.AggAttestation, error) {
	return p.fetch(ctx, sourceChainID, sourceHeight)
}

// Subscribe implements cchain.Provider.
func (p Provider) Subscribe(ctx context.Context, chainID uint64, height uint64, callback cchain.ProviderCallback) {
	// Start a async goroutine to fetch attestations until ctx is canceled.
	go func() {
		backoff, reset := p.backoffFunc(ctx) // Note that backoff returns immediately on ctx cancel.

		for ctx.Err() == nil {
			// Fetch next batch of attestations.
			atts, err := p.fetch(ctx, chainID, height)
			if ctx.Err() != nil {
				return // Don't backoff or log on ctx cancel, just return.
			} else if err != nil {
				log.Warn(ctx, "Failed fetching attestation; will retry", err)
				backoff()

				continue
			} else if len(atts) == 0 {
				// We reached the head of the chain, wait for new blocks.
				backoff() // Maybe do (consensus-block-period / N)

				continue
			}

			reset() // Reset fetch backoff

			// Call callback for each attestation
			for _, att := range atts {
				// Sanity checks
				if att.SourceChainID != chainID {
					log.Error(ctx, "Invalid attestation chain ID [BUG!]", nil)
					return
				} else if att.BlockHeight != height {
					log.Error(ctx, "Invalid attestation height [BUG!]", nil)
					return
				}

				// Retry callback on error
				for ctx.Err() == nil {
					err := callback(ctx, att)
					if ctx.Err() != nil {
						return // Don't backoff or log on ctx cancel, just return.
					} else if err != nil {
						log.Warn(ctx, "Failed processing attestation; will retry", err)
						callbackErrTotal.Inc()
						backoff()

						continue
					}
					streamHeight.Set(float64(height)) // Update stream height metric

					break // Success, stop retrying.
				}

				reset() // Reset callback backoff
				height++
			}
		}
	}()
}
