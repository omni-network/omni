package routerecon

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	ID             string         `json:"id"` // sha256("omni#{srcChainId}#{dstChainId}#{shardId}#{offset}")
	Type           string         `json:"type"`
	Status         string         `json:"status"`
	SrcChainID     chainID        `json:"srcChainId"`
	SrcTimestamp   time.Time      `json:"srcTimestamp"`
	SrcTxHash      common.Hash    `json:"srcTxHash"`
	SrcBlockNumber uint64         `json:"srcBlockNumber"`
	SrcBlockHash   common.Hash    `json:"srcBlockHash"`
	SrcGasLimit    uint64         `json:"srcGasLimit,string"`
	DstChainID     chainID        `json:"dstChainId"`
	DstTimestamp   time.Time      `json:"dstTimestamp"`
	DstTxHash      common.Hash    `json:"dstTxHash"`
	DstBlockNumber uint64         `json:"dstBlockNumber"`
	DstBlockHash   common.Hash    `json:"dstBlockHash"`
	DstGasLimit    uint64         `json:"dstGasLimit,string"`
	From           common.Address `json:"from"`
	To             common.Address `json:"to"`
	Data           struct {
		Relayer   common.Address `json:"relayer"`
		GasUsed   uint64         `json:"gasUsed,string"`
		Error     hexutil.Bytes  `json:"error"`
		Offset    uint64         `json:"offset,string"`
		ShardID   xchain.ShardID `json:"shardId,string"`
		ConfLevel string         `json:"confirmationLevel"` // ConfLevel of ShardID
		GasLimit  uint64         `json:"gasLimit,string"`
		Fees      string         `json:"fees"`
	} `json:"data"`
}

func (c crossTxJSON) IsCompleted() bool {
	return c.Status == "completed"
}

func (c crossTxJSON) IsPending() bool {
	return c.Status == "pending"
}

// Verify returns an error if fields are invalid.
func (c crossTxJSON) Verify() error {
	if c.Type != "omni" {
		return errors.New("invalid cross tx type", "got", c.Type)
	} else if !c.IsCompleted() && !c.IsPending() {
		return errors.New("invalid cross tx status", "got", c.Status)
	} else if c.Data.ConfLevel != "finalized" && c.Data.ConfLevel != "latest" {
		return errors.New("invalid conf level", "got", c.Data.ConfLevel)
	}

	return nil
}

// ExpectedID calculates the expected ID of the cross transaction from its fields.
func (c crossTxJSON) ExpectedID() string {
	s := fmt.Sprintf("omni#%s#%s#%d#%d", c.SrcChainID, c.DstChainID, c.Data.ShardID, c.Data.Offset)
	bz := sha256.Sum256([]byte(s))

	return hex.EncodeToString(bz[:])
}

// ConfLevel returns the confirmation level of the cross transaction. This matches the ShardID's ConfLevel.
func (c crossTxJSON) ConfLevel() (xchain.ConfLevel, error) {
	switch strings.ToLower(c.Data.ConfLevel) {
	case "finalized":
		return xchain.ConfFinalized, nil
	case "latest":
		return xchain.ConfLatest, nil
	default:
		return 0, errors.New("invalid conf level", "got", c.Data.ConfLevel)
	}
}

func (c crossTxJSON) Success() bool {
	return len(c.Data.Error) == 0
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
			ShardID:       c.Data.ShardID,
		},
		StreamOffset: c.Data.Offset,
	}, nil
}

type crossTxResponse struct {
	CrossTxs []crossTxJSON `json:"items"`
	Links    links         `json:"link"`
}

type links struct {
	Next      string `json:"next"`
	NextToken string `json:"nextToken"`
}

var _ filter = queryFilter{}

// queryFilter defines the cross transactions to query.
type queryFilter struct {
	Stream  xchain.StreamID // Defaults to any stream if zero
	Pending bool            // Default to Completed if false
}

func (f queryFilter) QueryParams(q url.Values) {
	if !f.HasStream() {
		return
	}

	q.Add("srcChainIds", routeScanChainID(f.Stream.SourceChainID))
	q.Add("dstChainIds", routeScanChainID(f.Stream.DestChainID))
}

func (f queryFilter) Match(crossTx crossTxJSON) (bool, error) {
	msgID, err := crossTx.MsgID()
	if err != nil {
		return false, errors.Wrap(err, "parse msg id")
	}

	if f.HasStream() && (f.Stream.DestChainID != msgID.DestChainID || f.Stream.SourceChainID != msgID.SourceChainID) {
		return false, errors.New("invalid dest or source chain", "filter", f.Stream, "msg", msgID)
	}

	// Client side filtering
	if f.HasStream() && f.Stream != msgID.StreamID {
		return false, nil
	} else if f.Pending && !crossTx.IsPending() {
		return false, nil
	} else if !f.Pending && !crossTx.IsCompleted() {
		return false, nil
	}

	return true, nil
}

func (f queryFilter) HasStream() bool {
	return f.Stream != (xchain.StreamID{})
}
