package ethclient

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/ethclient/headerdb"
	"github.com/omni-network/omni/lib/log"

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
// It caches headers by hash and height, not type.
type headerCache struct {
	Client
	db    *headerdb.DB
	limit int
}

func (c headerCache) HeaderByNumber(ctx context.Context, num *big.Int) (*types.Header, error) {
	// Avoid lookups for "dynamic" headers (h<=0, see rpc.BlockNumber)
	if num != nil && num.Sign() > 0 {
		if header, ok, err := c.db.ByHeight(ctx, num.Uint64()); err != nil {
			return nil, err
		} else if ok {
			cacheHits.WithLabelValues(c.Name()).Inc()
			return header, nil
		}
		cacheMisses.WithLabelValues(c.Name()).Inc()
	}

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

func (c headerCache) add(ctx context.Context, h *types.Header) {
	depth, err := c.db.AddAndReorg(ctx, h, c.Client.HeaderByHash)
	if err != nil {
		// Best effort, don't block on cache update.
		log.Warn(ctx, "Failed adding header to cache (will retry)", err)
		return
	}

	reorgTotal.WithLabelValues(c.Name()).Add(float64(depth))

	if _, err := c.db.MaybePrune(ctx, c.limit); err != nil {
		// Best effort, don't block on cache update.
		log.Warn(ctx, "Failed pruning cache [BUG]", err)
		return
	}
}
