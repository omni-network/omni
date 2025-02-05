package app

import (
	"testing"

	"github.com/omni-network/omni/lib/tutil"
)

// TestTokens ensures solver toke list does not change without explicit golden update.
func TestTokens(t *testing.T) {
	t.Parallel()

	golden := []map[string]any{}
	seen := make(map[Token]bool)

	for _, token := range tokens {
		if seen[token] {
			t.Errorf("duplicate token: %v", token)
		}

		seen[token] = true
		golden = append(golden, map[string]any{
			"name":        token.Name,
			"symbol":      token.Symbol,
			"address":     token.Address.Hex(),
			"chainId":     token.ChainID,
			"coingeckoId": token.CoingeckoID,
			"isMock":      token.IsMock,
		})
	}

	tutil.RequireGoldenJSON(t, golden)
}
