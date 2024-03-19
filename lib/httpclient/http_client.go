package httpclient

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

type Client struct {
	apiKey       string
	clientSecret string
	host         string
	http         http.Client
}

func NewClient(host string, apiKey string, clientSecret string) *Client {
	return &Client{
		host:         host,
		apiKey:       apiKey,       // pragma: allowlist secret
		clientSecret: clientSecret, // pragma: allowlist secret
	}
}

func (c *Client) SendRequest(ctx context.Context, endpoint string, httpMethod string, bodyJSON any, headers map[string]string, response any) (any, error) {
	endpoint = fmt.Sprintf("%s/%s", c.host, endpoint)
	bodyBytes, err := json.Marshal(bodyJSON)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling JSON")
	}

	req, err := http.NewRequestWithContext(
		ctx,
		httpMethod,
		endpoint,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "new http request")
	}

	if len(bodyBytes) > 0 {
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	req.Header = getHeaders(headers)

	resp, err := c.http.Do(req)
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
		return nil, errors.Wrap(err, "read response body", "body")
	}

	bodyString := string(body)
	log.Info(ctx, "Http: response body", "body", bodyString)
	if resp.StatusCode != http.StatusOK {
		return body, errors.Wrap(err, "http response code not okay", "status code", resp.StatusCode, "body", bodyString)
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal response JSON")
	}

	return response, nil
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
