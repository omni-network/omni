package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = (*Keeper)(nil)

const approvedFromLimit = 100

func (k *Keeper) AttestationsFrom(ctx context.Context, req *types.AttestationsFromRequest) (*types.AttestationsFromResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	atts, err := k.ListAttestationsFrom(ctx, req.ChainId, req.ConfLevel, req.FromOffset, approvedFromLimit)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.AttestationsFromResponse{Attestations: atts}, nil
}

func (k *Keeper) LatestAttestation(ctx context.Context, req *types.LatestAttestationRequest) (*types.LatestAttestationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	chainVer := xchain.ChainVersion{ID: req.ChainId, ConfLevel: xchain.ConfLevel(req.ConfLevel)}

	att, ok, err := k.latestAttestation(ctx, chainVer)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if !ok {
		return nil, status.Error(codes.NotFound, "no approved attestations for chain")
	}

	resp, err := k.formAttestationResponse(ctx, att)
	if err != nil {
		return nil, errors.Wrap(err, "form response")
	}

	return &types.LatestAttestationResponse{Attestation: resp}, nil
}

func (k *Keeper) EarliestAttestation(ctx context.Context, req *types.EarliestAttestationRequest) (*types.EarliestAttestationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	chainVer := xchain.ChainVersion{ID: req.ChainId, ConfLevel: xchain.ConfLevel(req.ConfLevel)}

	att, ok, err := k.earliestAttestation(ctx, chainVer)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else if !ok {
		return nil, status.Error(codes.NotFound, "no approved attestations for chain")
	}

	resp, err := k.formAttestationResponse(ctx, att)
	if err != nil {
		return nil, errors.Wrap(err, "form response")
	}

	return &types.EarliestAttestationResponse{Attestation: resp}, nil
}

func (k *Keeper) formAttestationResponse(ctx context.Context, att *Attestation) (*types.Attestation, error) {
	sigs, err := k.getSigs(ctx, att.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	consensusChainID, err := getConsensusChainID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get consensus chain id")
	}

	return AttestationFromDB(att, consensusChainID, sigs), nil
}

func (k *Keeper) ListAllAttestations(ctx context.Context, req *types.ListAllAttestationsRequest) (*types.ListAllAttestationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	chainVer := xchain.ChainVersion{ID: req.ChainId, ConfLevel: xchain.ConfLevel(req.ConfLevel)}

	atts, err := k.listAllAttestations(ctx, chainVer, Status(req.Status), req.FromOffset)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.ListAllAttestationsResponse{Attestations: atts}, nil
}

func (k *Keeper) WindowCompare(ctx context.Context, req *types.WindowCompareRequest) (*types.WindowCompareResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	cmp, err := k.windowCompare(ctx, req.XChainVersion(), req.GetBlockOffset())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.WindowCompareResponse{Cmp: int32(cmp)}, nil
}

func getConsensusChainID(ctx context.Context) (uint64, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return netconf.ConsensusChainIDStr2Uint64(sdkCtx.ChainID())
}
