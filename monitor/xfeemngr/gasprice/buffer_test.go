package gasprice_test

import (
	"context"
	"math/rand"
	"testing"

	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/monitor/xfeemngr/gasprice"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"

	"github.com/ethereum/go-ethereum"

	"github.com/stretchr/testify/require"
)

func TestBufferStream(t *testing.T) {
	t.Parallel()

	chainIDs := []uint64{1, 2, 3, 4, 5}

	// initial gas prices per chain
	initials := makePrices(chainIDs)

	// mock gas pricers per chain
	mocks := makeMockPricers(initials)

	thresh := 0.1
	tick := ticker.NewMock()
	ctx := context.Background()

	b := gasprice.NewBuffer(toEthGasPricers(mocks), gasprice.WithThresholdPct(thresh), gasprice.WithTicker(tick))

	// start streaming gas prices
	b.Stream(ctx)

	// tick once - initial prices should get buffered
	tick.Tick()

	// buffered price should be initial
	for chainID, price := range initials {
		require.Equal(t, price, b.GasPrice(chainID), "initial")
	}

	// just increase a little, but not above threshold
	for chainID, price := range initials {
		delta := umath.SubtractOrZero(uint64(float64(price)*thresh), 1)
		mocks[chainID].SetPrice(price + delta)
	}

	tick.Tick()

	// buffered price should still be initial
	for chainID, price := range initials {
		require.Equal(t, price, b.GasPrice(chainID), "within threshold")
	}

	// increase above threshold
	for chainID, price := range initials {
		mocks[chainID].SetPrice(price + uint64(float64(price)*thresh)*2)
	}

	tick.Tick()

	// buffered price should be updated
	for chainID, mock := range mocks {
		require.Equal(t, mock.Price(), b.GasPrice(chainID), "outside threshold")
	}

	// reset back to initial
	for chainID, price := range initials {
		mocks[chainID].SetPrice(price)
	}

	tick.Tick()

	// buffered price should be initial
	for chainID, price := range initials {
		require.Equal(t, price, b.GasPrice(chainID), "reset")
	}
}

// makePrices generates a map chainID -> gas price for each chainID.
func makePrices(chainIDs []uint64) map[uint64]uint64 {
	prices := make(map[uint64]uint64)

	for _, chainID := range chainIDs {
		prices[chainID] = randGasPrice()
	}

	return prices
}

// makeMockPricers generates a map of mock gas pricers for n chains.
func makeMockPricers(prices map[uint64]uint64) map[uint64]*gasprice.MockPricer {
	mocks := make(map[uint64]*gasprice.MockPricer)
	for chainID, price := range prices {
		mocks[chainID] = gasprice.NewMockPricer(price)
	}

	return mocks
}

// toEthGasPricers  transform map[uint64]*gasprice.MockPricer to map[uint64]ethereum.GasPricer (interface).
func toEthGasPricers(mocks map[uint64]*gasprice.MockPricer) map[uint64]ethereum.GasPricer {
	pricers := make(map[uint64]ethereum.GasPricer)
	for chainID, mock := range mocks {
		pricers[chainID] = mock
	}

	return pricers
}

// randGasPrice generates a random, reasonable gas price.
func randGasPrice() uint64 {
	oneGwei := 1_000_000_000 // i gwei
	return uint64(rand.Float64() * float64(oneGwei))
}
