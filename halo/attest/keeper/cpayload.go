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
	"golang.org/x/exp/maps"
)

var _ evmenginetypes.CPayloadProvider = (*Keeper)(nil)

func (k *Keeper) PreparePayload(ctx context.Context, height uint64, commit abci.ExtendedCommitInfo) ([]sdk.Msg, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err := baseapp.ValidateVoteExtensions(sdkCtx, k.skeeper, int64(height), sdkCtx.ChainID(), commit); err != nil {
		log.Error(ctx, "Cannot include invalid vote extensions in payload", err, "height", height)
		return nil, nil
	}

	msg, err := votesFromLastCommit(commit)
	if err != nil {
		// TODO(corver): Byzantine validators can stall block production by extending unaggregatable votes.
		//  We should probably aggregate based on uniq vote metadata hash, not just uniq block header.
		//  Same goes for Attestation table in keeper.
		return nil, err
	}

	return []sdk.Msg{msg}, nil
}

// votesFromLastCommit returns the aggregated votes contained in vote extensions
// of the last local commit.
func votesFromLastCommit(info abci.ExtendedCommitInfo) (*types.MsgAddVotes, error) {
	var allVotes []*types.Vote
	for _, vote := range info.Votes {
		if vote.BlockIdFlag != cmtproto.BlockIDFlagCommit {
			continue // Skip non block votes
		}
		votes, ok, err := votesFromExtension(vote.VoteExtension)
		if err != nil {
			return nil, err
		} else if !ok {
			continue
		}

		allVotes = append(allVotes, votes.Votes...)
	}

	aggVotes, err := aggregateVotes(allVotes)
	if err != nil {
		return nil, errors.Wrap(err, "aggregate votes")
	}

	return &types.MsgAddVotes{
		Authority: authtypes.NewModuleAddress(types.ModuleName).String(),
		Votes:     aggVotes,
	}, nil
}

// aggregateVotes aggregates the provided attestations by block header.
func aggregateVotes(votes []*types.Vote) ([]*types.AggVote, error) {
	aggsByHeader := make(map[xchain.BlockHeader]*types.AggVote) // map[BlockHash]Attestation
	for _, vote := range votes {
		header := vote.BlockHeader.ToXChain()
		agg, ok := aggsByHeader[header]
		if !ok {
			agg = &types.AggVote{
				BlockHeader:    vote.BlockHeader,
				BlockRoot:      vote.BlockRoot,
				MsgOffsets:     vote.MsgOffsets,
				ReceiptOffsets: vote.ReceiptOffsets,
			}
		} else if !bytes.Equal(agg.BlockRoot, vote.BlockRoot) {
			return nil, errors.New("conflicting vote block roots", log.Hex7("agg", agg.BlockRoot), log.Hex7("vote", vote.BlockRoot))
		} else if !msgOffsetsEqual(agg.MsgOffsets, vote.MsgOffsets) {
			return nil, errors.New("conflicting vote message offsets")
		} else if !receiptOffsetsEqual(agg.ReceiptOffsets, vote.ReceiptOffsets) {
			return nil, errors.New("conflicting vote receipt offsets")
		}

		agg.Signatures = append(agg.Signatures, vote.Signature)
		aggsByHeader[header] = agg
	}

	return sortAggregates(maps.Values(aggsByHeader)), nil
}

// sortAggregates returns the provided aggregates in a deterministic order.
// Note the provided slice is also sorted in-place.
func sortAggregates(aggs []*types.AggVote) []*types.AggVote {
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

// votesFromExtension returns the attestations contained in the vote extension, or false if none or an error.
func votesFromExtension(voteExtension []byte) (*types.Votes, bool, error) {
	if len(voteExtension) == 0 {
		return nil, false, nil
	}

	resp := new(types.Votes)
	if err := proto.Unmarshal(voteExtension, resp); err != nil {
		return nil, false, errors.Wrap(err, "decode vote extension")
	}

	return resp, true, nil
}

func msgOffsetsEqual(a, b []*types.MsgOffset) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i].DestChainId != b[i].DestChainId || a[i].StreamOffset != b[i].StreamOffset {
			return false
		}
	}

	return true
}

// receiptOffsetsEqual returns true if the provided receipt offsets are equal.
func receiptOffsetsEqual(a, b []*types.ReceiptOffset) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i].SourceChainId != b[i].SourceChainId || a[i].StreamOffset != b[i].StreamOffset {
			return false
		}
	}

	return true
}
