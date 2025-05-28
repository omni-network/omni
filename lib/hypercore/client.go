package hypercore

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/omni-network/omni/lib/errors"
)

const (
	MainnetAPI = "https://api.hyperliquid.xyz"
	TestnetAPI = "https://api.hyperliquid-testnet.xyz"
)

type Signer interface {
	Sign(ctx context.Context, digest []byte) ([65]byte, error)
}

// Client defines the interface for the hypercore client.
type Client interface {
	// UseBigBlocks enables the use of big ETH-RPC blocks for account associated with the signer.
	UseBigBlocks(ctx context.Context) error
}

type client struct {
	signer    Signer
	isTestnet bool
}

// NewClient creates a new hypercore mainnet client.
func NewClient(signer Signer) Client {
	return client{
		signer:    signer,
		isTestnet: false,
	}
}

// NewTestnetClient creates a new hypercore testnet client.
func NewTestnetClient(signer Signer) Client {
	return client{
		signer:    signer,
		isTestnet: true,
	}
}

// do signs and sends the action to the hypercore exchange API.
func (c client) do(ctx context.Context, action any) error {
	nonce := timestampMS()

	signature, err := signL1Action(ctx, c.signer, action, emptyAddress, nonce, 0, !c.isTestnet)
	if err != nil {
		return err
	}

	payload := ActionPayload{
		Action:    action,
		Nonce:     nonce,
		Signature: signature,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "marshal payload")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.exchangeURL(), bytes.NewReader(body))
	if err != nil {
		return errors.Wrap(err, "create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read response body")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected error response", "code", resp.StatusCode, "body", string(respBody))
	}

	return nil
}

func (c client) exchangeURL() string {
	return c.baseURL() + "/exchange"
}

func (c client) baseURL() string {
	if c.isTestnet {
		return TestnetAPI
	}

	return MainnetAPI
}
