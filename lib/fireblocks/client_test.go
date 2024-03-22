package fireblocks_test

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/fireblocks"
	"github.com/omni-network/omni/lib/k1util"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateTransactionGolden(t *testing.T) {
	ctx := context.Background()
	apiKey := uuid.New().String()
	t.Parallel()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id": "7931eb22-9901-4eae-8816-8c24e3db021a", "status": "SUBMITTED"}`))
	}))
	defer ts.Close()

	test4096Key := parseKey(t, testPrivateKey)
	client, err := fireblocks.NewDefaultClient(apiKey, test4096Key, ts.URL)
	require.NoError(t, err)

	opts := fireblocks.TransactionRequestOptions{
		Message: fireblocks.UnsignedRawMessage{
			Content: "test",
		},
	}
	resp, err := client.CreateTransaction(ctx, opts)
	require.NoError(t, err)
	cmp.Equal(expectSubmittedTransaction(resp.ID), resp)
}

func TestCreateAndWaitGolden(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	id := uuid.New().String()
	ctx := context.Background()
	apiKey := uuid.New().String()
	transactionsURL := "/v1/transactions"
	transactionIDURL := transactionsURL + "/" + id

	mux.HandleFunc(transactionIDURL, func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		response := expectCompletedTransaction(id)
		data, _ := json.Marshal(response)
		_, _ = res.Write(data)
	})

	mux.HandleFunc(transactionsURL, func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		response := expectSubmittedTransaction(id)
		data, _ := json.Marshal(response)
		_, _ = res.Write(data)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	test4096Key := parseKey(t, testPrivateKey)

	// Fast timeouts for testing
	cfg, err := fireblocks.NewConfig(time.Second, time.Millisecond, 10)
	require.NoError(t, err)

	client, err := fireblocks.NewClientWithConfig(apiKey, test4096Key, ts.URL, cfg)
	require.NoError(t, err)
	require.NotNil(t, client)

	resp, err := client.CreateAndWait(ctx, fireblocks.TransactionRequestOptions{
		Message: fireblocks.UnsignedRawMessage{
			Content:        "test",
			DerivationPath: []int{44, 60, 0, 0, 0},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func expectSubmittedTransaction(id string) *fireblocks.GetTransactionResponse {
	return &fireblocks.GetTransactionResponse{
		ID:     id,
		Status: "SUBMITTED",
	}
}

func expectCompletedTransaction(id string) *fireblocks.GetTransactionResponse {
	return &fireblocks.GetTransactionResponse{
		ID:     id,
		Status: "COMPLETED",
	}
}

func parseKey(t *testing.T, s string) *rsa.PrivateKey {
	t.Helper()

	p, _ := pem.Decode([]byte(s))
	k, err := x509.ParsePKCS8PrivateKey(p.Bytes)
	require.NoError(t, err)

	return k.(*rsa.PrivateKey) //nolint:forcetypeassert // parseKey is only used for testing
}

// copied from https://go.dev/src/crypto/rsa/rsa_test.go
const testPrivateKey = `-----BEGIN TESTING KEY-----
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
Axreud3rJ0rekUp6rI1joG717Wls
-----END TESTING KEY-----`

func TestSmoke(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	apiKey := ""
	privateTestKey := ""
	if apiKey == "" || privateTestKey == "" {
		t.Skip("API key and private key are required")
	}
	client, err := fireblocks.NewDefaultClient(apiKey, parseKey(t, privateTestKey), "https://sandbox-api.fireblocks.io")
	input := crypto.Keccak256([]byte("test"))
	require.NoError(t, err)
	resp, err := client.CreateAndWait(ctx, fireblocks.TransactionRequestOptions{
		Message: fireblocks.UnsignedRawMessage{
			Content:        string(input),
			DerivationPath: []int{44, 60, 0, 0},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	msg := resp.SignedMessages[0]
	content := msg.Content

	r := msg.Signature.R
	s := msg.Signature.S
	v := msg.Signature.V

	require.NotNil(t, r)
	require.NotNil(t, s)
	require.NotNil(t, v)

	require.NotEmpty(t, content)

	var addr common.Address
	publicKey := msg.PublicKey
	addr.SetBytes([]byte(publicKey))

	byteSig := []byte("0x" + msg.Signature.FullSig)
	require.NotNil(t, byteSig)

	_, err = k1util.Verify(addr, [32]byte(input), [65]byte(byteSig))
	require.NoError(t, err)
	// require.Truef(t, ok, "signature verification failed")
}
