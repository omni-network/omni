package resolvers

import (
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
	BlockHeight         hexutil.Big
	BlockHash           common.Hash
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

	// TODO: add paging for the messages.
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

type PageInfo struct {
	StartCursor hexutil.Big
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
