package types

import (
	"net/http"

	"github.com/omni-network/omni/lib/contracts/solvernet"
)

type JSONResponse interface {
	StatusCode() int
}

// JSONErrorResponse is a json response for http errors (e.g 4xx, 5xx), not used for rejections.
type JSONErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// CheckRequest is the expected request body for the /api/v1/check endpoint.
//
// NOTE: Check request / response types mirror SolvertNet.OrderData, built
// specifically for EVM -> EVM orders via SolverNetInbox / Outbox contracts,
// with ERC7683 type hash matching SolverNetInbox.ORDERDATA_TYPEHASH.
//
// To support multiple order types with this api (e.g. EVM -> Solana, Solana -> EVM)
// we'd need a more generic request / response format that discriminates on
// order type hash.
type CheckRequest struct {
	SourceChainID      uint64             `json:"sourceChainId"`
	DestinationChainID uint64             `json:"destChainId"`
	FillDeadline       uint32             `json:"fillDeadline"`
	Calls              solvernet.Calls    `json:"calls"`
	Expenses           solvernet.Expenses `json:"expenses"`
	Deposit            solvernet.Deposit  `json:"deposit"`
}

// CheckResponse is the response json for the /check endpoint.
type CheckResponse struct {
	Accepted          bool               `json:"accepted,omitempty"`
	Rejected          bool               `json:"rejected,omitempty"`
	RejectReason      string             `json:"rejectReason,omitempty"`
	RejectDescription string             `json:"rejectDescription,omitempty"`
	Error             *JSONErrorResponse `json:"error,omitempty"`
}

var _ JSONResponse = (*CheckResponse)(nil)

func (r CheckResponse) StatusCode() int {
	if r.Error != nil {
		return r.Error.Code
	}

	return http.StatusOK
}

// ContractsResponse is the response json for the /api/vi/contracts endpoint.
type ContractsResponse struct {
	Portal    string             `json:"portal,omitempty"`
	Inbox     string             `json:"inbox,omitempty"`
	Outbox    string             `json:"outbox,omitempty"`
	Middleman string             `json:"middleman,omitempty"`
	Error     *JSONErrorResponse `json:"error,omitempty"`
}

var _ JSONResponse = (*ContractsResponse)(nil)

func (r ContractsResponse) StatusCode() int {
	if r.Error != nil {
		return r.Error.Code
	}

	return http.StatusOK
}
