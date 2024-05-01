package tokenprice

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"
)

type Buffer struct {
	mu     sync.RWMutex
	buffer map[tokens.Token]float64 // map token to price
	pricer tokens.Pricer
	opts   *Opts
}

// NewBuffer creates a new token price buffer.
//
// A token price buffer maintains a buffered view of token prices for multiple
// tokens. Buffered token prices are not updated unless they are outside the
// threshold percentage. Start steaming token prices with Buffer.Stream(ctx).
func NewBuffer(price tokens.Pricer, opts ...func(*Opts)) (*Buffer, error) {
	o, err := makeOpts(opts)
	if err != nil {
		return nil, errors.Wrap(err, "invalid options")
	}

	return &Buffer{
		buffer: make(map[tokens.Token]float64),
		pricer: price,
		opts:   o,
	}, nil
}

// Price returns the buffered price for the given token.
// If the price is not known, returns 0.
func (b *Buffer) Price(token tokens.Token) float64 {
	p, _ := b.price(token)

	return p
}

// Stream starts streaming prices for all tokens into the buffer.
func (b *Buffer) Stream(ctx context.Context) {
	b.stream(ctx)
}

// stream starts streaming prices for all tokens into the buffer.
func (b *Buffer) stream(ctx context.Context) {
	ctx = log.WithCtx(ctx, "tokens", b.opts.tokens)
	tick := b.opts.ticker

	callback := func(ctx context.Context) {
		prices, err := b.pricer.Price(ctx, b.opts.tokens...)
		if err != nil {
			log.Error(ctx, "Failed to get prices - will retry", err)
			return
		}

		for token, price := range prices {
			buffed, ok := b.price(token)

			// if price is buffered, and is within threshold, skip
			if ok && inThreshold(price, buffed, b.opts.thresholdPct) {
				continue
			}

			b.setPrice(token, price)
		}
	}

	tick.Go(ctx, callback)
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
