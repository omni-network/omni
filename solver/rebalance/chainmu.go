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
	for _, chainID := range chainIDs {
		getMutex(chainID).Lock()
	}
}

// unlock releases the mutexes for the given chain IDs.
func unlock(chainIDs ...uint64) {
	for _, chainID := range chainIDs {
		getMutex(chainID).Unlock()
	}
}

func getMutex(chainID uint64) *sync.Mutex {
	mu.Lock()
	defer mu.Unlock()

	m, ok := chainMutexes[chainID]
	if !ok {
		m = &sync.Mutex{}
		chainMutexes[chainID] = m
	}

	return m
}
