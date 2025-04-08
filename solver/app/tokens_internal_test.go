package app

import (
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

// TestTokens ensures solver toke list does not change without explicit golden update.
func TestTokens(t *testing.T) {
	t.Parallel()

	golden := []map[string]any{}
	seen := make(map[tokens.Token]bool)

	for _, tkn := range tokens.All() {
		if !IsSupportedToken(tkn) {
			continue
		}

		if seen[tkn] {
			t.Errorf("duplicate token: %v", tkn)
		}

		bounds := GetSpendBounds(tkn)
		if !tkn.IsMock { // Require spend bounds for non-mock tokens
			require.NotNil(t, bounds.MaxSpend, "max spend should not be nil")
			require.NotNil(t, bounds.MinSpend, "min spend should not be nil")
		}

		seen[tkn] = true
		golden = append(golden, map[string]any{
			"name":        tkn.Name,
			"symbol":      tkn.Symbol,
			"address":     tkn.Address.Hex(),
			"maxSpend":    tkn.FormatAmt(bounds.MaxSpend),
			"minSpend":    tkn.FormatAmt(bounds.MinSpend),
			"chainId":     tkn.ChainID,
			"coingeckoId": tkn.CoingeckoID,
			"isMock":      tkn.IsMock,
		})
	}

	tutil.RequireGoldenJSON(t, golden)
}

func TestMaxSpendMinThreshold(t *testing.T) {
	t.Parallel()

	for _, token := range tokens.All() {
		if !IsSupportedToken(token) {
			continue
		}

		bounds, ok := tokenSpendBounds[token.Meta][token.ChainClass]
		if !ok {
			continue
		}

		thresh, ok := eoa.GetSolverNetThreshold(eoa.RoleSolver, netconf.Mainnet, token.ChainID, token.Meta)
		if !ok {
			continue
		}

		require.True(t, bi.GTE(thresh.MinBalance(), bounds.MaxSpend), "solver min balance must be greater than max spend: token=%s, min_bal=%s, max_spend=%s", token.Meta, thresh.MinBalance(), bounds.MaxSpend)
	}
}
