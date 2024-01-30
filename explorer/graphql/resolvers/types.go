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
}

type XBlock struct {
	UUID          graphql.ID
	SourceChainID hexutil.Big
	BlockHeight   hexutil.Big
	Timestamp     graphql.Time
	CreatedAt     graphql.Time
	BlockHash     common.Hash

	// TODO(Pavel): add paging for the messages.
	Messages []XMsg
}
