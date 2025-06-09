package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/solver/types"
)

const (
	endpointCheck = "/api/v1/check"
	endpointQuote = "/api/v1/quote"
	endpointPrice = "/api/v1/price"
	endpointRelay = "/api/v1/relay"
)

// New creates a new solver Client.
func New(host string, opts ...func(*Client)) Client {
	cl := Client{
		host:         host,
		debugReqBody: func([]byte) {},
		debugResBody: func([]byte) {},
	}

	for _, opt := range opts {
		opt(&cl)
	}

	return cl
}

func WithDebugBodies(debugReqBody, debugResBody func([]byte)) func(*Client) {
	return func(c *Client) {
		c.debugReqBody = debugReqBody
		c.debugResBody = debugResBody
	}
}

type Client struct {
	host         string
	debugReqBody func([]byte)
	debugResBody func([]byte)
}

// Check calls the check API endpoint returning whether the order will be rejected or not.
func (c Client) Check(ctx context.Context, req types.CheckRequest) (types.CheckResponse, error) {
	var res types.CheckResponse

	if err := c.do(ctx, endpointCheck, req, &res); err != nil {
		return types.CheckResponse{}, err
	}

	return res, nil
}

// Quote calls the quote API endpoint returning the required deposit or expense amounts.
func (c Client) Quote(ctx context.Context, req types.QuoteRequest) (types.QuoteResponse, error) {
	var res types.QuoteResponse

	if err := c.do(ctx, endpointQuote, req, &res); err != nil {
		return types.QuoteResponse{}, err
	}

	return res, nil
}

func (c Client) Price(ctx context.Context, req types.PriceRequest) (types.Price, error) {
	var res types.Price

	if err := c.do(ctx, endpointPrice, req, &res); err != nil {
		return types.Price{}, err
	}

	return res, nil
}

// Relay calls the relay API endpoint to submit a gasless order on behalf of a user.
func (c Client) Relay(ctx context.Context, req types.RelayRequest) (types.RelayResponse, error) {
	var res types.RelayResponse

	if err := c.do(ctx, endpointRelay, req, &res); err != nil {
		return types.RelayResponse{}, err
	}

	return res, nil
}

func (c Client) do(ctx context.Context, endpoint string, req any, res any) error {
	body, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "marshal request")
	}

	c.debugReqBody(body)

	uri, err := c.uri(endpoint)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewReader(body))
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

	c.debugResBody(respBody)

	if resp.StatusCode != http.StatusOK {
		var jsonError types.JSONErrorResponse
		if err := json.Unmarshal(respBody, &jsonError); err == nil {
			return errors.New(jsonError.Error.Message, "status", resp.StatusCode)
		}

		return errors.New("non-json-error response", "status", resp.StatusCode)
	}

	if err := json.Unmarshal(respBody, res); err != nil {
		return errors.Wrap(err, "decode response")
	}

	return nil
}

func (c Client) uri(path string) (string, error) {
	uri, err := url.JoinPath(c.host, path)
	if err != nil {
		return "", errors.Wrap(err, "join path", "base", c.host, "path", path)
	}

	return uri, nil
}
