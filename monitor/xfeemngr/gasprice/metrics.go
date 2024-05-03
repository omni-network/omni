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

	bufferedGasPrice = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xfeemngr",
		Name:      "buffered_gas_price",
		Help:      "Buffered gas price",
	}, []string{"chain"})
)
