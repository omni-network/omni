package engine

import (
	"math/big"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"

	fuzz "github.com/google/gofuzz"
)

// NewFuzzer returns a new fuzzer for valid Engine API types.
func NewFuzzer() *fuzz.Fuzzer {
	f := fuzz.New().NilChance(0)
	f.Funcs(
		func(h *types.Header, c fuzz.Continue) {
			c.FuzzNoCustom(h)
			h.Difficulty = common.Big0
			h.UncleHash = types.EmptyUncleHash
			h.WithdrawalsHash = nil
			h.ParentBeaconRoot = nil
			h.Nonce = types.BlockNonce{}
		},
		func(b *types.Block, c fuzz.Continue) {
			var header types.Header
			c.Fuzz(&header)

			block := types.NewBlock(&header, nil, nil, nil, trie.NewStackTrie(nil))

			*b = *block
		},
		func(b *engine.ExecutableData, c fuzz.Continue) {
			block := new(types.Block)
			c.Fuzz(block)

			env := engine.BlockToExecutableData(block, big.NewInt(0), nil)

			*b = *env.ExecutionPayload
		},
	)

	return f
}
