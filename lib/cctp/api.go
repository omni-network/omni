package cctp

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	MainnetAPI = "https://iris-api.circle.com"
	TestnetAPI = "https://iris-api-sandbox.circle.com"
)

// NewAPIClient creates a new CCTP API client.
func NewAPIClient(host string) APIClient {
	return APIClient{host}
}

// APIClient is a client for interacting with the CCTP API.
type APIClient struct {
	host string
}

// AttestationStatus is the status of an attestation, 'complete' or 'pending_confirmations'.
type AttestationStatus string

const (
	AttestationStatusComplete             AttestationStatus = "complete"
	AttestationStatusPendingConfirmations AttestationStatus = "pending_confirmations"
)

// Validate checks if the status is a known status.
func (s AttestationStatus) Validate() error {
	switch s {
	case AttestationStatusComplete, AttestationStatusPendingConfirmations:
		return nil
	default:
		return errors.New("unexpected attestation status", "status", s)
	}
}

// GetAttestation retrieves the attestation (and status) for a given message hash.
func (c APIClient) GetAttestation(ctx context.Context, messageHash string) ([]byte, AttestationStatus, error) {
	var res attestationResponse

	if err := c.do(ctx, "/v1/attestations/"+messageHash, &res); err != nil {
		return nil, "", err
	}

	status := AttestationStatus(res.Data.Status)
	if err := status.Validate(); err != nil {
		return nil, "", err
	}

	signature, err := hexutil.Decode(res.Data.Attestation)
	if err != nil {
		return nil, "", errors.Wrap(err, "decode attestation hex")
	}

	return signature, status, nil
}

func (c APIClient) do(ctx context.Context, path string, res any) error {
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
			return errors.New("get attestation", "error", jsonError.Data.Error)
		}

		return errors.New("unexpected status code", "code", resp.StatusCode)
	}

	if err := json.Unmarshal(respBody, res); err != nil {
		return errors.Wrap(err, "decode response")
	}

	return nil
}

func (c APIClient) uri(path string) (string, error) {
	uri, err := url.JoinPath(c.host, path)
	if err != nil {
		return "", errors.Wrap(err, "join path", "base", c.host, "path", path)
	}

	return uri, nil
}

type attestationResponse struct {
	Data struct {
		Attestation string `json:"attestation"`
		Status      string `json:"status"`
	} `json:"data"`
}

type errorResponse struct {
	Data struct {
		Error string `json:"error"`
	} `json:"data"`
}
