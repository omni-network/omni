package xchain

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
	MsgID                    // Unique ID of the message
	SourceMsgSender [20]byte // Sender on source chain, set to msg.Sender
	DestAddress     [20]byte // Target/To address to "call" on destination chain
	Data            []byte   // Data to provide to "call" on destination chain
	DestGasLimit    uint64   // Gas limit to use for "call" on destination chain
	TxHash          [32]byte // Hash of the source chain transaction that emitted the message
}

// Receipt is a cross-chain message receipt, the result of applying the Msg on the destination chain.
type Receipt struct {
	MsgID                   // Unique ID of the cross chain message that was applied.
	GasUsed        uint64   // Gas used during message "call"
	Success        bool     // Result, true for success, false for revert
	RelayerAddress [20]byte // Address of relayer that submitted the message
	TxHash         [32]byte // Hash of the relayer submission transaction
}

// BlockHeader uniquely identifies a cross chain block.
type BlockHeader struct {
	SourceChainID uint64   // Source chain ID as per https://chainlist.org
	BlockHeight   uint64   // Height of the source chain block
	BlockHash     [32]byte // Hash of the source chain block
}

// Block is a deterministic representation of the omni cross-chain properties of a source chain EVM block.
type Block struct {
	BlockHeader
	Msgs     []Msg     // All cross-chain messages sent/emittted in the block
	Receipts []Receipt // Receipts of all submitted cross-chain messages applied in the block
}

// Attestation by a validator of a cross-chain Block.
type Attestation struct {
	BlockHeader          // BlockHeader identifies the cross-chain Block
	BlockRoot   [32]byte // Merkle root of the cross-chain Block
	Signature   SigTuple // Validator signature and public key
}

// AggAttestation aggregates multiple attestation by a validator set of a cross-chain Block.
type AggAttestation struct {
	BlockHeader               // BlockHeader identifies the cross-chain Block
	ValidatorSetID uint64     // Unique identified of the validator set included in this aggregate.
	BlockRoot      [32]byte   // Merkle root of the cross-chain Block
	Signatures     []SigTuple // Validator signatures and public keys
}

// SigTuple is a validator signature and public key.
type SigTuple struct {
	ValidatorPubKey [33]byte // Validator public key; 33 bytes compressed secp256k1.
	Signature       [65]byte // Validator signature over XBlockRoot; Ethereum 65 bytes [R || S || V] format.
}
