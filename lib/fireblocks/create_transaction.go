package fireblocks

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/omni-network/omni/lib/errors"
)

type TransactionRequestOptions struct {
	Message UnsignedRawMessage
}

func (c Client) CreateTransaction(ctx context.Context, request CreateTransactionRequest) (*TransactionResponse, error) {
	res := TransactionResponse{}

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
	)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshaling response")
	}

	return &res, nil
}

func NewTransactionRequest(opt TransactionRequestOptions) (CreateTransactionRequest, error) {
	var rawMessageData RawMessageData
	h := sha256.New()
	wrappedMessage := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(opt.Message.Content)) + opt.Message.Content
	_, err := h.Write([]byte(wrappedMessage))
	if err != nil {
		return CreateTransactionRequest{}, errors.Wrap(err, "sha256")
	}
	hash := h.Sum(nil)
	str := hex.EncodeToString(hash)

	rawSigningMessage := UnsignedRawMessage{
		Content: str,
	}
	rawMessageData.Messages = append(rawMessageData.Messages, rawSigningMessage)

	req := CreateTransactionRequest{
		Operation: "RAW",
		Note:      "testing transaction",
		AssetID:   "ETH",
		Source: Source{
			Type: "VAULT_ACCOUNT",
			ID:   "0",
		},
		Destination: &Destination{
			Type: "VAULT_ACCOUNT",
		},
		CustomerRefID: "",
		ExtraParameters: &ExtraParameters{
			RawMessageData: rawMessageData,
		},
	}

	return req, nil
}
