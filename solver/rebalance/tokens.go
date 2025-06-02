package rebalance

import (
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tokens"
)

// swappable is a set of tokens that can be swapped to/from USDC.
var swappable = []tokens.Token{
	// USDC
	mustToken(evmchain.IDEthereum, tokens.USDC),
	mustToken(evmchain.IDBase, tokens.USDC),
	mustToken(evmchain.IDArbitrumOne, tokens.USDC),
	mustToken(evmchain.IDOptimism, tokens.USDC),

	// USDT
	mustToken(evmchain.IDEthereum, tokens.USDT),
	mustToken(evmchain.IDArbitrumOne, tokens.USDT),
	mustToken(evmchain.IDOptimism, tokens.USDT),

	// ETH
	mustToken(evmchain.IDEthereum, tokens.ETH),
	mustToken(evmchain.IDBase, tokens.ETH),
	mustToken(evmchain.IDArbitrumOne, tokens.ETH),
	mustToken(evmchain.IDOptimism, tokens.ETH),

	// WSETH
	mustToken(evmchain.IDEthereum, tokens.WSTETH),
	mustToken(evmchain.IDBase, tokens.WSTETH),
}

// SwappableTokens returns a list of swappable tokens.
func SwappableTokens() []tokens.Token {
	return swappable
}

// SwappableTokensByChain returns a list of swappable tokens for a given chain ID.
func SwappableTokensByChain(chainID uint64) []tokens.Token {
	var tkns []tokens.Token
	for _, t := range swappable {
		if t.ChainID == chainID {
			tkns = append(tkns, t)
		}
	}

	return tkns
}

func mustToken(chainID uint64, asset tokens.Asset) tokens.Token {
	tkn, ok := tokens.ByAsset(chainID, asset)
	if !ok {
		panic("token not found")
	}

	return tkn
}
