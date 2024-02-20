package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
)

type proposalServer struct {
	Keeper
	types.UnimplementedMsgServiceServer
}

// AddAggAttestations handles aggregate attestations proposed in a block.
func (s proposalServer) AddAggAttestations(_ context.Context, msg *types.MsgAggAttestations,
) (*types.AddAggAttestationsResponse, error) {
	localHeaders := headersByAddress(msg.Aggregates, s.attester.LocalAddress())
	if err := s.attester.SetProposed(localHeaders); err != nil {
		return nil, errors.Wrap(err, "set committed")
	}

	return &types.AddAggAttestationsResponse{}, nil
}

// NewProposalServer returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewProposalServer(keeper Keeper) types.MsgServiceServer {
	return &proposalServer{Keeper: keeper}
}

var _ types.MsgServiceServer = proposalServer{}
