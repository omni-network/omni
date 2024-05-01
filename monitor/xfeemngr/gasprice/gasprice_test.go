package gasprice_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/omni-network/omni/monitor/xfeemngr/gasprice"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"

	"github.com/ethereum/go-ethereum"

	"github.com/stretchr/testify/require"
)

type mockPricer struct {
	price   uint64
	initial uint64
}

func newMock(price uint64) *mockPricer {
	return &mockPricer{
		price:   price,
		initial: price,
	}
}

func (m *mockPricer) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(int64(m.price)), nil
}

func (m *mockPricer) setPrice(price uint64) {
	m.price = price
}

func TestBufferStream(t *testing.T) {
	t.Parallel()

	mocks := map[uint64]*mockPricer{
		1: newMock(100),
		2: newMock(200),
		3: newMock(300),
	}

	pricers := map[uint64]ethereum.GasPricer{
		1: mocks[1],
		2: mocks[2],
		3: mocks[3],
	}

	thresh := 0.1

	tick := ticker.NewMock()
	ctx := context.Background()

	b := gasprice.NewBuffer(pricers, gasprice.WithThresholdPct(thresh), gasprice.WithTicker(tick))

	b.Stream(ctx)
	tick.Tick(ctx)

	// buffered price should be initial
	for chainID, mock := range mocks {
		require.Equal(t, mock.initial, b.GasPrice(chainID), "initial")
	}

	// just increase a little, but not above threshold
	for _, mock := range mocks {
		mock.setPrice(mock.initial + uint64(float64(mock.initial)*thresh) - 1)
	}

	tick.Tick(ctx)

	// buffered price should still be initial
	for chainID, mock := range mocks {
		require.Equal(t, mock.initial, b.GasPrice(chainID), "within threshold")
	}

	// increase above threshold
	for _, mock := range mocks {
		mock.setPrice(mock.initial + uint64(float64(mock.initial)*thresh) + 10)
	}

	tick.Tick(ctx)

	// buffered price should be updated
	for chainID, mock := range mocks {
		require.Equal(t, mock.price, b.GasPrice(chainID), "outside threshold")
	}

	// reset back to initial
	for _, mock := range mocks {
		mock.setPrice(mock.initial)
	}

	tick.Tick(ctx)

	// buffered price should be initial
	for chainID, mock := range mocks {
		require.Equal(t, mock.initial, b.GasPrice(chainID), "reset")
	}
}
