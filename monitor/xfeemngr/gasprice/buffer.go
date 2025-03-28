package gasprice

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/omni-network/omni/lib/bi"
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
	buffer map[uint64]*big.Int

	// map chainID to provider
	pricers map[uint64]ethereum.GasPricer
}

type Buffer interface {
	GasPrice(chainID uint64) *big.Int
	Stream(ctx context.Context)
}

var _ Buffer = (*buffer)(nil)

// NewBuffer creates a new gas price buffer.
func NewBuffer(pricers map[uint64]ethereum.GasPricer, ticker ticker.Ticker) (Buffer, error) {
	return &buffer{
		mu:      sync.RWMutex{},
		once:    sync.Once{},
		buffer:  make(map[uint64]*big.Int),
		pricers: pricers,
		ticker:  ticker,
	}, nil
}

// GasPrice returns the buffered gas price for the given chainID.
// If the price is not known, returns 0.
func (b *buffer) GasPrice(chainID uint64) *big.Int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	resp, ok := b.buffer[chainID]
	if !ok {
		return bi.Zero()
	}

	return resp
}

// Stream starts streaming gas prices for all providers into the buffer.
func (b *buffer) Stream(ctx context.Context) {
	b.once.Do(func() {
		ctx := log.WithCtx(ctx, "component", "gasprice.Buffer")
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
		live, err := pricer.SuggestGasPrice(ctx)
		if err != nil {
			log.Warn(ctx, "Failed to get gas price (will retry)", err)
			return
		}

		guageLive(chainID, live)

		tiered := Tier(live)
		buffed := b.GasPrice(chainID)

		if bi.EQ(tiered, buffed) {
			return
		}

		b.setPrice(chainID, tiered)
		guageBuffered(chainID, tiered)
	}

	tick.Go(ctx, callback)
}

// guageLive updates "live" guages for chain's gas price.
func guageLive(chainID uint64, price *big.Int) {
	f, _ := price.Float64()
	liveGasPrice.WithLabelValues(chainName(chainID)).Set(f)
}

// guageBuffered updates "buffered" guages for a chain's gas price.
func guageBuffered(chainID uint64, price *big.Int) {
	f, _ := price.Float64()
	bufferedGasPrice.WithLabelValues(chainName(chainID)).Set(f)
}

// setPrice sets the buffered gas price for the given chainID.
func (b *buffer) setPrice(chainID uint64, price *big.Int) {
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
