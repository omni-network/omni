package tokenprice

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
)

type buffer struct {
	mu        sync.RWMutex
	once      sync.Once
	buffer    map[tokens.Asset]float64 // map token to price
	pricer    tokenpricer.Pricer
	tokens    []tokens.Asset
	threshold float64
	tick      ticker.Ticker
}

type Buffer interface {
	Price(token tokens.Asset) float64
	Stream(ctx context.Context)
}

var _ Buffer = (*buffer)(nil)

// NewBuffer creates a new token price buffer.
//
// A token price buffer maintains a buffered view of token prices for multiple
// tokens. Buffered token prices are not updated unless they are outside the
// threshold percentage. Start steaming token prices with Buffer.Stream(ctx).
func NewBuffer(price tokenpricer.Pricer, tkns []tokens.Asset, threshold float64, ticker ticker.Ticker) Buffer {
	return &buffer{
		mu:        sync.RWMutex{},
		buffer:    make(map[tokens.Asset]float64),
		pricer:    price,
		tokens:    tkns,
		threshold: threshold,
		tick:      ticker,
	}
}

// Price returns the buffered price for the given token.
// If the price is not known, returns 0.
func (b *buffer) Price(token tokens.Asset) float64 {
	p, _ := b.price(token)

	return p
}

// Stream starts streaming prices for all tokens into the buffer.
func (b *buffer) Stream(ctx context.Context) {
	b.once.Do(func() {
		ctx := log.WithCtx(ctx, "component", "tokenprice.Buffer")
		log.Info(ctx, "Streaming token prices into buffer")

		b.stream(ctx)
	})
}

// stream starts streaming prices for all tokens into the buffer.
func (b *buffer) stream(ctx context.Context) {
	callback := func(ctx context.Context) {
		prices, err := b.pricer.USDPrices(ctx, b.tokens...)
		if err != nil {
			log.Warn(ctx, "Failed to get prices (will retry)", err)
			return
		}

		guageLive(prices)

		// check if any prices have changed by more than the threshold
		refresh := false
		for token, price := range prices {
			buffed, ok := b.price(token)

			if ok && inThreshold(price, buffed, b.threshold) {
				continue
			}

			refresh = true
		}

		// if any outside threshold, update all
		if refresh {
			for token, price := range prices {
				b.setPrice(token, price)
			}
		}

		b.gaugeBuffered()
	}

	b.tick.Go(ctx, callback)
}

// guageLive updates "live" guages for token prices and conversion rates.
func guageLive(prices map[tokens.Asset]float64) {
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

// gaugeBuffered updates "buffered" gauges for token prices and conversion rates.
func (b *buffer) gaugeBuffered() {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for token, price := range b.buffer {
		if price == 0 {
			continue
		}

		bufferedTokenPrice.WithLabelValues(token.String()).Set(price)

		for otherToken, otherPrice := range b.buffer {
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
func (b *buffer) price(token tokens.Asset) (float64, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	p, ok := b.buffer[token]

	return p, ok
}

// setPrice sets the buffered price for the given token.
func (b *buffer) setPrice(token tokens.Asset, price float64) {
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
