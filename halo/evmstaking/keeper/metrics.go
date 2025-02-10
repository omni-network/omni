package keeper

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	eventDeliveries = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "halo",
		Subsystem: "evmstaking",
		Name:      "deliveries_total",
		Help:      "The number of deliveries of buffered staking events",
	})

	bufferedEvents = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "evmstaking",
		Name:      "buffered_events",
		Help:      "The number of buffered staking events to be delivered",
	})

	failedEvents = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "halo",
		Subsystem: "evmstaking",
		Name:      "failed_events_total",
		Help:      "The number of staking events that could not be delivered",
	})
)
