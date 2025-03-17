package app

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/bi"
	tokenslib "github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
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

		if token.Symbol == tokenslib.ETH.Symbol && !token.IsMock {
			require.Equal(t, token.MaxSpend, maxETHSpend)
			require.Equal(t, token.MinSpend, minETHSpend)
		}

		if token.Symbol == tokenslib.WSTETH.Symbol && !token.IsMock {
			require.Equal(t, token.MaxSpend, maxWSTETHSpend)
			require.Equal(t, token.MinSpend, minWSTETHSpend)
		}

		if token.Symbol == tokenslib.OMNI.Symbol && !token.IsMock {
			require.Equal(t, token.MaxSpend, maxOMNISpend)
			require.Equal(t, token.MinSpend, minOMNISpend)
		}

		seen[token] = true
		golden = append(golden, map[string]any{
			"name":        token.Name,
			"symbol":      token.Symbol,
			"address":     token.Address.Hex(),
			"maxSpend":    etherStr(token.MaxSpend),
			"minSpend":    etherStr(token.MinSpend),
			"chainId":     token.ChainID,
			"coingeckoId": token.CoingeckoID,
			"isMock":      token.IsMock,
		})
	}

	tutil.RequireGoldenJSON(t, golden)

	// check max / min
	require.Equal(t, "1.0000", etherStr(maxETHSpend))
	require.Equal(t, "0.0010", etherStr(minETHSpend))
	require.Equal(t, "1.0000", etherStr(maxWSTETHSpend))
	require.Equal(t, "0.0010", etherStr(minWSTETHSpend))
	require.Equal(t, "1000.0000", etherStr(maxOMNISpend))
	require.Equal(t, "0.1000", etherStr(minOMNISpend))
}

func etherStr(amount *big.Int) string {
	if amount == nil {
		return "nil"
	}

	return fmt.Sprintf("%.4f", bi.ToEtherF64(amount))
}
