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

	expected := make(map[string]bool)

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
						MsgOffsets: []*types.MsgOffset{
							{DestChainId: chain, StreamOffset: height},
							{DestChainId: chain + 1, StreamOffset: height + 1},
						},
						ReceiptOffsets: []*types.ReceiptOffset{
							{SourceChainId: chain, StreamOffset: height},
							{SourceChainId: chain + 1, StreamOffset: height + 1},
						},
					}

					if i != skipVal {
						expected[proto.MarshalTextString(vote)] = true
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

	resp, err := votesFromLastCommit(info)
	require.NoError(t, err)

	for _, agg := range resp.Votes {
		for _, sig := range agg.Signatures {
			actual := &types.Vote{
				BlockHeader:    agg.BlockHeader,
				BlockRoot:      agg.BlockRoot,
				Signature:      sig,
				MsgOffsets:     agg.MsgOffsets,
				ReceiptOffsets: agg.ReceiptOffsets,
			}

			str := proto.MarshalTextString(actual)

			require.True(t, expected[str], str)
			delete(expected, str)
		}
	}

	require.Empty(t, expected)
}
