package app

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	evmSynced = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "evm",
		Name:      "synced",
		Help:      "Constant gauge of 1 if attached the omni_evm is synced, 0 if syncing.",
	})

	evmHeight = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "evm",
		Name:      "height",
		Help:      "Latest block height of the attached omni_evm",
	})

	evmPeers = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "evm",
		Name:      "peers",
		Help:      "Number of execution P2P peers of the attached omni_evm",
	})

	cometSynced = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "comet",
		Name:      "synced",
		Help:      "Constant gauge of 1 if attached the cometBFT is synced, 0 if syncing.",
	})

	cometValidator = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "comet",
		Name:      "validator",
		Help:      "Constant gauge of 1 if local halo node is a cometBFT validator, 0 if not a validator.",
	})

	dbSize = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "db",
		Name:      "size_bytes",
		Help:      "Current size of the database directory in bytes.",
	})

	nodeReadiness = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "health",
		Name:      "ready",
		Help:      "Node readiness",
	})
)

// setConstantGauge sets the value of a gauge to 1 if b is true, 0 otherwise.
func setConstantGauge(gauge prometheus.Gauge, b bool) {
	var val float64
	if b {
		val = 1
	}
	gauge.Set(val)
}
