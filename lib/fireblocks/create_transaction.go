package fireblocks

import (
	"context"
	"net/http"

	"github.com/omni-network/omni/lib/errors"
)

type TransactionRequestOptions struct {
	Message UnsignedRawMessage
}

func (c Client) CreateTransaction(ctx context.Context, request CreateTransactionRequest) (*TransactionResponse, error) {
	var res TransactionResponse

	jwtToken, err := c.GenJWTToken(transactionEndpoint, request)
	if err != nil {
		return nil, err
	}

	response, err := c.http.SendRequest(
		ctx,
		transactionEndpoint,
		http.MethodPost,
		request,
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

func NewTransactionRequest(opt TransactionRequestOptions) CreateTransactionRequest {
	req := CreateTransactionRequest{
		Operation: "RAW",
		Note:      "testing transaction",
		Source: Source{
			Type: "VAULT_ACCOUNT",
		},
		Destination:   Destination{},
		CustomerRefID: "",
		ExtraParameters: RawMessageData{
			Messages: []UnsignedRawMessage{opt.Message},
		},
	}

	return req
}
