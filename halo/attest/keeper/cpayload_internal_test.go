//nolint:lll // Fixtures are long
package keeper

import (
	"context"
	"testing"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	types1 "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cosmos/gogoproto/proto"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestVotesFromCommitNonUnique(t *testing.T) {
	t.Parallel()

	const chainID = 100
	const height = 200
	var valAddr common.Address
	var valSig [65]byte

	newVote := func(hash, attRoot common.Hash) *types.Vote {
		return &types.Vote{
			BlockHeader: &types.BlockHeader{
				ChainId: chainID,
				Height:  height,
				Hash:    hash[:],
			},
			AttestationRoot: attRoot[:],
			Signature: &types.SigTuple{
				ValidatorAddress: valAddr[:],
				Signature:        valSig[:],
			},
		}
	}

	hash1 := common.BytesToHash([]byte("hash1"))
	hash2 := common.BytesToHash([]byte("hash2"))
	root1 := common.BytesToHash([]byte("root1"))
	root2 := common.BytesToHash([]byte("root2"))

	// Same chainID and Height, but different hash and attRoots combinations.
	aggs := aggregateVotes([]*types.Vote{
		newVote(hash1, root1),
		newVote(hash2, root2),
		newVote(hash1, root2),
	})

	// Result in different aggregates
	require.Len(t, aggs, 3)
}

func TestVotesFromCommit(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0)

	var blockHash common.Hash
	fuzzer.Fuzz(&blockHash)

	// Generate attestations for following matrix: chains, vals, height batches
	const skipVal = 2     // Skip this validator
	const skipChain = 300 // Skip this chain (out of window)
	chains := []uint64{100, 200, 300}
	vals := []k1.PrivKey{k1.GenPrivKey(), k1.GenPrivKey(), k1.GenPrivKey()}
	batches := [][]uint64{{1, 2}, {3}, { /*empty*/ }}

	expected := make(map[xchain.Vote]bool)
	total := 2 * 3 // 2 chains * 3 heights

	var evotes []abci.ExtendedVoteInfo
	for _, chain := range chains {
		for i, val := range vals {
			flag := types1.BlockIDFlagCommit
			if i == skipVal {
				flag = types1.BlockIDFlagAbsent
			}

			for _, batch := range batches {
				var votes []*types.Vote
				for _, height := range batch {
					addr, err := k1util.PubKeyToAddress(val.PubKey())
					require.NoError(t, err)

					var sig xchain.Signature65
					fuzzer.Fuzz(&sig)

					vote := &types.Vote{
						BlockHeader: &types.BlockHeader{
							ChainId: chain,
							Height:  height,
							Hash:    blockHash[:],
						},
						AttestationRoot: blockHash[:],
						Signature: &types.SigTuple{
							ValidatorAddress: addr[:],
							Signature:        sig[:],
						},
					}

					if i != skipVal && chain != skipChain {
						expected[vote.ToXChain()] = true
					}
					votes = append(votes, vote)
				}

				bz, err := proto.Marshal(&types.Votes{
					Votes: votes,
				})
				require.NoError(t, err)

				evotes = append(evotes, abci.ExtendedVoteInfo{
					VoteExtension: bz,
					BlockIdFlag:   flag,
				})
			}
		}
	}

	info := abci.ExtendedCommitInfo{
		Round: 99,
		Votes: evotes,
	}

	comparer := func(ctx context.Context, chainID uint64, height uint64) (int, error) {
		if chainID == skipChain {
			return 1, nil
		}

		return 0, nil
	}

	resp, err := votesFromLastCommit(context.Background(), comparer, info)
	require.NoError(t, err)

	require.Len(t, resp.Votes, total)

	for _, agg := range resp.Votes {
		for _, sig := range agg.Signatures {
			att := xchain.Vote{
				BlockHeader:     agg.BlockHeader.ToXChain(),
				AttestationRoot: common.Hash(agg.AttestationRoot),
				Signature:       sig.ToXChain(),
			}

			require.True(t, expected[att], att)
			delete(expected, att)
		}
	}

	require.Empty(t, expected)
}
