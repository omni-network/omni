package appv2

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	statusOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "solver_v2",
		Subsystem: "processor",
		Name:      "status_offset",
		Help:      "Last inbox offset processed by chain and status",
	}, []string{"chain", "target", "status"})

	processedEvents = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "solver_v2",
		Subsystem: "processor",
		Name:      "processed_events_total",
		Help:      "Total number of events processed by chain and status",
	}, []string{"chain", "target", "status"})
)
