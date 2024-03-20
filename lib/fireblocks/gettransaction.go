package fireblocks

import (
	"context"
	"net/http"
	"path/filepath"
)

// GetTransactionByID gets a transaction by its ID.
func (c Client) GetTransactionByID(ctx context.Context, transactionID string) (TransactionResponse, error) {
	endpoint := filepath.Join(transactionEndpoint, transactionID)
	jwtToken, err := c.genJWTToken(endpoint, nil)
	if err != nil {
		return TransactionResponse{}, err
	}

	var res TransactionResponse
	err = c.jsonHTTP.Send(
		ctx,
		endpoint,
		http.MethodGet,
		nil,
		c.getAuthHeaders(jwtToken),
		&res,
	)
	if err != nil {
		return TransactionResponse{}, err
	}

	return res, nil
}
