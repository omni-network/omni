package usdt0

import (
	"math/big"

	"github.com/omni-network/omni/lib/layerzero"

	"github.com/ethereum/go-ethereum/common"
)

// MsgSend is an instance of a USDT0 send between chains.
type MsgSend struct {
	TxHash      common.Hash         // Transaction hash
	BlockHeight uint64              // Height of source chain block
	SrcChainID  uint64              // Source chain ID
	DestChainID uint64              // Destination chain ID
	Amount      *big.Int            // Amount of USDT0 sent
	Status      layerzero.MsgStatus // Message status
}
