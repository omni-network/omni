package fireblocks

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

type HTTPClient struct {
	apiKey       string
	clientSecret string
	host         string
	http         http.Client
}

func NewHTTPClient(host string, apiKey string, clientSecret string) *HTTPClient {
	return &HTTPClient{
		host:         host,
		apiKey:       apiKey,       // pragma: allowlist secret
		clientSecret: clientSecret, // pragma: allowlist secret
	}
}

func (c *HTTPClient) Send(ctx context.Context, endpoint string, httpMethod string, data any, headers map[string]string) ([]byte, error) {
	endpoint, err := url.JoinPath(c.host, endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "joining endpoint")
	}

	req, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling JSON")
	}

	request, err := http.NewRequestWithContext(
		ctx,
		httpMethod,
		endpoint,
		io.NopCloser(bytes.NewReader(req)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "new http request")
	}

	request.Header = getHeaders(headers)

	resp, err := c.http.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "http Do")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Error(ctx, "Http: closing body failure", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, "http response code not okay", "status code", resp.StatusCode, "body", string(body))
	}

	return body, nil
}

func getHeaders(m map[string]string) http.Header {
	header := http.Header{}

	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/json")

	for k, v := range m {
		header.Set(k, v)
	}

	return header
}
