package gasprice

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	liveGasPrice = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xfeemngr",
		Name:      "live_gas_price",
		Help:      "Live gas price",
	}, []string{"chain"})

	bufferUpdates = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "xfeemngr",
		Name:      "buffer_updates_total",
		Help:      "The total number of buffer updates",
	}, []string{"chain"})

	bufferedGasPrice = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xfeemngr",
		Name:      "buffered_gas_price",
		Help:      "Buffered gas price",
	}, []string{"chain"})
)
