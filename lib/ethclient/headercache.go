package ethclient

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const defaultCacheLimit = 1000

func newHeaderCache(ethCl Client) *headerCache {
	return &headerCache{
		Client:  ethCl,
		limit:   defaultCacheLimit,
		headers: make(map[common.Hash]*types.Header),
	}
}

// headerCache extends/wraps a Client with a read-through FIFO cache for headers.
//
// It only caches headers by hash, not type or height.
type headerCache struct {
	Client
	limit int

	mu      sync.RWMutex
	fifo    []common.Hash
	headers map[common.Hash]*types.Header
}

func (c *headerCache) HeaderByNumber(ctx context.Context, num *big.Int) (*types.Header, error) {
	header, err := c.Client.HeaderByNumber(ctx, num)
	if err != nil {
		return nil, err
	}

	c.add(header)

	return header, nil
}

func (c *headerCache) HeaderByType(ctx context.Context, typ HeadType) (*types.Header, error) {
	header, err := c.Client.HeaderByType(ctx, typ)
	if err != nil {
		return nil, err
	}

	c.add(header)

	return header, nil
}

func (c *headerCache) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	if header, ok := c.get(hash); ok {
		return header, nil
	}

	header, err := c.Client.HeaderByHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	c.add(header)

	return header, nil
}

func (c *headerCache) add(h *types.Header) {
	c.mu.Lock()
	defer c.mu.Unlock()

	hash := h.Hash()
	c.fifo = append(c.fifo, hash)
	c.headers[hash] = h

	c.maybePruneUnsafe()
}

func (c *headerCache) get(hash common.Hash) (*types.Header, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	header, ok := c.headers[hash]

	return header, ok
}

// maybePrune removes the oldest header if over the limit.
// It is unsafe since it assumes the lock is held.
func (c *headerCache) maybePruneUnsafe() {
	for len(c.fifo) > c.limit {
		hash := c.fifo[0]
		delete(c.headers, hash)
		c.fifo = c.fifo[1:]
	}
}
