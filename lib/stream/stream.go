// Package stream provide a generic stream function.
package stream

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

type Callback[E any] func(ctx context.Context, elem E) error

type Deps[E any] struct {
	// Dependency functions

	// FetchBatch fetches the next batch elements from the provided height (inclusive).
	// The elements must be sequential, since the internal height cursors is incremented for each element returned.
	FetchBatch func(ctx context.Context, chainID uint64, height uint64) ([]E, error)
	// Backoff returns a backoff and reset function. See expbackoff package for the implementation.
	Backoff func(ctx context.Context) (func(), func())
	// Verify is a sanity check function, it ensures each element is valid.
	Verify func(ctx context.Context, elem E, height uint64) error

	// Config
	ElemLabel     string
	RetryCallback bool

	// Metrics
	IncFetchErr     func()
	IncCallbackErr  func()
	SetStreamHeight func(uint64)
}

// Stream streams elements from the provided height (inclusive) on a specific chain.
// It fetches the next batch of elements from the current height, then
// calls the callback function for each, then repeats (forever).
//
// It retries forever on fetch errors. It can either retry or return callback errors.
// It returns (nil) when the context is canceled.
//
//nolint:nilerr // The function contract states it returns nil on context errors.
func Stream[E any](ctx context.Context, deps Deps[E], srcChainID uint64, height uint64, callback Callback[E]) error {
	backoff, reset := deps.Backoff(ctx) // Note that backoff returns immediately on ctx cancel.

	for ctx.Err() == nil {
		// Fetch next batch of attestations.
		batch, err := deps.FetchBatch(ctx, srcChainID, height)
		if ctx.Err() != nil {
			return nil // Don't backoff or log on ctx cancel, just return nil.
		} else if err != nil {
			log.Warn(ctx, "Failed fetching "+deps.ElemLabel+" (will retry)", err, "height", height)
			deps.IncFetchErr()
			backoff()

			continue
		} else if len(batch) == 0 {
			// We reached the head of the chain, wait for new blocks.
			backoff()

			continue
		}

		reset() // Reset fetch backoff

		// Call callback for each element
		for _, elem := range batch {
			ectx := log.WithCtx(ctx, "height", height)
			if err := deps.Verify(ectx, elem, height); err != nil {
				return errors.Wrap(err, "verify")
			}

			// Retry callback on error
			for ectx.Err() == nil {
				err := callback(ectx, elem)
				if ectx.Err() != nil {
					return nil // Don't backoff or log on ctx cancel, just return nil.
				} else if err != nil && !deps.RetryCallback {
					deps.IncCallbackErr()
					return errors.Wrap(err, "callback")
				} else if err != nil {
					log.Warn(ectx, "Failed processing "+deps.ElemLabel+" (will retry)", err)
					deps.IncCallbackErr()
					backoff()

					continue
				}

				break // Success, stop retrying.
			}
			reset() // Reset callback backoff

			deps.SetStreamHeight(height)
			height++
		}
	}

	return nil // Context canceled.
}
