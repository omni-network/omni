package solve

import (
	"sync"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
)

type orderKey struct {
	orderID    solvernet.OrderID
	srcChainID uint64
}

type orderTracker struct {
	mu      sync.Mutex
	orders  map[orderKey]TestOrder
	status  map[orderKey]solvernet.OrderStatus
	tracked bool // tracked indicates when orders have been tracked. This prevents accidental misuse of the orderTracker API.
}

func newOrderTracker() *orderTracker {
	return &orderTracker{
		orders: make(map[orderKey]TestOrder),
		status: make(map[orderKey]solvernet.OrderStatus),
	}
}

func (t *orderTracker) TrackOrder(id solvernet.OrderID, order TestOrder) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.tracked {
		panic("all orders already tracked [BUG]")
	}

	key := orderKey{orderID: id, srcChainID: order.SourceChainID}
	t.orders[key] = order
}

func (t *orderTracker) AllTracked() {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.tracked {
		panic("all orders already tracked [BUG]")
	}
	t.tracked = true
}

func (t *orderTracker) UpdateStatus(id solvernet.OrderID, srcChainID uint64, status solvernet.OrderStatus) {
	t.mu.Lock()
	defer t.mu.Unlock()
	key := orderKey{orderID: id, srcChainID: srcChainID}
	t.status[key] = status
}

func (t *orderTracker) Done() (bool, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.tracked {
		return false, errors.New("not all orders tracked [BUG]")
	}

	for key, order := range t.orders {
		status, ok := t.status[key]
		if !ok {
			return false, nil
		}

		if order.ShouldReject {
			if status == solvernet.StatusFilled || status == solvernet.StatusClaimed {
				return false, errors.New("order should have been rejected", "order_id", key.orderID, "src_chain_id", key.srcChainID, "status", status)
			}

			if status != solvernet.StatusRejected {
				return false, nil
			}
		}

		if !order.ShouldReject {
			if status == solvernet.StatusRejected {
				return false, errors.New("order should have been filled", "order_id", key.orderID, "src_chain_id", key.srcChainID, "status", status)
			}

			if status != solvernet.StatusClaimed {
				return false, nil
			}
		}
	}

	return true, nil
}
