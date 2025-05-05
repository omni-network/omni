package coingecko

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var latency = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "lib",
	Subsystem: "coingecko",
	Name:      "latency_seconds",
	Help:      "Latency in seconds for coingecko requests by endpoint",
}, []string{"endpoint"})
