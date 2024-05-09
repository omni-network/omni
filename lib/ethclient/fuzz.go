package ethclient

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"

	fuzz "github.com/google/gofuzz"
)

// NewFuzzer returns a new fuzzer for valid ethereum types.
// If seed is zero, it uses current nano time as the seed.
func NewFuzzer(seed int64) *fuzz.Fuzzer {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	f := fuzz.NewWithSeed(seed).NilChance(0)
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

			block := types.NewBlock(&header, nil, nil, trie.NewStackTrie(nil))

			*b = *block //nolint:govet // Copy lock isn't a problem since we are creating a new object.
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
