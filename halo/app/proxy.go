package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

const (
	evmProxyMaxReq  = 256 * 1024 // 256 KiB
	evmProxyTimeout = 10 * time.Second
)

// startEVMProxy starts an HTTP server that proxies EVM JSON-RPC requests to the target.
// This is mostly for debugging purposes.
func startEVMProxy(ctx context.Context, abort chan<- error, listen string, target string) func(context.Context) error {
	if listen == "" || target == "" {
		log.Debug(ctx, "Not starting evm proxy")
		return func(context.Context) error { return nil }
	}

	log.Info(ctx, "Starting evm proxy", "lister", listen, "target", target)

	// Create a new HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := serveProxy(ctx, w, r, target); err != nil {
			log.DebugErr(ctx, "Proxy evm request error", err)
			http.Error(w, errors.Format(err), http.StatusInternalServerError)
		}
	})

	srv := &http.Server{
		Addr:              listen,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		Handler:           mux,
	}

	go func() {
		abort <- errors.Wrap(srv.ListenAndServe(), "serve api")
	}()

	return srv.Shutdown
}

func serveProxy(ctx context.Context, w http.ResponseWriter, r *http.Request, target string) error {
	// Limit read size to prevent abuse
	maxReader := http.MaxBytesReader(w, r.Body, evmProxyMaxReq)
	defer maxReader.Close()

	reqBody, err := io.ReadAll(maxReader)
	if err != nil {
		return errors.Wrap(err, "read body")
	}

	t0 := time.Now()
	id := t0.UnixNano() % 10_000 // Random ID for logging
	ip, _ := clientIP(r)
	var method, params string

	var msg jsonRPCMessage
	if err := json.Unmarshal(reqBody, &msg); err == nil {
		method = msg.Method
		params = string(msg.Params)
		if len(params) > 256 {
			params = fmt.Sprintf("%s...[%d]", params[:256], len(params))
		}
	}
	log.Debug(ctx, "Proxy evm request start",
		"id", id,
		"method", method,
		"params", params,
		"client_ip", ip,
		"req_body", len(reqBody),
	)

	ctx, cancel := context.WithTimeout(ctx, evmProxyTimeout)
	defer cancel()

	nextReq, err := http.NewRequestWithContext(
		ctx,
		r.Method,
		target,
		closeReader{bytes.NewReader(reqBody)},
	)
	if err != nil {
		return errors.Wrap(err, "create next request")
	}
	nextReq.Header = r.Header

	resp, err := new(http.Client).Do(nextReq)
	if err != nil {
		return errors.Wrap(err, "do request", "id", id)
	}
	defer resp.Body.Close()

	for k, vs := range resp.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}

	w.WriteHeader(resp.StatusCode)
	respBodyLen, _ := io.Copy(w, resp.Body)

	log.Debug(ctx, "Proxy evm request end",
		"id", id,
		"duration", time.Since(t0),
		"resp_len", respBodyLen,
		"status", resp.StatusCode,
	)

	return nil
}

// clientIP returns the client IP address from the request and the type/header used.
func clientIP(r *http.Request) (ip string, typ string) { //nolint:nonamedreturns // Disambiguate identical return types
	// first returns the first IP address in a comma-separated list.
	// Or just the string otherwise.
	first := func(ip string) string {
		return strings.Split(ip, ",")[0]
	}

	for _, header := range []string{
		"CF-Connecting-IP", // Use CloudFlare if present
		"X-Forwarded-For",  // Otherwise GCP / AWS LB
	} {
		if ip := r.Header.Get(header); ip != "" {
			return first(ip), header
		}
	}

	return first(r.RemoteAddr), "RemoteAddr" // Fallback to remote address
}

var _ io.ReadCloser = closeReader{}

type closeReader struct {
	io.Reader
}

func (closeReader) Close() error {
	return nil
}

// A value of this type can a JSON-RPC request, notification, successful response or
// error response. Which one it is depends on the fields.
type jsonRPCMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *jsonRPCError   `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

type jsonRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
