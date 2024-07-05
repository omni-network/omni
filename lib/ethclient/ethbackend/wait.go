package ethbackend

import (
	"context"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// NewWaiter returns a new Waiter.
func (b Backends) NewWaiter() *Waiter {
	return &Waiter{
		b:       b,
		waiting: make(chan struct{}),
		txs:     make(map[uint64][]*ethtypes.Transaction),
	}
}

// Waiter is a convenience struct to easily wait for multiple transactions to be mined.
// Adding is thread safe, but it panics if Add is called after Wait.
type Waiter struct {
	b       Backends
	mu      sync.Mutex
	waiting chan struct{}
	txs     map[uint64][]*ethtypes.Transaction
}

func (w *Waiter) Add(chainID uint64, tx *ethtypes.Transaction) {
	timer := time.NewTicker(time.Millisecond)
	defer timer.Stop()
	var locked bool
	for !locked {
		select {
		case <-timer.C:
			locked = w.mu.TryLock()
		case <-w.waiting:
			panic("waiting for a transaction already in progress")
		}
	}
	defer w.mu.Unlock()

	w.txs[chainID] = append(w.txs[chainID], tx)
}

func (w *Waiter) Wait(ctx context.Context) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	close(w.waiting)

	for chainID, txs := range w.txs {
		for _, tx := range txs {
			_, err := w.b.backends[chainID].WaitMined(ctx, tx)
			if err != nil {
				return errors.Wrap(err, "wait mined", "chain_id", chainID)
			}
		}
	}

	return nil
}
