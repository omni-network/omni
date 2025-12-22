//nolint:paralleltest // Parallel tests not supported since test-alias globals are used.
package expbackoff_test

import (
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/expbackoff"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	// Disable backoff for testing.
	expbackoff.SetAfterForT(t, func(d time.Duration) <-chan time.Time {
		ch := make(chan time.Time, 1)
		ch <- time.Now()

		return ch
	})

	ctx := t.Context()

	t.Run("default", func(t *testing.T) {
		var count int
		err := expbackoff.Retry(ctx, func() error {
			count++
			return io.EOF
		})
		require.ErrorIs(t, err, io.EOF)
		require.Equal(t, 3, count) // Default backoff count
	})

	t.Run("with count", func(t *testing.T) {
		const maxRetries = 5
		var count int
		err := expbackoff.Retry(ctx, func() error {
			count++
			return io.EOF
		}, expbackoff.WithRetryCount(maxRetries))
		require.ErrorIs(t, err, io.EOF)
		require.Equal(t, maxRetries, count) // Default backoff count
	})

	t.Run("context cancel", func(t *testing.T) {
		// Cancel the context
		ctx, cancel := context.WithCancel(ctx)
		cancel()

		var count int
		err := expbackoff.Retry(ctx, func() error {
			count++
			return io.EOF
		})
		require.ErrorIs(t, err, context.Canceled)
		require.Equal(t, 1, count) // No retries
	})

	t.Run("check", func(t *testing.T) {
		const expect = 2
		var count int
		check := func(err error) bool {
			return count < expect
		}
		err := expbackoff.Retry(ctx,
			func() error {
				count++
				return io.EOF
			},
			expbackoff.WithRetryCheck(check),
			expbackoff.WithRetryLabel("Test"),
		)
		require.ErrorIs(t, err, io.EOF)
		require.Equal(t, expect, count) // Not 3 retries
	})

	t.Run("retry error", func(t *testing.T) {
		const expect = 3
		var count int
		retryErr := expbackoff.RetryError{Cause: io.EOF}
		err := expbackoff.Retry(ctx, func() error {
			count++
			return retryErr
		})
		require.ErrorIs(t, err, retryErr)
		require.Equal(t, expect, count) // Not 3 retries
	})

	t.Run("custom retry error", func(t *testing.T) {
		const expect = 3
		var count int
		retryErr := customError{Cause: io.EOF}
		err := expbackoff.Retry(ctx, func() error {
			count++
			return retryErr
		})
		require.ErrorIs(t, err, retryErr)
		require.Equal(t, expect, count) // Not 3 retries
	})

}

type customError struct {
	Cause error
}

func (e customError) Error() string {
	return fmt.Sprintf("custom sentinel: %s", e.Cause)
}

func (e customError) As(target any) bool {
	if target == nil {
		return false
	}

	rErr, ok := target.(*expbackoff.RetryError)
	if !ok {
		return false
	}

	rErr.Cause = e.Cause

	return true
}
