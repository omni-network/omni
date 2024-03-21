package fireblocks

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

// TransactionRequestOptions are the options for creating a new transaction.
type TransactionRequestOptions struct {
	Message UnsignedRawMessage
}

// CreateTransaction creates a new transaction on the FireBlocks API.
func (c Client) CreateTransaction(ctx context.Context, opt TransactionRequestOptions) (*TransactionResponse, error) {
	request := newTransactionRequest(opt)
	jwtToken, err := c.genJWTToken(transactionEndpoint, request)
	if err != nil {
		return nil, err
	}

	var res TransactionResponse
	err = c.jsonHTTP.Send(
		ctx,
		transactionEndpoint,
		http.MethodPost,
		request,
		c.getAuthHeaders(jwtToken),
		&res,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// newTransactionRequest creates a new transaction request.
func newTransactionRequest(opt TransactionRequestOptions) createTransactionRequest {
	contentHash := sha256.Sum256([]byte(opt.Message.Content))

	return createTransactionRequest{
		Operation: "RAW",
		Note:      "testing transaction",
		AssetID:   "ETH",
		Source: source{
			Type: "VAULT_ACCOUNT",
			ID:   "0",
		},
		Destination: &destination{
			Type: "VAULT_ACCOUNT",
		},
		CustomerRefID: "",
		ExtraParameters: &extraParameters{
			RawMessageData: rawMessageData{
				Messages: []UnsignedRawMessage{{
					Content: hex.EncodeToString(contentHash[:]),
				}},
			},
		},
	}
}
