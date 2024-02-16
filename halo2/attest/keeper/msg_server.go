package keeper

import (
	"context"

	"github.com/omni-network/omni/halo2/attest/types"
	"github.com/omni-network/omni/lib/errors"
)

type msgServer struct {
	Keeper
	types.UnimplementedMsgServiceServer
}

func (s msgServer) AddAggAttestation(ctx context.Context, msg *types.MsgAggAttestation,
) (*types.AddAggAttestationResponse, error) {
	err := s.Keeper.Add(ctx, msg)
	if err != nil {
		return nil, errors.Wrap(err, "add aggregate attestation")
	}

	return &types.AddAggAttestationResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServiceServer = msgServer{}
