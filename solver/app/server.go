package app

import (
	"net/http"
	"time"

	"github.com/omni-network/omni/lib/errors"
)

// serveAPI starts the API server, returning a async error.
func serveAPI(address string, endpoints map[string]http.Handler) <-chan error {
	errChan := make(chan error)
	go func() {
		mux := http.NewServeMux()

		endpoints["/live"] = newLiveHandler()

		for endpoint, handler := range endpoints {
			mux.Handle(endpoint, instrumentHandler(endpoint, handler))
		}

		srv := &http.Server{
			Addr:              address,
			ReadHeaderTimeout: 5 * time.Second,
			IdleTimeout:       5 * time.Second,
			WriteTimeout:      5 * time.Second,
			Handler:           mux,
		}
		errChan <- errors.Wrap(srv.ListenAndServe(), "serve api")
	}()

	return errChan
}

// newLiveHandler returns a http handler that always return 200s.
// It indicates the API server is "live" (up and running).
func newLiveHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

// instrumentHandler wraps an http.Handler, instrumenting the latency and error rate.
func instrumentHandler(endpoint string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		iw := &instrumentWriter{ResponseWriter: w}

		handler.ServeHTTP(iw, r)

		apiLatency.WithLabelValues(endpoint).Observe(time.Since(t0).Seconds())
		apiErrors.WithLabelValues(endpoint).Add(iw.ErrorCount())
	})
}

// instrumentWriter wraps the response writer, tracking the status code written.
type instrumentWriter struct {
	http.ResponseWriter
	status int
}

func (w *instrumentWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *instrumentWriter) ErrorCount() float64 {
	if w.status < 400 {
		return 0 // 200 or unset is not an error
	}

	return 1
}
