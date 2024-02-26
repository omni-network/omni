package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/evmengine/types"
)

type proposalServer struct {
	*Keeper
	types.UnimplementedMsgServiceServer
}

// ExecutionPayload handles a new execution payload proposed in a block.
func (s proposalServer) ExecutionPayload(ctx context.Context, msg *types.MsgExecutionPayload,
) (*types.ExecutionPayloadResponse, error) {
	_, err := pushPayload(ctx, s.engineCl, msg)
	if err != nil {
		return nil, err
	}

	return &types.ExecutionPayloadResponse{}, nil
}

// NewProposalServer returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewProposalServer(keeper *Keeper) types.MsgServiceServer {
	return &proposalServer{Keeper: keeper}
}

var _ types.MsgServiceServer = proposalServer{}
