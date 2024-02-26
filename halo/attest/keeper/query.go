package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/attest/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = (*Keeper)(nil)

const approvedFromLimit = 100

func (k *Keeper) AttestationsFrom(ctx context.Context, req *types.AttestationsFromRequest) (*types.AttestationsFromResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	atts, err := k.attestationFrom(ctx, req.ChainId, req.FromHeight, approvedFromLimit)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.AttestationsFromResponse{Attestations: atts}, nil
}

func (k *Keeper) LatestAttestation(ctx context.Context, req *types.LatestAttestationRequest) (*types.LatestAttestationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	att, err := k.latestAttestation(ctx, req.ChainId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.LatestAttestationResponse{Attestation: att}, nil
}
