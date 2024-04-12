// Package expbackoff implements exponential backoff.
//
// It was copied from the GPL version of Obol which was
// originally copied from google.golang.org/grpc.
//
// See:
// - https://github.com/grpc/grpc-go/tree/master/backoff
// - https://github.com/ObolNetwork/charon/tree/v0.14.0/app/expbackoff
//
//nolint:gochecknoglobals // Default config and test-alias globals are ok.
package expbackoff

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

// Config defines the configuration options for backoff.
type Config struct {
	// BaseDelay is the amount of time to backoff after the first failure.
	BaseDelay time.Duration
	// Multiplier is the factor with which to multiply backoffs after a
	// failed retry. Should ideally be greater than 1.
	Multiplier float64
	// Jitter is the factor with which backoffs are randomized.
	Jitter float64
	// MaxDelay is the upper bound of backoff delay.
	MaxDelay time.Duration
}

// DefaultConfig is a backoff configuration with the default values specified
// at https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md.
//
// This should be useful for callers who want to configure backoff with
// non-default values only for a subset of the options.
//
// Copied from google.golang.org/grpc@v1.48.0/backoff/backoff.go.
var DefaultConfig = Config{
	BaseDelay:  1.0 * time.Second,
	Multiplier: 1.6,
	Jitter:     0.2,
	MaxDelay:   120 * time.Second,
}

// FastConfig is a common configuration for fast backoff.
var FastConfig = Config{
	BaseDelay:  100 * time.Millisecond,
	Multiplier: 1.6,
	Jitter:     0.2,
	MaxDelay:   5 * time.Second,
}

// WithPeriodicConfig configures the backoff with periodic backoff.
func WithPeriodicConfig(period time.Duration) func(*Config) {
	return func(config *Config) {
		config.BaseDelay = period
		config.Multiplier = 1
	}
}

// WithFastConfig configures the backoff with FastConfig.
func WithFastConfig() func(*Config) {
	return func(config *Config) {
		*config = FastConfig
	}
}

// With configures the backoff with the provided config.
func With(c Config) func(*Config) {
	return func(config *Config) {
		*config = c
	}
}

// New returns a backoff function configured via functional options applied to DefaultConfig.
// The backoff function will exponentially sleep longer each time it is called.
// The backoff function returns immediately after the context is canceled.
//
// Usage:
//
//	backoff := expbackoff.New(ctx)
//	for ctx.Err() == nil {
//	  resp, err := doThing(ctx)
//	  if err != nil {
//	    backoff()
//	    continue
//	  } else {
//	    return resp
//	  }
//	}
//
//nolint:nonamedreturns // Named returns used for clear code.
func New(ctx context.Context, opts ...func(*Config)) (backoff func()) {
	backoff, _ = NewWithReset(ctx, opts...)
	return backoff
}

// NewWithReset returns a backoff and a reset function configured via functional options applied to DefaultConfig.
// The backoff function will exponentially sleep longer each time it is called.
// Calling the reset function will reset the backoff sleep duration to Config.BaseDelay.
// The backoff function returns immediately after the context is canceled.
//
// Usage:
//
//	backoff, reset := expbackoff.NewWithReset(ctx)
//	for ctx.Err() == nil {
//	  resp, err := doThing(ctx)
//	  if err != nil {
//	    backoff()
//	    continue
//	  } else {
//	    reset()
//	    // Do something with the response.
//	  }
//	}
//
//nolint:nonamedreturns // Named returns used for clear code.
func NewWithReset(ctx context.Context, opts ...func(*Config)) (backoff func(), reset func()) {
	conf := DefaultConfig
	for _, opt := range opts {
		opt(&conf)
	}

	var retries int

	backoff = func() {
		if ctx.Err() != nil {
			return
		}

		select {
		case <-ctx.Done():
		case <-after(Backoff(conf, retries)):
		}
		retries++
	}

	reset = func() {
		retries = 0
	}

	return backoff, reset
}

// NewWithAutoReset returns a backoff function configured via functional options applied to DefaultConfig.
// The backoff function will exponentially sleep longer each time it is called.
// The backoff function is automatically reset if sufficient time has passed since the last backoff.
//
// This "sufficient delay" is the next backoff duration, so if the next backoff duration is 1s,
// the backoff will reset to initial duration after 1s of not being called.
//
//nolint:nonamedreturns // Named returns used for clear code.
func NewWithAutoReset(ctx context.Context, opts ...func(*Config)) (backoff func()) {
	conf := DefaultConfig
	for _, opt := range opts {
		opt(&conf)
	}

	var retries int

	lastBackoff := time.Now()
	backoff = func() {
		if ctx.Err() != nil {
			return
		}

		autoResetAt := lastBackoff.Add(Backoff(conf, retries))
		if time.Now().After(autoResetAt) {
			retries = 0
		}

		select {
		case <-ctx.Done():
		case <-after(Backoff(conf, retries)):
		}
		retries++
		lastBackoff = time.Now()
	}

	return backoff
}

// Backoff returns the amount of time to wait before the next retry given the
// number of retries.
// Copied from google.golang.org/grpc@v1.48.0/internal/backoff/backoff.go.
func Backoff(config Config, retries int) time.Duration {
	if retries == 0 {
		return config.BaseDelay
	}

	backoff := float64(config.BaseDelay)
	max := float64(config.MaxDelay)

	for backoff < max && retries > 0 {
		backoff *= config.Multiplier
		retries--
	}
	if backoff > max {
		backoff = max
	}
	// Randomize backoff delays so that if a cluster of requests start at
	// the same time, they won't operate in lockstep.
	backoff *= 1 + config.Jitter*(randFloat()*2-1)
	if backoff < 0 {
		return 0
	}

	return time.Duration(backoff)
}

// after is aliased for testing.
var after = time.After

// SetAfterForT sets the after internal function for testing.
func SetAfterForT(t *testing.T, fn func(d time.Duration) <-chan time.Time) {
	t.Helper()
	cached := after
	after = fn
	t.Cleanup(func() {
		after = cached
	})
}

// randFloat is aliased for testing.
var randFloat = rand.Float64

// SetRandFloatForT sets the random float internal function for testing.
func SetRandFloatForT(t *testing.T, fn func() float64) {
	t.Helper()
	cached := randFloat
	randFloat = fn
	t.Cleanup(func() {
		randFloat = cached
	})
}
