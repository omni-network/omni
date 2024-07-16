package app

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// JSONRPCMessage is a JSON-RPC request, notification, successful response or
// error response. Which one it is depends on the fields.
type JSONRPCMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *JSONRPCError   `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type proxy struct {
	target *url.URL
	mwares []Middleware
}

type Middleware func(context.Context, JSONRPCMessage) (JSONRPCMessage, error)

func newProxy(target string, mwares ...Middleware) (*proxy, error) {
	if len(mwares) == 0 {
		return nil, errors.New("no middlewares provided")
	}

	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, errors.Wrap(err, "parse target")
	}

	return &proxy{
		target: targetURL,
		mwares: mwares,
	}, nil
}

func (p *proxy) Proxy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = log.WithCtx(ctx, "remote_addr", r.RemoteAddr)

	err := p.proxy(ctx, w, r)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			http.Error(w, "proxy closed", http.StatusServiceUnavailable)
			return
		}

		log.Error(ctx, "Proxy error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *proxy) proxy(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "read body")
	}

	var reqMsg JSONRPCMessage
	if err := json.Unmarshal(reqBody, &reqMsg); err != nil {
		return errors.Wrap(err, "unmarshal request")
	}

	// apply middlewares
	for _, mware := range p.mwares {
		reqMsg, err = mware(ctx, reqMsg)
		if err != nil {
			return errors.Wrap(err, "middleware")
		}
	}

	reqBody, err = json.Marshal(reqMsg)
	if err != nil {
		return errors.Wrap(err, "marshal request")
	}

	nextReq, err := http.NewRequestWithContext(
		ctx,
		r.Method,
		p.target.String(),
		closeReader{bytes.NewReader(reqBody)},
	)
	if err != nil {
		return errors.Wrap(err, "create next request")
	}
	nextReq.Header = r.Header

	resp, err := new(http.Client).Do(nextReq)
	if err != nil {
		return errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)

	return nil
}

var _ io.ReadCloser = closeReader{}

type closeReader struct {
	io.Reader
}

func (closeReader) Close() error {
	return nil
}
