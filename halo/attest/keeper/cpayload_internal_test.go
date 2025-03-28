//nolint:lll // Fixtures are long
package keeper

import (
	"testing"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestVotesFromCommitNonUnique(t *testing.T) {
	t.Parallel()

	const chainID = 100
	const offset = 200
	const height = 300
	var valAddr common.Address
	var valSig [65]byte

	newVote := func(hash, msgRoot common.Hash) *types.Vote {
		return &types.Vote{
			AttestHeader: &types.AttestHeader{
				ConsensusChainId: 0,
				SourceChainId:    chainID,
				AttestOffset:     offset,
			},
			BlockHeader: &types.BlockHeader{
				ChainId:     chainID,
				BlockHeight: height,
				BlockHash:   hash[:],
			},
			MsgRoot: msgRoot[:],
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

	// Same chainID and Height, but different hash and msgRoots combinations.
	aggs, err := aggregateVotes([]*types.Vote{
		newVote(hash1, root1),
		newVote(hash2, root2),
		newVote(hash1, root2),
	})
	require.NoError(t, err)

	// Result in different aggregates
	require.Len(t, aggs, 3)
}

func TestAggregateVotes(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0)

	var blockHash common.Hash
	fuzzer.Fuzz(&blockHash)

	// Generate attestations for following matrix: chains, vals, offset batches
	chains := []uint64{100, 200}
	vals := []k1.PrivKey{k1.GenPrivKey(), k1.GenPrivKey(), k1.GenPrivKey()}
	batches := [][]uint64{{1, 2}, {3}, { /*empty*/ }}

	expected := make(map[[32]byte]map[xchain.SigTuple]bool)
	total := 2 * 3 // 2 chains * 3 heights

	var allVotes []*types.Vote
	for _, chain := range chains {
		for _, val := range vals {
			for _, batch := range batches {
				for _, offset := range batch {
					addr, err := k1util.PubKeyToAddress(val.PubKey())
					require.NoError(t, err)

					var sig xchain.Signature65
					fuzzer.Fuzz(&sig)

					vote := &types.Vote{
						AttestHeader: &types.AttestHeader{
							SourceChainId: chain,
							ConfLevel:     uint32(xchain.ConfFinalized),
							AttestOffset:  offset,
						},
						BlockHeader: &types.BlockHeader{
							ChainId:     chain,
							BlockHeight: offset * 2,
							BlockHash:   blockHash[:],
						},
						MsgRoot: blockHash[:],
						Signature: &types.SigTuple{
							ValidatorAddress: addr[:],
							Signature:        sig[:],
						},
					}

					attRoot, err := vote.AttestationRoot()
					require.NoError(t, err)

					if _, ok := expected[attRoot]; !ok {
						expected[attRoot] = make(map[xchain.SigTuple]bool)
					}
					expected[attRoot][xchain.SigTuple{
						ValidatorAddress: addr,
						Signature:        sig,
					}] = true

					allVotes = append(allVotes, vote)
				}
			}
		}
	}

	aggs, err := aggregateVotes(allVotes)
	require.NoError(t, err)

	require.Len(t, aggs, total)

	for _, agg := range aggs {
		attRoot, err := agg.AttestationRoot()
		require.NoError(t, err)
		for _, s := range agg.Signatures {
			sig, err := s.ToXChain()
			require.NoError(t, err)
			require.True(t, expected[attRoot][sig], "not found", agg, sig)
			delete(expected[attRoot], sig)
			if len(expected[attRoot]) == 0 {
				delete(expected, attRoot)
			}
		}
	}

	require.Empty(t, expected)
}
