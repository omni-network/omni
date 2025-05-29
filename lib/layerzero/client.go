package layerzero

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/omni-network/omni/lib/errors"
)

const (
	MainnetAPI = "https://scan.layerzero-api.com/v1"
	TestnetAPI = "https://scan-testnet.layerzero-api.com/v1"
)

type Client interface {
	GetMessagesByTx(ctx context.Context, txHash string) ([]Message, error)
}

// NewClient returns a LayerZero client for the given host.
func NewClient(host string) Client {
	return client{host}
}

type client struct {
	host string
}

var _ Client = (*client)(nil)

// GetMessagesByTx retrieves lz messages associated with a transaction hash.
func (c client) GetMessagesByTx(ctx context.Context, txHash string) ([]Message, error) {
	var response MessageResponse
	if err := c.do(ctx, path.Join("/messages/tx/", txHash), &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (c client) do(ctx context.Context, path string, res any) error {
	uri, err := c.uri(path)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return errors.Wrap(err, "create request")
	}

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return errors.Wrap(err, "http req")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read response")
	}

	if resp.StatusCode != http.StatusOK {
		var jsonError ErrorResponse
		if err := json.Unmarshal(respBody, &jsonError); err == nil {
			return errors.New("get messages", "code", jsonError.Code, "message", jsonError.Message)
		}

		return errors.New("unexpected status code", "code", resp.StatusCode, "body", string(respBody))
	}

	if err := json.Unmarshal(respBody, res); err != nil {
		return errors.Wrap(err, "decode response")
	}

	return nil
}

func (c client) uri(path string) (string, error) {
	uri, err := url.JoinPath(c.host, path)
	if err != nil {
		return "", errors.Wrap(err, "join path", "base", c.host, "path", path)
	}

	return uri, nil
}
