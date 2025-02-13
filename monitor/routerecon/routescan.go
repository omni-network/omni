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
	"github.com/omni-network/omni/lib/netconf"
)

const (
	baseURL                         = "https://api.routescan.io"
	crossTxURL                      = "/v2/network/%s/evm/cross-transactions"
	requestsPerDayLimitHeader       = "X-Ratelimit-Rpd-Limit"
	requestsPerMinuteLimitHeader    = "X-Ratelimit-Rpm-Limit"
	increasedRequestsPerDayLimit    = 200000
	increasedRequestsPerMinuteLimit = 600
)

func getCrossTxURL(network netconf.ID) string {
	net := "mainnet"
	if network == netconf.Omega {
		net = "testnet"
	}

	return fmt.Sprintf(crossTxURL, net)
}

func paginateLatestCrossTx(ctx context.Context, network netconf.ID, routeScanAPIKey string, filter filter) (crossTxJSON, error) {
	var (
		resp crossTxJSON
		next string
		err  error
	)
	for {
		resp, next, _, err = queryLatestCrossTx(ctx, network, routeScanAPIKey, filter, next)
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

func queryLatestCrossTx(ctx context.Context, network netconf.ID, routeScanAPIKey string, filter filter, next string) (crossTxJSON, string, bool, error) {
	url := baseURL + next
	if next == "" {
		// Build initial path
		url += getCrossTxURL(network)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return crossTxJSON{}, "", false, errors.Wrap(err, "new request")
	}

	const limit = 100

	if next == "" {
		// Build initial query params (server side filtering)
		q := req.URL.Query()
		q.Add("types", "omni")
		q.Add("limit", strconv.FormatUint(limit, 10))
		if routeScanAPIKey != "" {
			q.Add("apikey", routeScanAPIKey)
		}
		filter.QueryParams(q)
		req.URL.RawQuery = q.Encode()
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return crossTxJSON{}, "", false, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	requestsPerDayLimit, err := strconv.Atoi(resp.Header.Get(requestsPerDayLimitHeader))
	if err != nil {
		return crossTxJSON{}, "", false, errors.Wrap(err, "rate limit header", "header", requestsPerDayLimitHeader)
	}

	requestsPerMinuteLimit, err := strconv.Atoi(resp.Header.Get(requestsPerMinuteLimitHeader))
	if err != nil {
		return crossTxJSON{}, "", false, errors.Wrap(err, "rate limit header", "header", requestsPerMinuteLimitHeader)
	}

	isWithIncreasedRateLimit := requestsPerDayLimit >= increasedRequestsPerDayLimit && requestsPerMinuteLimit >= increasedRequestsPerMinuteLimit

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return crossTxJSON{}, "", isWithIncreasedRateLimit, errors.Wrap(err, "read response body")
	}

	if resp.StatusCode/http.StatusOK != 1 { // Checking for 2xx status code
		var errJSON errorJSON
		_ = json.Unmarshal(bz, &errJSON)

		return crossTxJSON{}, "", isWithIncreasedRateLimit, errors.New("bad response", "status", resp.Status, "err_code", errJSON.Code, "err_msg", errJSON.Message)
	}

	var crossTxResp crossTxResponse
	if err := json.Unmarshal(bz, &crossTxResp); err != nil {
		return crossTxJSON{}, "", isWithIncreasedRateLimit, errors.Wrap(err, "decode response")
	}

	if len(crossTxResp.CrossTxs) == 0 {
		return crossTxJSON{}, "", isWithIncreasedRateLimit, errors.New("empty response")
	} else if len(crossTxResp.CrossTxs) > limit {
		return crossTxJSON{}, "", isWithIncreasedRateLimit, errors.New("too many items in response")
	}

	for _, crossTx := range crossTxResp.CrossTxs {
		if err := crossTx.Verify(); err != nil {
			return crossTxJSON{}, "", isWithIncreasedRateLimit, errors.Wrap(err, "verify cross tx")
		}

		if ok, err := filter.Match(crossTx); err != nil {
			return crossTxJSON{}, "", isWithIncreasedRateLimit, errors.Wrap(err, "match filter")
		} else if !ok {
			continue
		}

		return crossTx, "", isWithIncreasedRateLimit, nil // Return found crossTx
	}

	// No matching crossTx found

	if crossTxResp.Links.Next == "" {
		return crossTxJSON{}, "", isWithIncreasedRateLimit, errors.New("no matching cross tx found")
	}

	return crossTxJSON{}, crossTxResp.Links.Next, isWithIncreasedRateLimit, nil // Return next page to query
}

const omegaResets = 4

func routeScanChainID(id uint64) string {
	if id == evmchain.IDOmniOmega {
		return fmt.Sprintf("%d_%d", id, omegaResets)
	}

	return strconv.FormatUint(id, 10)
}
