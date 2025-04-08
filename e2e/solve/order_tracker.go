package solve

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
)

type orderKey struct {
	orderID    solvernet.OrderID
	srcChainID uint64
}

func (k orderKey) SrcChain() string {
	return evmchain.Name(k.srcChainID)
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

func (t *orderTracker) Len() int {
	t.mu.Lock()
	defer t.mu.Unlock()

	return len(t.orders)
}

func (t *orderTracker) Remaining() (int, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.tracked {
		return 0, errors.New("not all orders tracked [BUG]")
	}

	var remaining int
	for key, order := range t.orders {
		status, ok := t.status[key]
		if !ok {
			remaining++
			continue
		}

		if order.ShouldReject {
			if status == solvernet.StatusFilled || status == solvernet.StatusClaimed {
				return 0, errors.New("order should have been rejected", "order_id", key.orderID, "src_chain_id", key.srcChainID, "status", status)
			}

			if status != solvernet.StatusRejected {
				remaining++
				continue
			}
		}

		if !order.ShouldReject {
			if status == solvernet.StatusRejected {
				return 0, errors.New("order should have been filled", "order_id", key.orderID, "src_chain_id", key.srcChainID, "status", status)
			}

			if status != solvernet.StatusClaimed {
				remaining++
				continue
			}
		}
	}

	return remaining, nil
}

func (t *orderTracker) DebugFirstPending(ctx context.Context) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for key := range t.orders {
		status, ok := t.status[key]
		if !ok || status == solvernet.StatusPending {
			log.Debug(ctx, "First pending order", "order_id", key.orderID, "src_chain", key.SrcChain(), "status", status)
			return
		}
	}
}
