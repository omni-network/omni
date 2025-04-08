package tokenprice_test

import (
	"math/rand"
	"testing"

	"github.com/omni-network/omni/lib/tokenmeta"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/monitor/xfeemngr/ticker"
	"github.com/omni-network/omni/monitor/xfeemngr/tokenprice"

	"github.com/stretchr/testify/require"
)

func TestBufferStream(t *testing.T) {
	t.Parallel()

	initial := map[tokenmeta.Meta]float64{
		tokenmeta.OMNI: randPrice(),
		tokenmeta.ETH:  randPrice(),
	}

	pricer := tokenpricer.NewMock(initial)

	thresh := 0.1
	tick := ticker.NewMock()
	ctx := t.Context()

	b := tokenprice.NewBuffer(pricer, []tokenmeta.Meta{tokenmeta.OMNI, tokenmeta.ETH}, thresh, tick)

	b.Stream(ctx)

	// tick once
	tick.Tick()

	// buffered price should be initial live
	for token, price := range initial {
		require.InEpsilon(t, price, b.Price(token), 0.001, "initial")
	}

	// 10 steps
	buffed := make(map[tokenmeta.Meta]float64)
	for i := 0; i < 10; i++ {
		for token := range initial {
			buffed[token] = b.Price(token)
			pricer.SetPrice(token, randPrice())
		}

		tick.Tick()

		live, err := pricer.Prices(ctx, tokenmeta.OMNI, tokenmeta.ETH)
		require.NoError(t, err)

		// check if any live price is outside threshold
		shouldRefresh := false
		for token, price := range live {
			if inThreshold(price, buffed[token], thresh) {
				continue
			}

			shouldRefresh = true
		}

		// if any price is outside threshold, all prices should be updated
		for token, price := range live {
			if shouldRefresh {
				require.InEpsilon(t, price, b.Price(token), 0.001, "should update")
			} else {
				require.InEpsilon(t, buffed[token], b.Price(token), 0.001, "should not update")
			}
		}
	}
}

// randPrice generates a random, reasonable token price.
func randPrice() float64 {
	return float64(rand.Intn(5000)) + rand.Float64()
}

// inThreshold returns true if a greater or less than b by pct.
func inThreshold(a, b, pct float64) bool {
	gt := a > b+(b*pct)
	lt := a < b-(b*pct)

	return !gt && !lt
}
