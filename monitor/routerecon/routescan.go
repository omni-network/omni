package routerecon

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
)

const (
	baseURL    = "https://api.routescan.io"
	crossTxURL = "/v2/network/testnet/evm/cross-transactions"
)

func paginateLatestCrossTx(ctx context.Context, filter filter) (crossTxJSON, error) {
	var (
		resp crossTxJSON
		next string
		err  error
	)
	for {
		resp, next, err = queryLatestCrossTx(ctx, filter, next)
		if err != nil {
			return crossTxJSON{}, errors.Wrap(err, "query latest cross tx")
		} else if next != "" {
			// Query next page
			continue
		}

		return resp, nil
	}
}

type filter interface {
	// QueryParams sets any server-side query param filters.
	QueryParams(q url.Values)
	// Match returns true if the item should be returned.
	Match(tx crossTxJSON) (bool, error)
}

func queryLatestCrossTx(ctx context.Context, filter filter, next string) (crossTxJSON, string, error) {
	url := baseURL + next
	if next == "" {
		// Build initial path
		url += crossTxURL
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return crossTxJSON{}, "", errors.Wrap(err, "new request")
	}

	const limit = 100

	if next == "" {
		// Build initial query params (server side filtering)
		q := req.URL.Query()
		q.Add("types", "omni")
		q.Add("limit", strconv.FormatUint(limit, 10))
		filter.QueryParams(q)
		req.URL.RawQuery = q.Encode()
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return crossTxJSON{}, "", errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return crossTxJSON{}, "", errors.Wrap(err, "read response body")
	}

	var crossTxResp crossTxResponse
	if err := json.Unmarshal(bz, &crossTxResp); err != nil {
		return crossTxJSON{}, "", errors.Wrap(err, "decode response")
	}

	if len(crossTxResp.CrossTxs) == 0 {
		return crossTxJSON{}, "", errors.New("empty response")
	} else if len(crossTxResp.CrossTxs) > limit {
		return crossTxJSON{}, "", errors.New("too many items in response")
	}

	for _, crossTx := range crossTxResp.CrossTxs {
		if err := crossTx.Verify(); err != nil {
			return crossTxJSON{}, "", errors.Wrap(err, "verify cross tx")
		}

		if ok, err := filter.Match(crossTx); err != nil {
			return crossTxJSON{}, "", errors.Wrap(err, "match filter")
		} else if !ok {
			continue
		}

		return crossTx, "", nil // Return found crossTx
	}

	// No matching crossTx found

	if crossTxResp.Links.Next == "" {
		return crossTxJSON{}, "", errors.New("no matching cross tx found")
	}

	return crossTxJSON{}, crossTxResp.Links.Next, nil // Return next page to query
}

const omegaResets = 4

func routeScanChainID(id uint64) string {
	if id == evmchain.IDOmniOmega {
		return fmt.Sprintf("%d_%d", id, omegaResets)
	}

	return strconv.FormatUint(id, 10)
}
