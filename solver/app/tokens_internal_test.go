package app

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/omni-network/omni/e2e/manifests"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/solver/client"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

// TestGenSupportTokensDoc generates the supported tokens docs.
func TestGenSupportTokensDoc(t *testing.T) {
	t.Parallel()
	genSupportedTokens(t, netconf.Omega, "../../../docs/docs/pages/sdk/assets/testnet.mdx")
	genSupportedTokens(t, netconf.Mainnet, "../../../docs/docs/pages/sdk/assets/mainnet.mdx")
}

func genSupportedTokens(t *testing.T, network netconf.ID, fileName string) {
	t.Helper()
	m, err := manifests.Manifest(network)
	require.NoError(t, err)

	metas, err := m.EVMChains()
	require.NoError(t, err)

	for _, chain := range solvernet.Chains(network) {
		meta, ok := evmchain.MetadataByID(chain.ID)
		require.True(t, ok, "chain %d not found", chain.ID)
		metas = append(metas, meta)
	}

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

			lines = append(lines, fmt.Sprintf("| %s | %d | %s | %s |", meta.PrettyName, meta.ChainID, token.Symbol, addr))
		}
	}

	sort.Strings(lines)

	var b bytes.Buffer
	b.WriteString("| Chain | Chain ID | Asset | Contract Address |\n")
	b.WriteString("| ----- | -------- | ----- | ---------------- |\n")

	for _, line := range lines {
		b.WriteString(line + "\n")
	}

	tutil.RequireGoldenBytes(t, b.Bytes(), tutil.WithFilename(fileName))
}

// TestTokens ensures solver token list does not change without explicit golden update.
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
		if tkn.ChainClass == tokens.ClassMainnet { // Require spend for mainnet tokens.
			require.True(t, ok, "missing spend bounds for token: %s", tkn)
			require.NotNil(t, bounds.MaxSpend, "max spend should not be nil")
			require.NotNil(t, bounds.MinSpend, "min spend should not be nil")
		}

		seen[tkn] = true
		golden = append(golden, map[string]any{
			"name":        tkn.Name,
			"symbol":      tkn.Symbol,
			"address":     tkn.UniAddress().String(),
			"maxSpend":    tkn.FormatAmt(bounds.MaxSpend),
			"minSpend":    tkn.FormatAmt(bounds.MinSpend),
			"chainId":     tkn.ChainID,
			"coingeckoId": tkn.CoingeckoID,
			"isMock":      tkn.IsMock,
		})
	}

	tutil.RequireGoldenJSON(t, golden)
}

// 12037465313470890

//go:generate go test . -run=TestTokenResponse -golden

func TestTokenResponse(t *testing.T) {
	t.Parallel()

	backends, mockClients := testBackends(t)
	for _, cl := range mockClients.clients {
		mockAnyBalance(t, cl, bi.Ether(100))
	}

	resp, err := tokensResponse(t.Context(), backends, common.Address{})
	require.NoError(t, err)

	tutil.RequireGoldenJSON(t, resp, tutil.WithFilename("TestTokens/tokens_response.json"))
}

func TestTokensEndpoint(t *testing.T) {
	t.Parallel()

	backends, mockClients := testBackends(t)
	for _, cl := range mockClients.clients {
		mockAnyBalance(t, cl, bi.Ether(100))
	}

	handler := handlerAdapter(newTokensHandler(backends, common.Address{}))

	srv := httptest.NewServer(handler)
	defer srv.Close()

	cl := client.New(srv.URL)
	resp, err := cl.Tokens(t.Context())
	require.NoError(t, err)
	require.NotEmpty(t, resp.Tokens)
}
