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

// TransactionRequestOptions are the options for creating a new transaction.
type TransactionRequestOptions struct {
	Message UnsignedRawMessage
}

// CreateTransaction creates a new transaction on the FireBlocks API.
func (c Client) CreateTransaction(ctx context.Context, opt TransactionRequestOptions) (*TransactionResponse, error) {
	request, err := newTransactionRequest(opt)
	if err != nil {
		return nil, errors.Wrap(err, "new transaction request")
	}
	jwtToken, err := c.genJWTToken(transactionEndpoint, request)
	if err != nil {
		return nil, err
	}

	response, err := c.apiRequest.Send(
		ctx,
		transactionEndpoint,
		http.MethodPost,
		request,
		c.getHeaders(jwtToken),
	)

	if err != nil {
		return nil, err
	}

	var res TransactionResponse
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling response")
	}

	return &res, nil
}

// newTransactionRequest creates a new transaction request.
func newTransactionRequest(opt TransactionRequestOptions) (createTransactionRequest, error) {
	h := sha256.New()
	wrappedMessage := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(opt.Message.Content)) + opt.Message.Content
	_, err := h.Write([]byte(wrappedMessage))
	if err != nil {
		return createTransactionRequest{}, errors.Wrap(err, "sha256")
	}
	hash := h.Sum(nil)

	var rawMessageData RawMessageData
	rawSigningMessage := UnsignedRawMessage{
		Content: hex.EncodeToString(hash),
	}
	rawMessageData.Messages = []UnsignedRawMessage{rawSigningMessage}

	req := createTransactionRequest{
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
			RawMessageData: rawMessageData,
		},
	}

	return req, nil
}
