// Package xtypes defines the types used by the omni cross-chain protocol.
package xtypes

// XStreamID uniquely identifies a cross-chain stream.
// A stream is a logical representation of a cross-chain connection between two chains.
type XStreamID struct {
	SourceChainID uint64 // Source chain ID as per https://chainlist.org/
	DestChainID   uint64 // Destination chain ID as per https://chainlist.org/
}

// XMsgID uniquely identifies a cross-chain message.
type XMsgID struct {
	XStreamID            // Unique ID of the XStream this message belongs to
	XStreamOffset uint64 // Monotonically incremented offset of XMsg in the XSteam
}

// XMsg is a cross-chain message.
type XMsg struct {
	XMsgID                   // Unique ID of the message
	SourceMsgSender [20]byte // Sender on source chain, set to msg.Sender
	DestAddress     [20]byte // Target/To address to "call" on destination chain
	Data            []byte   // Data to provide to "call" on destination chain
	DestGasLimit    uint64   // Gas limit to use for "call" on destination chain
}

// XReceipt is a cross-chain message receipt, the result of applying the XMsg on the destination chain.
type XReceipt struct {
	XMsgID                  // Unique ID of the cross chain message that was applied.
	GasUsed        uint64   // Gas used during message "call"
	Result         uint64   // 0 for success, 1 for revert
	RelayerAddress [20]byte // Address of relayer that submitted the message
}

// XBlockHeader uniquely identifies a cross chain block.
type XBlockHeader struct {
	SourceChainID uint64   // Source chain ID as per https://chainlist.org
	BlockHeight   uint64   // Height of the source chain block
	BlockHash     [32]byte // Hash of the source chain block
}

// XBlock is a deterministic representation of the omni cross-chain properties of a source chain EVM block.
type XBlock struct {
	XBlockHeader
	Msgs     []XMsg     // All cross-chain messages sent/emittted in the block
	Receipts []XReceipt // Receipts of all submitted cross-chain messages applied in the block
}
