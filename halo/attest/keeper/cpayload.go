package keeper

import (
	"bytes"
	"context"
	"sort"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

var _ evmenginetypes.VoteExtensionProvider = (*Keeper)(nil)

// PrepareVotes returns the cosmosSDK transaction MsgAddVotes that will include all the validator votes included
// in the previous block's vote extensions into the attest module.
//
// Note that the commit is expected to be valid and only contains valid VEs from the previous block as
// provided by a trusted cometBFT. Some votes (contained inside VEs) may however be invalid, they are discarded.
func (k *Keeper) PrepareVotes(ctx context.Context, commit abci.ExtendedCommitInfo, commitHeight uint64) ([]sdk.Msg, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// The VEs in LastLocalCommit is expected to be valid
	if err := baseapp.ValidateVoteExtensions(sdkCtx, k.skeeper, 0, "", commit); err != nil {
		return nil, errors.Wrap(err, "validate extensions [BUG]")
	}

	// Verify and discard invalid votes.
	// Votes inside the VEs are NOT guaranteed to be valid, since
	// VerifyVoteExtension isn't called after quorum is reached.
	var allVotes []*types.Vote
	for _, vote := range commit.Votes {
		if vote.BlockIdFlag != cmtproto.BlockIDFlagCommit {
			continue // Skip non-committed votes
		}

		selected, _, err := k.parseAndVerifyVoteExtension(sdkCtx, vote.Validator.Address, vote.VoteExtension, commitHeight) //nolint:contextcheck // sdkCtx passed
		if err != nil {
			log.Warn(ctx, "Discarding invalid vote extension", err, log.Hex7("validator", vote.Validator.Address))
			continue
		}

		allVotes = append(allVotes, selected...)
	}

	votes, err := aggregateVotes(allVotes)
	if err != nil {
		return nil, err
	}

	return []sdk.Msg{&types.MsgAddVotes{
		Authority: authtypes.NewModuleAddress(types.ModuleName).String(),
		Votes:     votes,
	}}, nil
}

// aggregateVotes aggregates the provided attestations by block header.
func aggregateVotes(votes []*types.Vote) ([]*types.AggVote, error) {
	uniqueAggs := make(map[common.Hash]*types.AggVote)
	for _, vote := range votes {
		attRoot, err := vote.AttestationRoot()
		if err != nil {
			return nil, err
		}
		agg, ok := uniqueAggs[attRoot]
		if !ok {
			agg = &types.AggVote{
				AttestHeader: vote.AttestHeader,
				BlockHeader:  vote.BlockHeader,
				MsgRoot:      vote.MsgRoot,
			}
		}

		agg.Signatures = append(agg.Signatures, vote.Signature)
		uniqueAggs[attRoot] = agg
	}

	return sortAggregates(flattenAggs(uniqueAggs)), nil
}

// flattenAggs returns the values of the provided map.
func flattenAggs(aggsByRoot map[common.Hash]*types.AggVote) []*types.AggVote {
	aggs := make([]*types.AggVote, 0, len(aggsByRoot))
	for _, agg := range aggsByRoot {
		aggs = append(aggs, agg)
	}

	return aggs
}

// sortAggregates returns the provided aggregates in a deterministic order.
// Note the provided slice is also sorted in-place.
func sortAggregates(aggs []*types.AggVote) []*types.AggVote {
	sort.Slice(aggs, func(i, j int) bool {
		if aggs[i].AttestHeader.AttestOffset != aggs[j].AttestHeader.AttestOffset {
			return aggs[i].AttestHeader.AttestOffset < aggs[j].AttestHeader.AttestOffset
		}
		if aggs[i].BlockHeader.ChainId != aggs[j].BlockHeader.ChainId {
			return aggs[i].BlockHeader.ChainId < aggs[j].BlockHeader.ChainId
		}

		return bytes.Compare(aggs[i].BlockHeader.BlockHash, aggs[j].BlockHeader.BlockHash) < 0
	})

	return aggs
}
