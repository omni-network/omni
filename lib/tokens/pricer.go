package tokens

import (
	"context"
	"sync"
	"time"
)

// Pricer is the token price provider interface.
type Pricer interface {
	// Price returns the price of each provided token in USD.
	Price(ctx context.Context, tokens ...Token) (map[Token]float64, error)
}

type CachedPricer struct {
	p     Pricer
	mu    sync.Mutex
	cache map[Token]float64
}

func NewCachedPricer(p Pricer) *CachedPricer {
	return &CachedPricer{
		p:     p,
		cache: make(map[Token]float64),
	}
}

func (c *CachedPricer) Price(ctx context.Context, tokens ...Token) (map[Token]float64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	prices := make(map[Token]float64)

	var uncached []Token

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

	newPrices, err := c.p.Price(ctx, uncached...)
	if err != nil {
		return nil, err
	}

	for token, price := range newPrices {
		prices[token] = price
		c.cache[token] = price
	}

	return prices, nil
}

func (c *CachedPricer) ClearCache() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[Token]float64)
}

func (c *CachedPricer) ClearCacheForever(ctx context.Context, evictInterval time.Duration) {
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
