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

	// Network is the environment that we have deployed in, either testnet or mainnet
	Network Environment
}

type Environment int

const (
	TestNet Environment = iota + 1 // EnumIndex = 1
	MainNet                        // EnumIndex = 2
)

// String - Creating common behavior - give the type a String function.
func (e Environment) String() string {
	return [...]string{"testnet", "mainnet"}[e-1]
}

// DefaultConfig returns a Config with default values.
func DefaultConfig() Config {
	cfg, err := NewConfig(
		time.Duration(30)*time.Second,
		time.Duration(2)*time.Second,
		2,
		// TODO: Change this to MainNet in the future or ensure calls to NewClientWithConfig
		TestNet,
	)
	if err != nil {
		panic("invalid default config")
	}

	return cfg
}

// NewConfig returns a new Config with the given parameters.
func NewConfig(networkTimeout time.Duration, queryInterval time.Duration, logFreqFactor int, env Environment) (Config, error) {
	cfg := Config{
		NetworkTimeout: networkTimeout,
		QueryInterval:  queryInterval,
		LogFreqFactor:  logFreqFactor,
		Network:        env,
	}
	err := cfg.check()
	if err != nil {
		return Config{}, errors.New("invalid config", err)
	}

	return cfg, nil
}

// check validates the Config.
func (m Config) check() error {
	if m.LogFreqFactor <= 0 {
		return errors.New("must provide LogFreqFactor")
	}

	if m.NetworkTimeout <= 0 {
		return errors.New("must provide NetworkTimeout")
	}

	if m.QueryInterval <= 0 {
		return errors.New("must provide QueryInterval")
	}

	if m.Network != TestNet && m.Network != MainNet {
		return errors.New("must provide valid Network")
	}

	return nil
}
