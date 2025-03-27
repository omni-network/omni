package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/solver/types"
)

func newAPIError(err error, statusCode int) APIError {
	return APIError{Err: err, StatusCode: statusCode}
}

// APIError wraps an error with a non-200 HTTP status code.
type APIError struct {
	Err        error
	StatusCode int
}

func (e APIError) Error() string {
	return e.Err.Error()
}

type Handler struct {
	// Endpoint is the http endpoint path.
	Endpoint string
	// ZeroReq returns a zero struct pointer of the request type used for marshaling incoming requests.
	ZeroReq func() any
	// HandleFunc is the function that handles the request and returns a response.
	// The request will be a pointer (same at type returned by ZeroReq).
	// The response must be a struct and optional error.
	HandleFunc func(context.Context, any) (any, error)

	// SkipInstrument skips the handler instrumentation.
	SkipInstrument bool
}

// newContractsHandler returns a http handler that returns the contract address for `network`.
func newContractsHandler(addrs contracts.Addresses) Handler {
	return Handler{
		Endpoint:       endpointContracts,
		SkipInstrument: true, // Reduce noise as this endpoint returns static data.
		ZeroReq:        func() any { return nil },
		HandleFunc: func(context.Context, any) (any, error) {
			return types.ContractsResponse{
				Portal:    addrs.Portal,
				Inbox:     addrs.SolverNetInbox,
				Outbox:    addrs.SolverNetOutbox,
				Middleman: addrs.SolverNetMiddleman,
				Executor:  addrs.SolverNetExecutor,
			}, nil
		},
	}
}

// newCheckHandler returns a handler for the /check endpoint.
// It is responsible for http request / response handling, and delegates
// logic to a checkFunc.
func newCheckHandler(checkFunc checkFunc) Handler {
	return Handler{
		Endpoint: endpointCheck,
		ZeroReq:  func() any { return &types.CheckRequest{} },
		HandleFunc: func(ctx context.Context, request any) (any, error) {
			req, ok := request.(*types.CheckRequest)
			if !ok {
				return nil, errors.New("invalid request type [BUG]", "type", fmt.Sprintf("%T", request))
			}

			err := checkFunc(ctx, *req)
			if r := new(RejectionError); errors.As(err, &r) {
				return types.CheckResponse{
					Rejected:          true,
					RejectReason:      r.Reason.String(),
					RejectDescription: r.Err.Error(),
				}, nil
			} else if err != nil {
				return types.CheckResponse{}, err
			}

			return types.CheckResponse{Accepted: true}, nil
		},
	}
}

// newQuoteHandler returns a handler for the /quote endpoint.
// It is responsible to http request / response handling, and delegates
// logic to a quoteFunc.
func newQuoteHandler(quoteFunc quoteFunc) Handler {
	return Handler{
		Endpoint: endpointQuote,
		ZeroReq:  func() any { return &types.QuoteRequest{} },
		HandleFunc: func(ctx context.Context, request any) (any, error) {
			req, ok := request.(*types.QuoteRequest)
			if !ok {
				return nil, errors.New("invalid request type [BUG]", "type", fmt.Sprintf("%T", request))
			}

			return quoteFunc(ctx, *req)
		},
	}
}

var gatewayTimeout = time.Second * 10

func handlerAdapter(h Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, rr *http.Request) {
		defer rr.Body.Close()
		ctx := rr.Context()
		ctx, cancel := context.WithTimeout(ctx, gatewayTimeout)
		defer cancel()

		req := h.ZeroReq()
		if req == nil { //nolint:revive // noop if-block for readability
			// Skip request unmarshalling if ZeroReq returns nil.
		} else if err := json.NewDecoder(rr.Body).Decode(req); err != nil {
			writeErrResponse(ctx, w, newAPIError(err, http.StatusBadRequest))
			return
		}

		res, err := h.HandleFunc(ctx, req)
		if err != nil {
			writeErrResponse(ctx, w, err)
			return
		}

		writeJSONResponse(ctx, w, http.StatusOK, res)
	})
}

func writeErrResponse(ctx context.Context, w http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError
	if errors.Is(ctx.Err(), context.Canceled) {
		// If request context is canceled, return a 408 instead of 500.
		statusCode = http.StatusRequestTimeout
	} else if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		// If request context deadline exceeded, return a 504 instead of 500.
		statusCode = http.StatusGatewayTimeout
	}

	var apiErr APIError
	if errors.As(err, &apiErr) {
		statusCode = apiErr.StatusCode
	}

	log.DebugErr(ctx, "Serving API error", err, "status", statusCode)

	writeJSONResponse(ctx, w, statusCode, types.JSONErrorResponse{
		Error: types.JSONError{
			Code:    statusCode,
			Status:  http.StatusText(statusCode),
			Message: removeBUG(err.Error()),
		},
	})
}

func writeJSONResponse(ctx context.Context, w http.ResponseWriter, status int, resp any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Error(ctx, "Failed writing api response [BUG]", err)
	}
}

// removeBUG removes [BUG] from the error messages, so they are not included in responses to users.
func removeBUG(s string) string { return strings.ReplaceAll(s, "[BUG]", "") }
