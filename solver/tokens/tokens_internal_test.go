package tokens

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/bi"
	tokenslib "github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

// TestTokens ensures solver toke list does not change without explicit golden update.
func TestTokens(t *testing.T) {
	t.Parallel()

	golden := []map[string]any{}
	seen := make(map[Token]bool)

	for _, token := range tokens {
		if seen[token] {
			t.Errorf("duplicate token: %v", token)
		}

		chainClass := mustChainClass(token.ChainID)

		if token.Symbol == tokenslib.ETH.Symbol && !token.IsMock {
			require.Equal(t, token.MaxSpend, spendBounds[tokenslib.ETH][chainClass].MaxSpend)
			require.Equal(t, token.MinSpend, spendBounds[tokenslib.ETH][chainClass].MinSpend)
		}

		if token.Symbol == tokenslib.WSTETH.Symbol && !token.IsMock {
			require.Equal(t, token.MaxSpend, spendBounds[tokenslib.WSTETH][chainClass].MaxSpend)
			require.Equal(t, token.MinSpend, spendBounds[tokenslib.WSTETH][chainClass].MinSpend)
		}

		if token.Symbol == tokenslib.OMNI.Symbol && !token.IsMock {
			require.Equal(t, token.MaxSpend, spendBounds[tokenslib.OMNI][chainClass].MaxSpend)
			require.Equal(t, token.MinSpend, spendBounds[tokenslib.OMNI][chainClass].MinSpend)
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
	require.Equal(t, "1.0000", etherStr(spendBounds[tokenslib.ETH]["mainnet"].MaxSpend))
	require.Equal(t, "0.0010", etherStr(spendBounds[tokenslib.ETH]["mainnet"].MinSpend))
	require.Equal(t, "1.0000", etherStr(spendBounds[tokenslib.ETH]["testnet"].MaxSpend))
	require.Equal(t, "0.0010", etherStr(spendBounds[tokenslib.ETH]["testnet"].MinSpend))

	require.Equal(t, "4.0000", etherStr(spendBounds[tokenslib.WSTETH]["mainnet"].MaxSpend))
	require.Equal(t, "0.0010", etherStr(spendBounds[tokenslib.WSTETH]["mainnet"].MinSpend))
	require.Equal(t, "1.0000", etherStr(spendBounds[tokenslib.WSTETH]["testnet"].MaxSpend))
	require.Equal(t, "0.0010", etherStr(spendBounds[tokenslib.WSTETH]["testnet"].MinSpend))

	require.Equal(t, "120000.0000", etherStr(spendBounds[tokenslib.OMNI]["mainnet"].MaxSpend))
	require.Equal(t, "0.1000", etherStr(spendBounds[tokenslib.OMNI]["mainnet"].MinSpend))
	require.Equal(t, "1000.0000", etherStr(spendBounds[tokenslib.OMNI]["testnet"].MaxSpend))
	require.Equal(t, "0.1000", etherStr(spendBounds[tokenslib.OMNI]["testnet"].MinSpend))
}

func etherStr(amount *big.Int) string {
	if amount == nil {
		return "nil"
	}

	return fmt.Sprintf("%.4f", bi.ToEtherF64(amount))
}
