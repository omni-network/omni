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

// AddVotes is called with all aggregated votes included in a new finalized block.
func (s msgServer) AddVotes(ctx context.Context, msg *types.MsgAddVotes,
) (*types.AddVotesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if sdkCtx.ExecMode() != sdk.ExecModeFinalize {
		return nil, errors.New("only allowed in finalize mode")
	}

	for _, aggVote := range msg.Votes {
		if err := aggVote.Verify(); err != nil {
			return nil, errors.Wrap(err, "verify aggVote")
		}
	}

	// Not verifying votes here since this block is now finalized, so it is too late to reject votes.
	err := s.Keeper.Add(ctx, msg)
	if err != nil {
		return nil, errors.Wrap(err, "add votes")
	}

	// Update the voter state with the local headers.
	localHeaders := headersByAddress(msg.Votes, s.voter.LocalAddress())
	logLocalVotes(ctx, localHeaders, "committed")
	if err := s.voter.SetCommitted(localHeaders); err != nil {
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

func logLocalVotes(ctx context.Context, headers []*types.AttestHeader, typ string) {
	if len(headers) == 0 {
		return
	}

	const limit = 5
	offsets := make(map[uint64][]string)
	for _, header := range headers {
		offset := offsets[header.SourceChainId]
		if len(offset) == limit {
			offset = append(offset, "...")
		} else if len(offset) < limit {
			offset = append(offset, strconv.FormatUint(header.AttestOffset, 10))
		} else {
			continue
		}
		offsets[header.SourceChainId] = offset
	}
	attrs := []any{
		slog.Int("votes", len(headers)),
	}
	for cid, hs := range offsets {
		attrs = append(attrs, slog.String(
			strconv.FormatUint(cid, 10),
			fmt.Sprint(hs),
		))
	}

	log.Debug(ctx, "Marked local votes as "+typ, attrs...)
}
