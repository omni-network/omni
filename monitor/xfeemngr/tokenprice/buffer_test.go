package tokenprice_test

import (
	"context"
	"math/rand"
	"testing"

	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
	"github.com/omni-network/omni/monitor/xfeemngr/tokenprice"

	"github.com/stretchr/testify/require"
)

func TestBufferStream(t *testing.T) {
	t.Parallel()

	initial := map[tokens.Token]float64{
		tokens.OMNI: randPrice(),
		tokens.ETH:  randPrice(),
	}

	pricer := tokens.NewMockPricer(initial)

	thresh := 0.1
	tick := ticker.NewMock()
	ctx := context.Background()

	b := tokenprice.NewBuffer(pricer, tokens.OMNI, tokens.ETH, tokenprice.WithThresholdPct(thresh), tokenprice.WithTicker(tick))

	b.Stream(ctx)
	tick.Tick()

	// buffered price should be initial
	for token, price := range initial {
		require.InEpsilon(t, price, b.Price(token), 0.01, "initial")
	}

	// just increase a little, but not above threshold
	for token, price := range initial {
		pricer.SetPrice(token, price+(price*thresh)-1)
	}

	tick.Tick()

	// buffered price should still be initial
	for token, price := range initial {
		require.InEpsilon(t, price, b.Price(token), 0.01, "within threshold")
	}

	// increase above threshold
	for token, price := range initial {
		pricer.SetPrice(token, price+(price*thresh)*2)
	}

	tick.Tick()

	// buffered price should be updated
	for token := range initial {
		prices, err := pricer.Price(ctx, token)
		require.NoError(t, err)
		price := prices[token]
		require.InEpsilon(t, price, b.Price(token), 0.01, "outside threshold")
	}

	// reset back to initial
	for token, price := range initial {
		pricer.SetPrice(token, price)
	}

	tick.Tick()

	// buffered price should be initial
	for token, price := range initial {
		require.InEpsilon(t, price, b.Price(token), 0.01, "reset")
	}
}

// randPrice generates a random, reasonable token price.
func randPrice() float64 {
	return float64(rand.Intn(5000)) + rand.Float64()
}
