package app

import (
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

	rejectedOrders = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "solver",
		Subsystem: "processor",
		Name:      "rejected_orders_total",
		Help:      "Total number of rejected orders by chain and reason",
	}, []string{"chain", "reason"})

	orderAge = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "solver",
		Subsystem: "processor",
		Name:      "order_age_seconds",
		Help:      "Order age (from creation) in seconds by chain and status",
		Buckets:   prometheus.ExponentialBucketsRange(1, 60, 5),
	}, []string{"chain", "status"})

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
