package gasprice_test

import (
	"context"
	"math/rand"
	"testing"

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

	offset := 0.3    // 30% offset
	tolerance := 0.5 // 50% tolerance
	tick := ticker.NewMock()
	ctx := context.Background()

	withOffset := func(price uint64) uint64 {
		return uint64(float64(price) * (1 + offset))
	}

	atTolerance := func(price uint64) uint64 {
		return uint64(float64(price) * (1 - tolerance))
	}

	b, err := gasprice.NewBuffer(toEthGasPricers(mocks), offset, tolerance, tick)
	require.NoError(t, err)

	b.Stream(ctx)

	// tick once
	tick.Tick()

	// buffered price should be initial live + offset
	for chainID, price := range initials {
		require.Equal(t, withOffset(price), b.GasPrice(chainID), "initial")
	}

	// 10 steps
	buffed := make(map[uint64]uint64)
	for i := 0; i < 10; i++ {
		for chainID, mock := range mocks {
			buffed[chainID] = b.GasPrice(chainID)
			mock.SetPrice(randGasPrice())
		}

		tick.Tick()

		// for each step, we check if buffer properly updates (or doesn't)
		for chainID, mock := range mocks {
			tooLow := mock.Price() > buffed[chainID]
			tooHigh := mock.Price() < atTolerance(buffed[chainID])

			if tooHigh || tooLow {
				require.Equal(t, withOffset(mock.Price()), b.GasPrice(chainID), 0.01, "should change")
			} else {
				require.Equal(t, buffed[chainID], b.GasPrice(chainID), 0.01, "should not change")
			}
		}
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
