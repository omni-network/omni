package coingecko

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
)

const (
	endpointSimplePrice = "/api/v3/simple/price"
	defaultProdHost     = "https://api.coingecko.com"
	proProdHost         = "https://pro-api.coingecko.com"
	apikeyHeader        = "x-cg-pro-api-key" //nolint:gosec // This is the header
	currencyUSD         = "usd"
)

type Client struct {
	host   string
	apikey string
}

var _ tokenpricer.Pricer = Client{}

// New creates a new coingecko Client with the given options.
func New(opts ...func(*options)) Client {
	o := defaultOpts()
	for _, opt := range opts {
		opt(&o)
	}

	return Client{
		host:   o.Host,
		apikey: o.APIKey,
	}
}

// Price returns the price of the base asset denominated in the quote asset.
func (c Client) Price(ctx context.Context, base, quote tokens.Asset) (float64, error) {
	// Coingecko only supports a limited amount of "quote currencies",
	// So convert to USD first, then to the quote currency.

	m, err := c.getPrice(ctx, currencyUSD, base)
	if err != nil {
		return 0, errors.Wrap(err, "get price")
	}
	basePrice := m[base]

	m, err = c.getPrice(ctx, currencyUSD, quote)
	if err != nil {
		return 0, errors.Wrap(err, "get price")
	}
	quotePrice := m[quote]

	return basePrice / quotePrice, nil
}

// USDPrice returns the price of the token in USD.
func (c Client) USDPrice(ctx context.Context, tkn tokens.Asset) (float64, error) {
	prices, err := c.USDPrices(ctx, tkn)
	if err != nil {
		return 0, err
	}

	return prices[tkn], nil
}

// USDPrices returns the price of each coin in USD.
func (c Client) USDPrices(ctx context.Context, tkns ...tokens.Asset) (map[tokens.Asset]float64, error) {
	return c.getPrice(ctx, currencyUSD, tkns...)
}

// simplePriceResponse is the response from the simple/price endpoint.
// It maps coin id to currency to price.
type simplePriceResponse map[string]map[string]float64

// GetPrice returns the price of each coin in the given currency.
// See supported currencies: https://api.coingecko.com/api/v3/simple/supported_vs_currencies
func (c Client) getPrice(ctx context.Context, currency string, tkns ...tokens.Asset) (map[tokens.Asset]float64, error) {
	ids := make([]string, len(tkns))
	for i, t := range tkns {
		ids[i] = t.CoingeckoID
	}

	params := url.Values{
		"ids":           {strings.Join(ids, ",")},
		"vs_currencies": {currency},
	}

	var resp simplePriceResponse
	if err := c.doReq(ctx, endpointSimplePrice, params, &resp); err != nil {
		return nil, errors.Wrap(err, "do req", "endpoint", "get_price")
	}

	prices := make(map[tokens.Asset]float64)

	for _, tkn := range tkns {
		priceByCurrency, ok := resp[tkn.CoingeckoID]
		if !ok {
			return nil, errors.New("missing token in response", "token", tkn)
		}

		price, ok := priceByCurrency[currency]
		if !ok {
			return nil, errors.New("missing price in response", "token", tkn, "currency", currency)
		}

		if price <= 0 {
			return nil, errors.New("invalid price in response", "token", tkn, "price", price)
		}

		prices[tkn] = price
	}

	return prices, nil
}

// doReq makes a GET request to the given path & params, and decodes the response into response.
func (c Client) doReq(ctx context.Context, path string, params url.Values, response any) error {
	uri, err := c.uri(path, params)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), nil)
	if err != nil {
		return errors.Wrap(err, "create request", "url", uri.String())
	}

	if c.apikey != "" {
		req.Header.Set(apikeyHeader, c.apikey) //nolint:canonicalheader // As per CoinGacko docs
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "get req", "url", uri.String())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("non-200 response", "url", uri.String(), "status", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return errors.Wrap(err, "decode response")
	}

	return nil
}

func (c Client) uri(path string, params url.Values) (*url.URL, error) {
	uri, err := url.Parse(c.host + path + "?" + params.Encode())
	if err != nil {
		return nil, errors.Wrap(err, "parse url", "host", c.host, "path", path, "params", params.Encode())
	}

	return uri, nil
}
