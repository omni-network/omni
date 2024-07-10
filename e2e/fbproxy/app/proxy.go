package app

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// A value of this type can a JSON-RPC request, notification, successful response or
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
	mu     sync.RWMutex
	target *url.URL
	mwares []Middleware
}

type Middleware func(context.Context, JSONRPCMessage) (JSONRPCMessage, error)

func newProxy(target string, mwares ...Middleware) (*proxy, error) {
	if len(mwares) == 0 {
		return nil, errors.New("no middlewares provided")
	}

	resp := new(proxy)

	if err := resp.setTarget(target); err != nil {
		return nil, err
	}

	resp.mwares = mwares

	return resp, nil
}

func (p *proxy) Proxy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = log.WithCtx(ctx, "remote_addr", r.RemoteAddr)

	err := p.proxy(ctx, w, r)
	if err != nil {
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
	for _, mware := range p.getMiddlewares() {
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
		p.getTarget().String(),
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

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		_, _ = io.Copy(w, resp.Body)

		return nil
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read response")
	}

	_, _ = w.Write(respBytes)

	return nil
}

func (p *proxy) getTarget() *url.URL {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.target
}

func (p *proxy) getMiddlewares() []Middleware {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.mwares
}

func (p *proxy) setTarget(target string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	u, err := url.Parse(target)
	if err != nil {
		return errors.Wrap(err, "parse target url")
	}

	p.target = u

	return nil
}

var _ io.ReadCloser = closeReader{}

type closeReader struct {
	io.Reader
}

func (closeReader) Close() error {
	return nil
}
