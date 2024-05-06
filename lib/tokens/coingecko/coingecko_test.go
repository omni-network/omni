package coingecko_test

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/coingecko"

	"github.com/stretchr/testify/require"
)

func TestGetPrice(t *testing.T) {
	t.Parallel()

	// map token id -> currency -> price
	// set during request handler
	testPrices := make(map[string]map[string]float64)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/api/v3/simple/price", r.URL.Path)

		q := r.URL.Query()
		ids := strings.Split(q.Get("ids"), ",")
		currencies := strings.Split(q.Get("vs_currencies"), ",")

		resp := make(map[string]map[string]float64)
		for _, id := range ids {
			resp[id] = make(map[string]float64)

			if _, ok := testPrices[id]; !ok {
				testPrices[id] = make(map[string]float64)
			}

			for _, currency := range currencies {
				resp[id][currency] = randPrice()

				// also store the price, so we can assert against it
				testPrices[id][currency] = resp[id][currency]
			}
		}

		bz, _ := json.Marshal(resp)
		_, _ = w.Write(bz)
	}))

	defer ts.Close()

	c := coingecko.New(coingecko.WithHost(ts.URL))
	prices, err := c.Price(context.Background(), tokens.OMNI, tokens.ETH)
	require.NoError(t, err)
	require.InEpsilon(t, prices[tokens.OMNI], testPrices[tokens.OMNI.CoingeckoID()]["usd"], 0.01)
	require.InEpsilon(t, prices[tokens.ETH], testPrices[tokens.ETH.CoingeckoID()]["usd"], 0.01)
}

func randPrice() float64 {
	return float64(int(rand.Float64()*10000)) / 100
}
