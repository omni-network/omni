package expbackoff

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

const defaultRetries = 3

type retryConfig struct {
	// Count is the max number of retries to attempt.
	Count int
	// Check is a function that returns true if the error should be retried.
	Check func(error) bool
}

func defaultRetryConfig() retryConfig {
	return retryConfig{
		Count: defaultRetries,
		Check: func(error) bool {
			return true
		},
	}
}

// WithRetryCount sets the number of retries to n.
// Note this is only applicable for use with Retry.
func WithRetryCount(n int) func(config *Config) {
	return func(c *Config) {
		c.retryConfig.Count = n
	}
}

// WithRetryCheck provides a custom error check function for retrying.
// Note this is only applicable for use with Retry.
func WithRetryCheck(check func(error) bool) func(config *Config) {
	return func(c *Config) {
		c.retryConfig.Check = check
	}
}

// Retry calls the provided function multiple times (default=3) with backoff until:
// - The function returns nil (Retry returns nil)
// - The context is canceled (Retry returns the context error)
// - The retry count is exhausted (Retry returns the last error).
// - The optional check function returns false (Retry returns the function error).
func Retry(ctx context.Context, fn func() error, opts ...func(*Config)) error {
	var remaining int // Workaround to extract retry config from options.
	var check func(error) bool
	opts = append(opts, func(c *Config) {
		remaining = c.retryConfig.Count
		check = c.retryConfig.Check
	})

	backoff := New(ctx, opts...)
	for {
		remaining--

		err := fn()
		if err == nil {
			return nil
		} else if !check(err) {
			return err
		} else if remaining <= 0 {
			return errors.Wrap(err, "max retries")
		} // else log error, backoff and retry

		log.Warn(ctx, "Will retry error after backoff", err, "remaining", remaining)
		backoff()

		if ctx.Err() != nil {
			return errors.Wrap(ctx.Err(), "retry timeout")
		}
	}
}
