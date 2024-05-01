package tokenprice_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
	"github.com/omni-network/omni/monitor/xfeemngr/tokenprice"

	"github.com/stretchr/testify/require"
)

type mockPricer struct {
	prices  map[tokens.Token]float64
	initial map[tokens.Token]float64
}

func newMock(prices map[tokens.Token]float64) *mockPricer {
	cloned := make(map[tokens.Token]float64)
	for k, v := range prices {
		cloned[k] = v
	}

	return &mockPricer{
		prices:  prices,
		initial: cloned,
	}
}

func (m *mockPricer) Price(ctx context.Context, tkns ...tokens.Token) (map[tokens.Token]float64, error) {
	resp := make(map[tokens.Token]float64)
	for _, t := range tkns {
		resp[t] = m.prices[t]
	}

	return resp, nil
}

func (m *mockPricer) setPrice(token tokens.Token, price float64) {
	m.prices[token] = price
}

func TestBufferStream(t *testing.T) {
	t.Parallel()

	prices := map[tokens.Token]float64{
		tokens.OMNI: 20,
		tokens.ETH:  3000,
	}

	pricer := newMock(prices)

	thresh := 0.1
	tick := ticker.NewMock()
	ctx := context.Background()

	b, err := tokenprice.NewBuffer(
		pricer,
		tokenprice.WithTokens(tokens.OMNI, tokens.ETH),
		tokenprice.WithThresholdPct(thresh),
		tokenprice.WithTicker(tick))

	require.NoError(t, err)

	b.Stream(ctx)
	tick.Tick(ctx)

	// buffered price should be initial
	for token, price := range pricer.initial {
		require.InEpsilon(t, price, b.Price(token), 0.01, "initial")
	}

	// just increase a little, but not above threshold
	for token, price := range pricer.initial {
		pricer.setPrice(token, price+(price*thresh)-1)
	}

	tick.Tick(ctx)

	// buffered price should still be initial
	for token, price := range pricer.initial {
		require.InEpsilon(t, price, b.Price(token), 0.01, "within threshold")
	}

	// increase above threshold
	for token, price := range pricer.prices {
		pricer.setPrice(token, price+(price*thresh)+10)
	}

	tick.Tick(ctx)

	// buffered price should be updated
	for token, price := range pricer.prices {
		require.InEpsilon(t, price, b.Price(token), 0.01, "outside threshold")
	}

	// reset back to initial
	for token, price := range pricer.initial {
		pricer.setPrice(token, price)
	}

	tick.Tick(ctx)

	// buffered price should be initial
	for token, price := range pricer.initial {
		require.InEpsilon(t, price, b.Price(token), 0.01, "reset")
	}
}
