package rpc

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/solver/types"
)

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

var gatewayTimeout = time.Second * 10

func handlerAdapter(h Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, rr *http.Request) {
		defer rr.Body.Close()
		ctx := rr.Context()
		ctx, cancel := context.WithTimeout(ctx, gatewayTimeout)
		defer cancel()

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			writeErrResponse(ctx, w, errors.Wrap(err, "read body", StatusAttr(http.StatusBadRequest)))
			return
		}

		req := h.ZeroReq()
		if req == nil {
			// Skip request unmarshalling if ZeroReq returns nil.
		} else if err := json.Unmarshal(body, req); err != nil {
			// TODO(corver): remove once issue identified
			log.DebugErr(ctx, "Failed to unmarshal request", err, "body", string(body))
			writeErrResponse(ctx, w, errors.Wrap(err, "unmarshal", StatusAttr(http.StatusBadRequest)))

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

	if sc, ok := getStatusCode(err); ok {
		statusCode = sc
	}

	log.DebugErr(ctx, "Serving API error", err, "status", statusCode)

	writeJSONResponse(ctx, w, statusCode, types.JSONErrorResponse{
		Error: types.JSONError{
			Code:    statusCode,
			Status:  http.StatusText(statusCode),
			Message: errors.Format(err),
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
