package fireblocks

import (
	"time"

	"github.com/omni-network/omni/lib/errors"
)

// Config houses parameters for altering the behavior of a SimpleTxManager.
type Config struct {

	// NetworkTimeout is the allowed duration for a single network request.
	// This is intended to be used for network requests that can be replayed.
	NetworkTimeout time.Duration

	// QueryInterval is the interval at which the FireBlocks client will
	// call the get transaction by id to check for confirmations after a txn
	// has been sent
	QueryInterval time.Duration

	// LogFreqFactor is the frequency at which the FireBlocks client will
	// log a warning message if the transaction has not been signed yet
	LogFreqFactor int
}

func NewConfig(networkTimeout time.Duration, queryInterval time.Duration) (Config, error) {
	cfg := Config{
		NetworkTimeout: networkTimeout,
		QueryInterval:  queryInterval,
	}
	err := cfg.Check()
	if err != nil {
		return Config{}, errors.New("invalid config", err)
	}

	return cfg, nil
}

func (m Config) Check() error {
	if m.LogFreqFactor <= 0 {
		return errors.New("must provide LogFreqFactor")
	}

	if m.NetworkTimeout <= 0 {
		return errors.New("must provide NetworkTimeout")
	}

	if m.QueryInterval <= 0 {
		return errors.New("must provide QueryInterval")
	}

	return nil
}
