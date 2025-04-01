package tokens

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/netconf"
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

		if !token.IsMock {
			require.Equal(t, token.MaxSpend, spendBounds[token.Token][chainClass].MaxSpend)
			require.Equal(t, token.MinSpend, spendBounds[token.Token][chainClass].MinSpend)
		}

		seen[token] = true
		golden = append(golden, map[string]any{
			"name":        token.Name,
			"symbol":      token.Symbol,
			"address":     token.Address.Hex(),
			"maxSpend":    primaryStr(token.Token, token.MaxSpend),
			"minSpend":    primaryStr(token.Token, token.MinSpend),
			"chainId":     token.ChainID,
			"coingeckoId": token.CoingeckoID,
			"isMock":      token.IsMock,
		})
	}

	tutil.RequireGoldenJSON(t, golden)
}

func primaryStr(token tokenslib.Token, amount *big.Int) string {
	if amount == nil {
		return "nil"
	}

	return fmt.Sprintf("%.4f", tokenslib.ToPrimaryF64(token, amount))
}

func TestMaxSpendMinThreshold(t *testing.T) {
	t.Parallel()

	for _, token := range tokens {
		bounds, ok := spendBounds[token.Token][token.ChainClass]
		if !ok {
			continue
		}

		thresh, ok := eoa.GetSolverNetThreshold(eoa.RoleSolver, netconf.Mainnet, token.ChainID, token.Token)
		if !ok {
			continue
		}

		require.True(t, bi.GTE(thresh.MinBalance(), bounds.MaxSpend), "solver min balance must be greater than max spend: token=%s, min_bal=%s, max_spend=%s", token.Token, thresh.MinBalance(), bounds.MaxSpend)
	}
}
