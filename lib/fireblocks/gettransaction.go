package fireblocks

import (
	"context"
	"net/http"
	"path/filepath"
)

// GetTransactionByID gets a transaction by its ID.
func (c Client) GetTransactionByID(ctx context.Context, transactionID string) (GetTransactionResponse, error) {
	endpoint := filepath.Join(transactionEndpoint, transactionID)
	jwtToken, err := c.token(endpoint, nil)
	if err != nil {
		return GetTransactionResponse{}, err
	}

	var res GetTransactionResponse
	err = c.jsonHTTP.Send(
		ctx,
		endpoint,
		http.MethodGet,
		nil,
		c.getAuthHeaders(jwtToken),
		&res,
	)
	if err != nil {
		return GetTransactionResponse{}, err
	}

	return res, nil
}
