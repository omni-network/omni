package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/fireblocks/http"
)

func (c *Client) GetTransactionByID(ctx context.Context, transactionID string, opts http.JWTOpts) (*TransactionResponse, error) {
	var res TransactionResponse

	endpoint := fmt.Sprintf("%s/%s", TransactionEndpoint, transactionID)
	opts.URI = endpoint

	httpReq, err := c.http.CreateGetRequest(ctx, TransactionEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "getPostRequest")
	}

	response, err := c.http.SendRequest(
		ctx,
		httpReq,
		opts,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal response JSON")
	}

	return &res, nil
}
