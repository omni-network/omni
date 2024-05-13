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

	atts, err := k.ListAttestationsFrom(ctx, req.ChainId, req.FromOffset, approvedFromLimit)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.AttestationsFromResponse{Attestations: atts}, nil
}

func (k *Keeper) LatestAttestation(ctx context.Context, req *types.LatestAttestationRequest) (*types.LatestAttestationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	att, ok, err := k.latestAttestation(ctx, req.ChainId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if !ok {
		return nil, status.Error(codes.NotFound, "no approved attestations for chain")
	}

	sigs, err := k.getSigs(ctx, att.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	respSigs := make([]*types.SigTuple, 0, len(sigs))
	for _, sig := range sigs {
		respSigs = append(respSigs, &types.SigTuple{
			ValidatorAddress: sig.GetValidatorAddress(),
			Signature:        sig.GetSignature(),
		})
	}

	return &types.LatestAttestationResponse{Attestation: &types.Attestation{
		BlockHeader: &types.BlockHeader{
			ChainId: att.GetChainId(),
			Offset:  att.GetOffset(),
			Height:  att.GetHeight(),
			Hash:    att.GetHash(),
		},
		AttestationRoot: att.GetAttestationRoot(),
		ValidatorSetId:  att.GetValidatorSetId(),
		Signatures:      respSigs,
	}}, nil
}

func (k *Keeper) WindowCompare(ctx context.Context, req *types.WindowCompareRequest) (*types.WindowCompareResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	cmp, err := k.windowCompare(ctx, req.ChainId, req.GetBlockOffset())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.WindowCompareResponse{Cmp: int32(cmp)}, nil
}
