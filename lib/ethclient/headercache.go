package ethclient

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/ethclient/headerdb"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	dbm "github.com/cosmos/cosmos-db"
)

const defaultCacheLimit = 1000

func newHeaderCache(ethCl Client) (*headerCache, error) {
	db, err := headerdb.New(dbm.NewMemDB())
	if err != nil {
		return nil, err
	}

	return &headerCache{
		Client: ethCl,
		db:     db,
		limit:  defaultCacheLimit,
	}, nil
}

// headerCache extends/wraps a Client with a read-through cache for headers.
//
// It caches headers by hash only, not height or type since that has non-zero risk of returning reorged blocks.
type headerCache struct {
	Client
	db    *headerdb.DB
	limit int
}

func (c headerCache) HeaderByNumber(ctx context.Context, num *big.Int) (*types.Header, error) {
	header, err := c.Client.HeaderByNumber(ctx, num)
	if err != nil {
		return nil, err
	}

	c.add(ctx, header)

	return header, nil
}

func (c headerCache) HeaderByType(ctx context.Context, typ HeadType) (*types.Header, error) {
	header, err := c.Client.HeaderByType(ctx, typ)
	if err != nil {
		return nil, err
	}

	c.add(ctx, header)

	return header, nil
}

func (c headerCache) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	if header, ok, err := c.db.ByHash(ctx, hash); err != nil {
		return nil, err
	} else if ok {
		cacheHits.WithLabelValues(c.Name()).Inc()
		return header, nil
	}
	cacheMisses.WithLabelValues(c.Name()).Inc()

	header, err := c.Client.HeaderByHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	c.add(ctx, header)

	return header, nil
}

// SubscribeNewHead subscribes to new headers and caches them.
// Note that the Subscription.Unsubscribe MUST be called to avoid leaking resources (if error is nil).
func (c headerCache) SubscribeNewHead(ctx context.Context, outer chan<- *types.Header) (ethereum.Subscription, error) {
	inner := make(chan *types.Header)
	sub, err := c.Client.SubscribeNewHead(ctx, inner)
	if err != nil {
		return nil, err
	}

	cancelSub := newCancelSub(sub)

	// Not using ctx for scheduling since the contract of SubscribeNewHead
	// states that context is only applicable to initial setup, not subscription lifetime.

	go func() {
		for {
			select {
			case <-cancelSub.cancel:
				return
			case header := <-inner:
				instrumentWebsocket(c.Name(), header)

				c.add(ctx, header)

				select {
				case outer <- header:
				case <-cancelSub.cancel:
					return
				}
			}
		}
	}()

	return cancelSub, nil
}

func (c headerCache) add(ctx context.Context, h *types.Header) {
	depth, err := c.db.AddAndReorg(ctx, h, c.Client.HeaderByHash)
	if err != nil {
		// Best effort, don't block on cache update.
		log.Warn(ctx, "Failed adding header to cache (will retry)", err)
		return
	}

	reorgTotal.WithLabelValues(c.Name()).Inc()
	// Increment counter instead of adding depth, since one cannot easily distinguish counter
	// increases when close together. Log the depth instead.
	if depth > 0 {
		log.Debug(ctx, "Chain reorg detected", "chain", c.Name(), "depth", depth, "height", h.Number)
	}

	if _, err := c.db.MaybePrune(ctx, c.limit); err != nil {
		// Best effort, don't block on cache update.
		log.Warn(ctx, "Failed pruning cache [BUG]", err)
		return
	}
}

func newCancelSub(sub ethereum.Subscription) *cancelSub {
	return &cancelSub{
		Subscription: sub,
		cancel:       make(chan struct{}),
	}
}

// cancelSub wraps an ethereum.Subscription and adds a cancel channel that is closed on Unsubscribe.
type cancelSub struct {
	ethereum.Subscription
	cancelOnce sync.Once
	cancel     chan struct{}
}

func (s *cancelSub) Unsubscribe() {
	s.cancelOnce.Do(func() {
		close(s.cancel)
	})
	s.Subscription.Unsubscribe()
}

func instrumentWebsocket(name string, header *types.Header) {
	epochSecs, err := umath.ToInt64(header.Time)
	if err != nil {
		return
	}
	latency := time.Since(time.Unix(epochSecs, 0))
	websocketLatency.WithLabelValues(name).Set(latency.Seconds())
}
