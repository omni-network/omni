package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/solver/types"
)

const (
	endpointCheck = "/api/v1/check"
)

type Client struct {
	host string
}

// New creates a new solver Client.
func New(host string) Client {
	return Client{
		host: host,
	}
}

// Check runs solver check on the provided request and returns solver's response.
func (c Client) Check(ctx context.Context, req types.CheckRequest) (types.CheckResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return types.CheckResponse{}, errors.Wrap(err, "marshal request")
	}

	uri, err := c.uri(endpointCheck)
	if err != nil {
		return types.CheckResponse{}, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewReader(body))
	if err != nil {
		return types.CheckResponse{}, errors.Wrap(err, "create request")
	}

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return types.CheckResponse{}, errors.Wrap(err, "http req")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var jsonError types.JSONErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&jsonError); err == nil {
			return types.CheckResponse{}, errors.New(jsonError.Error.Message, "status", resp.StatusCode)
		}

		return types.CheckResponse{}, errors.New("non-json-error response", "status", resp.StatusCode)
	}

	var response types.CheckResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return types.CheckResponse{}, errors.Wrap(err, "decode response")
	}

	return response, nil
}

func (c Client) uri(path string) (string, error) {
	uri, err := url.JoinPath(c.host, path)
	if err != nil {
		return "", errors.Wrap(err, "join path", "base", c.host, "path", path)
	}

	return uri, nil
}
