package app

import (
	"bytes"
	"fmt"
	"sort"
	"testing"

	"github.com/omni-network/omni/e2e/manifests"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

// TestGenSupportTokensDoc generates the supported tokens docs.
func TestGenSupportTokensDoc(t *testing.T) {
	t.Parallel()
	genSupportedTokens(t, netconf.Omega, "Testnet", "../../../docs/docs/pages/sdk/assets/testnet.mdx")
	genSupportedTokens(t, netconf.Mainnet, "Mainnet", "../../../docs/docs/pages/sdk/assets/mainnet.mdx")
}

func genSupportedTokens(t *testing.T, network netconf.ID, networkName string, fileName string) {
	t.Helper()
	m, err := manifests.Manifest(network)
	require.NoError(t, err)

	metas, err := m.EVMChains()
	require.NoError(t, err)

	var lines []string
	for _, meta := range metas {
		if meta.ChainID == network.Static().OmniExecutionChainID {
			continue
		}

		// Add native ETH asset
		for asset := range supportedAssets {
			token, ok := tokens.ByAsset(meta.ChainID, asset)
			if !ok || token.IsMock {
				continue
			}

			addr := token.Address.Hex()
			if token.IsNative() {
				addr = "Native"
			}

			lines = append(lines, fmt.Sprintf("| %s | %s | %d | %s | %s |", networkName, meta.PrettyName, meta.ChainID, token.Symbol, addr))
		}
	}

	sort.Strings(lines)

	var b bytes.Buffer
	b.WriteString("| Network | Chain | Chain ID | Asset | Contract Address |\n")
	b.WriteString("| ------- | ----- | -------- | ----- | ---------------- |\n")

	for _, line := range lines {
		b.WriteString(line + "\n")
	}

	tutil.RequireGoldenBytes(t, b.Bytes(), tutil.WithFilename(fileName))
}

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
