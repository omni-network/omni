package cctp

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// MsgSendUSDC is a CCTP message of an xchain USDC transfer.
type MsgSendUSDC struct {
	TxHash       common.Hash
	SrcChainID   uint64
	DestChainID  uint64
	Amount       *big.Int
	MessageBytes []byte
	MessageHash  common.Hash
	Recipient    common.Address
}

// DB is a stub DB interface for storing CCTP messages, required to that
// receiving messages on destination chain is robust to restarts / crashes.
type DB interface {
	Insert(ctx context.Context, msg MsgSendUSDC) error
}
