//nolint:paralleltest // Parallel tests not supported since test-alias globals are used.
package expbackoff_test

import (
	"context"
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
		err := expbackoff.Retry(ctx, func() error {
			count++
			return io.EOF
		}, expbackoff.WithRetryCheck(check))
		require.ErrorIs(t, err, io.EOF)
		require.Equal(t, expect, count) // Not 3 retries
	})
}
