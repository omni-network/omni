package coingecko_test

import (
	"encoding/json"
	"flag"
	"math/big"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/tokenpricer/coingecko"
	"github.com/omni-network/omni/lib/tokens"
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
	require.False(t, ok)

	c := coingecko.New(coingecko.WithAPIKey(apikey))

	// USD prices for omni and eth
	usdPrices, err := c.USDPrices(t.Context(), tokens.OMNI, tokens.ETH)
	tutil.RequireNoError(t, err)
	require.NotEmpty(t, usdPrices)

	// eth price in omni
	price1, err := c.Price(t.Context(), tokens.ETH, tokens.OMNI)
	tutil.RequireNoError(t, err)
	t.Logf("ETH/OMNI: %v", price1)

	// omni price in eth
	price2, err := c.Price(t.Context(), tokens.OMNI, tokens.ETH)
	tutil.RequireNoError(t, err)
	t.Logf("OMNI/ETH: %v", price2)

	// alternate way to get omni price in eth (since eth is a supported currency)
	prices, err := c.GetPrice(t.Context(), "eth", tokens.OMNI)
	tutil.RequireNoError(t, err)
	require.NotEmpty(t, prices)
	price3 := prices[tokens.OMNI]
	t.Logf("OMNI/eth: %v", price3)

	require.Equal(t, new(big.Rat).Inv(price2), price1)
	require.Equal(t, price2, price3)

	// Test NOM price derivation (should be OMNI / 75)
	nomPrices, err := c.USDPrices(t.Context(), tokens.NOM, tokens.OMNI)
	tutil.RequireNoError(t, err)
	require.NotEmpty(t, nomPrices)

	expectedNOMPrice := nomPrices[tokens.OMNI] / 75.0
	require.InEpsilon(t, expectedNOMPrice, nomPrices[tokens.NOM], 0.001, "NOM price should be OMNI/75")

	// Test that OMNI pricing still works independently
	omniOnlyPrices, err := c.USDPrices(t.Context(), tokens.OMNI)
	tutil.RequireNoError(t, err)
	require.NotEmpty(t, omniOnlyPrices)
	require.InEpsilon(t, nomPrices[tokens.OMNI], omniOnlyPrices[tokens.OMNI], 0.001, "OMNI price should be consistent")
}

type testCase struct {
	name         string
	invalid      bool         // invalid response
	empty        bool         // empty response
	omitToken    tokens.Asset // omit a requested token
	renameToken  tokens.Asset // rename a requested token
	omitCurrency string       // omit a requested currency
	zeros        bool         // include zero prices
	negatives    bool         // include negative prices
}

func TestGetUSDPrice(t *testing.T) {
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
			test.omitToken != tokens.Asset{} ||
			test.renameToken != tokens.Asset{} ||
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
			prices, err := c.USDPrices(t.Context(), tokens.OMNI, tokens.ETH)

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

func TestNOMDerivation(t *testing.T) {
	t.Parallel()

	// Create a mock server that only serves OMNI price
	omniPrice := 10.0 // $10 OMNI
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]map[string]float64{
			"omni-network": {
				"usd": omniPrice,
			},
		}
		bz, _ := json.Marshal(resp)
		_, _ = w.Write(bz)
	}))
	defer server.Close()

	c := coingecko.New(coingecko.WithHost(server.URL))

	// Test NOM-only request
	prices, err := c.USDPrices(t.Context(), tokens.NOM)
	require.NoError(t, err)

	expectedNOMPrice := omniPrice / 75.0
	require.InEpsilon(t, expectedNOMPrice, prices[tokens.NOM], 0.001)

	// Test mixed request (NOM + OMNI)
	mixedPrices, err := c.USDPrices(t.Context(), tokens.NOM, tokens.OMNI)
	require.NoError(t, err)

	require.InEpsilon(t, omniPrice, mixedPrices[tokens.OMNI], 0.001)
	require.InEpsilon(t, expectedNOMPrice, mixedPrices[tokens.NOM], 0.001)

	// Test that NOM price is exactly OMNI/75
	require.InEpsilon(t, mixedPrices[tokens.OMNI]/75.0, mixedPrices[tokens.NOM], 0.001)

	// Test that requesting OMNI alone still works (no NOM derivation needed)
	omniOnly, err := c.USDPrices(t.Context(), tokens.OMNI)
	require.NoError(t, err)
	require.InEpsilon(t, omniPrice, omniOnly[tokens.OMNI], 0.001)
}

func randPrice() float64 {
	return float64(int((rand.Float64()+0.01)*10000)) / 100
}
