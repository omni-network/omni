package coingecko_test

import (
	"encoding/json"
	"flag"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/tokenmeta"
	"github.com/omni-network/omni/lib/tokenpricer/coingecko"
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
	prices, err := c.Prices(t.Context(), tokenmeta.OMNI, tokenmeta.ETH)
	tutil.RequireNoError(t, err)
	require.NotEmpty(t, prices)
}

type testCase struct {
	name         string
	invalid      bool           // invalid response
	empty        bool           // empty response
	omitToken    tokenmeta.Meta // omit a requested token
	renameToken  tokenmeta.Meta // rename a requested token
	omitCurrency string         // omit a requested currency
	zeros        bool           // include zero prices
	negatives    bool           // include negative prices
}

func TestGetPrice(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		{name: "success"},
		{name: "empty", empty: true},
		{name: "omit eth", omitToken: tokenmeta.ETH},
		{name: "rename eth", renameToken: tokenmeta.ETH},
		{name: "omit omni", omitToken: tokenmeta.OMNI},
		{name: "rename omni", renameToken: tokenmeta.OMNI},
		{name: "omit usd", omitCurrency: "usd"},
		{name: "zeros", zeros: true},
		{name: "negatives", negatives: true},
	}

	shouldErr := func(t *testing.T, test testCase) bool {
		t.Helper()
		return (test.invalid ||
			test.empty ||
			test.omitToken != tokenmeta.Meta{} ||
			test.renameToken != tokenmeta.Meta{} ||
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
			prices, err := c.Prices(t.Context(), tokenmeta.OMNI, tokenmeta.ETH)

			if shouldErr(t, test) {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.InEpsilon(t, prices[tokenmeta.OMNI], servedPrices[tokenmeta.OMNI.CoingeckoID]["usd"], 0.01)
			require.InEpsilon(t, prices[tokenmeta.ETH], servedPrices[tokenmeta.ETH.CoingeckoID]["usd"], 0.01)
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
