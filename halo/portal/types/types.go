package types

import (
	"github.com/omni-network/omni/lib/errors"

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

type Portal interface {
	CreateMsg(ctx sdk.Context, typ MsgType, msgTypeID uint64) error
}
