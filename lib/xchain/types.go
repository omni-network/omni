package xchain

import (
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// BroadcastChainID is the chain ID used by broadcast messages.
const BroadcastChainID uint64 = 0

//go:generate stringer -type=ConfLevel -linecomment

// ChainVersion defines a version of a source chain; either some draft (fuzzy) version or finalized.
type ChainVersion struct {
	ID        uint64    // Source chain ID as per https://chainlist.org/
	ConfLevel ConfLevel // ConfLevel defines the block "version"; either some fuzzy version or finalized.
}

func NewChainVersion(chainID uint64, confLevel ConfLevel) ChainVersion {
	return ChainVersion{
		ID:        chainID,
		ConfLevel: confLevel,
	}
}

// ConfLevel defines a xblock confirmation level.
// This is similar to a "version"; with ConfFinalized being the final version and fuzzy conf levels being drafts.
type ConfLevel byte

// Valid returns true if this confirmation level is valid.
func (c ConfLevel) Valid() bool {
	return c > ConfUnknown && c < confSentinel && !strings.Contains(c.String(), "ConfLevel")
}

// IsFuzzy returns true if this confirmation level is not ConfFinalized.
func (c ConfLevel) IsFuzzy() bool {
	return c == ConfLatest
}

// Label returns a short label for the confirmation level.
// IT is the uppercase first letter of the confirmation level.
func (c ConfLevel) Label() string {
	return strings.ToUpper(c.String()[:1])
}

// ConfLevel values MUST never change as they are persisted on-chain.
const (
	ConfUnknown   ConfLevel = 0 // unknown
	ConfLatest    ConfLevel = 1 // latest
	_             ConfLevel = 2 // reserved
	_             ConfLevel = 3 // reserved
	ConfFinalized ConfLevel = 4 // final
	confSentinel  ConfLevel = 5 // sentinel must always be last
)

// FuzzyConfLevels returns a list of all fuzzy confirmation levels.
func FuzzyConfLevels() []ConfLevel {
	return []ConfLevel{ConfLatest}
}

type ShardID uint64

const (
	// ShardFinalized0 is the default finalized confirmation level shard.
	ShardFinalized0 = ShardID(ConfFinalized)

	// ShardLatest0 is the default latest confirmation level shard.
	ShardLatest0 = ShardID(ConfLatest)

	// ShardBroadcast0 is the default broadcast shard. It uses the finalized confirmation level.
	ShardBroadcast0 = ShardID(ConfFinalized) | 0x0100
)

// ConfLevel returns confirmation level encoded in the
// last 8 bits of the shardID.
func (s ShardID) ConfLevel() ConfLevel {
	return ConfLevel(byte(s & 0xFF))
}

// Flags returns flags encoded in the 2nd-to-last byte of the shardID.
func (s ShardID) Flags() byte {
	return byte((s >> 8) & 0xFF)
}

// Label returns a short label for the shard.
// IT is the uppercase first letter of the confirmation level.
func (s ShardID) Label() string {
	resp, ok := map[ShardID]string{
		ShardFinalized0: "F",
		ShardLatest0:    "L",
		ShardBroadcast0: "B",
	}[s]
	if ok {
		return resp
	}

	return strconv.FormatUint(uint64(s), 10)
}

// Broadcast returns the value of the 8th flag (least significant bit).
func (s ShardID) Broadcast() bool {
	return s.Flags()&0b00000001 == 1
}

// Signature65 is a 65 byte Ethereum signature [R || S || V] format.
type Signature65 [65]byte

// StreamID uniquely identifies a cross-chain stream.
// A stream is a logical representation of a cross-chain connection between two chains.
type StreamID struct {
	SourceChainID uint64  // Source chain ID as per https://chainlist.org/
	DestChainID   uint64  // Destination chain ID as per https://chainlist.org/
	ShardID       ShardID // ShardID identifies a sequence of xmsgs (and maps to ConfLevel).
}

func (s StreamID) ConfLevel() ConfLevel {
	return ConfLevel(s.ShardID)
}

func (s StreamID) ChainVersion() ChainVersion {
	return ChainVersion{ID: s.SourceChainID, ConfLevel: s.ConfLevel()}
}

// MsgID uniquely identifies a cross-chain message.
type MsgID struct {
	StreamID            // Unique ID of the Stream this message belongs to
	StreamOffset uint64 // Monotonically incremented offset of Msg in the Stream (1-indexed)
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
	ConfLevel      ConfLevel      // Confirmation level of submitted attestation
	GasUsed        uint64         // Gas used during message "call"
	Success        bool           // Result, true for success, false for revert
	Error          []byte         // Error message if the message failed
	RelayerAddress common.Address // Address of relayer that submitted the message
	TxHash         common.Hash    // Hash of the relayer submission transaction
}

// BlockHeader uniquely identifies a cross chain block.
type BlockHeader struct {
	ChainID     uint64      // Source chain ID as per https://chainlist.org
	BlockHeight uint64      // Height of the source-chain block
	BlockHash   common.Hash // Hash of the source-chain block
}

// Block is a deterministic representation of the omni cross-chain properties of a source chain EVM block.
type Block struct {
	BlockHeader
	Msgs       []Msg       // All cross-chain messages sent/emittted in the block
	Receipts   []Receipt   // Receipts of all submitted cross-chain messages applied in the block
	ParentHash common.Hash // ParentHash is the hash of the parent block.
	Timestamp  time.Time   // Timestamp of the source chain block
}

// ShouldAttest returns true if the xblock should be attested by the omni consensus chain validators.
// All "non-empty" xblocks should be attested to.
// Every Nth block based on the chain's attest interval should be attested to.
// Attested blocks are assigned an incremented XBlockOffset.
func (b Block) ShouldAttest(attestInterval uint64) bool {
	if len(b.Msgs) > 0 {
		return true
	}

	if attestInterval == 0 {
		return false // Avoid empty attestations for zero interval.
	}

	return b.BlockHeight%attestInterval == 0
}

// AttestHeader uniquely identifies an attestation that require quorum vote.
// This is used to determine duplicate votes.
type AttestHeader struct {
	ConsensusChainID uint64       // Omni consensus chain ID this attestation/vote belangs to. Used for replay-protection.
	ChainVersion     ChainVersion // ChainVersion defines a "version" of a chain being attested to ; either some fuzzy version or finalized.
	AttestOffset     uint64       // Monotonically increasing offset of this vote per chain version. 1-indexed.
}

// Vote by a validator for a cross-chain Block.
type Vote struct {
	AttestHeader             // AttestHeader identifies the attestation this vote should be included in.
	BlockHeader              // BlockHeader identifies the cross-chain Block being voted for.
	MsgRoot      common.Hash // Merkle root of all messages in the cross-chain Block
	Signature    SigTuple    // Validator signature and public key
}

// Attestation containing quorum votes by the validator set of a cross-chain Block.
type Attestation struct {
	AttestHeader               // AttestHeader uniquely identifies the attestation.
	BlockHeader                // BlockHeader identifies the cross-chain Block
	ValidatorSetID uint64      // Validator set that approved this attestation.
	MsgRoot        common.Hash // Merkle root of all messages in the cross-chain Block
	Signatures     []SigTuple  // Validator signatures and public keys
}

func (a Attestation) AttestationRoot() ([32]byte, error) {
	return AttestationRoot(a.AttestHeader, a.BlockHeader, a.MsgRoot)
}

// SigTuple is a validator signature and address.
type SigTuple struct {
	ValidatorAddress common.Address // Validator Ethereum address
	Signature        Signature65    // Validator signature over XBlockRoot; Ethereum 65 bytes [R || S || V] format.
}

// Submission is a cross-chain submission of a set of messages and their proofs.
type Submission struct {
	AttestationRoot common.Hash  // Attestation merkle root of the cross-chain Block
	ValidatorSetID  uint64       // Validator set that approved the attestation.
	AttHeader       AttestHeader // AttestHeader identifies the attestation this submission belongs to.
	BlockHeader     BlockHeader  // BlockHeader identifies the cross-chain Block
	Msgs            []Msg        // Messages to be submitted
	Proof           [][32]byte   // Merkle multi proofs of the messages
	ProofFlags      []bool       // Flags indicating whether the proof is a left or right proof
	Signatures      []SigTuple   // Validator signatures and public keys
	DestChainID     uint64       // Destination chain ID, for internal use only
}

// SubmitCursor is a cursor that tracks the progress of a cross-chain stream on destination portal contracts.
type SubmitCursor struct {
	StreamID           // Stream ID of the Stream this cursor belongs to
	MsgOffset   uint64 // Latest submitted Msg offset of the Stream
	BlockOffset uint64 // Latest submitted cross chain block offset
}

// EmitCursor is a cursor that tracks the progress of a cross-chain stream on source portal contracts.
type EmitCursor struct {
	StreamID         // Stream ID of the Stream this cursor belongs to
	MsgOffset uint64 // Latest emitted Msg offset of the Stream
}
