package provider

import (
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

	streamHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "stream_height",
		Help:      "Latest streamed xblock height per worker per source chain. Alert if not growing.",
	}, []string{"worker", "chain"})
)
