package cctp

import (
	"time"
)

type Option func(*options)

type options struct {
	mintInterval  time.Duration
	purgeInterval time.Duration
}

// WithMintInterval sets the cadence for the mint loop.
func WithMintInterval(interval time.Duration) Option {
	return func(c *options) {
		c.mintInterval = interval
	}
}

// WithPurgeInterval sets the cadence for the purge loop.
func WithPurgeInterval(interval time.Duration) Option {
	return func(c *options) {
		c.purgeInterval = interval
	}
}

func defaultOpts() *options {
	return &options{
		mintInterval:  30 * time.Second,
		purgeInterval: 20 * time.Minute,
	}
}
