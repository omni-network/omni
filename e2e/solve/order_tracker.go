package solve

import (
	"sync"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
)

type orderKey struct {
	id         solvernet.OrderID
	srcChainID uint64
}

type orderTracker struct {
	mu     sync.Mutex
	orders map[orderKey]TestOrder
	status map[orderKey]solvernet.OrderStatus
}

func newOrderTracker() *orderTracker {
	return &orderTracker{
		orders: make(map[orderKey]TestOrder),
		status: make(map[orderKey]solvernet.OrderStatus),
	}
}

func (t *orderTracker) add(id solvernet.OrderID, order TestOrder) {
	t.mu.Lock()
	defer t.mu.Unlock()
	key := orderKey{id: id, srcChainID: order.SourceChainID}
	t.orders[key] = order
}

func (t *orderTracker) setStatus(id solvernet.OrderID, srcChainID uint64, status solvernet.OrderStatus) {
	t.mu.Lock()
	defer t.mu.Unlock()
	key := orderKey{id: id, srcChainID: srcChainID}
	t.status[key] = status
}

func (t *orderTracker) done() (bool, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for key, order := range t.orders {
		status, ok := t.status[key]
		if !ok {
			return false, nil
		}

		if order.ShouldReject {
			if status == solvernet.StatusFilled || status == solvernet.StatusClaimed {
				return false, errors.New("order should have been rejected", "id", key.id, "src_chain_id", key.srcChainID, "status", status)
			}

			if status != solvernet.StatusRejected {
				return false, nil
			}
		}

		if !order.ShouldReject {
			if status == solvernet.StatusRejected {
				return false, errors.New("order should have been filled", "id", key.id, "src_chain_id", key.srcChainID, "status", status)
			}

			if status != solvernet.StatusClaimed {
				return false, nil
			}
		}
	}

	return true, nil
}
