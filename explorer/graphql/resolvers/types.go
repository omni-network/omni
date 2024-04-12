package resolvers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/graph-gophers/graphql-go"
)

type XMsg struct {
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
