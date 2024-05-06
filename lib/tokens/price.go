package tokens

import (
	"context"
)

// Pricer is the token price provider interface.
type Pricer interface {
	// Price returns the price of each provided token in USD.
	Price(ctx context.Context, tokens ...Token) (map[Token]float64, error)
}

type CachedPricer struct {
	p     Pricer
	cache map[Token]float64
}

func NewCachedPricer(p Pricer) *CachedPricer {
	return &CachedPricer{
		p:     p,
		cache: make(map[Token]float64),
	}
}

func (c *CachedPricer) Price(ctx context.Context, tokens ...Token) (map[Token]float64, error) {
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
	c.cache = make(map[Token]float64)
}
