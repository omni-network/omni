package relayer

import (
	"context"

	"github.com/omni-network/omni/lib/chaos"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"golang.org/x/sync/semaphore"
)

// activeBuffer links the output of each worker's cprovider/creators (one per chain version)
// to the destination chain async sender. Fan-in buffer.
//
// It limits the number of concurrent transactions it forwards to the async sender
// to limit the mempool size.
//
// While the mempool limit is reached, calls to AddInput block.
//
// If stops processing on any error.
type activeBuffer struct {
	chainName    string
	buffer       chan xchain.Submission
	mempoolLimit int64
	errChan      chan error
	sendAsync    SendAsync
}

func newActiveBuffer(chainName string, mempoolLimit int64, sendAsync SendAsync) *activeBuffer {
	return &activeBuffer{
		chainName:    chainName,
		buffer:       make(chan xchain.Submission),
		mempoolLimit: mempoolLimit,
		errChan:      make(chan error, 1),
		sendAsync:    sendAsync,
	}
}

// AddInput adds a new submission to the buffer. It blocks while mempoolLimit is reached.
func (b *activeBuffer) AddInput(ctx context.Context, submission xchain.Submission) error {
	select {
	case <-ctx.Done():
		b.submitErr(errors.Wrap(ctx.Err(), "context canceled"))
	case b.buffer <- submission: // Unbuffered, will block until a reader is ready. We don't want to restart the worker.
	}

	bufferLen.WithLabelValues(b.chainName).Set(float64(len(b.buffer)))

	return nil
}

// Run processes the buffer, sending submissions to the async sender.
func (b *activeBuffer) Run(ctx context.Context) error {
	sema := semaphore.NewWeighted(b.mempoolLimit)
	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context canceled")
		case err := <-b.errChan:
			return err
		case submission := <-b.buffer:
			if err := sema.Acquire(ctx, 1); err != nil {
				return errors.Wrap(err, "acquire semaphore")
			}
			mempoolLen.WithLabelValues(b.chainName).Inc()

			// Trigger async send synchronously (for ordered nonces), but wait for response async.
			response := b.sendAsync(ctx, submission)
			go func() {
				err := <-response
				if err != nil {
					b.submitErr(errors.Wrap(err, "send submission"))
				}

				mempoolLen.WithLabelValues(b.chainName).Dec()
				sema.Release(1)
			}()

			// Chaos test this worker with random errors.
			if err := chaos.MaybeError(ctx); err != nil {
				return err
			}
		}
	}
}

func (b *activeBuffer) submitErr(err error) {
	select {
	case b.errChan <- err:
	default:
	}
}
