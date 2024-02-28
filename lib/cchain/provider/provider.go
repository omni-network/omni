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

// FetchFunc abstracts fetching attestations from the consensus chain.
type FetchFunc func(ctx context.Context, chainID uint64, fromHeight uint64,
) ([]xchain.Attestation, error)

// LatestFunc abstracts fetching the latest attestation from the consensus chain.
type LatestFunc func(ctx context.Context, chainID uint64) (xchain.Attestation, bool, error)
type WindowFunc func(ctx context.Context, chainID uint64, height uint64) (int, error)

// Provider implements cchain.Provider.
type Provider struct {
	fetch       FetchFunc
	latest      LatestFunc
	window      WindowFunc
	backoffFunc func(context.Context) (func(), func())
	chainNames  map[uint64]string
}

// NewProviderForT creates a new provider for testing.
func NewProviderForT(_ *testing.T, fetch FetchFunc, latest LatestFunc, window WindowFunc,
	backoffFunc func(context.Context) (func(), func()),
) Provider {
	return Provider{
		latest:      latest,
		fetch:       fetch,
		window:      window,
		backoffFunc: backoffFunc,
	}
}

func (p Provider) AttestationsFrom(ctx context.Context, sourceChainID uint64, sourceHeight uint64,
) ([]xchain.Attestation, error) {
	return p.fetch(ctx, sourceChainID, sourceHeight)
}

func (p Provider) LatestAttestation(ctx context.Context, sourceChainID uint64,
) (xchain.Attestation, bool, error) {
	return p.latest(ctx, sourceChainID)
}

func (p Provider) WindowCompare(ctx context.Context, sourceChainID uint64, height uint64) (int, error) {
	return p.window(ctx, sourceChainID, height)
}

// Subscribe implements cchain.Provider.
func (p Provider) Subscribe(in context.Context, srcChainID uint64, height uint64, workerName string,
	callback cchain.ProviderCallback,
) {
	srcChain := p.chainNames[srcChainID]
	ctx := log.WithCtx(in, "src_chain", srcChain)

	// Start a async goroutine to fetch attestations until ctx is canceled.
	go func() {
		backoff, reset := p.backoffFunc(ctx) // Note that backoff returns immediately on ctx cancel.

		for ctx.Err() == nil {
			// Fetch next batch of attestations.
			atts, err := p.fetch(ctx, srcChainID, height)
			if ctx.Err() != nil {
				return // Don't backoff or log on ctx cancel, just return.
			} else if err != nil {
				log.Warn(ctx, "Failed fetching attestation; will retry", err, "height", height)
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
				actx := log.WithCtx(ctx, "height", att.BlockHeight)
				// Sanity checks
				if att.SourceChainID != srcChainID {
					log.Error(actx, "Invalid attestation srcChain ID [BUG!]", nil)
					return
				} else if att.BlockHeight != height {
					log.Error(actx, "Invalid attestation height [BUG!]", nil)
					return
				}

				// Retry callback on error
				for actx.Err() == nil {
					err := callback(actx, att)
					if actx.Err() != nil {
						return // Don't backoff or log on ctx cancel, just return.
					} else if err != nil {
						log.Warn(actx, "Failed processing attestation; will retry", err)
						callbackErrTotal.WithLabelValues(workerName, srcChain).Inc()
						backoff()

						continue
					}
					streamHeight.WithLabelValues(workerName, srcChain).Set(float64(height))

					break // Success, stop retrying.
				}

				reset() // Reset callback backoff
				height++
			}
		}
	}()
}
