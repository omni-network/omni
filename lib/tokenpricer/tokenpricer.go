package tokenpricer

import (
	"context"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokenmeta"
)

// Pricer is the token price provider interface.
type Pricer interface {
	// Price returns the price of the token in USD.
	Price(ctx context.Context, tokens tokenmeta.Meta) (float64, error)
	// Prices returns the price of each provided token in USD.
	Prices(ctx context.Context, tokens ...tokenmeta.Meta) (map[tokenmeta.Meta]float64, error)
}

type Cached struct {
	p     Pricer
	mu    sync.Mutex
	cache map[tokenmeta.Meta]float64
}

func NewCached(p Pricer) *Cached {
	return &Cached{
		p:     p,
		cache: make(map[tokenmeta.Meta]float64),
	}
}

func (c *Cached) Price(ctx context.Context, token tokenmeta.Meta) (float64, error) {
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

func (c *Cached) Prices(ctx context.Context, tokens ...tokenmeta.Meta) (map[tokenmeta.Meta]float64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	prices := make(map[tokenmeta.Meta]float64)

	var uncached []tokenmeta.Meta

	for _, token := range tokens {
		if price, ok := c.cache[token]; ok {
			prices[token] = price
		} else {
			uncached = append(uncached, token)
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

	c.cache = make(map[tokenmeta.Meta]float64)
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
