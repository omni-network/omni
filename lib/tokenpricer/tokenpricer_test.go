package tokenpricer_test

import (
	"testing"

	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/stretchr/testify/require"
)

func TestCachedPricer(t *testing.T) {
	t.Parallel()
	const epsilon = 1e-6

	ETH := tokens.ETH
	NOM := tokens.NOM

	pricer := tokenpricer.NewUSDMock(map[tokens.Asset]float64{
		ETH: 100,
		NOM: 200,
	})

	cached := tokenpricer.NewCached(pricer)

	prices, err := cached.USDPrices(t.Context(), ETH, NOM)
	require.NoError(t, err)
	require.InEpsilon(t, 100.0, prices[ETH], epsilon)
	require.InEpsilon(t, 200.0, prices[NOM], epsilon)

	// change prices
	pricer.SetUSDPrice(ETH, 150)
	pricer.SetUSDPrice(NOM, 250)

	// prices should still be cached
	prices, err = cached.USDPrices(t.Context(), ETH, NOM)
	require.NoError(t, err)
	require.InEpsilon(t, 100.0, prices[ETH], epsilon)
	require.InEpsilon(t, 200.0, prices[NOM], epsilon)

	// clear cache
	cached.ClearCache()

	// prices should be updated
	prices, err = cached.USDPrices(t.Context(), ETH, NOM)
	require.NoError(t, err)
	require.InEpsilon(t, 150.0, prices[ETH], epsilon)
	require.InEpsilon(t, 250.0, prices[NOM], epsilon)
}
