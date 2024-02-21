package keeper

import (
	"bytes"
	"context"
	"sort"

	"github.com/omni-network/omni/halo/attest/types"
	evmenginetypes "github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/gogoproto/proto"
)

var _ evmenginetypes.CPayloadProvider = Keeper{}

func (k Keeper) PreparePayload(ctx context.Context, height uint64, commit abci.ExtendedCommitInfo) ([]sdk.Msg, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err := baseapp.ValidateVoteExtensions(sdkCtx, k.skeeper, int64(height), sdkCtx.ChainID(), commit); err != nil {
		log.Error(ctx, "Cannot include invalid vote extensions in payload", err, "height", height)
		return nil, nil
	}

	aggAtts, ok, err := aggregatesFromLastCommit(commit)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, nil
	}

	return []sdk.Msg{aggAtts}, nil
}

// aggregatesFromLastCommit returns the aggregateAtts attestations contained in vote extensions
// of the last local commit.
func aggregatesFromLastCommit(info abci.ExtendedCommitInfo) (*types.MsgAggAttestations, bool, error) {
	var allAtts []*types.Attestation
	for _, vote := range info.Votes {
		if vote.BlockIdFlag != cmtproto.BlockIDFlagCommit {
			continue // Skip non block votes
		}
		atts, ok, err := attestationsFromVoteExt(vote.VoteExtension)
		if err != nil {
			return nil, false, err
		} else if !ok {
			continue
		}

		allAtts = append(allAtts, atts.Attestations...)
	}

	return &types.MsgAggAttestations{
		Authority:  authtypes.NewModuleAddress(types.ModuleName).String(),
		Aggregates: aggregateAtts(allAtts),
	}, len(allAtts) > 0, nil
}

// aggregateAtts aggregates the provided attestations by block header.
func aggregateAtts(atts []*types.Attestation) []*types.AggAttestation {
	aggsByHeader := make(map[xchain.BlockHeader]*types.AggAttestation) // map[BlockHash]AggAttestation
	for _, att := range atts {
		header := att.BlockHeader.ToXChain()
		agg, ok := aggsByHeader[header]
		if !ok {
			agg = &types.AggAttestation{
				BlockHeader: att.BlockHeader,
				BlockRoot:   att.BlockRoot,
			}
		}

		agg.Signatures = append(agg.Signatures, att.Signature)
		aggsByHeader[header] = agg
	}

	return flattenAggsByHeader(aggsByHeader)
}

// flattenAggsByHeader returns the provided map of aggregates by header as a slice in a deterministic order.
func flattenAggsByHeader(aggsByHeader map[xchain.BlockHeader]*types.AggAttestation) []*types.AggAttestation {
	aggs := make([]*types.AggAttestation, 0, len(aggsByHeader))
	for _, agg := range aggsByHeader {
		aggs = append(aggs, agg)
	}

	return sortAggregates(aggs)
}

// sortAggregates returns the provided aggregates in a deterministic order.
// Note the provided slice is also sorted in-place.
func sortAggregates(aggs []*types.AggAttestation) []*types.AggAttestation {
	sort.Slice(aggs, func(i, j int) bool {
		if aggs[i].BlockHeader.Height != aggs[j].BlockHeader.Height {
			return aggs[i].BlockHeader.Height < aggs[j].BlockHeader.Height
		}
		if aggs[i].BlockHeader.ChainId != aggs[j].BlockHeader.ChainId {
			return aggs[i].BlockHeader.ChainId < aggs[j].BlockHeader.ChainId
		}

		return bytes.Compare(aggs[i].BlockHeader.Hash, aggs[j].BlockHeader.Hash) < 0
	})

	return aggs
}

// attestationsFromVoteExt returns the attestations contained in the vote extension, or false if none or an error.
func attestationsFromVoteExt(voteExtension []byte) (*types.Attestations, bool, error) {
	if len(voteExtension) == 0 {
		return nil, false, nil
	}

	resp := new(types.Attestations)
	if err := proto.Unmarshal(voteExtension, resp); err != nil {
		return nil, false, errors.Wrap(err, "decode vote extension")
	}

	return resp, true, nil
}
