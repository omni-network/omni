package fireblocks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/omni-network/omni/lib/errors"
)

func (c *Client) GetTransactionByID(ctx context.Context, transactionID string) (*TransactionResponse, error) {
	var res TransactionResponse
	endpoint := fmt.Sprintf("%s/%s", TransactionEndpoint, transactionID)

	httpReq, err := c.createGetRequest(ctx, TransactionEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "getPostRequest")
	}
	response, err := c.sendRequest(
		ctx,
		httpReq,
		endpoint,
		nil,
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
