package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/uni"
	"github.com/omni-network/omni/lib/unibackend"
	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
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

func (e APIError) Unwrap() error {
	return e.Err
}

func (e APIError) Format(st fmt.State, verb rune) {
	if fmter, ok := e.Err.(fmt.Formatter); ok {
		fmter.Format(st, verb)
	} else {
		_, _ = io.WriteString(st, e.Err.Error())
	}
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
				Portal: uni.EVMAddress(addrs.Portal),
				Inbox:  uni.EVMAddress(addrs.SolverNetInbox),
				Outbox: uni.EVMAddress(addrs.SolverNetOutbox),
				// Middleman deprecated and logic moved to executor, temporarily retained for backwards compatibility.
				Middleman: uni.EVMAddress(addrs.SolverNetExecutor),
				Executor:  uni.EVMAddress(addrs.SolverNetExecutor),
			}, nil
		},
	}
}

// newCheckHandler returns a handler for the /check endpoint.
// It is responsible for http request / response handling, and delegates
// logic to a checkFunc.
func newCheckHandler(checkFunc checkFunc, traceFunc traceFunc) Handler {
	return Handler{
		Endpoint: endpointCheck,
		ZeroReq:  func() any { return &types.CheckRequest{} },
		HandleFunc: func(ctx context.Context, request any) (any, error) {
			req, ok := request.(*types.CheckRequest)
			if !ok {
				return nil, errors.New("invalid request type [BUG]", "type", fmt.Sprintf("%T", request))
			}

			// Returns trace result if debug == true, else nil.
			maybeTrace := func() map[string]any {
				if !req.Debug {
					return nil
				}

				trace, err := traceFunc(ctx, *req)
				if err != nil {
					return map[string]any{"error": errors.Format(err)}
				}

				return trace.Map()
			}

			err := checkFunc(ctx, *req)
			if r := new(RejectionError); errors.As(err, &r) {
				return types.CheckResponse{
					Rejected:          true,
					RejectCode:        r.Reason,
					RejectReason:      r.Reason.String(),
					RejectDescription: errors.Format(r.Err),
					Trace:             maybeTrace(),
				}, nil
			} else if err != nil {
				return types.CheckResponse{}, err
			}

			return types.CheckResponse{
				Accepted: true,
				Trace:    maybeTrace(),
			}, nil
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

			res, err := quoteFunc(ctx, *req)
			if r := new(RejectionError); errors.As(err, &r) {
				return types.QuoteResponse{
					// include quoted response, even if rejected (useful for min/max rejections)
					Deposit:           res.Deposit,
					Expense:           res.Expense,
					Rejected:          true,
					RejectCode:        r.Reason,
					RejectReason:      r.Reason.String(),
					RejectDescription: errors.Format(r.Err),
				}, nil
			} else if err != nil {
				return types.QuoteResponse{}, err
			}

			return res, nil
		},
	}
}

func newPriceHandler(priceFunc priceHandlerFunc) Handler {
	return Handler{
		Endpoint: endpointPrice,
		ZeroReq:  func() any { return &types.PriceRequest{} },
		HandleFunc: func(ctx context.Context, request any) (any, error) {
			req, ok := request.(*types.PriceRequest)
			if !ok {
				return nil, errors.New("invalid request type [BUG]", "type", fmt.Sprintf("%T", request))
			}

			res, err := priceFunc(ctx, *req)
			if err != nil {
				return nil, err
			}

			return res, nil
		},
	}
}

func newTokensHandler(backends unibackend.Backends, solverAddr common.Address) Handler {
	return Handler{
		Endpoint:       endpointTokens,
		SkipInstrument: true, // Reduce noise as this endpoint returns static data.
		ZeroReq:        func() any { return nil },
		HandleFunc: func(ctx context.Context, _ any) (any, error) {
			return tokensResponse(ctx, backends, solverAddr)
		},
	}
}

// newRelayHandler returns a handler for the /relay endpoint.
// It validates gasless orders and signatures, then submits them via openFor.
func newRelayHandler(relayFunc relayFunc) Handler {
	return Handler{
		Endpoint: endpointRelay,
		ZeroReq:  func() any { return &types.RelayRequest{} },
		HandleFunc: func(ctx context.Context, request any) (any, error) {
			req, ok := request.(*types.RelayRequest)
			if !ok {
				return nil, errors.New("invalid request type [BUG]", "type", fmt.Sprintf("%T", request))
			}

			res, err := relayFunc(ctx, *req)
			if r := new(RelayError); errors.As(err, &r) {
				return types.RelayResponse{
					Success: false,
					Error: &types.RelayError{
						Code:        r.Code,
						Message:     r.Message,
						Description: r.Description,
					},
				}, nil
			} else if err != nil {
				return types.RelayResponse{}, err
			}

			return res, nil
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

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			writeErrResponse(ctx, w, newAPIError(err, http.StatusBadRequest))
			return
		}

		req := h.ZeroReq()
		if req == nil { //nolint:revive // noop if-block for readability
			// Skip request unmarshalling if ZeroReq returns nil.
		} else if err := json.Unmarshal(body, req); err != nil {
			// TODO(corver): remove once issue identified
			log.DebugErr(ctx, "Failed to unmarshal request", err, "body", string(body))
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
			Message: removeBUG(errors.Format(err)),
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
