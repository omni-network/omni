package fireblocks

import (
	"crypto/rsa"

	"github.com/omni-network/omni/lib/errors"
)

const transactionEndpoint string = "v1/transactions"

type Client struct {
	cfg        Config
	apiKey     string
	privateKey *rsa.PrivateKey
	host       string
	apiRequest HTTPClient
}

// NewClientWithConfig creates a new FireBlocks client with a custom configuration.
func NewClientWithConfig(apiKey string, privateKey *rsa.PrivateKey, host string, cfg Config) (*Client, error) {
	apiRequestClient := NewHTTPClient(host, apiKey, "")
	client := &Client{
		apiKey:     apiKey,     // pragma: allowlist secret
		privateKey: privateKey, // pragma: allowlist secret
		host:       host,
		apiRequest: *apiRequestClient,
		cfg:        cfg,
	}

	err := client.check()
	if err != nil {
		return nil, errors.Wrap(err, "client check")
	}

	return client, nil
}

// NewDefaultClient creates a new FireBlocks client with default configuration.
func NewDefaultClient(apiKey string, privateKey *rsa.PrivateKey, host string) (*Client, error) {
	client, err := NewClientWithConfig(apiKey, privateKey, host, DefaultConfig())
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

	return nil
}
