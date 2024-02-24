package backend

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// NewWaiter returns a new Waiter.
func (b Backends) NewWaiter() *Waiter {
	return &Waiter{
		b:   b,
		txs: make(map[uint64][]*ethtypes.Transaction),
	}
}

// Waiter is a convenience struct to easily wait for multiple transactions to be mined.
type Waiter struct {
	b   Backends
	txs map[uint64][]*ethtypes.Transaction
}

func (w *Waiter) Add(chainID uint64, tx *ethtypes.Transaction) {
	w.txs[chainID] = append(w.txs[chainID], tx)
}

func (w *Waiter) Wait(ctx context.Context) error {
	for chainID, txs := range w.txs {
		for _, tx := range txs {
			rec, err := bind.WaitMined(ctx, w.b.backends[chainID], tx)
			if err != nil {
				return errors.Wrap(err, "wait mined", "chain_id", chainID)
			} else if rec.Status != ethtypes.ReceiptStatusSuccessful {
				return errors.New("tx status unsuccessful", "chain_id", chainID)
			}
		}
	}

	return nil
}
