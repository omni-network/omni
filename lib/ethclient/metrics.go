package ethclient

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	latencyHist = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "lib",
		Subsystem: "ethclient",
		Name:      "latency_seconds",
		Help:      "Latency in seconds for ethereum JSON-RPC requests by chain and endpoint",
	}, []string{"chain", "endpoint"})

	errorCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "ethclient",
		Name:      "errors_total",
		Help:      "Total number of errors returned by a Ethereum JSON-RPC by chain and endpoint",
	}, []string{"chain", "endpoint"})
)

// latency returns a function that records the latency of an RPC call.
func latency(chain string, endpoint string) func() {
	start := time.Now()
	return func() {
		latencyHist.WithLabelValues(chain, endpoint).Observe(time.Since(start).Seconds())
	}
}

// incError increments the error count for a given chain and endpoint.
func incError(chain, endpoint string) {
	errorCount.WithLabelValues(chain, endpoint).Inc()
}

func spanName(endpoint string) string {
	return "ethclient/" + endpoint
}
