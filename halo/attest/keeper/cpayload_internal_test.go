//nolint:lll // Fixtures are long
package keeper

import (
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

func TestVotesFromCommit(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0)

	var blockHash common.Hash
	fuzzer.Fuzz(&blockHash)

	// Generate attestations for following matrix: chains, vals, height batches
	const skipVal = 2 // Skip this validator
	chains := []uint64{100, 200}
	vals := []k1.PrivKey{k1.GenPrivKey(), k1.GenPrivKey(), k1.GenPrivKey()}
	batches := [][]uint64{{1, 2}, {3}, { /*empty*/ }}

	expected := make(map[xchain.Attestation]bool)

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
						BlockRoot: blockHash[:],
						Signature: &types.SigTuple{
							ValidatorAddress: addr[:],
							Signature:        sig[:],
						},
					}

					if i != skipVal {
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

	resp, ok, err := votesFromLastCommit(info)
	require.NoError(t, err)
	require.True(t, ok)

	for _, agg := range resp.Votes {
		for _, sig := range agg.Signatures {
			att := xchain.Attestation{
				BlockHeader: agg.BlockHeader.ToXChain(),
				BlockRoot:   common.Hash(agg.BlockRoot),
				Signature:   sig.ToXChain(),
			}

			require.True(t, expected[att], att)
			delete(expected, att)
		}
	}

	require.Empty(t, expected)
}
