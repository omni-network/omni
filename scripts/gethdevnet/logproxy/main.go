package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

//nolint:gochecknoglobals // Simple flags for a simple script
var (
	flagUpstream = flag.String("upstream", "http://geth:8551", "Upstream Geth RPC endpoint")
	flagAddr     = flag.String("addr", ":9551", "Address to listen on")
)

func main() {
	flag.Parse()

	if err := runProxy(*flagUpstream, *flagAddr); err != nil {
		slog.Error("Fatal error", slog.String("err", err.Error()))
		os.Exit(1)
	}

	slog.Info("Done")
}

func runProxy(upstreamAddr string, listenAddr string) error {
	target, err := url.Parse(upstreamAddr)
	if err != nil {
		return err
	}

	slog.Info("Serving logproxy", "listen", listenAddr, "upstream ", upstreamAddr)

	return http.ListenAndServe(listenAddr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewReader(body))

		rpcReq := struct {
			ID     int    `json:"id"`
			Method string `json:"method"`
			Params []any  `json:"params"`
		}{}
		if err := json.Unmarshal(body, &rpcReq); err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		wrapper := responseWriterWrapper{ResponseWriter: w}
		httputil.NewSingleHostReverseProxy(target).ServeHTTP(&wrapper, r)

		respBody := wrapper.body
		if wrapper.header.Get("Content-Encoding") == "gzip" {
			respBody, err = gunzip(respBody)
			if err != nil {
				slog.Error("Error gunzipping response body", "err", err)
			}
		}

		rpcResp := struct {
			Result any `json:"result"`
			Error  struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
				Data    any    `json:"data"`
			} `json:"error"`
		}{}
		if err := json.Unmarshal(respBody, &rpcResp); err != nil {
			slog.Error("Error parsing response body", "err", err, "body", string(respBody))
		}

		var respErr string
		if rpcResp.Error.Message != "" {
			respErr = fmt.Sprint(rpcResp.Error)
		}

		slog.Info("🌊 Proxied request",
			"id", rpcReq.ID,
			"method", rpcReq.Method,
			"params", rpcReq.Params,
			"response_code", wrapper.status,
			"response_result", rpcResp.Result,
			"response_error", respErr,
		)
	}))
}

func gunzip(body []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return io.ReadAll(r)
}

type responseWriterWrapper struct {
	http.ResponseWriter
	body   []byte
	status int
	header http.Header
}

func (w *responseWriterWrapper) Write(b []byte) (int, error) {
	w.body = append(w.body, b...)

	return w.ResponseWriter.Write(b)
}

func (w *responseWriterWrapper) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *responseWriterWrapper) Header() http.Header {
	w.header = w.ResponseWriter.Header()
	return w.header
}
