package data

import (
	"math/big"
	"time"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/graphql/resolvers"

	"github.com/ethereum/go-ethereum/common"

	"github.com/graph-gophers/graphql-go"
)

type Provider struct {
	EntClient ent.Client
}

func (Provider) XBlock(sourceChainID uint64, height uint64) (*resolvers.XBlock, bool, error) {
	h := common.Hash{}
	h.SetBytes([]byte{1, 3, 23, 111, 27, 45, 98, 103, 94, 55, 1, 3, 23, 111, 27, 45, 98, 103, 94, 55})
	var chainID big.Int
	chainID.SetUint64(sourceChainID)
	var blockHeight big.Int
	blockHeight.SetUint64(height)

	res := resolvers.XBlock{
		SourceChainIDRaw: resolvers.BigInt{Int: chainID},
		BlockHeightRaw:   resolvers.BigInt{Int: blockHeight},
		BlockHashRaw:     h,
		Timestamp:        graphql.Time{Time: time.Now()},
		Messages:         dummyMessages(),
	}

	return &res, true, nil
}

func dummyMessages() []resolvers.XMsg {
	var a, b, c, d big.Int
	a.SetUint64(2)
	a.SetUint64(3)
	c.SetInt64(5)
	d.SetInt64((100_000))

	destAddr := common.Address{}
	destAddr.SetBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9})
	senderAddr := common.Address{}
	senderAddr.SetBytes([]byte{34, 73, 84, 82, 12, 15, 43, 24, 76, 3, 6, 0, 0, 0, 3, 2, 4, 5})

	txHash := common.Hash{}
	txHash.SetBytes([]byte{5, 0, 0, 4, 0, 0, 1})

	res := []resolvers.XMsg{
		{
			SourceChainID:          resolvers.BigInt{Int: a},
			DestChainID:            resolvers.BigInt{Int: b},
			StreamOffset:           resolvers.BigInt{Int: c},
			SourceMessageSenderRaw: destAddr,
			DestAddressRaw:         destAddr,
			DestGasLimit:           resolvers.BigInt{Int: d},
			TxHashRaw:              txHash,
		},
	}

	return res
}
