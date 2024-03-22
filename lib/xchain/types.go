package xchain

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Signature65 is a 65 byte Ethereum signature [R || S || V] format.
type Signature65 [65]byte

// StreamID uniquely identifies a cross-chain stream.
// A stream is a logical representation of a cross-chain connection between two chains.
type StreamID struct {
	SourceChainID uint64 // Source chain ID as per https://chainlist.org/
	DestChainID   uint64 // Destination chain ID as per https://chainlist.org/
}

// MsgID uniquely identifies a cross-chain message.
type MsgID struct {
	StreamID            // Unique ID of the Stream this message belongs to
	StreamOffset uint64 // Monotonically incremented offset of Msg in the Steam
}

// Msg is a cross-chain message.
type Msg struct {
	MsgID                          // Unique ID of the message
	SourceMsgSender common.Address // Sender on source chain, set to msg.Sender
	DestAddress     common.Address // Target/To address to "call" on destination chain
	Data            []byte         // Data to provide to "call" on destination chain
	DestGasLimit    uint64         // Gas limit to use for "call" on destination chain
	TxHash          common.Hash    // Hash of the source chain transaction that emitted the message
}

// Receipt is a cross-chain message receipt, the result of applying the Msg on the destination chain.
type Receipt struct {
	MsgID                         // Unique ID of the cross chain message that was applied.
	GasUsed        uint64         // Gas used during message "call"
	Success        bool           // Result, true for success, false for revert
	RelayerAddress common.Address // Address of relayer that submitted the message
	TxHash         common.Hash    // Hash of the relayer submission transaction
}

// BlockHeader uniquely identifies a cross chain block.
type BlockHeader struct {
	SourceChainID uint64      // Source chain ID as per https://chainlist.org
	BlockHeight   uint64      // Height of the source chain block
	BlockHash     common.Hash // Hash of the source chain block
}

// Block is a deterministic representation of the omni cross-chain properties of a source chain EVM block.
type Block struct {
	BlockHeader
	Msgs      []Msg     // All cross-chain messages sent/emittted in the block
	Receipts  []Receipt // Receipts of all submitted cross-chain messages applied in the block
	Timestamp time.Time // Timestamp of the source chain block
}

// Vote by a validator of a cross-chain Block.
type Vote struct {
	BlockHeader                 // BlockHeader identifies the cross-chain Block
	AttestationRoot common.Hash // Attestation merkle root of the cross-chain Block
	Signature       SigTuple    // Validator signature and public key
}

// Attestation containing quorum votes by the validator set of a cross-chain Block.
type Attestation struct {
	BlockHeader                 // BlockHeader identifies the cross-chain Block
	ValidatorSetID  uint64      // Validator set that approved this attestation.
	AttestationRoot common.Hash // Attestation merkle root of the cross-chain Block
	Signatures      []SigTuple  // Validator signatures and public keys
}

// SigTuple is a validator signature and address.
type SigTuple struct {
	ValidatorAddress common.Address // Validator Ethereum address
	Signature        Signature65    // Validator signature over XBlockRoot; Ethereum 65 bytes [R || S || V] format.
}

// Submission is a cross-chain submission of a set of messages and their proofs.
type Submission struct {
	AttestationRoot common.Hash // Attestation merkle root of the cross-chain Block
	ValidatorSetID  uint64      // Validator set that approved the attestation.
	BlockHeader     BlockHeader // BlockHeader identifies the cross-chain Block
	Msgs            []Msg       // Messages to be submitted
	Proof           [][32]byte  // Merkle multi proofs of the messages
	ProofFlags      []bool      // Flags indicating whether the proof is a left or right proof
	Signatures      []SigTuple  // Validator signatures and public keys
	DestChainID     uint64      // Destination chain ID, for internal use only
}

// StreamCursor is a cursor that tracks the progress of a cross-chain stream on destination portal contracts.
type StreamCursor struct {
	StreamID                 // Stream ID of the Stream this cursor belongs to
	Offset            uint64 // Latest applied Msg offset of the Stream
	SourceBlockHeight uint64 // Height of the source chain block
}
