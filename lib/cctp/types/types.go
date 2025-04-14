package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// MsgSendUSDC represents a USDC transfer message between chains.
type MsgSendUSDC struct {
	MessageHash  common.Hash
	TxHash       common.Hash
	SrcChainID   uint64
	DestChainID  uint64
	Amount       *big.Int
	MessageBytes []byte
	Recipient    common.Address
}
