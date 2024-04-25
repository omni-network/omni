package resolvers

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
)

type XMsg struct {
	ID                  graphql.ID
	UUID                graphql.ID
	SourceMessageSender common.Address
	DestAddress         common.Address
	Data                []byte
	DestGasLimit        hexutil.Big
	SourceChainID       hexutil.Big
	DestChainID         hexutil.Big
	StreamOffset        hexutil.Big
	TxHash              common.Hash
	Status              string
	BlockHeight         hexutil.Big
	BlockHash           common.Hash
	SourceBlockTime     graphql.Time
	ReceiptTxHash       *common.Hash
	Block               XBlock
	Receipts            []XReceipt
}

type XBlock struct {
	ID            graphql.ID
	UUID          graphql.ID
	SourceChainID hexutil.Big
	BlockHeight   hexutil.Big
	Timestamp     graphql.Time
	CreatedAt     graphql.Time
	BlockHash     common.Hash

	Messages []XMsg
	Receipts []XReceipt
}

type XReceipt struct {
	ID             graphql.ID
	UUID           graphql.ID
	Success        graphql.NullBool
	GasUsed        hexutil.Big
	RelayerAddress common.Address
	SourceChainID  hexutil.Big
	DestChainID    hexutil.Big
	StreamOffset   hexutil.Big
	TxHash         common.Hash
	Timestamp      graphql.Time
	BlockHeight    hexutil.Big
	BlockHash      common.Hash
	Block          XBlock
	Messages       []XMsg
}

type Chain struct {
	Name    string
	ChainID hexutil.Big
}

type SearchResult struct {
	BlockHeight   hexutil.Big
	TxHash        common.Hash
	SourceChainID hexutil.Big
	Type          SearchResultType
}

type SearchResultType string

const (
	BLOCK   SearchResultType = "BLOCK"
	MESSAGE SearchResultType = "MESSAGE"
	RECEIPT SearchResultType = "RECEIPT"
	ADDRESS SearchResultType = "ADDRESS"
)

type XMsgStatus string

const (
	PENDING XMsgStatus = "PENDING"
	SUCCESS XMsgStatus = "SUCCESS"
	FAILED  XMsgStatus = "FAILED"
)

type PageInfo struct {
	PrevCursor  hexutil.Big
	NextCursor  hexutil.Big
	HasNextPage bool
	HasPrevPage bool
}

type XMsgResult struct {
	TotalCount hexutil.Big
	Edges      []XMsgEdge
	PageInfo   PageInfo
}

type XMsgEdge struct {
	Cursor hexutil.Big
	Node   XMsg
}

var (
	XMsgStatusMap = map[string]XMsgStatus{
		"success": SUCCESS,
		"pending": PENDING,
		"failed":  FAILED,
	}
)

func ParseString(str string) (XMsgStatus, bool) {
	c, ok := XMsgStatusMap[strings.ToLower(str)]
	return c, ok
}
