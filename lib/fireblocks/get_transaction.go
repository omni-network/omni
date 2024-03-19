package fireblocks

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/omni-network/omni/lib/errors"
)

func (c Client) GetTransactionByID(ctx context.Context, transactionID string) (*TransactionResponse, error) {
	var res TransactionResponse

	endpoint := filepath.Join(transactionEndpoint, transactionID)
	jwtToken, err := c.GenJWTToken(endpoint, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.http.SendRequest(
		ctx,
		endpoint,
		http.MethodGet,
		nil,
		c.getHeaders(jwtToken),
		res,
	)
	if err != nil {
		return nil, err
	}

	res, ok := response.(TransactionResponse)
	if !ok {
		return nil, errors.New("response is not a TransactionResponse")
	}

	return &res, nil
}
