package keeper

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	*Keeper
	types.UnimplementedMsgServiceServer
}

func (s msgServer) AddVotes(ctx context.Context, msg *types.MsgAddVotes,
) (*types.AddVotesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if sdkCtx.ExecMode() != sdk.ExecModeFinalize {
		return nil, errors.New("only allowed in finalize mode")
	}

	// Update the voter state with the local headers.
	localHeaders := headersByAddress(msg.Votes, s.voter.LocalAddress())
	if err := s.voter.SetCommitted(localHeaders); err != nil {
		return nil, errors.Wrap(err, "set committed")
	}

	err := s.Keeper.Add(ctx, msg)
	if err != nil {
		return nil, errors.Wrap(err, "add votes")
	}

	if len(msg.Votes) == 0 {
		return &types.AddVotesResponse{}, nil
	}

	// Make nice logs
	heights := make(map[uint64][]uint64)
	for _, header := range localHeaders {
		heights[header.ChainId] = append(heights[header.ChainId], header.Height)
	}
	attrs := []any{
		slog.Int("attestations", len(localHeaders)),
		log.Hex7("validator", s.voter.LocalAddress().Bytes()),
	}
	for cid, hs := range heights {
		attrs = append(attrs, slog.String(
			strconv.FormatUint(cid, 10),
			fmt.Sprint(hs),
		))
	}

	log.Debug(ctx, "Marked local votes as committed", attrs...)

	return &types.AddVotesResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServiceServer = msgServer{}

// headersByAddress returns the attestations for the provided address.
func headersByAddress(aggregates []*types.AggVote, address common.Address) []*types.BlockHeader {
	var filtered []*types.BlockHeader
	for _, agg := range aggregates {
		for _, sig := range agg.Signatures {
			if bytes.Equal(sig.ValidatorAddress, address[:]) {
				filtered = append(filtered, agg.BlockHeader)
				break // Continue to the next aggregate.
			}
		}
	}

	return filtered
}
