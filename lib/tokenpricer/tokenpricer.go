package tokenpricer

import (
	"context"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokens"
)

// Pricer is the token price provider interface.
type Pricer interface {
	// Price returns the price of the token in USD.
	Price(ctx context.Context, tokens tokens.Asset) (float64, error)
	// Prices returns the price of each provided token in USD.
	Prices(ctx context.Context, tokens ...tokens.Asset) (map[tokens.Asset]float64, error)
}

type Cached struct {
	p     Pricer
	mu    sync.Mutex
	cache map[tokens.Asset]float64
}

func NewCached(p Pricer) *Cached {
	return &Cached{
		p:     p,
		cache: make(map[tokens.Asset]float64),
	}
}

func (c *Cached) Price(ctx context.Context, token tokens.Asset) (float64, error) {
	prices, err := c.Prices(ctx, token)
	if err != nil {
		return 0, err
	}

	price, ok := prices[token]
	if !ok {
		return 0, errors.New("missing token price [BUG]")
	}

	return price, nil
}

func (c *Cached) Prices(ctx context.Context, assets ...tokens.Asset) (map[tokens.Asset]float64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	prices := make(map[tokens.Asset]float64)

	var uncached []tokens.Asset

	for _, asset := range assets {
		if price, ok := c.cache[asset]; ok {
			prices[asset] = price
		} else {
			uncached = append(uncached, asset)
		}
	}

	if len(uncached) == 0 {
		return prices, nil
	}

	newPrices, err := c.p.Prices(ctx, uncached...)
	if err != nil {
		return nil, err
	}

	for token, price := range newPrices {
		prices[token] = price
		c.cache[token] = price
	}

	return prices, nil
}

func (c *Cached) ClearCache() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[tokens.Asset]float64)
}

func (c *Cached) ClearCacheForever(ctx context.Context, evictInterval time.Duration) {
	ticker := time.NewTicker(evictInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.ClearCache()
		}
	}
}
