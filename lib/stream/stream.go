// Package stream provide a generic stream function.
package stream

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"go.opentelemetry.io/otel/trace"
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
	IncFetchErr        func()
	IncCallbackErr     func()
	SetStreamHeight    func(uint64)
	SetCallbackLatency func(time.Duration)
	StartTrace         func(ctx context.Context, height uint64, spanName string) (context.Context, trace.Span)
}

// Stream streams elements from the provided height (inclusive) on a specific chain.
// It fetches the next batch of elements from the current height, then
// calls the callback function for each, then repeats (forever).
//
// It retries forever on fetch errors. It can either retry or return callback errors.
// It returns (nil) when the context is canceled.
//
//nolint:nilerr // The function contract states it returns nil on context errors.
func Stream[E any](rootCtx context.Context, deps Deps[E], srcChainID uint64, height uint64, callback Callback[E]) error {
	backoff, reset := deps.Backoff(rootCtx) // Note that backoff returns immediately on ctx cancel.

	for rootCtx.Err() == nil {
		// Fetch next batch of elements.
		fetchCtx, span := deps.StartTrace(rootCtx, height, "fetch")
		batch, err := deps.FetchBatch(fetchCtx, srcChainID, height)
		span.End()
		if rootCtx.Err() != nil {
			return nil // Don't backoff or log on ctx cancel, just return nil.
		} else if err != nil {
			log.Warn(rootCtx, "Failed fetching "+deps.ElemLabel+" (will retry)", err, "height", height)
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
		err = forEach(rootCtx, batch, func(ctx context.Context, elem E) error {
			ctx, span := deps.StartTrace(ctx, height, "callback")
			defer span.End()
			ctx = log.WithCtx(ctx, "height", height)

			if err := deps.Verify(ctx, elem, height); err != nil {
				return errors.Wrap(err, "verify")
			}

			// Retry callback on error
			for ctx.Err() == nil {
				t0 := time.Now()
				err := callback(ctx, elem)
				deps.SetCallbackLatency(time.Since(t0))
				if ctx.Err() != nil {
					return nil // Don't backoff or log on ctx cancel, just return nil.
				} else if err != nil && !deps.RetryCallback {
					deps.IncCallbackErr()
					return errors.Wrap(err, "callback")
				} else if err != nil {
					log.Warn(ctx, "Failed processing "+deps.ElemLabel+" (will retry)", err)
					deps.IncCallbackErr()
					backoff()

					continue
				}

				break // Success, stop retrying.
			}
			reset() // Reset callback backoff

			deps.SetStreamHeight(height)
			height++

			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil // Context canceled.
}

func forEach[E any](ctx context.Context, slice []E, fn func(context.Context, E) error) error {
	for _, elem := range slice {
		err := fn(ctx, elem)
		if err != nil {
			return err
		}
	}

	return nil
}
