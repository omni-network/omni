package app

import (
	"bytes"
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

type proxy struct {
	mu        sync.RWMutex
	instance  *anvilInstance
	target    *url.URL
	fuzzyHead bool
}

func newProxy(instance anvilInstance) (*proxy, error) {
	resp := new(proxy)
	if err := resp.setTarget(instance); err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *proxy) Proxy(w http.ResponseWriter, r *http.Request) {
	err := p.proxy(w, r)
	if err != nil {
		log.Error(r.Context(), "Proxy error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *proxy) proxy(w http.ResponseWriter, r *http.Request) error {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "read body")
	}

	var reqMsg jsonRPCMessage
	if err := json.Unmarshal(reqBody, &reqMsg); err != nil {
		return errors.Wrap(err, "unmarshal request")
	}

	nextReq, err := http.NewRequestWithContext(
		r.Context(),
		r.Method,
		p.GetTarget().String(),
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

	if resp.StatusCode != http.StatusOK || !p.IsFuzzyEnabled() {
		w.WriteHeader(resp.StatusCode)
		_, _ = io.Copy(w, resp.Body)

		return nil
	}

	shouldFuzz, height, err := isFuzzyXMsgLogFilter(r.Context(), p.GetTarget().String(), reqMsg)
	if err != nil {
		return errors.Wrap(err, "check for fuzzy log filter")
	}

	if !shouldFuzz {
		w.WriteHeader(resp.StatusCode)
		_, _ = io.Copy(w, resp.Body)

		return nil
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read response")
	}

	respBytes, _, err = parseAndFuzzXMsgs(r.Context(), height, respBytes)
	if err != nil {
		return errors.Wrap(err, "fuzz xmsgs")
	}

	_, _ = w.Write(respBytes)

	return nil
}

func (p *proxy) EnableFuzzyHead(_ http.ResponseWriter, r *http.Request) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.fuzzyHead = true

	log.Info(r.Context(), "Fuzzy head enabled")
}

func (p *proxy) DisableFuzzyHead(_ http.ResponseWriter, r *http.Request) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.fuzzyHead = false

	log.Info(r.Context(), "Fuzzy head disabled")
}

func (p *proxy) IsFuzzyEnabled() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.fuzzyHead
}

func (p *proxy) GetTarget() *url.URL {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.target
}

func (p *proxy) getInstance() *anvilInstance {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.instance
}

func (p *proxy) setTarget(target anvilInstance) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	u, err := url.Parse(target.URL())
	if err != nil {
		return errors.Wrap(err, "parse target url")
	}

	if p.instance != nil {
		p.instance.stop()
	}

	p.instance = &target
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
