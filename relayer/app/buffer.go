package relayer

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"golang.org/x/sync/semaphore"
)

// activeBuffer links the output of cprovider/creator to the opsender.
// It has an large activeBuffer allowing many submissions to be queued up.
// It however limits the number of concurrent transactions it forwards to opsender
// to limiting our mempool size.
// If stops processing on any error.
type activeBuffer struct {
	chainName    string
	buffer       chan xchain.Submission
	mempoolLimit int64
	errChan      chan error
	sender       SendFunc
}

func newActiveBuffer(chainName string, mempoolLimit int64, size int, sender SendFunc) *activeBuffer {
	return &activeBuffer{
		chainName:    chainName,
		buffer:       make(chan xchain.Submission, size),
		mempoolLimit: mempoolLimit,
		errChan:      make(chan error, 1),
		sender:       sender,
	}
}

func (b *activeBuffer) AddInput(_ context.Context, submission xchain.Submission) error {
	select {
	case b.buffer <- submission:
	default:
		b.submitErr(errors.New("async send activeBuffer overflow"))
	}

	bufferLen.WithLabelValues(b.chainName).Set(float64(len(b.buffer)))

	return nil
}

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

			go func() {
				if err := b.sender(ctx, submission); err != nil {
					b.submitErr(err)
				}
				sema.Release(1)
				mempoolLen.WithLabelValues(b.chainName).Dec()
			}()
		}
	}
}

func (b *activeBuffer) submitErr(err error) {
	select {
	case b.errChan <- err:
	default:
	}
}
