package fireblocks

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/omni-network/omni/lib/errors"
)

type TransactionRequestOptions struct {
	Amount  int64
	Message UnsignedRawMessage
}

func (c Client) CreateTransaction(ctx context.Context, request CreateTransactionRequest) (*TransactionResponse, error) {
	var res TransactionResponse

	req, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	jwtToken, err := c.genJWTToken(transactionEndpoint, req)
	if err != nil {
		return nil, err
	}

	response, err := c.http.SendRequest(
		ctx,
		transactionEndpoint,
		http.MethodPost,
		req,
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
	// TODO: remove fields as needed
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
