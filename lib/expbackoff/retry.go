package expbackoff

import (
	"context"
	"fmt"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// RetryError allows immediate retry without warns.
type RetryError struct {
	Cause error
}

func (e RetryError) Error() string {
	return fmt.Sprintf("retry sentinel: %s", e.Cause)
}

const defaultRetries = 3

type retryConfig struct {
	// Label is used for logging.
	Label string
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

// WithRetryLabel sets the retry prefix label for logging.
// Note this is only applicable for use with Retry.
func WithRetryLabel(labelPrefix string) func(config *Config) {
	return func(c *Config) {
		c.retryConfig.Label = labelPrefix
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
	var cfg retryConfig // Workaround to extract config from opts
	opts = append(opts, func(c *Config) {
		cfg = c.retryConfig
	})

	backoff := New(ctx, opts...)
	remaining := cfg.Count
	prefix := cfg.Label
	if prefix != "" {
		prefix += ": "
	}

	for {
		remaining--

		err := fn()
		if err == nil {
			return nil
		} else if !cfg.Check(err) {
			return err
		} else if remaining <= 0 {
			return errors.Wrap(err, "max retries")
		}

		// If retry sentinel, retry immediately without warn logging or backoff.
		var rErr RetryError
		if errors.As(err, &rErr) {
			log.DebugErr(ctx, prefix+"Retrying error sentinel immediately", err, "remaining", remaining)
			continue
		}

		// else log error, backoff and retry

		log.Warn(ctx, prefix+"Will retry error after backoff", err, "remaining", remaining)
		backoff()

		if ctx.Err() != nil {
			return errors.Wrap(ctx.Err(), "retry timeout")
		}
	}
}
