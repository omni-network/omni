// Package cctp provides functionality for working with the Circle Cross-Chain Transfer Protocol (CCTP).
package cctp

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	MainnetAPI = "https://iris-api.circle.com"
	TestnetAPI = "https://iris-api-sandbox.circle.com"
)

// Client is the interface for CCTP clients.
type Client interface {
	GetAttestation(ctx context.Context, messageHash common.Hash) ([]byte, AttestationStatus, error)
}

// NewClient returns a CCTP client for the given host.
func NewClient(host string) Client {
	return client{host}
}

type client struct {
	host string
}

var _ Client = (*client)(nil)

// GetAttestation retrieves the attestation (and status) for a given message hash.
func (c client) GetAttestation(ctx context.Context, messageHash common.Hash) ([]byte, AttestationStatus, error) {
	var res attestationResponse

	if err := c.do(ctx, "/v1/attestations/"+messageHash.Hex(), &res); err != nil {
		return nil, "", err
	}

	status := AttestationStatus(res.Status)
	if err := status.Validate(); err != nil {
		return nil, "", err
	}

	if status == AttestationStatusPendingConfirmations {
		return nil, status, nil
	}

	attestation, err := hexutil.Decode(res.Attestation)
	if err != nil {
		return nil, "", errors.Wrap(err, "decode attestation hex")
	}

	return attestation, status, nil
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

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("attestation not found")
	}

	if resp.StatusCode != http.StatusOK {
		var jsonError errorResponse
		if err := json.Unmarshal(respBody, &jsonError); err == nil {
			return errors.New("get attestation", "error", jsonError.Error)
		}

		return errors.New("unexpected status code", "code", resp.StatusCode)
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

type attestationResponse struct {
	Attestation string `json:"attestation"`
	Status      string `json:"status"`
}

type errorResponse struct {
	Error string `json:"error"`
}
