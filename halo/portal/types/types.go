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

const (
	MsgTypeUnknown  MsgType = 0
	MsgTypeValSet   MsgType = 1
	msgTypeSentinel MsgType = 2 // Must be last
)

type EmitPortal interface {
	CreateMsg(ctx sdk.Context, typ MsgType, msgTypeID uint64, destChainID uint64, shardID xchain.ShardID) error
}
