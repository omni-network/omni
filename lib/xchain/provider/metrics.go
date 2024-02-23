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
		Help:      "Total number of callback errors per source chain. Alert if growing.",
	}, []string{"chain"})

	streamHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "xprovider",
		Name:      "stream_height",
		Help:      "Latest streamed xblock height per source chain. Alert if not growing.",
	}, []string{"chain"})
)
