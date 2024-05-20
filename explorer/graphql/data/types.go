package data

import (
	"fmt"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
)

const StatusParseErr = StatusErr("Invalid status")

type StatusErr string

func (s StatusErr) Error() string {
	return string(s)
}

// Define the Go struct for the Status enum type.
type Status string

func ParseStatus(s string) (Status, error) {
	switch s {
	case string(StatusFailed):
		return StatusFailed, nil
	case string(StatusPending):
		return StatusPending, nil
	case string(StatusSuccess):
		return StatusSuccess, nil
	default:
		return Status(""), StatusParseErr
	}
}

func MustParseStatus(s string) Status {
	st, err := ParseStatus(s)
	if err != nil {
		panic(err)
	}

	return st
}

const (
	StatusFailed  Status = "FAILED"
	StatusPending Status = "PENDING"
	StatusSuccess Status = "SUCCESS"
)

type Chainer interface {
	Chain(id string) (Chain, bool)
}

// Define the Go struct for the XMsg type.
type XMsg struct {
	Chainer

	ID            graphql.ID
	Block         XBlock
	To            common.Address
	ToURL         string
	Data          hexutil.Bytes
	DestChainID   hexutil.Big
	GasLimit      hexutil.Big
	DisplayID     string
	Offset        hexutil.Big
	Receipt       *XReceipt
	Sender        common.Address
	SenderURL     string
	SourceChainID hexutil.Big
	Status        Status
	TxHash        common.Hash
	TxURL         string
}

func (m XMsg) SourceChain() (Chain, error) {
	c, ok := m.Chain(m.SourceChainID.String())
	if !ok {
		return Chain{}, errors.New("chain not found", "chain_id", m.SourceChainID.String())
	}

	return c, nil
}

func (m XMsg) DestChain() (Chain, error) {
	c, ok := m.Chain(m.DestChainID.String())
	if !ok {
		return Chain{}, errors.New("chain not found", "chain_id", m.DestChainID.String())
	}

	return c, nil
}

// Define the Go struct for the XBlock type.
type XBlock struct {
	Chainer
	ID        graphql.ID
	ChainID   hexutil.Big
	Height    hexutil.Big
	Hash      common.Hash
	Messages  []XMsg
	Timestamp graphql.Time
	URL       string
}

func (b XBlock) Chain() (Chain, error) {
	c, ok := b.Chainer.Chain(b.ChainID.String())
	if !ok {
		return Chain{}, fmt.Errorf("chain not found: %s", b.ChainID.String())
	}

	return c, nil
}

// Define the Go struct for the XReceipt type.
type XReceipt struct {
	Chainer
	ID            graphql.ID
	GasUsed       hexutil.Big
	Success       bool
	Relayer       common.Address
	SourceChainID hexutil.Big
	DestChainID   hexutil.Big
	Offset        hexutil.Big
	TxHash        common.Hash
	TxURL         string
	Timestamp     graphql.Time
	RevertReason  *string
}

func (r *XReceipt) SourceChain() (Chain, error) {
	c, ok := r.Chain(r.SourceChainID.String())
	if !ok {
		return Chain{}, fmt.Errorf("chain not found: %s", r.SourceChainID.String())
	}

	return c, nil
}

func (r *XReceipt) DestChain() (Chain, error) {
	c, ok := r.Chain(r.DestChainID.String())
	if !ok {
		return Chain{}, fmt.Errorf("chain not found: %s", r.DestChainID.String())
	}

	return c, nil
}

// Define the Go struct for the Chain type.
type Chain struct {
	ID        graphql.ID
	ChainID   hexutil.Big
	DisplayID Long
	Name      string
	LogoURL   string
}

// Define the Go struct for the XMsgConnection type.
type XMsgConnection struct {
	TotalCount Long
	Edges      []XMsgEdge
	PageInfo   PageInfo
}

// Define the Go struct for the XMsgEdge type.
type XMsgEdge struct {
	Cursor graphql.ID
	Node   XMsg
}

// Define the Go struct for the PageInfo type.
type PageInfo struct {
	HasNextPage bool
	HasPrevPage bool
	TotalPages  Long
	CurrentPage Long
}

// Define the Go struct for the StatsResult type.
type StatsResult struct {
	TotalMsgs  Long
	TopStreams []StreamStats
}

// Define the Go struct for the StreamStats type.
type StreamStats struct {
	SourceChain Chain
	DestChain   Chain
	MsgCount    Long
}
