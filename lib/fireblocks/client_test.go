package fireblocks_test

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/fireblocks"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	_ "embed"
)

//go:embed testdata/test_private_key.pem
var testPrivateKey []byte

func TestSignOK(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	apiKey := uuid.New().String()
	txID := uuid.New().String()

	// Create a private key and sign an expected message
	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr := crypto.PubkeyToAddress(privKey.PublicKey)
	digest := crypto.Keccak256([]byte("test"))
	expectSig, err := crypto.Sign(digest, privKey)
	require.NoError(t, err)

	// Start a test http server that serves the expected responses
	var count int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		if count <= 2 {
			// Just return txID and "submitted" on first two attempts
			bz, _ := json.Marshal(struct {
				ID     string `json:"id"`
				Status string `json:"status"`
			}{
				ID:     txID,
				Status: "SUBMITTED",
			})
			_, _ = w.Write(bz)

			return
		}

		// Then return the signed transaction
		bz, _ := json.Marshal(fireblocks.TransactionResponseForT(t, txID, [65]byte(expectSig), &privKey.PublicKey))
		_, _ = w.Write(bz)
	}))
	defer ts.Close()

	client, err := fireblocks.New(netconf.Simnet, apiKey, parseKey(t, testPrivateKey),
		fireblocks.WithHost(ts.URL),                    // Use the test server for all requests.
		fireblocks.WithQueryInterval(time.Millisecond), // Fast timeout and interval for testing
		fireblocks.WithLogFreqFactor(1),
		fireblocks.WithTestAccount(addr, 0),
	)
	require.NoError(t, err)

	actualSig, err := client.Sign(ctx, [32]byte(digest), addr)
	require.NoError(t, err)

	require.Equal(t, [65]byte(expectSig), actualSig)
}

func parseKey(t *testing.T, data []byte) *rsa.PrivateKey {
	t.Helper()

	p, _ := pem.Decode(data)
	k, err := x509.ParsePKCS8PrivateKey(p.Bytes)
	require.NoError(t, err)

	return k.(*rsa.PrivateKey)
}

// Populate this or run TestSmoke via terminal
// func init() {
//	os.Setenv("FIREBLOCKS_APIKEY", "")
//	os.Setenv("FIREBLOCKS_KEY_PATH", "")
//}

func TestSmoke(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	apiKey, ok := os.LookupEnv("FIREBLOCKS_APIKEY")
	if !ok {
		t.Skip("FIREBLOCKS_APIKEY not set")
	}
	privKeyFile, ok := os.LookupEnv("FIREBLOCKS_KEY_PATH")
	if !ok {
		t.Skip("FIREBLOCKS_KEY_PATH not set")
	}
	privKey, err := os.ReadFile(privKeyFile)
	require.NoError(t, err)

	client, err := fireblocks.New(netconf.Staging, apiKey, parseKey(t, privKey))
	require.NoError(t, err)

	addr := common.BytesToAddress(hexutil.MustDecode("0x7a6cF389082dc698285474976d7C75CAdE08ab7e"))

	t.Run("assets", func(t *testing.T) {
		t.Parallel()

		assets, err := client.GetSupportedAssets(ctx)
		require.NoError(t, err)

		for i, asset := range assets {
			if !strings.Contains(asset.ID, "ETH_TEST") {
				continue
			}
			t.Logf("asset %d: %#v", i, asset)
		}
	})

	t.Run("accounts", func(t *testing.T) {
		t.Parallel()

		accounts, err := client.Accounts(ctx)
		tutil.RequireNoError(t, err)

		t.Logf("accounts: %#v", accounts)
	})

	t.Run("pubkey", func(t *testing.T) {
		t.Parallel()

		pubkey, err := client.GetPublicKey(ctx, 0)
		tutil.RequireNoError(t, err)

		t.Logf("address: %#v", crypto.PubkeyToAddress(*pubkey))
	})

	t.Run("sign", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
		defer cancel()

		_, err = client.Sign(ctx, tutil.RandomHash(), addr)
		tutil.RequireNoError(t, err)
	})
}
