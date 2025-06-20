package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/uni"
	"github.com/omni-network/omni/scripts/trade/users"
	"github.com/omni-network/omni/solver/types"

	"github.com/google/uuid"
)

var _ users.Service = Client{}

type Client struct {
	baseURL string
}

func New(baseURL string) Client {
	return Client{
		baseURL: baseURL,
	}
}

func (c Client) Create(ctx context.Context, req users.RequestCreate) (users.User, error) {
	var u users.User
	if err := c.do(ctx, "/api/v1/users/create", req, &u); err != nil {
		return users.User{}, err
	}

	return u, nil
}

func (c Client) GetByID(ctx context.Context, id uuid.UUID) (users.User, error) {
	req := struct {
		ID uuid.UUID `json:"id"`
	}{
		ID: id,
	}

	var u users.User
	if err := c.do(ctx, "/api/v1/users/get_by_id", req, &u); err != nil {
		return users.User{}, err
	}

	return u, nil
}

func (c Client) GetByPrivyID(ctx context.Context, privyID string) (users.User, error) {
	req := struct {
		PrivyID string `json:"privy_id"`
	}{
		PrivyID: privyID,
	}

	var u users.User
	if err := c.do(ctx, "/api/v1/users/get_by_privy_id", req, &u); err != nil {
		return users.User{}, err
	}

	return u, nil
}

func (c Client) GetByAddress(ctx context.Context, address uni.Address) (users.User, error) {
	req := struct {
		Address string `json:"address"`
	}{
		Address: address.String(),
	}

	var u users.User
	if err := c.do(ctx, "/api/v1/users/get_by_address", req, &u); err != nil {
		return users.User{}, err
	}

	return u, nil
}

func (c Client) ListAll(ctx context.Context) ([]users.User, error) {
	var us []users.User
	if err := c.do(ctx, "/api/v1/users/list_all", nil, &us); err != nil {
		return nil, err
	}

	return us, nil
}

func (c Client) do(ctx context.Context, endpoint string, req any, res any) error {
	var body []byte
	if req != nil {
		var err error
		body, err = json.Marshal(req)
		if err != nil {
			return errors.Wrap(err, "marshal request")
		}
	}

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
	uri, err := url.JoinPath(c.baseURL, path)
	if err != nil {
		return "", errors.Wrap(err, "join path", "base", c.baseURL, "path", path)
	}

	return uri, nil
}
