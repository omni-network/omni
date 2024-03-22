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
func (c Client) CreateTransaction(ctx context.Context, opt TransactionRequestOptions) (CreateTransactionResponse, error) {
	request := newTransactionRequest(opt)
	jwtToken, err := c.token(transactionEndpoint, request)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	var res CreateTransactionResponse
	err = c.jsonHTTP.Send(
		ctx,
		transactionEndpoint,
		http.MethodPost,
		request,
		c.getAuthHeaders(jwtToken),
		&res,
	)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	return res, nil
}

// newTransactionRequest creates a new transaction request.
func newTransactionRequest(opt TransactionRequestOptions) createTransactionRequest {
	contentHash := sha256.Sum256([]byte(opt.Message.Content))

	return createTransactionRequest{
		Operation: "RAW",
		Source: source{
			Type: "VAULT_ACCOUNT",
			ID:   "0",
		},
		AssetID: "ETH_TEST3",
		ExtraParameters: &extraParameters{
			RawMessageData: rawMessageData{
				// Algorithm: "MPC_EDDSA_ED25519",
				Messages: []UnsignedRawMessage{{
					Content: hex.EncodeToString(contentHash[:]),
					// DerivationPath: opt.Message.DerivationPath,
				}},
			},
		},
	}
}
