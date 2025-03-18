package gasprice_test

import (
	"context"
	"math/big"
	"math/rand"
	"testing"

	"github.com/omni-network/omni/lib/bi"
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

	tick := ticker.NewMock()
	ctx := context.Background()

	b, err := gasprice.NewBuffer(toEthGasPricers(mocks), tick)
	require.NoError(t, err)

	b.Stream(ctx)

	// tick once
	tick.Tick()

	// buffered price should be initial live + offset
	for chainID, price := range initials {
		require.Equal(t, gasprice.Tier(price), b.GasPrice(chainID), "initial")
	}

	// 10 steps
	live := make(map[uint64]*big.Int)
	for i := 0; i < 10; i++ {
		for chainID, mock := range mocks {
			live[chainID] = randGasPrice()
			mock.SetPrice(live[chainID])
		}

		tick.Tick()

		// for each step, we check if buffer properly updates (or doesn't)
		for chainID, mock := range mocks {
			tier := gasprice.Tier(mock.Price())
			require.True(t, bi.GTE(tier, live[chainID]), "tier greater than live")
			require.Equal(t, tier, b.GasPrice(chainID), "buffer equal to tier")
		}
	}
}

// makePrices generates a map chainID -> gas price for each chainID.
func makePrices(chainIDs []uint64) map[uint64]*big.Int {
	prices := make(map[uint64]*big.Int)

	for _, chainID := range chainIDs {
		prices[chainID] = randGasPrice()
	}

	return prices
}

// makeMockPricers generates a map of mock gas pricers for n chains.
func makeMockPricers(prices map[uint64]*big.Int) map[uint64]*gasprice.MockPricer {
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
func randGasPrice() *big.Int {
	oneGwei := int64(1_000_000_000) // i gwei
	return bi.N(rand.Int63n(oneGwei))
}
