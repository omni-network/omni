package types

import (
	"bytes"
	"fmt"
	"math/big"
	"reflect"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//go:generate stringer -type=MsgStatus -trimprefix=MsgStatus

// It matches status proto enum defined in lib/cctp/db/dp.proto.
type MsgStatus int32

const (
	MsgStatusUnknown   MsgStatus = 0
	MsgStatusSubmitted MsgStatus = 1
	MsgStatusMinted    MsgStatus = 2
)

func (s MsgStatus) Validate() error {
	switch s {
	case MsgStatusSubmitted, MsgStatusMinted:
		return nil
	case MsgStatusUnknown:
		return errors.New("unknown message status")
	default:
		return errors.New("invalid message status")
	}
}

// MsgSendUSDC represents a USDC transfer message between chains.
type MsgSendUSDC struct {
	TxHash       common.Hash
	BlockHeight  uint64
	MessageHash  common.Hash
	SrcChainID   uint64
	DestChainID  uint64
	Amount       *big.Int
	MessageBytes []byte
	Recipient    common.Address
	Status       MsgStatus
}

func (m MsgSendUSDC) Validate() error {
	emptyAddr := common.Address{}
	emptyHash := common.Hash{}

	if m.Amount == nil {
		return errors.New("nil amount")
	}

	if m.Amount.Sign() <= 0 {
		return errors.New("non-positive amount")
	}

	if m.SrcChainID == 0 {
		return errors.New("zero source chain ID")
	}

	if m.DestChainID == 0 {
		return errors.New("zero destination chain ID")
	}

	if m.Recipient == emptyAddr {
		return errors.New("empty recipient address")
	}

	if len(m.MessageBytes) == 0 {
		return errors.New("empty message bytes")
	}

	if m.TxHash == emptyHash {
		return errors.New("empty transaction hash")
	}

	if m.MessageHash == emptyHash {
		return errors.New("empty message hash")
	}

	if m.MessageHash != crypto.Keccak256Hash(m.MessageBytes) {
		return errors.New("message hash != hash of message bytes")
	}

	if m.BlockHeight == 0 {
		return errors.New("zero block height")
	}

	return m.Status.Validate()
}

func (m MsgSendUSDC) Equals(n MsgSendUSDC) bool {
	return reflect.DeepEqual(m, n)
}

// abbrevBz returns a truncated hex representation of the bytes.
// For bytes <= 8 bytes, returns the full hex string.
// For longer bytes, returns "0x" + first 4 bytes + "..." + last 4 bytes.
func abbrevBz(b []byte) string {
	if len(b) <= 8 {
		return "0x" + common.Bytes2Hex(b)
	}

	return "0x" + common.Bytes2Hex(b[:4]) + "..." + common.Bytes2Hex(b[len(b)-4:])
}

// Diff return json object describing the diff between two MsgSendUSDCs.
// This should be used for logging / debugging purposes only.
func (m MsgSendUSDC) Diff(n MsgSendUSDC) map[string]string {
	diff := make(map[string]string)
	notEq := func(a, b any) string { return fmt.Sprintf("%v != %v", a, b) }

	if m.TxHash != n.TxHash {
		diff["tx_hash"] = notEq(m.TxHash, n.TxHash)
	}
	if m.BlockHeight != n.BlockHeight {
		diff["block_height"] = notEq(m.BlockHeight, n.BlockHeight)
	}
	if m.MessageHash != n.MessageHash {
		diff["message_hash"] = notEq(m.MessageHash, n.MessageHash)
	}
	if m.SrcChainID != n.SrcChainID {
		diff["src_chain_id"] = notEq(m.SrcChainID, n.SrcChainID)
	}
	if m.DestChainID != n.DestChainID {
		diff["dest_chain_id"] = notEq(m.DestChainID, n.DestChainID)
	}
	if m.Amount.Cmp(n.Amount) != 0 {
		diff["amount"] = notEq(m.Amount, n.Amount)
	}
	if m.Recipient != n.Recipient {
		diff["recipient"] = notEq(m.Recipient, n.Recipient)
	}
	if m.Status != n.Status {
		diff["status"] = notEq(m.Status, n.Status)
	}
	if !bytes.Equal(m.MessageBytes, n.MessageBytes) {
		diff["message_bytes"] = notEq(abbrevBz(m.MessageBytes), abbrevBz(n.MessageBytes))
	}

	return diff
}
