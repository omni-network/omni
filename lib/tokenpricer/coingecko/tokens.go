package coingecko

import (
	"context"
	"net/url"

	"github.com/omni-network/omni/lib/errors"
)

type Coin struct {
	ID        string            `json:"id"`
	Symbol    string            `json:"symbol"`
	Name      string            `json:"name"`
	Platforms map[string]string `json:"platforms"` // Addresses by asset-platform
}

// AllCoins returns a list of all supported active coins.
func (c Client) AllCoins(ctx context.Context) ([]Coin, error) {
	params := url.Values{
		"include_platform": {"true"},
	}
	var resp []Coin
	if err := c.doReq(ctx, endpointCoinsList, params, &resp); err != nil {
		return nil, errors.Wrap(err, "do req", "endpoint", "list_tokens")
	}

	return resp, nil
}
