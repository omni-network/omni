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

	reorgTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "ethclient",
		Name:      "cache_reorg_total",
		Help:      "Total number of reorgs detected by chain",
	}, []string{"chain"})

	cacheHits = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "ethclient",
		Name:      "cache_hits_total",
		Help:      "Total number of cache hits by chain",
	}, []string{"chain"})

	cacheMisses = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "ethclient",
		Name:      "cache_misses_total",
		Help:      "Total number of cache misses by chain",
	}, []string{"chain"})

	// Using gauge for websocket latency is good enough (avoid expensive histogram).
	websocketLatency = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "ethclient",
		Name:      "wss_header_latency_seconds",
		Help:      "Last header age in seconds received via websockets by chain",
	}, []string{"chain"})
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
