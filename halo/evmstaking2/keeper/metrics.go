package keeper

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	eventDeliveryHeight = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "evmstaking",
		Name:      "event_delivery_height",
		Help:      "The height at which all scheduled staking events were delivered",
	})

	scheduledEvents = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "evmstaking",
		Name:      "pending_events",
		Help:      "The number of pending staking events to be delivered",
	})

	failedEvents = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "evmstaking",
		Name:      "failed_events",
		Help:      "The number of staking events that could not be delivered",
	})
)
