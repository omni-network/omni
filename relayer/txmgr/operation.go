package txmgr

import (
	"context"
	"fmt"
	"time"

	"github.com/omni-network/omni/lib/errors"
)

// ErrFailedPermanently is an error raised by Do when the
// underlying Operation has been retried maxAttempts times.
type ErrFailedPermanently struct {
	attempts int
	LastErr  error
}

func (e *ErrFailedPermanently) Error() string {
	return fmt.Sprintf("operation failed permanently after %d attempts: %v", e.attempts, e.LastErr)
}

func (e *ErrFailedPermanently) Unwrap() error {
	return e.LastErr
}

// Strategy is used to calculate how long a particular Operation
// should wait between attempts.
type Strategy interface {
	// Duration returns how long to wait for a given retry attempt.
	Duration(attempt int) time.Duration
}

// Do perform the provided Operation up to maxAttempts times
// with delays in between each retry according to the provided
// Strategy.
func Do[T any](ctx context.Context, maxAttempts int, strategy Strategy, op func() (T, error)) (T, error) {
	var empty, ret T
	var err error
	if maxAttempts < 1 {
		return empty, fmt.Errorf("need at least 1 attempt to run op, but have %d max attempts", maxAttempts)
	}

	for i := 0; i < maxAttempts; i++ {
		if ctx.Err() != nil {
			return empty, errors.Wrap(ctx.Err(), "context canceled")
		}
		ret, err = op()
		if err == nil {
			return ret, nil
		}
		// Don't sleep when we are about to exit the loop & return ErrFailedPermanently
		if i != maxAttempts-1 {
			time.Sleep(strategy.Duration(i))
		}
	}
	return empty, &ErrFailedPermanently{
		attempts: maxAttempts,
		LastErr:  err,
	}
}

type FixedStrategy struct {
	Dur time.Duration
}

func (f *FixedStrategy) Duration(_ int) time.Duration {
	return f.Dur
}

func Fixed(dur time.Duration) Strategy {
	return &FixedStrategy{
		Dur: dur,
	}
}
