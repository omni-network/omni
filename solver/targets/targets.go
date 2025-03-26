// Package targets defines list of targets supported by Omni's v1 solver. Targets are
// restricted to reduce attack surface area, and keep order flow predictable.
// Targets restriction will be removed / lessened in future versions.
package targets

import (
	"context"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

type Target struct {
	Name      string
	Addresses func(chainID uint64) map[common.Address]bool
}

var (
	// targetsRestricted maps each network to whether targets should be restricted to the allowed set.
	targetsRestricted = map[netconf.ID]bool{
		netconf.Staging: true,
		netconf.Omega:   true,
		netconf.Mainnet: true,
	}

	// static are known, static targets.
	static = []Target{
		eigen,
		staking,
	}

	// targets is the list of targets all targets, set during Init.
	targets []Target
	mu      sync.RWMutex
)

// InitStatic initializes onlystatic targets.
func InitStatic() {
	targets = []Target{}
	targets = append(targets, static...)
}

// Init initializes the targets.
func Init(ctx context.Context) error {
	symbiotic, err := getSymbiotic(ctx)
	if err != nil {
		return errors.Wrap(err, "symbiotic target")
	}

	mu.Lock()
	defer mu.Unlock()

	targets = []Target{}
	targets = append(targets, static...)
	targets = append(targets, symbiotic)

	return nil
}

// TryInitRefreshForever tries to initialize the targets forever, with exponential backoff.
// Once initialized, targets are refreshed every hour.
func TryInitRefreshForever(ctx context.Context) {
	backoff := expbackoff.New(ctx)
	for ctx.Err() == nil {
		if ctx.Err() != nil {
			return
		}

		err := Init(ctx)
		if err == nil {
			log.Info(ctx, "Targets initialized")
			go refreshForever(ctx)

			return
		}

		log.Warn(ctx, "Failed to init targets, will retry", err)
		backoff()
	}
}

// refreshForever refreshes the targets every hour.
func refreshForever(ctx context.Context) {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := Init(ctx)
			if err != nil {
				log.Warn(ctx, "Failed to refresh targets, will retry", err)
			}

			log.Info(ctx, "Targets refreshed")
		}
	}
}

func networkChainAddrs(m map[uint64]map[common.Address]bool) func(uint64) map[common.Address]bool {
	return func(chainID uint64) map[common.Address]bool {
		return m[chainID]
	}
}

// IsRestricted returns true if the given network restricts targets.
func IsRestricted(network netconf.ID) bool {
	return targetsRestricted[network]
}

// Get returns the allowed target for the given chain and address.
func Get(chainID uint64, target common.Address) (Target, bool) {
	mu.RLock()
	defer mu.RUnlock()

	for _, t := range targets {
		if _, ok := t.Addresses(chainID)[target]; ok {
			return t, true
		}
	}

	return Target{}, false
}

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}

func set(addrs ...common.Address) map[common.Address]bool {
	s := make(map[common.Address]bool)
	for _, addr := range addrs {
		s[addr] = true
	}

	return s
}
