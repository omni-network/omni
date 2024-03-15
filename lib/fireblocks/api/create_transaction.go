package fireblocks

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/omni-network/omni/lib/errors"
)

type CreateTransactionRequest struct {
	Operation          string         `json:"operation"`
	Note               string         `json:"note"`
	ExternalTxID       string         `json:"externalTxId"`
	AssetID            string         `json:"assetId"`
	Source             any            `json:"source"`
	Destination        any            `json:"destination"`
	Destinations       any            `json:"destinations"`
	Amount             string         `json:"amount"`
	TreatAsGrossAmount bool           `json:"treatAsGrossAmount"`
	ForceSweep         bool           `json:"forceSweep"`
	FeeLevel           string         `json:"feeLevel"`
	Fee                string         `json:"fee"`
	PriorityFee        string         `json:"priorityFee"`
	FailOnLowFee       bool           `json:"failOnLowFee"`
	MaxFee             string         `json:"maxFee"`
	GasLimit           string         `json:"gasLimit"`
	GasPrice           string         `json:"gasPrice"`
	NetworkFee         string         `json:"networkFee"`
	ReplaceTxByHash    string         `json:"replaceTxByHash"`
	CustomerRefID      string         `json:"customerRefId"`
	ExtraParameters    RawMessageData `json:"extraParameters"`
}

type RawMessageData struct {
	Messages  []UnsignedRawMessage `json:"messages"`
	Algorithm string               `json:"algorithm"`
}

type UnsignedRawMessage struct {
	Content           string    `json:"content"`
	Bip44addressIndex float64   `json:"bip44AddressIndex"`
	Bip44change       float64   `json:"bip44Change"`
	DerivationPath    []float64 `json:"derivationPath"`
}

type TransactionResponse struct {
	ID             string         `json:"id"`
	Status         string         `json:"status"`
	SystemMessages SystemMessages `json:"systemMessages"`
}
type SystemMessages struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (c *Client) CreateTransaction(ctx context.Context, opts TransactionRequestOptions) (*TransactionResponse, error) {
	var res TransactionResponse

	request, err := NewTransactionRequest(opts)
	if err != nil {
		return nil, errors.Wrap(err, "newTransactionRequest")
	}

	req, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	httpReq, err := c.createPostRequest(ctx, TransactionEndpoint, req)
	if err != nil {
		return nil, errors.Wrap(err, "getPostRequest")
	}
	response, err := c.sendRequest(
		ctx,
		httpReq,
		TransactionEndpoint,
		req,
		opts.Nonce,
	)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &res)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal response JSON")
	}

	return &res, nil
}

type TransactionRequestOptions struct {
	Nonce    string
	Amount   int64
	Messages []UnsignedRawMessage
}

func NewTransactionRequest(opt TransactionRequestOptions) (*CreateTransactionRequest, error) {
	req := &CreateTransactionRequest{
		Operation:          "RAW",
		Note:               "",
		ExternalTxID:       "",
		AssetID:            "",
		Source:             nil,
		Destination:        nil,
		Amount:             strconv.Itoa(int(opt.Amount)),
		TreatAsGrossAmount: false,
		ForceSweep:         false,
		FeeLevel:           "",
		Fee:                "",
		PriorityFee:        "",
		FailOnLowFee:       false,
		MaxFee:             "",
		GasLimit:           "",
		GasPrice:           "",
		NetworkFee:         "",
		ReplaceTxByHash:    "",
		CustomerRefID:      "",
		ExtraParameters:    RawMessageData{},
	}

	return req, nil
}
