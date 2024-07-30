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
)
