package types

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgType uint32

func (m MsgType) Validate() error {
	if m == MsgTypeUnknown || m >= msgTypeSentinel {
		return errors.New("invalid message type")
	}

	return nil
}

//go:generate stringer -type=MsgType --trimprefix=MsgType

const (
	MsgTypeUnknown  MsgType = 0
	MsgTypeValSet   MsgType = 1
	MsgTypeNetwork  MsgType = 2
	msgTypeSentinel MsgType = 3 // Must be last
)

// EmitPortal provides an interface for modules to emit cross chain messages.
type EmitPortal interface {
	// EmitMsg emits a cross chain message in the current block returning the xblock ID/Height/Offset.
	EmitMsg(ctx sdk.Context, typ MsgType, msgTypeID uint64, destChainID uint64, shardID xchain.ShardID) (uint64, error)
}

func (m *Msg) MsgType() MsgType {
	return MsgType(m.Type)
}

func (m *Msg) ShardID() xchain.ShardID {
	return xchain.ShardID(m.ShardId)
}
