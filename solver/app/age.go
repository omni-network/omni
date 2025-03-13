package app

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/umath"
)

const (
	maxAgeCache = 10000 // Max orders to track in age cache

	// statusDestFilled is an un-official status when an order was filled on the destination chain.
	// This contrasts with the official "filled" status which is emitted on the source chain after xmsg received.
	statusDestFilled = "dest_filled"
)

type instrDestFilledFunc func(ctx context.Context, dstChainID uint64, dstHeight uint64, order Order)

func newAgeCache(backends ethbackend.Backends) *ageCache {
	return &ageCache{
		createdAts: make(map[solvernet.OrderID]time.Time),
		backends:   backends,
	}
}

// ageCache enables best-effort tracking of order ages.
// Since on-chain state doesn't contain "created_height", a cache is used.
type ageCache struct {
	mu         sync.Mutex
	backends   ethbackend.Backends
	createdAts map[solvernet.OrderID]time.Time
}

func (a *ageCache) blockMeta(ctx context.Context, chainID uint64, height uint64) (string, time.Time, error) {
	backend, err := a.backends.Backend(chainID)
	if err != nil {
		return "", time.Time{}, err
	}

	header, err := backend.HeaderByNumber(ctx, umath.NewBigInt(height))
	if err != nil {
		return "", time.Time{}, err
	}
	timeI64, err := umath.ToInt64(header.Time)
	if err != nil {
		return "", time.Time{}, err
	}

	name, _ := backend.Chain()

	return name, time.Unix(timeI64, 0), nil
}

// InstrumentAge instruments the age of an order event.
func (a *ageCache) InstrumentAge(ctx context.Context, srcChainID uint64, srcHeight uint64, order Order) slog.Attr {
	a.mu.Lock()
	defer a.mu.Unlock()

	age, err := a.instrumentUnsafe(ctx, srcChainID, srcHeight, order.ID, order.Status.String())
	if err != nil {
		log.Warn(ctx, "Failed instrumenting order event (will retry)", err)
	}

	if a.maybePurge() {
		log.Warn(ctx, "Purged overflowing age cache", nil)
	}

	if age == 0 {
		return slog.Attr{}
	}

	return slog.Duration("age", age)
}

// InstrumentDestFilled instruments the age of an order filled on the destination chain.
func (a *ageCache) InstrumentDestFilled(ctx context.Context, dstChainID uint64, dstHeight uint64, order Order) {
	a.mu.Lock()
	defer a.mu.Unlock()

	_, err := a.instrumentUnsafe(ctx, dstChainID, dstHeight, order.ID, statusDestFilled)
	if err != nil {
		log.Warn(ctx, "Failed instrumenting order filled (will retry)", err)
	}
}

func (a *ageCache) instrumentUnsafe(ctx context.Context, chainID uint64, height uint64, orderID OrderID, status string) (time.Duration, error) {
	chainName, timestamp, err := a.blockMeta(ctx, chainID, height)
	if err != nil {
		return 0, err
	}

	if status == solvernet.StatusPending.String() {
		// Order is created in the block that emits pending status
		a.createdAts[orderID] = timestamp

		return 0, nil
	}

	createdAt, ok := a.createdAts[orderID]
	if !ok {
		// Pending event not seen or purged, best-effort ignore
		return 0, nil
	}

	age := timestamp.Sub(createdAt)
	orderAge.WithLabelValues(chainName, status).Observe(age.Seconds())

	// Remove from cache once final status is reached
	if status == solvernet.StatusRejected.String() ||
		status == solvernet.StatusClaimed.String() {
		delete(a.createdAts, orderID)
	}

	return age, nil
}

// maybePurge returns true if the cache was purged.
// This is required to prevent memory leaks.
func (a *ageCache) maybePurge() bool {
	if len(a.createdAts) < maxAgeCache {
		return false
	}

	a.createdAts = make(map[solvernet.OrderID]time.Time)

	return true
}
