package resolvers

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/graph-gophers/graphql-go"
)

type Msg struct {
	UUID                   BigInt
	SourceMessageSenderRaw common.Address
	DestAddressRaw         common.Address
	Data                   []byte
	DestGasLimit           BigInt
	SourceChainID          BigInt
	DestChainID            BigInt
	StreamOffset           BigInt
	TxHashRaw              common.Hash
}

type Block struct {
	UUID          BigInt
	SourceChainID BigInt
	BlockHeight   BigInt
	Timestamp     graphql.Time
	CreatedAt     graphql.Time
	BlockHashRaw  common.Hash

	// TODO(Pavel): add paging for the messages.
	Messages []Msg
}
