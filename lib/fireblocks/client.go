package fireblocks

import (
	"context"
	"crypto/rsa"
	"os"

	"github.com/omni-network/omni/lib/apirequestor"
	"github.com/omni-network/omni/lib/errors"

	"github.com/golang-jwt/jwt/v5"
)

const transactionEndpoint string = "v1/transactions"

type FireBlocks interface {
	// CreateTransaction creates a new transaction on the FireBlocks API.
	// We use raw signing by default
	CreateTransaction(ctx context.Context, opt TransactionRequestOptions) (string, error)

	// GetTransactionById retrieves a transaction by its ID.
	GetTransactionById(ctx context.Context, transactionID string) (*TransactionResponse, error)

	// CreateAndWait waits for a transaction to be signed.
	CreateAndWait(ctx context.Context, opts TransactionRequestOptions) (*TransactionResponse, error)
}

type Client struct {
	cfg            Config
	apiKey         string
	privateKeyPath string
	privateKey     *rsa.PrivateKey
	host           string
	apiRequest     apirequestor.Client
}

// NewClientWithConfig creates a new FireBlocks client with a custom configuration.
func NewClientWithConfig(apiKey string, privateKeyPath string, host string, cfg Config) (*Client, error) {
	privateKey, err := genPrivateKey(privateKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "genPrivateKey")
	}
	apiRequestClient := apirequestor.NewClient(host, apiKey, "")
	client := &Client{
		apiKey:         apiKey,         // pragma: allowlist secret
		privateKeyPath: privateKeyPath, // pragma: allowlist secret
		privateKey:     privateKey,     // pragma: allowlist secret
		host:           host,
		apiRequest:     *apiRequestClient,
		cfg:            cfg,
	}

	err = client.check()
	if err != nil {
		return nil, errors.Wrap(err, "client check")
	}

	return client, nil
}

// NewDefaultClient creates a new FireBlocks client with default configuration.
func NewDefaultClient(apiKey string, clientSecret string, host string) (*Client, error) {
	client, err := NewClientWithConfig(apiKey, clientSecret, host, DefaultConfig())
	return client, err
}

// getHeaders returns the headers for the FireBlocks API.
func (c Client) getHeaders(jwtToken string) map[string]string {
	header := make(map[string]string)
	header["X-API-KEY"] = c.apiKey
	header["Authorization"] = "Bearer " + jwtToken

	return header
}

// check checks if the client is properly configured.
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

// genPrivateKey generates a private key from a file.
func genPrivateKey(privateKeyPath string) (*rsa.PrivateKey, error) {
	// TODO: use github secrets instead of reading from file
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "reading private key", "private key path", privateKeyPath)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing RSA private key")
	}

	return privateKey, nil
}
