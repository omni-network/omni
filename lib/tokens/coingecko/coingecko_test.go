package coingecko_test

import (
	"context"
	"encoding/json"
	"flag"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/coingecko"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "run integration tests")

func TestIntegration(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	apikey, ok := os.LookupEnv("COINGECKO_APIKEY")
	require.True(t, ok)

	c := coingecko.New(coingecko.WithAPIKey(apikey))
	prices, err := c.Prices(context.Background(), tokens.OMNI, tokens.ETH)
	tutil.RequireNoError(t, err)
	require.NotEmpty(t, prices)
}

type testCase struct {
	name         string
	invalid      bool         // invalid response
	empty        bool         // empty response
	omitToken    tokens.Token // omit a requested token
	renameToken  tokens.Token // rename a requested token
	omitCurrency string       // omit a requested currency
	zeros        bool         // include zero prices
	negatives    bool         // include negative prices
}

func TestGetPrice(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		{name: "success"},
		{name: "empty", empty: true},
		{name: "omit eth", omitToken: tokens.ETH},
		{name: "rename eth", renameToken: tokens.ETH},
		{name: "omit omni", omitToken: tokens.OMNI},
		{name: "rename omni", renameToken: tokens.OMNI},
		{name: "omit usd", omitCurrency: "usd"},
		{name: "zeros", zeros: true},
		{name: "negatives", negatives: true},
	}

	shouldErr := func(t *testing.T, test testCase) bool {
		t.Helper()
		return (test.invalid ||
			test.empty ||
			test.omitToken != tokens.Token{} ||
			test.renameToken != tokens.Token{} ||
			test.omitCurrency != "" ||
			test.zeros ||
			test.negatives)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			server, servedPrices, token := makeTestServer(t, test)
			defer server.Close()

			c := coingecko.New(coingecko.WithHost(server.URL), coingecko.WithAPIKey(token))
			prices, err := c.Prices(context.Background(), tokens.OMNI, tokens.ETH)

			if shouldErr(t, test) {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.InEpsilon(t, prices[tokens.OMNI], servedPrices[tokens.OMNI.CoingeckoID]["usd"], 0.01)
			require.InEpsilon(t, prices[tokens.ETH], servedPrices[tokens.ETH.CoingeckoID]["usd"], 0.01)
		})
	}
}

func makeTestServer(t *testing.T, test testCase) (*httptest.Server, map[string]map[string]float64, string) {
	t.Helper()

	// set during request handler
	servedPrices := make(map[string]map[string]float64)

	apikey := tutil.RandomHash().Hex()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/api/v3/simple/price", r.URL.Path)
		require.Equal(t, "GET", r.Method)
		require.Equal(t, apikey, r.Header.Get(coingecko.GetAPIKeyHeader()))

		resp := make(map[string]map[string]float64)

		if test.invalid {
			_, _ = w.Write([]byte("invalid json"))
			return
		}

		if test.empty {
			bz, err := json.Marshal(resp)
			require.NoError(t, err)
			_, _ = w.Write(bz)

			return
		}

		q := r.URL.Query()
		ids := strings.Split(q.Get("ids"), ",")
		currencies := strings.Split(q.Get("vs_currencies"), ",")

		for _, id := range ids {
			if id == test.omitToken.CoingeckoID {
				continue
			}

			if id == test.renameToken.CoingeckoID {
				id = "renamed"
			}

			resp[id] = make(map[string]float64)

			if _, ok := servedPrices[id]; !ok {
				servedPrices[id] = make(map[string]float64)
			}

			for _, currency := range currencies {
				if currency == test.omitCurrency {
					continue
				}

				price := randPrice()

				if test.zeros {
					price = 0
				}

				if test.negatives {
					price = -price
				}

				resp[id][currency] = price

				// also store the price, so we can assert against it
				servedPrices[id][currency] = resp[id][currency]
			}
		}

		bz, _ := json.Marshal(resp)
		_, _ = w.Write(bz)
	}))

	return server, servedPrices, apikey
}

func randPrice() float64 {
	return float64(int((rand.Float64()+0.01)*10000)) / 100
}
