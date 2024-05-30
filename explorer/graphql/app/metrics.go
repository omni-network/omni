package app

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	inFlight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "in_flight_requests",
		Subsystem: "http_server",
		Help:      "A gauge of requests currently being served by the wrapped handler.",
	}, []string{"handler"})

	counter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name:      "requests_total",
			Subsystem: "http_server",
			Help:      "A counter for requests to the wrapped handler.",
		},
		[]string{"code", "handler", "method"},
	)

	// duration is partitioned by the HTTP method and handler. It uses custom
	// buckets based on the expected request duration.
	duration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:      "request_duration_seconds",
			Subsystem: "http_server",
			Help:      "A histogram of latencies for requests.",
			Buckets:   []float64{0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		},
		[]string{"handler", "method"},
	)

	requestSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:      "request_size_bytes",
			Subsystem: "http_server",
			Help:      "A histogram of request sizes for requests.",
			Buckets:   prometheus.ExponentialBucketsRange(128, 1024*1024, 8),
		},
		[]string{"handler"},
	)

	responseSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:      "response_size_bytes",
			Subsystem: "http_server",
			Help:      "A histogram of response sizes for requests.",
			Buckets:   prometheus.ExponentialBucketsRange(128, 1024*1024, 8),
		},
		[]string{"handler"},
	)
)

// Middleware is a function that wraps an http.Handler and returns a new http.Handler.
type Middleware func(http.Handler) http.Handler

// instrumentHandler wraps an http.Handler and provides Prometheus metrics for it.
func instrumentHandler(name string) Middleware {
	return func(next http.Handler) http.Handler {
		// Instrument the handlers with all the metrics, injecting the "handler" label by currying.
		return promhttp.InstrumentHandlerInFlight(inFlight.With(prometheus.Labels{"handler": name}),
			promhttp.InstrumentHandlerDuration(duration.MustCurryWith(prometheus.Labels{"handler": name}),
				promhttp.InstrumentHandlerCounter(counter.MustCurryWith(prometheus.Labels{"handler": name}),
					promhttp.InstrumentHandlerRequestSize(requestSize.MustCurryWith(prometheus.Labels{"handler": name}),
						promhttp.InstrumentHandlerResponseSize(responseSize.MustCurryWith(prometheus.Labels{"handler": name}), next),
					),
				),
			),
		)
	}
}
