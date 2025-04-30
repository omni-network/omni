package rebalance

import (
	"time"
)

type Options func(*options)

// WithInterval sets the interval at which to rebalance the solver's balance.
func WithInterval(interval time.Duration) Options {
	return func(o *options) {
		o.interval = interval
	}
}

type options struct {
	// interval at which to rebalance the solver'balance.
	interval time.Duration
}

func defaultOps() options {
	return options{
		interval: 30 * time.Minute,
	}
}
