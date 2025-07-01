package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tracer"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

const (
	endpointSimplePrice = "/api/v3/simple/price"
	endpointCoinsList   = "/api/v3/coins/list"
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
// Note that for canonical solver prices, base=deposit and quote=expense.
func (c Client) Price(ctx context.Context, base, quote tokens.Asset) (*big.Rat, error) {
	// Coingecko only supports a limited amount of "quote currencies",
	// So convert to USD first, then to the quote currency.

	var basePrice, quotePrice *big.Rat

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		m, err := c.getPrice(ctx, currencyUSD, base)
		if err != nil {
			return errors.Wrap(err, "get price")
		}
		basePrice = m[base]

		return nil
	})
	eg.Go(func() error {
		m, err := c.getPrice(ctx, currencyUSD, quote)
		if err != nil {
			return errors.Wrap(err, "get price")
		}
		quotePrice = m[quote]

		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, errors.Wrap(err, "get price")
	}

	return new(big.Rat).Quo(basePrice, quotePrice), nil
}

// USDPrice returns the price of the asset in USD.
func (c Client) USDPrice(ctx context.Context, asset tokens.Asset) (float64, error) {
	prices, err := c.USDPrices(ctx, asset)
	if err != nil {
		return 0, err
	}

	return prices[asset], nil
}

// USDPrices returns the price of each asset in USD.
func (c Client) USDPrices(ctx context.Context, assets ...tokens.Asset) (map[tokens.Asset]float64, error) {
	rats, err := c.getPrice(ctx, currencyUSD, assets...)
	if err != nil {
		return nil, errors.Wrap(err, "get price")
	}

	prices := make(map[tokens.Asset]float64)
	for asset, price := range rats {
		f, _ := price.Float64()
		prices[asset] = f
	}

	return prices, nil
}

// USDPricesRat returns the price of each asset in USD.
func (c Client) USDPricesRat(ctx context.Context, assets ...tokens.Asset) (map[tokens.Asset]*big.Rat, error) {
	return c.getPrice(ctx, currencyUSD, assets...)
}

// USDPriceRat returns the price of each asset in USD.
func (c Client) USDPriceRat(ctx context.Context, asset tokens.Asset) (*big.Rat, error) {
	rats, err := c.USDPricesRat(ctx, asset)
	if err != nil {
		return nil, errors.Wrap(err, "get price")
	}

	return rats[asset], nil
}

// map[asset.CoingeckoID]map[currency]price.
type simplePriceResponse map[string]map[string]json.Number

// GetPrice returns the price of each asset in the given currency.
// See supported currencies: https://api.coingecko.com/api/v3/simple/supported_vs_currencies
func (c Client) getPrice(ctx context.Context, currency string, assets ...tokens.Asset) (map[tokens.Asset]*big.Rat, error) {
	ctx, span := tracer.Start(ctx, "coingecko/price", trace.WithAttributes(
		attribute.String("currency", currency),
		attribute.String("assets", fmt.Sprint(assets)),
	))
	defer span.End()

	ids := make([]string, len(assets))
	for i, t := range assets {
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

	prices := make(map[tokens.Asset]*big.Rat)

	for _, asset := range assets {
		priceByCurrency, ok := resp[asset.CoingeckoID]
		if !ok {
			return nil, errors.New("missing asset in response", "asset", asset)
		}

		priceStr, ok := priceByCurrency[currency]
		if !ok {
			return nil, errors.New("missing price in response", "asset", asset, "currency", currency)
		}

		// Parsing the json.Number as big.Rat floating point number
		// This avoid going through float64 which can lose precision
		price, ok := new(big.Rat).SetString(priceStr.String()) //nolint:gosec // CVE-2022-23772 fixed in go1.17
		if !ok {
			return nil, errors.New("invalid price string", "asset", asset, "price", priceStr)
		} else if price.Sign() <= 0 {
			return nil, errors.New("proic enot positive", "asset", asset, "price", price)
		}

		prices[asset] = price
	}

	return prices, nil
}

// doReq makes a GET request to the given path & params, and decodes the response into response.
func (c Client) doReq(ctx context.Context, path string, params url.Values, response any) error {
	defer func(t0 time.Time) {
		latency.WithLabelValues(path).Observe(time.Since(t0).Seconds())
	}(time.Now())

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
	uri, err := url.Parse(c.host)
	if err != nil {
		return nil, errors.Wrap(err, "parse host", "host", c.host)
	}
	uri = uri.JoinPath(path)
	uri.RawQuery = params.Encode()

	return uri, nil
}
