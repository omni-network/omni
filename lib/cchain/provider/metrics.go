package provider

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	callbackErrTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "callback_error_total",
		Help:      "Total number of callback errors per worker per source chain. Alert if growing.",
	}, []string{"worker", "chain"})

	fetchErrTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "fetch_error_total",
		Help:      "Total number of fetch errors per worker per source chain. Alert if growing.",
	}, []string{"worker", "chain"})

	streamHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "stream_height",
		Help:      "Latest streamed xblock height per worker per source chain. Alert if not growing.",
	}, []string{"worker", "chain"})

	callbackLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "callback_latency_seconds",
		Help:      "Callback latency in seconds per worker per source chain. Alert if growing.",
		Buckets:   []float64{.001, .002, .005, .01, .025, .05, .1, .25, .5, 1, 2.5},
	}, []string{"worker", "chain"})

	queryLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "query_latency_seconds",
		Help:      "Latency (in seconds) of halo ABCI queries per endpoint.",
	}, []string{"endpoint"})

	// TODO(corver): This is very similar to fetchErrTotal, maybe this is sufficient and we can remove fetchErrTotal?
	queryErrTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "query_error_total",
		Help:      "Total number of query errors per endpoint. Alert if growing.",
	}, []string{"endpoint"})
)

func latency(endpoint string) func() {
	start := time.Now()
	return func() {
		queryLatency.WithLabelValues(endpoint).Observe(time.Since(start).Seconds())
	}
}

func incQueryErr(endpoint string) {
	queryErrTotal.WithLabelValues(endpoint).Inc()
}
