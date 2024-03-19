package fireblocks

import (
	"context"
	"crypto/rsa"
	"os"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/httpclient"

	"github.com/golang-jwt/jwt/v5"
)

const transactionEndpoint string = "/v1/transactions"

type FireBlocks interface {
	NewClientWithConfig(apiKey string, clientSecret string, host string, cfg Config) *Client

	NewDefaultClient(apiKey string, clientSecret string, host string) *Client

	// CreateTransaction creates a new transaction on the FireBlocks API.
	// We use raw signing by default
	CreateTransaction(ctx context.Context, request CreateTransactionRequest) (string, error)

	// GetTransactionById retrieves a transaction by its ID.
	GetTransactionById(ctx context.Context, transactionID string) (*TransactionResponse, error)

	// WaitSigned waits for a transaction to be signed.
	WaitSigned(ctx context.Context, opts TransactionRequestOptions) (*TransactionResponse, error)
}

type Client struct {
	FireBlocks
	cfg            Config
	apiKey         string
	privateKeyPath string
	privateKey     *rsa.PrivateKey
	host           string
	http           httpclient.Client
}

func NewClientWithConfig(apiKey string, privateKeyPath string, host string, cfg Config) (*Client, error) {
	privateKey, err := genPrivateKey(privateKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "genPrivateKey")
	}
	httpClient := httpclient.NewClient(host, apiKey, "")
	client := &Client{
		apiKey:         apiKey,         // pragma: allowlist secret
		privateKeyPath: privateKeyPath, // pragma: allowlist secret
		privateKey:     privateKey,     // pragma: allowlist secret
		host:           host,
		http:           *httpClient,
		cfg:            cfg,
	}

	err = client.check()
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
	header["X-API-KEY"] = c.apiKey
	header["Authorization"] = "Bearer " + jwtToken

	return header
}

func (c Client) check() error {
	if c.host == "" {
		return errors.New("host is required")
	}
	if c.apiKey == "" {
		return errors.New("apiKey is required")
	}
	if c.privateKeyPath == "" {
		return errors.New("private key path required")
	}

	return nil
}

func genPrivateKey(privateKeyPath string) (*rsa.PrivateKey, error) {
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "error reading private key from %s: %w", "private key path", privateKeyPath)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing RSA private key")
	}

	return privateKey, nil
}
