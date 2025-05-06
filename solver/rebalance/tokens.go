package rebalance

import (
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tokens"
)

// rebalanceTokens is a list of tokens supported by the rebalance service.
var rebalanceTokens = []tokens.Token{
	// USDC
	mustToken(evmchain.IDEthereum, tokens.USDC),
	mustToken(evmchain.IDBase, tokens.USDC),

	// WSETH
	mustToken(evmchain.IDEthereum, tokens.WSTETH),
	mustToken(evmchain.IDBase, tokens.WSTETH),
}

func Tokens() []tokens.Token {
	return rebalanceTokens
}

func TokensByChain(chainID uint64) []tokens.Token {
	var tokens []tokens.Token
	for _, t := range rebalanceTokens {
		if t.ChainID == chainID {
			tokens = append(tokens, t)
		}
	}

	return tokens
}

func mustToken(chainID uint64, asset tokens.Asset) tokens.Token {
	tkn, ok := tokens.ByAsset(chainID, asset)
	if !ok {
		panic("token not found")
	}

	return tkn
}
