package tokenpricer_test

import (
	"testing"

	"github.com/omni-network/omni/lib/tokenmeta"
	"github.com/omni-network/omni/lib/tokenpricer"

	"github.com/stretchr/testify/require"
)

func TestCachedPricer(t *testing.T) {
	t.Parallel()
	const epsilon = 1e-6

	ETH := tokenmeta.ETH
	OMNI := tokenmeta.OMNI

	pricer := tokenpricer.NewMock(map[tokenmeta.Meta]float64{
		ETH:  100,
		OMNI: 200,
	})

	cached := tokenpricer.NewCached(pricer)

	prices, err := cached.Prices(t.Context(), ETH, OMNI)
	require.NoError(t, err)
	require.InEpsilon(t, 100.0, prices[ETH], epsilon)
	require.InEpsilon(t, 200.0, prices[OMNI], epsilon)

	// change prices
	pricer.SetPrice(ETH, 150)
	pricer.SetPrice(OMNI, 250)

	// prices should still be cached
	prices, err = cached.Prices(t.Context(), ETH, OMNI)
	require.NoError(t, err)
	require.InEpsilon(t, 100.0, prices[ETH], epsilon)
	require.InEpsilon(t, 200.0, prices[OMNI], epsilon)

	// clear cache
	cached.ClearCache()

	// prices should be updated
	prices, err = cached.Prices(t.Context(), ETH, OMNI)
	require.NoError(t, err)
	require.InEpsilon(t, 150.0, prices[ETH], epsilon)
	require.InEpsilon(t, 250.0, prices[OMNI], epsilon)
}
