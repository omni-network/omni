package rebalance

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestChainMutexes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		locks [][]uint64
	}{
		{
			name:  "single chain lock",
			locks: [][]uint64{{1}},
		},
		{
			name:  "multiple chains - separate locks",
			locks: [][]uint64{{1}, {2}},
		},
		{
			name:  "multiple chains - single locks",
			locks: [][]uint64{{1, 2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var wg sync.WaitGroup
			var counter int
			var counterMu sync.Mutex // counter mutex

			incr := func() {
				counterMu.Lock()
				defer counterMu.Unlock()
				counter++
			}

			// Lock/ unlock, increment counter
			for _, locks := range tt.locks {
				wg.Add(1)
				go func(locks []uint64) {
					defer wg.Done()

					// Acquire all locks
					lock(locks...)
					defer unlock(locks...)

					// Wait a bit
					time.Sleep(100 * time.Millisecond)

					incr()
				}(locks)
			}

			// Lock / unlock again, confirm no deadlock
			for _, locks := range tt.locks {
				wg.Add(1)
				go func(locks []uint64) {
					defer wg.Done()

					lock(locks...)
					defer unlock(locks...)
				}(locks)
			}

			wg.Wait()
			require.Equal(t, len(tt.locks), counter, "All goroutines should have executed")
		})
	}
}
