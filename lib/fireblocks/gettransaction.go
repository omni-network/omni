package fireblocks

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/omni-network/omni/lib/errors"
)

// getTransactionByID gets a transaction by its ID.
func (c Client) getTransactionByID(ctx context.Context, transactionID string) (transaction, error) {
	endpoint := filepath.Join(endpointTransactions, transactionID)
	headers, err := c.authHeaders(endpoint, nil)
	if err != nil {
		return transaction{}, err
	}

	var res transaction
	var errRes errorResponse
	ok, err := c.jsonHTTP.Send(
		ctx,
		endpoint,
		http.MethodGet,
		nil,
		headers,
		&res,
		&errRes,
	)
	if err != nil {
		return transaction{}, err
	} else if !ok {
		return transaction{}, errors.New("failed to get transaction", "resp_msg", errRes.Message, "resp_code", errRes.Code)
	}

	return res, nil
}
