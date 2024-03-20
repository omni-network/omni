package fireblocks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

func (c *HTTPClient) Send(ctx context.Context, endpoint string, httpMethod string, bodyJSON any, headers map[string]string) (string, error) {
	endpoint = fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequestWithContext(
		ctx,
		httpMethod,
		endpoint,
		nil,
	)
	if err != nil {
		return "", errors.Wrap(err, "new http request")
	}

	bodyBytes, err := json.Marshal(bodyJSON)
	if err != nil {
		return "", errors.Wrap(err, "marshaling JSON")
	}

	if bodyJSON != nil {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	req.Header = getHeaders(headers)

	resp, err := c.http.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "http Do")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Error(ctx, "Http: closing body failure", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "read response body", "body")
	}

	bodyString := string(body)
	if resp.StatusCode != http.StatusOK {
		return bodyString, errors.Wrap(err, "http response code not okay", "status code", resp.StatusCode, "body", bodyString)
	}

	return bodyString, nil
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
