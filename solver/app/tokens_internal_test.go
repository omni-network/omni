package app

import (
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/manifests"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/evmchain"
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

		bounds, ok := GetSpendBounds(tkn)
		if !tkn.IsMock { // Require spend bounds for non-mock tokens
			require.True(t, ok, "missing spend bounds for token: %s", tkn)
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

func TestTokenResponse(t *testing.T) {
	t.Parallel()

	mainnet, err := manifests.Mainnet()
	require.NoError(t, err)

	chains := []uint64{
		evmchain.IDOmniMainnet,
	}
	for _, name := range mainnet.PublicChains {
		chain, ok := evmchain.MetadataByName(name)
		require.True(t, ok, "chain %s not found", name)
		chains = append(chains, chain.ChainID)
	}

	resp, err := tokensResponse(chains)
	require.NoError(t, err)

	tutil.RequireGoldenJSON(t, resp, tutil.WithFilename("TestTokens/tokens_response.json"))
}

func TestMaxSpendMinThreshold(t *testing.T) {
	t.Parallel()

	for _, token := range tokens.All() {
		if !IsSupportedToken(token) {
			continue
		}

		bounds, ok := tokenSpendBounds[token.Asset][token.ChainClass]
		if !ok {
			continue
		}

		thresh, ok := eoa.GetSolverNetThreshold(eoa.RoleSolver, netconf.Mainnet, token.ChainID, token.Asset)
		if !ok {
			continue
		}

		require.True(t, bi.GTE(thresh.MinBalance(), bounds.MaxSpend), "solver min balance must be greater than max spend: token=%s, min_bal=%s, max_spend=%s", token.Asset, thresh.MinBalance(), bounds.MaxSpend)
	}
}
