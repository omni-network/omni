package tokens_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/tokens"

	"github.com/stretchr/testify/require"
)

func TestCachedPricer(t *testing.T) {
	t.Parallel()
	const epsilon = 1e-6

	pricer := tokens.NewMockPricer(map[tokens.Token]float64{
		tokens.ETH:  100,
		tokens.OMNI: 200,
	})

	cached := tokens.NewCachedPricer(pricer)

	prices, err := cached.Price(context.Background(), tokens.ETH, tokens.OMNI)
	require.NoError(t, err)
	require.InEpsilon(t, 100.0, prices[tokens.ETH], epsilon)
	require.InEpsilon(t, 200.0, prices[tokens.OMNI], epsilon)

	// change prices
	pricer.SetPrice(tokens.ETH, 150)
	pricer.SetPrice(tokens.OMNI, 250)

	// prices should still be cached
	prices, err = cached.Price(context.Background(), tokens.ETH, tokens.OMNI)
	require.NoError(t, err)
	require.InEpsilon(t, 100.0, prices[tokens.ETH], epsilon)
	require.InEpsilon(t, 200.0, prices[tokens.OMNI], epsilon)

	// clear cache
	cached.ClearCache()

	// prices should be updated
	prices, err = cached.Price(context.Background(), tokens.ETH, tokens.OMNI)
	require.NoError(t, err)
	require.InEpsilon(t, 150.0, prices[tokens.ETH], epsilon)
	require.InEpsilon(t, 250.0, prices[tokens.OMNI], epsilon)
}
