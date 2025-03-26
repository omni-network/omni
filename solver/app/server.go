package app

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tracer"

	"github.com/rs/cors"
	"go.opentelemetry.io/otel/attribute"
)

const (
	endpointQuote     = "/api/v1/quote"
	endpointContracts = "/api/v1/contracts"
	endpointCheck     = "/api/v1/check"
)

// serveAPI starts the API server, returning a async error.
func serveAPI(address string, handlers ...Handler) <-chan error {
	errChan := make(chan error)
	go func() {
		mux := http.NewServeMux()

		// Add all handlers
		for _, handler := range handlers {
			fn := handlerAdapter(handler)
			if !handler.SkipInstrument {
				fn = instrumentHandler(handler.Endpoint, fn)
			}
			mux.Handle(handler.Endpoint, fn)
		}

		// Add health check endpoints (not instrumented)
		mux.Handle("/live", newLiveHandler())
		mux.Handle("/", newLiveHandler()) // Also serve live from root for easy health checks

		c := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Origin", "Content-Type", "Accept"},
		})

		srv := &http.Server{
			Addr:              address,
			ReadHeaderTimeout: 5 * time.Second,
			IdleTimeout:       5 * time.Second,
			WriteTimeout:      5 * time.Second,
			Handler:           c.Handler(mux),
		}
		errChan <- errors.Wrap(srv.ListenAndServe(), "serve api")
	}()

	return errChan
}

// newLiveHandler returns a http handler that always return 200s.
// It indicates the API server is "live" (up and running).
// It isn't necessarily "ready" or "healthy" since downstream dependencies may be down.
func newLiveHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

// instrumentHandler wraps an http.Handler, instrumenting the latency and error rate.
func instrumentHandler(endpoint string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiConcurrent.Inc()
		defer apiConcurrent.Dec()

		t0 := time.Now()
		iw := &instrumentWriter{ResponseWriter: w}

		// Start trace
		ctx, span := tracer.Start(r.Context(), "api")
		span.SetAttributes(attribute.String("endpoint", endpoint))
		defer span.End()
		traceID := span.SpanContext().TraceID()

		ctx = log.WithCtx(ctx, log.Hex7("tid", traceID[:]))
		r = r.WithContext(ctx)

		handler.ServeHTTP(iw, r)

		latency := time.Since(t0)
		apiLatency.WithLabelValues(endpoint).Observe(latency.Seconds())
		apiResponses.WithLabelValues(endpoint, iw.StatusClass()).Inc()

		ip, typ := clientIP(r)
		log.Debug(ctx, "Served API request",
			"endpoint", endpoint,
			"status", iw.StatusCode(),
			"latency_millis", latency.Milliseconds(),
			"client_ip", ip,
			"client_ip_type", typ,
			"user_agent", r.UserAgent(),
		)
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

// StatusClass returns the status class of the response ("2xx", "4xx", "5xx") .
func (w *instrumentWriter) StatusClass() string {
	if w.status == 0 {
		return "2xx" // unset is 200
	}

	return fmt.Sprintf("%dxx", w.status/100)
}

// StatusCode returns the status code of the response.
func (w *instrumentWriter) StatusCode() int {
	if w.status == 0 {
		return http.StatusOK
	}

	return w.status
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
