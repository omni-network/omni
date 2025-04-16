package tokenpricer

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokens"
)

// Pricer is the token price provider interface.
type Pricer interface {
	// USDPrice returns the price of the token in USD.
	USDPrice(ctx context.Context, tokens tokens.Asset) (float64, error)
	// USDPriceRat returns the price of the token in USD.
	USDPriceRat(ctx context.Context, tokens tokens.Asset) (*big.Rat, error)
	// USDPrices returns the price of each provided token in USD.
	USDPrices(ctx context.Context, tokens ...tokens.Asset) (map[tokens.Asset]float64, error)
	// USDPricesRat returns the price of each provided token in USD.
	USDPricesRat(ctx context.Context, tokens ...tokens.Asset) (map[tokens.Asset]*big.Rat, error)
	// Price returns the price of the base asset denominated in the quote asset.
	// Note that for canonical solver prices, base=deposit and quote=expense.
	Price(ctx context.Context, base, quote tokens.Asset) (*big.Rat, error)
}

type Cached struct {
	p     Pricer
	mu    sync.RWMutex
	cache map[pair]*big.Rat
}

type pair struct {
	Base  tokens.Asset
	Quote tokens.Asset
}

func NewCached(p Pricer) *Cached {
	return &Cached{
		p:     p,
		cache: make(map[pair]*big.Rat),
	}
}

func (c *Cached) get(base, quote tokens.Asset) (*big.Rat, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	price, ok := c.cache[pair{Base: base, Quote: quote}]

	return price, ok
}

func (c *Cached) set(base, quote tokens.Asset, price *big.Rat) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[pair{Base: base, Quote: quote}] = price
}

// Price returns the price of the base asset denominated in the quote asset.
// Note that for canonical solver prices, base=deposit and quote=expense.
func (c *Cached) Price(ctx context.Context, base, quote tokens.Asset) (*big.Rat, error) {
	if base == quote {
		return big.NewRat(1, 1), nil
	}

	price, ok := c.get(base, quote)
	if ok {
		return price, nil
	}

	price, err := c.p.Price(ctx, base, quote)
	if err != nil {
		return nil, err
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

// USDPrices returns the price of each coin in USD.
func (c *Cached) USDPrices(ctx context.Context, assets ...tokens.Asset) (map[tokens.Asset]float64, error) {
	rats, err := c.USDPricesRat(ctx, assets...)
	if err != nil {
		return nil, errors.Wrap(err, "get price")
	}

	prices := make(map[tokens.Asset]float64)
	for asset, price := range rats {
		f, _ := price.Float64()
		prices[asset] = f
	}

	return prices, nil
}

func (c *Cached) USDPriceRat(ctx context.Context, asset tokens.Asset) (*big.Rat, error) {
	prices, err := c.USDPricesRat(ctx, asset)
	if err != nil {
		return nil, err
	}

	return prices[asset], nil
}

func (c *Cached) USDPricesRat(ctx context.Context, assets ...tokens.Asset) (map[tokens.Asset]*big.Rat, error) {
	prices := make(map[tokens.Asset]*big.Rat)

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

	newPrices, err := c.p.USDPricesRat(ctx, uncached...)
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

	c.cache = make(map[pair]*big.Rat)
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
