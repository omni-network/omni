package usdt0

import (
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tokens"
)

var (
	// tokenByChain maps chain ID to TetherZero token (or canonical USDT, for mainnet).
	tokenByChain = map[uint64]tokens.Token{
		evmchain.IDEthereum:    mustToken(evmchain.IDEthereum, tokens.USDT),
		evmchain.IDOptimism:    mustToken(evmchain.IDOptimism, tokens.USDT0),
		evmchain.IDArbitrumOne: mustToken(evmchain.IDArbitrumOne, tokens.USDT0),
		evmchain.IDHyperEVM:    mustToken(evmchain.IDHyperEVM, tokens.USDT0),
	}
)

func mustToken(chainID uint64, asset tokens.Asset) tokens.Token {
	tkn, ok := tokens.ByAsset(chainID, asset)
	if !ok {
		panic("token not found")
	}

	return tkn
}
