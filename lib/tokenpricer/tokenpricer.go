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
	// USDPrice returns the price of the token in USD.
	USDPrice(ctx context.Context, tokens tokens.Asset) (float64, error)
	// USDPrices returns the price of each provided token in USD.
	USDPrices(ctx context.Context, tokens ...tokens.Asset) (map[tokens.Asset]float64, error)
	// Price returns the price of the base asset denominated in the quote asset.
	Price(ctx context.Context, base, quote tokens.Asset) (float64, error)
}

type Cached struct {
	p     Pricer
	mu    sync.RWMutex
	cache map[pair]float64
}

type pair struct {
	Base  tokens.Asset
	Quote tokens.Asset
}

func NewCached(p Pricer) *Cached {
	return &Cached{
		p:     p,
		cache: make(map[pair]float64),
	}
}

func (c *Cached) get(base, quote tokens.Asset) (float64, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	price, ok := c.cache[pair{Base: base, Quote: quote}]

	return price, ok
}

func (c *Cached) set(base, quote tokens.Asset, price float64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[pair{Base: base, Quote: quote}] = price
}

func (c *Cached) Price(ctx context.Context, base, quote tokens.Asset) (float64, error) {
	if base == quote {
		return 1, nil
	}

	price, ok := c.get(base, quote)
	if ok {
		return price, nil
	}

	price, err := c.p.Price(ctx, base, quote)
	if err != nil {
		return 0, err
	}

	c.set(base, quote, price)

	return price, nil
}

func (c *Cached) USDPrice(ctx context.Context, token tokens.Asset) (float64, error) {
	prices, err := c.USDPrices(ctx, token)
	if err != nil {
		return 0, err
	}

	price, ok := prices[token]
	if !ok {
		return 0, errors.New("missing token price [BUG]")
	}

	return price, nil
}

func (c *Cached) USDPrices(ctx context.Context, assets ...tokens.Asset) (map[tokens.Asset]float64, error) {
	prices := make(map[tokens.Asset]float64)

	var uncached []tokens.Asset

	for _, asset := range assets {
		if price, ok := c.get(asset, tokens.USDC); ok {
			prices[asset] = price
		} else {
			uncached = append(uncached, asset)
		}
	}

	if len(uncached) == 0 {
		return prices, nil
	}

	newPrices, err := c.p.USDPrices(ctx, uncached...)
	if err != nil {
		return nil, err
	}

	for token, price := range newPrices {
		prices[token] = price
		c.set(token, tokens.USDC, price)
	}

	return prices, nil
}

func (c *Cached) ClearCache() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[pair]float64)
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
