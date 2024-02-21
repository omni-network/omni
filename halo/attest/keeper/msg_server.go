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
	Keeper
	types.UnimplementedMsgServiceServer
}

func (s msgServer) AddAggAttestations(ctx context.Context, msg *types.MsgAggAttestations,
) (*types.AddAggAttestationsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if sdkCtx.ExecMode() != sdk.ExecModeFinalize {
		return nil, errors.New("only allowed in finalize mode")
	}

	// Update the attester state with the local headers.
	localHeaders := headersByAddress(msg.Aggregates, s.attester.LocalAddress())
	if err := s.attester.SetCommitted(localHeaders); err != nil {
		return nil, errors.Wrap(err, "set committed")
	}

	err := s.Keeper.Add(ctx, msg)
	if err != nil {
		return nil, errors.Wrap(err, "add aggregate attestation")
	}

	if len(msg.Aggregates) == 0 {
		return &types.AddAggAttestationsResponse{}, nil
	}

	// Make nice logs
	heights := make(map[uint64][]uint64)
	for _, header := range localHeaders {
		heights[header.ChainId] = append(heights[header.ChainId], header.Height)
	}
	attrs := []any{
		slog.Int("attestations", len(localHeaders)),
		log.Hex7("validator", s.attester.LocalAddress().Bytes()),
	}
	for cid, hs := range heights {
		attrs = append(attrs, slog.String(
			strconv.FormatUint(cid, 10),
			fmt.Sprint(hs),
		))
	}

	log.Debug(ctx, "Marked local attestations as committed", attrs...)

	return &types.AddAggAttestationsResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServiceServer = msgServer{}

// headersByAddress returns the attestations for the provided address.
func headersByAddress(aggregates []*types.AggAttestation, address common.Address) []*types.BlockHeader {
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
