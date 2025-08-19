package app

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/unibackend"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

// TestCheckSameChainIntegration tests the /check endpoint for a same-chain order against mainnet RPC endpoints.
func TestCheckSameChainIntegration(t *testing.T) {
	if !*integration {
		t.Skip("Skipping integration test, set -integration to run")
	}

	t.Parallel()

	// Get RPC URLs from environment
	ethRPC := os.Getenv("ETH_RPC")
	baseRPC := os.Getenv("BASE_RPC")
	require.NotEmpty(t, ethRPC, "ETH_RPC environment variable must be set")
	require.NotEmpty(t, baseRPC, "BASE_RPC environment variable must be set")

	// Create clients for Ethereum and Base
	ethClient, err := ethclient.DialContext(t.Context(), "ethereum", ethRPC)
	require.NoError(t, err)
	baseClient, err := ethclient.DialContext(t.Context(), "base", baseRPC)
	require.NoError(t, err)

	// solver private key does not matter
	pk, err := crypto.GenerateKey()
	require.NoError(t, err)

	// Create backends from clients
	clients := map[uint64]ethclient.Client{
		evmchain.IDEthereum: ethClient,
		evmchain.IDBase:     baseClient,
	}
	backends, err := ethbackend.BackendsFromClients(clients, pk)
	require.NoError(t, err)

	outbox := common.HexToAddress("0x084b603269a8fd0d0f7037e591665c025ce3549b")
	solver := common.HexToAddress("0x8cC81c5C09394CEaCa7a53be5f547AE719D75dFC")

	// Create check handler
	handler := newCheckHandler(
		newChecker(unibackend.EVMBackends(backends), func(_ uint64, _ common.Address, _ []byte) bool { return true }, unaryPrice, solver, outbox, nil),
		newTracer(backends, solver, outbox),
	)

	// Create test server
	srv := httptest.NewServer(handlerAdapter(handler))
	defer srv.Close()

	rawReq := `{
		"orderId": "0x0000000000000000000000000000000000000000",
		"sourceChainId": 8453,
		"destChainId": 8453,
		"fillDeadline": 1849042175,
		"calls": [
			{
				"target": "0x62398788692aDed44638F8b9F3eE4B977C78ff46",
				"value": "0x879f3115648c42",
				"data": "0x"
			}
		],
		"deposit": {
			"amount": "0x889f3115648c42",
			"token": "0x0000000000000000000000000000000000000000"
		},
		"expenses": [
			{
				"amount": "0x879f3115648c42",
				"token": "0x0000000000000000000000000000000000000000",
				"spender": "0x0000000000000000000000000000000000000000"
			}
		],
		"debug": true
	}`

	// Make the request
	req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, srv.URL, bytes.NewBufferString(rawReq))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Read and parse response
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var checkResp types.CheckResponse
	err = json.Unmarshal(body, &checkResp)
	require.NoError(t, err)

	// Verify the response
	require.NotNil(t, checkResp)
	require.True(t, checkResp.Accepted, "Check response should be accepted")

	// Pretty print the response
	prettyJSON, err := json.MarshalIndent(checkResp, "", "  ")
	require.NoError(t, err)
	t.Logf("Check response:\n%s", prettyJSON)
}
