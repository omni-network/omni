package provider

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	callbackErrTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "xprovider",
		Name:      "callback_error_total",
		Help:      "Total number of callback errors per source chain version. Alert if growing.",
	}, []string{"chain_version"})

	fetchErrTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "xprovider",
		Name:      "fetch_error_total",
		Help:      "Total number of fetch errors per source chain version. Alert if growing.",
	}, []string{"chain_version"})

	streamHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "xprovider",
		Name:      "stream_height",
		Help:      "Latest streamed xblock height per source chain version. Alert if not growing.",
	}, []string{"chain_version"})

	callbackLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "lib",
		Subsystem: "xprovider",
		Name:      "callback_latency_seconds",
		Help:      "Callback latency in seconds per source chain version. Alert if growing.",
		Buckets:   []float64{.001, .002, .005, .01, .025, .05, .1, .25, .5, 1, 2.5},
	}, []string{"chain_version"})
)
