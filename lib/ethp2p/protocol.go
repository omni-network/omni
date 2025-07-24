package ethp2p

import (
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rlp"
)

// Unexported devp2p message codes from p2p package.
const (
	handshakeMsg = 0x00
	discMsg      = 0x01
	// pongMsg      = 0x03.
)

// Unexported devp2p protocol lengths from p2p package.
const (
	baseProtoLen = 16
	ethProtoLen  = 17
	snapProtoLen = 8
)

// Unexported handshake structure from p2p package.
type protoHandshake struct {
	Version    uint64
	Name       string
	Caps       []p2p.Cap
	ListenPort uint64
	ID         []byte
	Rest       []rlp.RawValue `rlp:"tail"`
}

// proto is an enum representing devp2p protocol types.
type proto string

const (
	protoUnknown proto = "unknown"
	protoBase    proto = "base"
	protoEth     proto = "eth"
	protoSnap    proto = "snap"
)

// (assuming the negotiated capabilities are exactly {eth,snap}).
func parseProto(code uint64) proto {
	switch {
	case code < baseProtoLen:
		return protoBase
	case code < baseProtoLen+ethProtoLen:
		return protoEth
	case code < baseProtoLen+ethProtoLen+snapProtoLen:
		return protoSnap
	default:
		return protoUnknown
	}
}

// MsgCode returns the message code for a given protocol and message number.
func (p proto) MsgCode(msg uint64) uint64 {
	return p.Offset() + msg
}

// Offset returns the offset at which the specified protocol's messages
// begin.
func (p proto) Offset() uint64 {
	switch p {
	case protoBase:
		return 0
	case protoEth:
		return baseProtoLen
	case protoSnap:
		return baseProtoLen + ethProtoLen
	default:
		return 33333
	}
}
