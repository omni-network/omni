package gasprice

import (
	"context"
	"fmt"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"

	"github.com/ethereum/go-ethereum"
)

type buffer struct {
	mu     sync.RWMutex
	once   sync.Once
	ticker ticker.Ticker

	// map chainID to buffered gas price (not changed if outside threshold)
	buffer map[uint64]uint64

	// map chainID to provider
	pricers map[uint64]ethereum.GasPricer

	// pct to offset live -> buffer
	// ex. with offset 0.5, live=100, buffered=150 (50% higher)
	// if live increases above 150, offset buffer will update to 225 (50% higher)
	offset float64

	// pct below the buffer the live value be to trigger a buffer decrease
	// ex. with tolerance 0.5, buffered=150, live must be below 75 to decrease buffer
	tolerance float64
}

type Buffer interface {
	GasPrice(chainID uint64) uint64
	Stream(ctx context.Context)
}

var _ Buffer = (*buffer)(nil)

// NewBuffer creates a new gas price buffer.
//
// A gas price buffer maintains a buffered view of gas prices for multiple
// chains. Buffered gas prices exceed live prices by an offset. They decrease
// if live prices fall below tolerance.
func NewBuffer(pricers map[uint64]ethereum.GasPricer, offset, tolerance float64, ticker ticker.Ticker) (Buffer, error) {
	if offset < 0 {
		return nil, errors.New("offset must be >= 0")
	}

	if tolerance < 0 {
		return nil, errors.New("tolerance must be >= 0")
	}

	if (1+offset)*(1-tolerance) >= 1 {
		return nil, errors.New("applying offset would trigger tolerance")
	}

	return &buffer{
		mu:        sync.RWMutex{},
		once:      sync.Once{},
		buffer:    make(map[uint64]uint64),
		pricers:   pricers,
		offset:    offset,
		tolerance: tolerance,
		ticker:    ticker,
	}, nil
}

// GasPrice returns the buffered gas price for the given chainID.
// If the price is not known, returns 0.
func (b *buffer) GasPrice(chainID uint64) uint64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.buffer[chainID]
}

// Stream starts streaming gas prices for all providers into the buffer.
func (b *buffer) Stream(ctx context.Context) {
	b.once.Do(func() {
		ctx = log.WithCtx(ctx, "component", "gasprice.Buffer")
		log.Info(ctx, "Streaming gas prices into buffer")

		b.streamAll(ctx)
	})
}

// streamAll starts streaming gas prices for all providers into the buffer.
func (b *buffer) streamAll(ctx context.Context) {
	for chainID := range b.pricers {
		b.streamOne(ctx, chainID)
	}
}

// streamOne starts streaming gas prices for the given chainID into the buffer.
func (b *buffer) streamOne(ctx context.Context, chainID uint64) {
	ctx = log.WithCtx(ctx, "chainID", chainID)
	pricer := b.pricers[chainID]
	tick := b.ticker

	callback := func(ctx context.Context) {
		liveBn, err := pricer.SuggestGasPrice(ctx)
		if err != nil {
			log.Warn(ctx, "Failed to get gas price (will retry)", err)
			return
		}

		live := liveBn.Uint64()
		guageLive(chainID, live)

		buffed := b.GasPrice(chainID)
		tooLow := live > buffed
		tooHigh := live < uint64(float64(buffed)*(1-b.tolerance))

		log.Debug(ctx, "Checking buffer",
			"live", live,
			"buffered", buffed,
			"too_low", tooLow,
			"too_high", tooHigh,
			"offset", b.offset,
			"tolerance", b.tolerance,
		)

		// do nothing
		if !tooLow && !tooHigh {
			log.Debug(ctx, "No update needed")
			return
		}

		log.Info(ctx, "Updating buffer",
			"live", live,
			"buffered", buffed,
			"too_low", tooLow,
			"too_high", tooHigh,
		)

		corrected := uint64(float64(live) * (1 + b.offset))
		b.setPrice(chainID, corrected)
		guageBuffered(chainID, corrected)
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
	bufferUpdates.WithLabelValues(chainName(chainID)).Inc()
}

// setPrice sets the buffered gas price for the given chainID.
func (b *buffer) setPrice(chainID, price uint64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.buffer[chainID] = price
}

// chainName returns the name of the chain with the given chainID.
func chainName(chainID uint64) string {
	meta, ok := evmchain.MetadataByID(chainID)
	if !ok {
		return fmt.Sprintf("chain-%d", chainID)
	}

	return meta.Name
}
