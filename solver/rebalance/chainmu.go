package rebalance

import (
	"sync"
)

var (
	// chainMutexes provides mutex locking per chain ID.
	chainMutexes = make(map[uint64]*sync.Mutex)

	// mu protects access to the chainMutexes map.
	mu sync.Mutex
)

// lock acquires the mutexes for the given chain IDs.
func lock(chainIDs ...uint64) {
	mu.Lock()
	defer mu.Unlock()

	for _, chainID := range chainIDs {
		mu, ok := chainMutexes[chainID]
		if !ok {
			mu = &sync.Mutex{}
			chainMutexes[chainID] = mu
		}
		mu.Lock()
	}
}

// unlock releases the mutexes for the given chain IDs.
func unlock(chainIDs ...uint64) {
	mu.Lock()
	defer mu.Unlock()

	for _, chainID := range chainIDs {
		if mu, ok := chainMutexes[chainID]; ok {
			mu.Unlock()
		}
	}
}
