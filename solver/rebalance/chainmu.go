package rebalance

import (
	"sync"
)

var mutexes sync.Map // map[chainID]*sync.Mutex

// lock acquires the mutexes for the given chain IDs, it returns an unlock function.
// Usage:
//
//	defer lock(evmchain.IDEthereum, evmchain.IDHyperEVM)()
func lock(chainIDs ...uint64) func() {
	var unlocks []func()
	for _, chainID := range chainIDs {
		m, _ := mutexes.LoadOrStore(chainID, new(sync.Mutex))
		mu := m.(*sync.Mutex) //nolint:revive,forcetypeassert // Known type
		mu.Lock()
		unlocks = append(unlocks, mu.Unlock)
	}

	return func() {
		for _, unlock := range unlocks {
			unlock()
		}
	}
}
