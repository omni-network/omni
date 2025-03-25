// Package targets defines list of targets supported by Omni's v1 solver. Targets are
// restricted to reduce attack surface area, and keep order flow predictable.
// Targets restriction will be removed / lessened in future versions.
package targets

import (
	"context"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
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
		netconf.Mainnet: true,
	}

	// static are known, static targets.
	static = []Target{
		eigen,
		staking,
	}

	// dynamic is the list of dynamic targets, see RefreshForever.
	dynamic   []Target
	dynamicMu sync.RWMutex
)

// refreshOnce refreshes the dynamic targets once.
func refreshOnce(ctx context.Context) error {
	symbiotic, err := getSymbiotic(ctx)
	if err != nil {
		return errors.Wrap(err, "symbiotic target")
	}

	dynamicMu.Lock()
	defer dynamicMu.Unlock()
	dynamic = []Target{symbiotic}

	return nil
}

// RefreshForever refreshes dynamic targets forever.
// It blocks forever, only returning when the context is closed,
//
// Note that technically refreshing is only required on mainnet, but testing it
// on all networks is useful.
func RefreshForever(ctx context.Context) {
	ticker := time.NewTimer(0) // Immediately refresh on startup
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
		ticker.Reset(time.Hour) // Then refresh every hour

		if err := refreshOnce(ctx); err != nil {
			log.Warn(ctx, "Failed to refresh targets (will retry)", err)
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
	for _, t := range static {
		if _, ok := t.Addresses(chainID)[target]; ok {
			return t, true
		}
	}

	dynamicMu.RLock()
	defer dynamicMu.RUnlock()

	for _, t := range dynamic {
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
