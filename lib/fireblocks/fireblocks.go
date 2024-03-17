package fireblocks

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/fireblocks/http"
)

const TransactionEndpoint string = "v1/transactions"

type FireBlocks interface {
	// CreateTransaction creates a new transaction on the FireBlocks API.
	// We use raw signing by default
	CreateTransaction(ctx context.Context, request CreateTransactionRequest) (string, error)

	// GetTransactionById retrieves a transaction by its ID.
	GetTransactionById(ctx context.Context, transactionID string) (*CreateTransactionRequest, error)

	// WaitSigned waits for a transaction to be signed.
	WaitSigned(ctx context.Context, opts TransactionRequestOptions, jwtOpts http.JWTOpts) (*TransactionResponse, error)
}

type Client struct {
	FireBlocks
	cfg          Config
	apiKey       string
	clientSecret string
	host         string
	http         *http.Client
}

func NewClient(apiKey string, clientSecret string, host string) *Client {
	httpClient := http.NewClient(apiKey, clientSecret, host)
	cfg := Config{
		NetworkTimeout: time.Duration(300) * time.Second,
		QueryInterval:  time.Duration(2) * time.Second,
		LogFreqFactor:  2,
	}

	return &Client{
		apiKey:       apiKey,       // pragma: allowlist secret
		clientSecret: clientSecret, // pragma: allowlist secret
		host:         host,
		http:         httpClient,
		cfg:          cfg,
	}
}
