package rpc

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	apiLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "trade",
		Subsystem: "rpc",
		Name:      "latency",
		Help:      "RPC server request latency in seconds per endpoint",
	}, []string{"endpoint"})

	apiResponses = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "trade",
		Subsystem: "rpc",
		Name:      "response_total",
		Help:      "Total responses served by the API server per endpoint per status code class (2xx, 4xx, 5xx)",
	}, []string{"endpoint", "class"})

	apiConcurrent = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "trade",
		Subsystem: "rpc",
		Name:      "concurrent_requests",
		Help:      "Number of concurrent requests being served by the API server (at scrape time)",
	})
)
