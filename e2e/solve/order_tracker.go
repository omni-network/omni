package solve

import (
	"sync"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
)

type orderTracker struct {
	mu     sync.Mutex
	orders map[solvernet.OrderID]TestOrder
	status map[solvernet.OrderID]solvernet.OrderStatus
}

func newOrderTracker() *orderTracker {
	return &orderTracker{
		orders: make(map[solvernet.OrderID]TestOrder),
		status: make(map[solvernet.OrderID]solvernet.OrderStatus),
	}
}

func (t *orderTracker) add(id solvernet.OrderID, order TestOrder) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.orders[id] = order
}

func (t *orderTracker) setStatus(id solvernet.OrderID, status solvernet.OrderStatus) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.status[id] = status
}

func (t *orderTracker) done() (bool, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for id, order := range t.orders {
		status, ok := t.status[id]
		if !ok {
			return false, nil
		}

		if order.ShouldReject {
			if status == solvernet.StatusFilled || status == solvernet.StatusClaimed {
				return false, errors.New("order should have been rejected", "id", id, "status", status)
			}

			if status != solvernet.StatusRejected {
				return false, nil
			}
		}

		if !order.ShouldReject {
			if status == solvernet.StatusRejected {
				return false, errors.New("order should have been filled", "id", id, "status", status)
			}

			if status != solvernet.StatusClaimed {
				return false, nil
			}
		}
	}

	return true, nil
}
