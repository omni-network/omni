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
		Help:      "Last inbox offset processed by chain and status",
	}, []string{"chain", "status"})

	processedEvents = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "solver",
		Subsystem: "processor",
		Name:      "processed_events_total",
		Help:      "Total number of events processed by chain and status",
	}, []string{"chain", "status"})
)
