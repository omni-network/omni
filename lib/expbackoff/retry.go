package expbackoff

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
)

const defaultRetries = 3

// WithRetryCount sets the number of retries to n.
// Note this is only applicable for use with Retry.
func WithRetryCount(n int) func(config *Config) {
	return func(c *Config) {
		c.retryCount = n
	}
}

// Retry calls the provided function multiple times (default=3) with backoff until:
// - The function returns nil (Retry returns nil)
// - The context is canceled (Retry returns the context error)
// - The retry count is exhausted (Retry returns the last error).
func Retry(ctx context.Context, fn func() error, opts ...func(*Config)) error {
	var remaining int // Workaround to extract retry count from options.
	opts = append(opts, func(c *Config) {
		remaining = c.retryCount
	})

	backoff := New(ctx, opts...)
	for {
		remaining--

		err := fn()
		if err == nil {
			return nil
		} else if remaining <= 0 {
			return errors.Wrap(err, "max retries")
		}

		backoff()

		if ctx.Err() != nil {
			return errors.Wrap(ctx.Err(), "retry timeout")
		}
	}
}
