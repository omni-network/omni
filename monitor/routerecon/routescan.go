package routerecon

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	baseURL    = "https://api.routescan.io"
	crossTxURL = "/v2/network/testnet/evm/cross-transactions?types=omni&limit=100"
)

type chainID string

func (c chainID) ID() (uint64, error) {
	if strings.Contains(string(c), "_") {
		var resp, dummy uint64
		_, err := fmt.Sscanf(string(c), "%d_%d", &resp, &dummy)
		if err != nil {
			return 0, errors.Wrap(err, "parse chain id")
		}

		return resp, nil
	}

	resp, err := strconv.ParseUint(string(c), 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "parse chain id")
	}

	return resp, nil
}

func (c *chainID) UnmarshalJSON(bz []byte) error {
	var str string
	if err := json.Unmarshal(bz, &str); err == nil {
		*c = chainID(str)
		return nil
	}

	var i uint64
	if err := json.Unmarshal(bz, &i); err != nil {
		return errors.Wrap(err, "unmarshal chain id")
	}

	*c = chainID(strconv.FormatUint(i, 10))

	return nil
}

type crossTxJSON struct {
	ID             string    `json:"id"`
	Type           string    `json:"type"`
	Status         string    `json:"status"`
	SrcChainID     chainID   `json:"srcChainId"`
	SrcTimestamp   time.Time `json:"srcTimestamp"`
	SrcTxHash      string    `json:"srcTxHash"`
	SrcBlockNumber uint64    `json:"srcBlockNumber"`
	SrcBlockHash   string    `json:"srcBlockHash"`
	SrcGasLimit    string    `json:"srcGasLimit"`
	DstChainID     chainID   `json:"dstChainId"`
	DstTimestamp   time.Time `json:"dstTimestamp"`
	DstTxHash      string    `json:"dstTxHash"`
	DstBlockNumber uint64    `json:"dstBlockNumber"`
	DstBlockHash   string    `json:"dstBlockHash"`
	DstGasLimit    string    `json:"dstGasLimit"`
	From           string    `json:"from"`
	To             string    `json:"to"`
	Data           struct {
		Relayer           string        `json:"relayer"`
		GasUsed           uint64        `json:"gasUsed,string"`
		Error             hexutil.Bytes `json:"error"`
		Offset            uint64        `json:"offset,string"`
		ShardID           uint64        `json:"shardId,string"`
		ConfirmationLevel string        `json:"confirmationLevel"`
		GasLimit          uint64        `json:"gasLimit,string"`
		Fees              string        `json:"fees"`
	} `json:"data"`
}

func (c crossTxJSON) MsgID() (xchain.MsgID, error) {
	srcChainID, err := c.SrcChainID.ID()
	if err != nil {
		return xchain.MsgID{}, errors.Wrap(err, "parse src chain id")
	}

	dstChainID, err := c.DstChainID.ID()
	if err != nil {
		return xchain.MsgID{}, errors.Wrap(err, "parse dst chain id")
	}

	return xchain.MsgID{
		StreamID: xchain.StreamID{
			SourceChainID: srcChainID,
			DestChainID:   dstChainID,
			ShardID:       xchain.ShardID(c.Data.ShardID),
		},
		StreamOffset: c.Data.Offset,
	}, nil
}

type crossTxResponse struct {
	Items []crossTxJSON `json:"items"`
	Links links         `json:"link"`
}

type links struct {
	Next      string `json:"next"`
	NextToken string `json:"nextToken"`
}

func paginateLatestCrossTx(ctx context.Context) (crossTxJSON, error) {
	next := crossTxURL
	for {
		var resp crossTxJSON
		var err error
		resp, next, err = queryLatestCrossTx(ctx, next)
		if err != nil {
			return crossTxJSON{}, errors.Wrap(err, "query latest cross tx")
		} else if next == "" {
			return resp, nil
		}
	}
}

func queryLatestCrossTx(ctx context.Context, next string) (crossTxJSON, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+next, nil)
	if err != nil {
		return crossTxJSON{}, "", errors.Wrap(err, "new request")
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
	}

	for _, item := range crossTxResp.Items {
		if item.Type != "omni" {
			return crossTxJSON{}, "", errors.New("invalid cross tx type")
		}

		if item.Status == "completed" {
			return item, "", nil
		}
	}

	if crossTxResp.Links.Next == "" {
		return crossTxJSON{}, "", errors.New("no more cross tx")
	}

	return crossTxJSON{}, crossTxResp.Links.Next, nil
}
