//nolint:lll // Fixtures are long
package keeper

import (
	"encoding/hex"
	"encoding/json"
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestAggsFromCommit(t *testing.T) {
	t.Parallel()
	// Got this fixture from SmokeTest
	infoBZ := `{
 "votes": [
  {
   "validator": {
    "address": "9F1DAr20DN7s2Unon0tEhoJ4y4g=",
    "power": 1
   },
   "vote_extension": "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",
   "extension_signature": "yb/xqE8JHvlXEqcIaC+tZ3CpTIp4EoKMBR8t13dT/IRknZnYGK+sda+92TXzF8Fjk7RlHkfmz2AMg0PAvQ7h5A==",
   "block_id_flag": 2
  }
 ]
}
`
	info := new(abci.ExtendedCommitInfo)
	err := json.Unmarshal([]byte(infoBZ), info)
	require.NoError(t, err)

	aggAtts, ok, err := aggregatesFromLastCommit(*info)
	require.NoError(t, err)
	require.True(t, ok)

	type block struct {
		ChainID uint64
		Height  uint64
	}
	expected := map[block]bool{
		{ChainID: 100, Height: 0}: true,
		{ChainID: 100, Height: 1}: true,
		{ChainID: 100, Height: 2}: true,
		{ChainID: 200, Height: 0}: true,
		{ChainID: 200, Height: 1}: true,
		{ChainID: 200, Height: 2}: true,
		{ChainID: 999, Height: 0}: true,
		{ChainID: 999, Height: 1}: true,
		{ChainID: 999, Height: 2}: true,
	}

	require.Equal(t, len(expected), len(aggAtts.Aggregates))
	for _, agg := range aggAtts.Aggregates {
		require.Len(t, agg.Signatures, 1)
		b := block{
			ChainID: agg.BlockHeader.ChainId,
			Height:  agg.BlockHeader.Height,
		}
		require.True(t, expected[b], b)
		delete(expected, b)
	}

	require.Empty(t, expected)
}

func TestAggsFromCommit2(t *testing.T) {
	t.Parallel()
	// Got this fixture from E2E tests
	infoBZ := `{"votes":[{"validator":{"address":"Bcn9WsCUB1EubaiqVurLh7ifqbY=","power":1},"vote_extension":"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","extension_signature":"1spgeubk7rihfX1HxeYb0A8yF6FrnEpAM6h/YBHdZdhelEZ3B2zqacgCIM0zhN8GUMeI0mNf0lU+7MJbNmHciQ==","block_id_flag":2},{"validator":{"address":"12SOKLu1Z+ewBXLrClwbDH2S8zY=","power":1},"vote_extension":"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","extension_signature":"y6f+Qzug2mc6rivalsAKQxMZKMEkMMsnAi5bSoQe8Q47ChKcUm7uyvCLq19JvIMojbdZO5yfymj93onBjGwCEg==","block_id_flag":2}]}`

	info := new(abci.ExtendedCommitInfo)
	err := json.Unmarshal([]byte(infoBZ), info)
	require.NoError(t, err)

	val1, err := hex.DecodeString("edb2037ff27974dfccac23e8bc230f7c1c13fb37")
	require.NoError(t, err)
	val2, err := hex.DecodeString("37457f8f6d6b67ebba5282417a14a75d55a23830")
	require.NoError(t, err)

	aggAtts, ok, err := aggregatesFromLastCommit(*info)
	require.NoError(t, err)
	require.True(t, ok)

	type att struct {
		ChainID uint64
		Height  uint64
		Val     common.Address
	}

	// Fixture contains attestations from both vals for chain 100+200 for blocks 0-5.
	expected := make(map[att]bool)
	makeExpected := func(val []byte) {
		for _, chainID := range []uint64{100, 200} {
			for _, height := range []uint64{0, 1, 2, 3, 4, 5} {
				expected[att{ChainID: chainID, Height: height, Val: common.BytesToAddress(val)}] = true
			}
		}
	}
	makeExpected(val1)
	makeExpected(val2)
	delete(expected, att{ChainID: 100, Height: 5, Val: common.BytesToAddress(val2)}) // Except val2 for this block.

	for _, agg := range aggAtts.Aggregates {
		for _, sig := range agg.Signatures {
			a := att{
				Val:     common.BytesToAddress(sig.ValidatorAddress),
				ChainID: agg.BlockHeader.ChainId,
				Height:  agg.BlockHeader.Height,
			}
			require.True(t, expected[a], a)
			delete(expected, a)
		}
	}

	require.Empty(t, expected)
}
