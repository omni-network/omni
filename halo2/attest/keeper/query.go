package keeper

import (
	"context"

	"github.com/omni-network/omni/halo2/attest/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

const approvedFromLimit = 100

func (k Keeper) ApprovedFrom(ctx context.Context, req *types.ApprovedFromRequest) (*types.ApprovedFromResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	aggs, err := k.approvedFrom(ctx, req.ChainId, req.FromHeight, approvedFromLimit)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.ApprovedFromResponse{Aggregates: aggs}, nil
}
