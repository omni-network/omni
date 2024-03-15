package http

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

type Client struct {
	apiKey       string
	clientSecret string
	host         string
	http         http.Client
}

func NewClient(apiKey string, clientSecret string, host string) *Client {
	return &Client{
		apiKey:       apiKey,       // pragma: allowlist secret
		clientSecret: clientSecret, // pragma: allowlist secret
		host:         host,
	}
}

func (c *Client) SendRequest(ctx context.Context, req *http.Request, opts JWTOpts) ([]byte, error) {
	authToken, err := genJWTToken(opts)
	if err != nil {
		return nil, errors.Wrap(err, "genJWTToken")
	}
	req.Header = getHeader(authToken, c.apiKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, errors.New("http Do")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Error(ctx, "FireBlocks: closing body failure", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	if resp.StatusCode != http.StatusOK {
		return body, errors.Wrap(err, "http response code: %d", resp.StatusCode)
	}

	return body, nil
}

func (c *Client) CreatePostRequest(ctx context.Context, endpoint string, request []byte) (*http.Request, error) {
	endpoint = fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		endpoint,
		bytes.NewReader(request),
	)
	if err != nil {
		return nil, errors.Wrap(err, "new post request")
	}

	return req, nil
}

func (c *Client) CreateGetRequest(ctx context.Context, endpoint string) (*http.Request, error) {
	endpoint = fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "new get request")
	}

	return req, nil
}

func getHeader(jwtToken string, apiKey string) http.Header {
	header := http.Header{}

	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json")
	header.Add("X-API-Key", apiKey)
	header.Add("Authorization", fmt.Sprintf("Bearer %x", jwtToken))

	return header
}
