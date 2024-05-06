package gasprice

import (
	"context"
	"fmt"
	"sync"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum"
)

type Buffer struct {
	mu      sync.RWMutex
	once    sync.Once
	buffer  map[uint64]uint64             // map chainID to buffered gas price (not changed if outside threshold)
	pricers map[uint64]ethereum.GasPricer // map chainID to provider
	opts    *Opts
}

// NewBuffer creates a new gas price buffer.
//
// A gas price buffer maintains a buffered view of gas prices for multiple
// chains. Buffered gas prices are not updated unless they are outside the
// threshold percentage. Start steaming gas prices with Buffer.Stream(ctx).
func NewBuffer(pricers map[uint64]ethereum.GasPricer, opts ...func(*Opts)) *Buffer {
	return &Buffer{
		mu:      sync.RWMutex{},
		once:    sync.Once{},
		buffer:  make(map[uint64]uint64),
		pricers: pricers,
		opts:    makeOpts(opts),
	}
}

// GasPrice returns the buffered gas price for the given chainID.
// If the price is not known, returns 0.
func (b *Buffer) GasPrice(chainID uint64) uint64 {
	p, _ := b.price(chainID)
	return p
}

// Stream starts streaming gas prices for all providers into the buffer.
func (b *Buffer) Stream(ctx context.Context) {
	b.once.Do(func() {
		ctx = log.WithCtx(ctx, "component", "gasprice.Buffer")
		log.Info(ctx, "Streaming gas prices into buffer")

		b.streamAll(ctx)
	})
}

// streamAll starts streaming gas prices for all providers into the buffer.
func (b *Buffer) streamAll(ctx context.Context) {
	for chainID := range b.pricers {
		b.streamOne(ctx, chainID)
	}
}

// streamOne starts streaming gas prices for the given chainID into the buffer.
func (b *Buffer) streamOne(ctx context.Context, chainID uint64) {
	ctx = log.WithCtx(ctx, "chainID", chainID)
	pricer := b.pricers[chainID]
	tick := b.opts.ticker

	callback := func(ctx context.Context) {
		gpriceBig, err := pricer.SuggestGasPrice(ctx)
		if err != nil {
			log.Error(ctx, "Failed to get gas price - will retry", err)
			return
		}

		gprice := gpriceBig.Uint64()
		guageLive(chainID, gprice)

		// if price is buffed, and within threshold, return
		buffed, ok := b.price(chainID)
		if ok && inThreshold(gprice, buffed, b.opts.thresholdPct) {
			return
		}

		b.setPrice(chainID, gprice)
		guageBuffered(chainID, gprice)
	}

	tick.Go(ctx, callback)
}

// guageLive updates "live" guages for chain's gas price.
func guageLive(chainID uint64, price uint64) {
	liveGasPrice.WithLabelValues(chainName(chainID)).Set(float64(price))
}

// guageBuffered updates "buffered" guages for a chain's gas price.
func guageBuffered(chainID uint64, price uint64) {
	bufferedGasPrice.WithLabelValues(chainName(chainID)).Set(float64(price))
}

// setPrice sets the buffered gas price for the given chainID.
func (b *Buffer) setPrice(chainID, price uint64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.buffer[chainID] = price
}

// price returns the buffered gas price for the given chainID.
// If the price is not found, returns 0 and false.
func (b *Buffer) price(chainID uint64) (uint64, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	price, ok := b.buffer[chainID]

	return price, ok
}

// inThreshold returns true if a greater or less than b by pct.
func inThreshold(a, b uint64, pct float64) bool {
	bf := float64(b)
	gt := a > uint64(bf+(bf*pct))
	lt := a < uint64(bf-(bf*pct))

	return !gt && !lt
}

// chainName returns the name of the chain with the given chainID.
func chainName(chainID uint64) string {
	meta, ok := evmchain.MetadataByID(chainID)
	if !ok {
		return fmt.Sprintf("chain-%d", chainID)
	}

	return meta.Name
}
