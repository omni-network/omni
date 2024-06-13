package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/registry/types"
	"github.com/omni-network/omni/lib/errors"

	"cosmossdk.io/orm/types/ormerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = (*Keeper)(nil)

func (k Keeper) Network(ctx context.Context, req *types.NetworkRequest) (*types.NetworkResponse, error) {
	networkID := req.Id
	if req.Latest {
		var err error
		networkID, err = k.networkTable.LastInsertedSequence(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	network, err := k.networkTable.Get(ctx, networkID)
	if errors.Is(err, ormerrors.NotFound) {
		return nil, status.Error(codes.NotFound, "no network found for id")
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	portals := make([]*types.Portal, 0, len(network.GetPortals()))
	for _, portal := range network.GetPortals() {
		portals = append(portals, &types.Portal{
			ChainId:      portal.GetChainId(),
			Address:      portal.GetAddress(),
			DeployHeight: portal.GetDeployHeight(),
			ShardIds:     portal.GetShardIds(),
		})
	}

	return &types.NetworkResponse{
		Id:            network.GetId(),
		CreatedHeight: network.GetCreatedHeight(),
		Portals:       portals,
	}, nil
}
