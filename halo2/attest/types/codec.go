package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), new(MsgAggAttestation))
	registry.RegisterImplementations((*tx.MsgResponse)(nil), new(AddAggAttestationResponse))

	msgservice.RegisterMsgServiceDesc(registry, &_MsgService_serviceDesc)
}
