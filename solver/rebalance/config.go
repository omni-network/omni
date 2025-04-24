package rebalance

import (
	"time"
)

type Config struct {
	// Interval at which to rebalance each chain
	Interval time.Duration
}

func DefaultConfig() Config {
	return Config{
		Interval: 30 * time.Minute,
	}
}
