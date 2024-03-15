package fireblocks

import (
	"context"

	"github.com/omni-network/omni/lib/fireblocks/http"
)

const TransactionEndpoint string = "v1/transactions"

type FireBlocks interface {
	// CreateTransaction creates a new transaction on the FireBlocks API.
	// We use raw signing by default
	CreateTransaction(ctx context.Context, request CreateTransactionRequest) (string, error)

	// GetTransactionById retrieves a transaction by its ID.
	GetTransactionById(ctx context.Context, transactionID string) (*CreateTransactionRequest, error)
}

type Client struct {
	FireBlocks
	apiKey       string
	clientSecret string
	host         string
	http         *http.Client
}

func NewClient(apiKey string, clientSecret string, host string) *Client {
	httpClient := http.NewClient(apiKey, clientSecret, host)
	return &Client{
		apiKey:       apiKey,       // pragma: allowlist secret
		clientSecret: clientSecret, // pragma: allowlist secret
		host:         host,
		http:         httpClient,
	}
}
