package fireblocks

import (
	"context"
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/omni-network/omni/lib/errors"
)

// GetTransactionByID gets a transaction by its ID.
func (c Client) GetTransactionByID(ctx context.Context, transactionID string) (*TransactionResponse, error) {
	endpoint := filepath.Join(transactionEndpoint, transactionID)
	jwtToken, err := c.genJWTToken(endpoint, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.apiRequest.Send(
		ctx,
		endpoint,
		http.MethodGet,
		nil,
		c.getHeaders(jwtToken),
	)
	if err != nil {
		return nil, err
	}

	var res TransactionResponse
	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshaling response")
	}

	return &res, nil
}
