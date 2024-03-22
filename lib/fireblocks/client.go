package fireblocks

import (
	"crypto/rsa"

	"github.com/omni-network/omni/lib/errors"
)

const transactionEndpoint string = "v1/transactions"
const supportedAssetsEndpoint string = "v1/supported_assets"

// Client is a JSON HTTP client for the FireBlocks API.
type Client struct {
	cfg        Config
	apiKey     string
	privateKey *rsa.PrivateKey
	jsonHTTP   jsonHTTP
}

// NewClientWithConfig creates a new FireBlocks client with a custom configuration.
func NewClientWithConfig(apiKey string, privateKey *rsa.PrivateKey, host string, cfg Config) (*Client, error) {
	jsonHTTP := newJSONHTTP(host, apiKey, "")
	client := &Client{
		apiKey:     apiKey,
		privateKey: privateKey,
		jsonHTTP:   jsonHTTP,
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
	return NewClientWithConfig(apiKey, privateKey, host, DefaultConfig())
}

// getAuthHeaders returns the authentication headers for the FireBlocks API.
func (c Client) getAuthHeaders(jwtToken string) map[string]string {
	header := make(map[string]string)
	header["X-API-KEY"] = c.apiKey
	header["Authorization"] = "Bearer " + jwtToken

	return header
}

// check checks if the client is properly configured.
func (c Client) check() error {
	if c.apiKey == "" {
		return errors.New("apiKey is required")
	}
	if c.privateKey == nil {
		return errors.New("privateKey is required")
	}

	return nil
}
