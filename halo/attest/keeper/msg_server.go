package keeper

import (
	"bytes"
	"context"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type msgServer struct {
	*Keeper
	types.UnimplementedMsgServiceServer
}

// AddVotes is called with all aggregated votes included in a new finalized block.
func (s msgServer) AddVotes(ctx context.Context, msg *types.MsgAddVotes,
) (*types.AddVotesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if sdkCtx.ExecMode() != sdk.ExecModeFinalize {
		return nil, errors.New("only allowed in finalize mode")
	}

	if msg.Authority != authtypes.NewModuleAddress(types.ModuleName).String() {
		return nil, errors.New("invalid authority")
	}

	for _, aggVote := range msg.Votes {
		if err := aggVote.Verify(); err != nil {
			return nil, errors.Wrap(err, "verify aggVote")
		}
	}

	// Not verifying votes here since this block is now finalized, so it is too late to reject votes.
	err := s.Add(ctx, msg)
	if err != nil {
		return nil, errors.Wrap(err, "add votes")
	}

	// Update the voter state with the local headers.
	localHeaders := headersByAddress(msg.Votes, s.voter.LocalAddress())
	if err := s.voter.SetCommitted(ctx, localHeaders); err != nil {
		return nil, errors.Wrap(err, "set committed")
	}

	return &types.AddVotesResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServiceServer = msgServer{}

// headersByAddress returns the attestations for the provided address.
func headersByAddress(aggregates []*types.AggVote, address common.Address) []*types.AttestHeader {
	var filtered []*types.AttestHeader
	for _, agg := range aggregates {
		for _, sig := range agg.Signatures {
			if bytes.Equal(sig.ValidatorAddress, address[:]) {
				filtered = append(filtered, agg.AttestHeader)
				break // Continue to the next aggregate.
			}
		}
	}

	return filtered
}
