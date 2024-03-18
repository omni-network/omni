package fireblocks

import (
	"context"
	"fmt"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/httpclient"
)

const transactionEndpoint string = "v1/transactions"

type FireBlocks interface {
	NewClientWithConfig(apiKey string, clientSecret string, host string, cfg Config) *Client

	NewDefaultClient(apiKey string, clientSecret string, host string) *Client

	// CreateTransaction creates a new transaction on the FireBlocks API.
	// We use raw signing by default
	CreateTransaction(ctx context.Context, request CreateTransactionRequest, jwtOpts JWTOpts) (string, error)

	// GetTransactionById retrieves a transaction by its ID.
	GetTransactionById(ctx context.Context, transactionID string, jwtOpts JWTOpts) (*TransactionResponse, error)

	// WaitSigned waits for a transaction to be signed.
	WaitSigned(ctx context.Context, opts TransactionRequestOptions, jwtOpts JWTOpts) (*TransactionResponse, error)
}

type Client struct {
	FireBlocks
	cfg          Config
	apiKey       string
	clientSecret string
	host         string
	http         httpclient.Client
}

func NewClientWithConfig(apiKey string, clientSecret string, host string, cfg Config) (*Client, error) {
	httpClient := httpclient.NewClient(apiKey, clientSecret, host)
	client := &Client{
		apiKey:       apiKey,       // pragma: allowlist secret
		clientSecret: clientSecret, // pragma: allowlist secret
		host:         host,
		http:         *httpClient,
		cfg:          cfg,
	}
	err := client.check()
	if err != nil {
		return nil, errors.Wrap(err, "client check")
	}

	return client, nil
}

func NewDefaultClient(apiKey string, clientSecret string, host string) (*Client, error) {
	client, err := NewClientWithConfig(apiKey, clientSecret, host, DefaultConfig())
	return client, err
}

func (c Client) getHeaders(jwtToken string) map[string]string {
	header := make(map[string]string)
	header["X-API-Key"] = c.apiKey
	header["Authorization"] = fmt.Sprintf("Bearer %x", jwtToken)

	return header
}

func (c Client) check() error {
	if c.host == "" {
		return errors.New("host is required")
	}
	if c.apiKey == "" {
		return errors.New("apiKey is required")
	}
	if c.clientSecret == "" {
		return errors.New("clientSecret is required")
	}

	return nil
}
