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
		Help:      "Total number of callback errors per worker per source chain version. Alert if growing.",
	}, []string{"worker", "chain_version"})

	fetchErrTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "fetch_error_total",
		Help:      "Total number of fetch errors per worker per source chain version. Alert if growing.",
	}, []string{"worker", "chain_version"})

	streamHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "stream_offset",
		Help:      "Latest streamed xblock offset per worker per source chain version. Alert if not growing.",
	}, []string{"worker", "chain_version"})

	callbackLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "callback_latency_seconds",
		Help:      "Callback latency in seconds per worker per source chain version. Alert if growing.",
		Buckets:   []float64{.001, .002, .005, .01, .025, .05, .1, .25, .5, 1, 2.5},
	}, []string{"worker", "chain_version"})

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

	fetchLookbackSteps = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "fetch_lookback_steps",
		Buckets:   []float64{0, 1, 2, 4, 8, 16, 32, 64, 128, 256},
		Help:      "Number of steps in the exponential backoff process to find a start for binary search",
	}, []string{"chain_version"})

	fetchBinarySearchSteps = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "fetch_binary_search_steps",
		Buckets:   []float64{0, 1, 2, 4, 8, 16, 32, 64, 128, 256},
		Help:      "Number of steps in the binary search process to find the right height",
	}, []string{"chain_version"})
)

func lookbackStepsMetric(chainName string, steps int) {
	fetchLookbackSteps.WithLabelValues(chainName).Observe(float64(steps))
}

func binarySearchStepsMetric(chainName string, steps int) {
	fetchBinarySearchSteps.WithLabelValues(chainName).Observe(float64(steps))
}

func latency(endpoint string) func() {
	start := time.Now()
	return func() {
		queryLatency.WithLabelValues(endpoint).Observe(time.Since(start).Seconds())
	}
}

func incQueryErr(endpoint string) {
	queryErrTotal.WithLabelValues(endpoint).Inc()
}
