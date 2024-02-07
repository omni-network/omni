package provider

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//nolint:gochecknoglobals // Promauto metrics are global.
var (
	callbackErrTotal = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "callback_error_total",
		Help:      "Total number of callback errors per source chain. Alert if growing.",
	})

	streamHeight = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "cprovider",
		Name:      "stream_height",
		Help:      "Latest streamed xblock height per source chain. Alert if not growing.",
	})
)
