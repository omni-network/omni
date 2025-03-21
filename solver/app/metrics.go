package app

import (
	"github.com/omni-network/omni/lib/promutil"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	statusOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "solver",
		Subsystem: "processor",
		Name:      "status_offset",
		Help:      "Last inbox offset processed by processor and status",
	}, []string{"proc", "status"})

	processedEvents = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "solver",
		Subsystem: "processor",
		Name:      "processed_events_total",
		Help:      "Total number of events processed by processor and status",
	}, []string{"proc", "status"})

	processorLag = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "solver",
		Subsystem: "processor",
		Name:      "lag_seconds",
		Help:      "The elapsed seconds since latest processed block timestamp per processor",
	}, []string{"proc"})

	rejectedOrders = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "solver",
		Subsystem: "processor",
		Name:      "rejected_orders_total",
		Help:      "Total number of rejected orders by source chain and reason",
	}, []string{"chain", "reason"})

	filledOrders = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "solver",
		Subsystem: "processor",
		Name:      "filled_orders_total",
		Help:      "Total number of filled orders by source chain, destination chain and target",
	}, []string{"src_chain", "dst_chain", "target"})

	orderAge = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "solver",
		Subsystem: "processor",
		Name:      "order_age_seconds",
		Help:      "Order age (from creation) in seconds by chain and status",
		Buckets:   prometheus.ExponentialBucketsRange(1, 60*60, 8),
	}, []string{"chain", "status"})

	oldestOrder = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "solver",
		Subsystem: "processor",
		Name:      "agecache_oldest_order_seconds",
		Help:      "Oldest order in age cache per chain in seconds",
	}, []string{"chain"})

	apiLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "solver",
		Subsystem: "api",
		Name:      "latency",
		Help:      "API server request latency in seconds per endpoint",
	}, []string{"endpoint"})

	apiResponses = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "solver",
		Subsystem: "api",
		Name:      "response_total",
		Help:      "Total responses served by the API server per endpoint per status code class (2xx, 4xx, 5xx)",
	}, []string{"endpoint", "class"})

	apiConcurrent = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "solver",
		Subsystem: "api",
		Name:      "concurrent_requests",
		Help:      "Number of concurrent requests being served by the API server (at scrape time)",
	})
)
