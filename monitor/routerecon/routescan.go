package routerecon

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
)

const (
	baseURL    = "https://api.routescan.io"
	crossTxURL = "/v2/network/testnet/evm/cross-transactions"
)

func paginateLatestCrossTx(ctx context.Context, filter queryFilter) (crossTxJSON, error) {
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

func queryLatestCrossTx(ctx context.Context, filter queryFilter, next string) (crossTxJSON, string, error) {
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
		if filter.HasStream() {
			q.Add("srcChainIds", routeScanChainID(filter.Stream.SourceChainID))
			q.Add("dstChainIds", routeScanChainID(filter.Stream.DestChainID))
		}
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

	if len(crossTxResp.Items) == 0 {
		return crossTxJSON{}, "", errors.New("empty response")
	} else if len(crossTxResp.Items) > limit {
		return crossTxJSON{}, "", errors.New("too many items in response")
	}

	for _, item := range crossTxResp.Items {
		// Validate server side filtering
		if item.Type != "omni" {
			return crossTxJSON{}, "", errors.New("invalid cross tx type")
		}

		msgID, err := item.MsgID()
		if err != nil {
			return crossTxJSON{}, "", errors.Wrap(err, "parse msg id")
		}

		if filter.HasStream() && (filter.Stream.DestChainID != msgID.DestChainID || filter.Stream.SourceChainID != msgID.SourceChainID) {
			return crossTxJSON{}, "", errors.New("invalid dest or source chain", "filter", filter.Stream, "msg", msgID)
		}

		// Client side filtering
		if filter.HasStream() && filter.Stream != msgID.StreamID {
			continue
		} else if filter.Pending && item.Status != "pending" {
			continue
		} else if !filter.Pending && item.Status != "completed" {
			continue
		}

		return item, "", nil // Return found item
	}

	// No matching item found

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
