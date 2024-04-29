package coingecko

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/omni-network/omni/lib/errors"
)

type Currency string

const (
	USD Currency = "usd"

	endpointSimplePrice = "/api/v3/simple/price"
	prodHost            = "https://api.coingecko.com"
)

type Client struct {
	host string
}

// New creates a new goingecko Client with the given options.
func New(opts ...func(*options)) Client {
	o := defaultOpts()
	for _, opt := range opts {
		opt(&o)
	}

	return Client{
		host: o.Host,
	}
}

// GetPriceUSD returns the price of each coin in USD.
func (c Client) GetPriceUSD(ctx context.Context, ids ...string) (map[string]float64, error) {
	return c.GetPrice(ctx, USD, ids...)
}

// simplePriceResponse is the response from the simple/price endpoint.
// It mapes coin id to currency to price.
type simplePriceResponse map[string]map[string]float64

// GetPrice returns the price of each coin in the given currency.
func (c Client) GetPrice(ctx context.Context, currency Currency, ids ...string) (map[string]float64, error) {
	params := url.Values{
		"ids":           {strings.Join(ids, ",")},
		"vs_currencies": {string(currency)},
	}

	var resp simplePriceResponse
	if err := c.doReq(ctx, endpointSimplePrice, params, &resp); err != nil {
		return nil, err
	}

	prices := make(map[string]float64)
	for id, price := range resp {
		prices[id] = price[string(currency)]
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

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "get", "url", uri.String())
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("get", "url", uri.String(), "status", resp.Status)
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
