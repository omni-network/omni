package tokenprice

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
)

type Buffer struct {
	mu     sync.RWMutex
	once   sync.Once
	buffer map[tokens.Token]float64 // map token to price
	pricer tokens.Pricer
	tokens []tokens.Token
	opts   *Opts
}

// NewBuffer creates a new token price buffer.
//
// A token price buffer maintains a buffered view of token prices for multiple
// tokens. Buffered token prices are not updated unless they are outside the
// threshold percentage. Start steaming token prices with Buffer.Stream(ctx).
func NewBuffer(price tokens.Pricer, optOrTokens ...any) *Buffer {
	opts := defaultOpts()
	tkns := make([]tokens.Token, 0)

	for _, optOrToken := range optOrTokens {
		if o, ok := optOrToken.(func(*Opts)); ok {
			o(opts)
		}

		if token, ok := optOrToken.(tokens.Token); ok {
			tkns = append(tkns, token)
		}
	}

	return &Buffer{
		mu:     sync.RWMutex{},
		buffer: make(map[tokens.Token]float64),
		pricer: price,
		tokens: tkns,
		opts:   opts,
	}
}

// Price returns the buffered price for the given token.
// If the price is not known, returns 0.
func (b *Buffer) Price(token tokens.Token) float64 {
	p, _ := b.price(token)

	return p
}

// Stream starts streaming prices for all tokens into the buffer.
func (b *Buffer) Stream(ctx context.Context) {
	b.once.Do(func() {
		ctx = log.WithCtx(ctx, "component", "tokenprice.Buffer")
		log.Info(ctx, "Streaming token prices into buffer")

		b.stream(ctx)
	})
}

// stream starts streaming prices for all tokens into the buffer.
func (b *Buffer) stream(ctx context.Context) {
	tick := b.opts.ticker

	callback := func(ctx context.Context) {
		prices, err := b.pricer.Price(ctx, b.tokens...)
		if err != nil {
			log.Error(ctx, "Failed to get prices - will retry", err)
			return
		}

		guageLive(prices)

		// update buffered prices, if necessary
		for token, price := range prices {
			buffed, ok := b.price(token)

			// if price is buffered, and is within threshold, skip
			if ok && inThreshold(price, buffed, b.opts.thresholdPct) {
				continue
			}

			b.setPrice(token, price)
		}

		guageBuffered(b.buffer)
	}

	tick.Go(ctx, callback)
}

// guageLive updates "live" guages for token prices and conversion rates.
func guageLive(prices map[tokens.Token]float64) {
	for token, price := range prices {
		if price == 0 {
			continue
		}

		liveTokenPrice.WithLabelValues(token.String()).Set(price)

		for otherToken, otherPrice := range prices {
			if otherToken == token {
				continue
			}

			// rate "token / other" is "price other / price token"
			rate := otherPrice / price

			liveConversionRate.WithLabelValues(token.String(), otherToken.String()).Set(rate)
		}
	}
}

// guageBuffered updates "buffered" guages for token prices and conversion rates.
func guageBuffered(prices map[tokens.Token]float64) {
	for token, price := range prices {
		if price == 0 {
			continue
		}

		bufferedTokenPrice.WithLabelValues(token.String()).Set(price)

		for otherToken, otherPrice := range prices {
			if otherToken == token {
				continue
			}

			// rate "token / other" is "price other / price token"
			rate := otherPrice / price

			bufferedConversionRate.WithLabelValues(token.String(), otherToken.String()).Set(rate)
		}
	}
}

// price returns the buffered price for the given token.
// If the price is not known, returns 0 and false.
func (b *Buffer) price(token tokens.Token) (float64, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	p, ok := b.buffer[token]

	return p, ok
}

// setPrice sets the buffered price for the given token.
func (b *Buffer) setPrice(token tokens.Token, price float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.buffer[token] = price
}

// inThreshold returns true if a greater or less than b by pct.
func inThreshold(a, b, pct float64) bool {
	gt := a > b+(b*pct)
	lt := a < b-(b*pct)

	return !gt && !lt
}
