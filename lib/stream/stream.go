// Package stream provide a generic stream function.
package stream

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"go.opentelemetry.io/otel/trace"
)

type Callback[E any] func(ctx context.Context, elem E) error

// Cache abstracts a simple element cache.
type Cache[E any] interface {
	// Get returns strictly sequential elements from the provided height (inclusive) or nil.
	Get(from uint64) []E
	// Set sets elements from the provided height (inclusive).
	// The elements must be strictly sequential.
	Set(from uint64, elems []E)
}

type Deps[E any] struct {
	// Dependency functions

	// FetchBatch fetches the next batch of elements from the provided height (inclusive).
	// The elements must be sequential, since the internal height cursors is incremented for each element returned.
	FetchBatch func(ctx context.Context, chainID uint64, height uint64) ([]E, error)
	// Backoff returns a backoff function. See expbackoff package for the implementation.
	Backoff func(ctx context.Context) func()
	// Verify is a sanity check function, it ensures each element is valid.
	Verify func(ctx context.Context, elem E, height uint64) error
	// Height returns the height of an element.
	Height func(elem E) uint64
	// Cache of elements.
	Cache Cache[E]

	// Config
	FetchWorkers  uint64
	ElemLabel     string
	HeightLabel   string
	RetryCallback bool

	// Metrics
	IncFetchErr        func()
	IncCallbackErr     func()
	SetStreamHeight    func(uint64)
	IncCacheHit        func()
	IncCacheMiss       func()
	SetCallbackLatency func(time.Duration)
	StartTrace         func(ctx context.Context, height uint64, spanName string) (context.Context, trace.Span)
}

// Stream streams elements from the provided height (inclusive) of a specific chain.
// It fetches the batches of elements from the current height, and
// calls the callback function for each element in strictly-sequential order.
//
// It supports concurrent fetching of single-element-batches only.
// It retries forever on fetch errors.
// It can either retry or return callback errors.
// It returns (nil) when the context is canceled.
func Stream[E any](ctx context.Context, deps Deps[E], srcChainID uint64, startHeight uint64, callback Callback[E]) error {
	if deps.FetchWorkers == 0 {
		return errors.New("invalid zero fetch worker count")
	}

	// Define a robust fetch function that fetches a batch of elements from a height (inclusive).
	// It only returns an empty list if the context is canceled.
	// It retries forever on error or if no elements found.
	fetchFunc := func(ctx context.Context, height uint64) []E {
		backoff := deps.Backoff(ctx) // Note that backoff returns immediately on ctx cancel.
		for {
			if ctx.Err() != nil {
				return nil
			}

			if elems := deps.Cache.Get(height); len(elems) > 0 {
				deps.IncCacheHit()
				return elems
			}

			fetchCtx, span := deps.StartTrace(ctx, height, "fetch")
			elems, err := deps.FetchBatch(fetchCtx, srcChainID, height)
			span.End()

			if ctx.Err() != nil {
				return nil
			} else if err != nil {
				log.Warn(ctx, "Failed fetching "+deps.ElemLabel+" (will retry)", err, deps.HeightLabel, height)
				deps.IncFetchErr()
				backoff()

				continue
			} else if len(elems) == 0 {
				// We reached the head of the chain, wait for new blocks.
				backoff()

				continue
			}

			deps.IncCacheMiss() // Only count non-empty fetches as cache misses.
			deps.Cache.Set(height, elems)

			return elems
		}
	}

	// Define a robust callback function that retries on error.
	callbackFunc := func(ctx context.Context, elem E) error {
		height := deps.Height(elem)
		ctx, span := deps.StartTrace(ctx, height, "callback")
		defer span.End()
		ctx = log.WithCtx(ctx, deps.HeightLabel, height)

		backoff := deps.Backoff(ctx)

		if err := deps.Verify(ctx, elem, height); err != nil {
			return errors.Wrap(err, "verify")
		}

		// Retry callback on error
		for {
			if ctx.Err() != nil {
				return nil // Don't backoff or log on ctx cancel, just return nil.
			}

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

			deps.SetStreamHeight(height)

			return nil
		}
	}

	// Sorting buffer connects the concurrent fetch workers to the callback
	sorter := newSortingBuffer(startHeight, deps, callbackFunc)

	// Ensure that fetch workers are stopped when streaming / processing is done.
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Start fetching workers
	startFetchWorkers(ctx, deps, fetchFunc, sorter, startHeight)

	// Sort fetch results and call callback
	return sorter.Process(ctx)
}

// startFetchWorkers starts <deps.FetchWorkers> worker goroutines
// that fetch all batches concurrently from the provided <startHeight>.
//
// Concurrent fetching is only supported for single-element-batches, since
// each worker fetches: startHeight + Ith-iteration + Nth-worker.
//
// For multi-element-batches, only a single worker is supported.
func startFetchWorkers[E any](
	ctx context.Context,
	deps Deps[E],
	work func(ctx context.Context, height uint64) []E,
	sorter *sortingBuffer[E],
	startHeight uint64,
) {
	for i := uint64(0); i < deps.FetchWorkers; i++ {
		go func(workerID uint64, height uint64) {
			for {
				// Work function MUST be robust, always returning a non-empty strictly-sequential batch
				// or nil if the context was canceled.
				batch := work(ctx, height)
				if ctx.Err() != nil {
					return
				} else if len(batch) == 0 {
					log.Error(ctx, "Work function returned an empty batch [BUG]", nil)
					return
				} else if len(batch) > 1 && deps.FetchWorkers > 1 {
					log.Error(ctx, "Concurrent fetching only supported for single element batches [BUG]", nil)
					return
				}

				var last uint64
				for i, e := range batch {
					last = deps.Height(e)
					if last != height+uint64(i) {
						log.Error(ctx, "Invalid batch [BUG]", nil)
						return
					}
				}

				sorter.Add(ctx, workerID, batch)

				// Calculate next height to fetch
				height = last + deps.FetchWorkers
			}
		}(i, startHeight+i) // Initialize a height to fetch per worker
	}
}

// sortingBuffer buffers unordered batches of elements (one batch per worker),
// providing elements to the callback in strictly-sequential sorted order.
type sortingBuffer[E any] struct {
	deps        Deps[E]
	callback    func(ctx context.Context, elem E) error
	startHeight uint64

	mu      sync.Mutex
	buffer  map[uint64]workerElem[E] // Worker elements by height
	counts  map[uint64]int           // Count of elements per worker
	signals map[uint64]chan struct{} // Processes <> Worker comms
}

const processorID = math.MaxUint64

func newSortingBuffer[E any](
	startHeight uint64,
	deps Deps[E],
	callback func(ctx context.Context, elem E) error,
) *sortingBuffer[E] {
	signals := make(map[uint64]chan struct{})
	signals[processorID] = make(chan struct{}, 1)
	for i := uint64(0); i < deps.FetchWorkers; i++ {
		signals[i] = make(chan struct{}, 1)
	}

	return &sortingBuffer[E]{
		startHeight: startHeight,
		deps:        deps,
		callback:    callback,
		buffer:      make(map[uint64]workerElem[E]),
		counts:      make(map[uint64]int),
		signals:     signals,
	}
}

// signal signals the ID to wakeup.
func (m *sortingBuffer[E]) signal(signalID uint64) {
	select {
	case m.signals[signalID] <- struct{}{}:
	default:
	}
}

// retryLock repeatedly obtains the lock and calls the callback while it returns false.
// It returns once the callback returns true or an error.
func (m *sortingBuffer[E]) retryLock(ctx context.Context, signalID uint64, fn func(ctx context.Context) (bool, error)) error {
	timer := time.NewTicker(time.Nanosecond) // Initial timer is instant
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-m.signals[signalID]:
		case <-timer.C:
		}

		m.mu.Lock()
		done, err := fn(ctx)
		m.mu.Unlock()
		if err != nil {
			return err
		} else if done {
			return nil
		}

		// Not done, so retry again, much later
		timer.Reset(time.Second)
	}
}

func (m *sortingBuffer[E]) Add(ctx context.Context, workerID uint64, batch []E) {
	_ = m.retryLock(ctx, workerID, func(_ context.Context) (bool, error) {
		// Wait for any previous batch this worker added to be processed before adding this batch.
		// This results in backpressure to workers, basically only buffering a single batch per worker.
		if m.counts[workerID] > 0 {
			return false, nil // Previous batch still in buffer, retry a bit later
		}

		// Add the batch
		for _, e := range batch {
			height := m.deps.Height(e)

			// Invariant check (error handling in workerFunc)
			if _, ok := m.buffer[height]; ok {
				return false, errors.New("duplicate element [BUG]")
			}

			m.buffer[height] = workerElem[E]{WorkerID: workerID, E: e}
		}

		m.counts[workerID] = len(batch)
		m.signal(processorID) // Signal the processor

		return true, nil // Don't retry lock again, we are done.
	})
}

// Process calls the callback function in strictly-sequential order from <height> (inclusive)
// as elements become available in the buffer.
func (m *sortingBuffer[E]) Process(ctx context.Context) error {
	next := m.startHeight
	return m.retryLock(ctx, processorID, func(ctx context.Context) (bool, error) {
		elem, ok := m.buffer[next]
		if !ok {
			return false, nil // Next height not in buffer, retry a bit later
		}
		delete(m.buffer, next)

		err := m.callback(ctx, elem.E)
		if err != nil {
			return false, err // Don't retry again
		}

		m.counts[elem.WorkerID]--

		if m.counts[elem.WorkerID] == 0 {
			m.signal(elem.WorkerID) // Signal the worker that it can add another batch
		}

		next++

		if _, ok := m.buffer[next]; ok {
			m.signal(processorID) // Signal ourselves if next elements already in buffer.
		}

		return false, nil // Retry again with next height
	})
}

// workerElem represents an element processed by a worker.
type workerElem[E any] struct {
	WorkerID uint64
	E        E
}
