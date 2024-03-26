package fireblocks

import (
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

// options houses parameters for altering the behavior of a SimpleTxManager.
type options struct {
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

	// SignNote is a note to include in the sign request
	SignNote string

	// Host is the base URL for the FireBlocks API.
	Host string

	// TestAccounts overrides dynamic account
	TestAccounts map[common.Address]uint64
}

// defaultOptions returns a options with default values.
func defaultOptions() options {
	return options{
		NetworkTimeout: time.Duration(30) * time.Second,
		QueryInterval:  time.Second,
		LogFreqFactor:  10,
		Host:           hostProd,
		TestAccounts:   make(map[common.Address]uint64),
		SignNote:       "omni sign note not set",
	}
}

func WithQueryInterval(interval time.Duration) func(*options) {
	return func(cfg *options) {
		cfg.QueryInterval = interval
	}
}

func WithLogFreqFactor(factor int) func(*options) {
	return func(cfg *options) {
		cfg.LogFreqFactor = factor
	}
}

func WithTestAccount(addr common.Address, accID uint64) func(*options) {
	return func(cfg *options) {
		cfg.TestAccounts[addr] = accID
	}
}

func WithHost(host string) func(*options) {
	return func(cfg *options) {
		cfg.Host = host
	}
}

func WithSignNote(note string) func(*options) {
	return func(cfg *options) {
		cfg.SignNote = note
	}
}

// check validates the options.
func (c options) check() error {
	if c.LogFreqFactor <= 0 {
		return errors.New("must provide LogFreqFactor")
	}

	if c.NetworkTimeout <= 0 {
		return errors.New("must provide NetworkTimeout")
	}

	if c.QueryInterval <= 0 {
		return errors.New("must provide QueryInterval")
	}

	return nil
}
