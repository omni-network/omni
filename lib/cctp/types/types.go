package types

import (
	"math/big"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

func (msg MsgSendUSDC) Validate() error {
	emptyAddr := common.Address{}
	emptyHash := common.Hash{}

	if msg.Amount == nil {
		return errors.New("nil amount")
	}

	if msg.Amount.Sign() <= 0 {
		return errors.New("non-positive amount")
	}

	if msg.SrcChainID == 0 {
		return errors.New("zero source chain ID")
	}

	if msg.DestChainID == 0 {
		return errors.New("zero destination chain ID")
	}

	if msg.Recipient == emptyAddr {
		return errors.New("empty recipient address")
	}

	if len(msg.MessageBytes) == 0 {
		return errors.New("empty message bytes")
	}

	if msg.TxHash == emptyHash {
		return errors.New("empty transaction hash")
	}

	if msg.MessageHash == emptyHash {
		return errors.New("empty message hash")
	}

	if msg.MessageHash != crypto.Keccak256Hash(msg.MessageBytes) {
		return errors.New("message hash != hash of message bytes")
	}

	return nil
}
