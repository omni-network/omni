package app

import (
	"context"
	"log/slog"
	"maps"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/svmutil"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/unibackend"

	"github.com/gagliardetto/solana-go/rpc"
)

const (
	maxAgeCache = 10000 // Max orders to track in age cache

	// statusDestFilled is an un-official status when an order was filled on the destination chain.
	// This contrasts with the official "filled" status which is emitted on the source chain after xmsg received.
	statusDestFilled = "dest_filled"
)

type cacheVal struct {
	CreatedAt  time.Time
	SrcChainID uint64
}

type destFilledAge func(ctx context.Context, dstChainID uint64, dstHeight uint64, order Order) slog.Attr

func newAgeCache(backends unibackend.Backends) *ageCache {
	return &ageCache{
		createdAts: make(map[solvernet.OrderID]cacheVal),
		backends:   backends,
	}
}

// ageCache enables best-effort tracking of order ages.
// Since on-chain state doesn't contain "created_height", a cache is used.
type ageCache struct {
	mu         sync.Mutex
	backends   unibackend.Backends
	createdAts map[solvernet.OrderID]cacheVal
}

func (a *ageCache) blockTime(ctx context.Context, chainID uint64, height uint64) (time.Time, error) {
	backend, err := a.backends.Backend(chainID)
	if err != nil {
		return time.Time{}, err
	}

	if backend.IsSVM() {
		block, ok, err := svmutil.GetBlock(ctx, backend.SVMClient(), height, rpc.TransactionDetailsNone)
		if err != nil {
			return time.Time{}, svmutil.WrapRPCError(err, "getBlock")
		} else if !ok {
			return time.Time{}, errors.New("block not found", "height", height)
		}

		return block.BlockTime.Time(), nil
	}

	header, err := backend.EVMBackend().HeaderByNumber(ctx, bi.N(height))
	if err != nil {
		return time.Time{}, err
	}
	timeI64, err := umath.ToInt64(header.Time)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(timeI64, 0), nil
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

	return age
}

// InstrumentDestFilled instruments the age of an order filled on the destination chain.
func (a *ageCache) InstrumentDestFilled(ctx context.Context, dstChainID uint64, dstHeight uint64, order Order) slog.Attr {
	a.mu.Lock()
	defer a.mu.Unlock()

	age, err := a.instrumentUnsafe(ctx, dstChainID, dstHeight, order.ID, statusDestFilled)
	if err != nil {
		log.Warn(ctx, "Failed instrumenting order filled (will retry)", err)
		return slog.Attr{}
	}

	return age
}

func (a *ageCache) instrumentUnsafe(ctx context.Context, chainID uint64, height uint64, orderID OrderID, status string) (slog.Attr, error) {
	timestamp, err := a.blockTime(ctx, chainID, height)
	if err != nil {
		return slog.Attr{}, err
	}

	if status == solvernet.StatusPending.String() {
		// Order is created in the block that emits pending status
		a.createdAts[orderID] = cacheVal{
			CreatedAt:  timestamp,
			SrcChainID: chainID,
		}

		return slog.Attr{}, nil
	}

	v, ok := a.createdAts[orderID]
	if !ok {
		// Pending event not seen or purged, best-effort ignore
		return slog.Attr{}, nil
	}

	age := timestamp.Sub(v.CreatedAt)
	orderAge.WithLabelValues(evmchain.Name(chainID), status).Observe(age.Seconds())

	// Remove from cache once final status is reached
	if status == solvernet.StatusRejected.String() ||
		status == solvernet.StatusClaimed.String() {
		delete(a.createdAts, orderID)
	}

	return slog.Float64("age_s", age.Seconds()), nil
}

// maybePurge returns true if the cache was purged.
// This is required to prevent memory leaks.
func (a *ageCache) maybePurge() bool {
	if len(a.createdAts) < maxAgeCache {
		return false
	}

	a.createdAts = make(map[solvernet.OrderID]cacheVal)

	return true
}

func (a *ageCache) Clone() map[solvernet.OrderID]cacheVal {
	a.mu.Lock()
	defer a.mu.Unlock()

	return maps.Clone(a.createdAts)
}

// monitorAgeCacheForever monitors the age cache instrumenting the oldest order per chain.
func monitorAgeCacheForever(ctx context.Context, network netconf.Network, cache *ageCache) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	const tooOld = time.Hour

	// OrderAge is tuple containing order and age
	type OrderAge struct {
		OrderID solvernet.OrderID
		Age     time.Duration
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Reset oldest order for all chains (to remove stale data)
			for _, chain := range network.EVMChains() {
				oldestOrder.Reset(chain.Name)
			}

			// Calculate oldest order per chain
			oldest := make(map[string]OrderAge)
			for orderID, v := range cache.Clone() {
				chain := network.ChainName(v.SrcChainID)
				age := time.Since(v.CreatedAt)
				if oldest[chain].Age < age {
					oldest[chain] = OrderAge{
						OrderID: orderID,
						Age:     age,
					}
				}
			}

			// Instrument
			for chain, o := range oldest {
				oldestOrder.WithLabelValues(chain).Set(o.Age.Seconds())

				if o.Age > tooOld {
					log.Warn(ctx, "Age cache has very old order", nil, "src_chain", chain, "order_id", o.OrderID, "age", o.Age)
				}
			}
		}
	}
}
