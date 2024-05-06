package xfeemngr

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	onChainGasPrice = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xfeemngr",
		Name:      "on_chain_gas_price",
		Help:      "Gas price for the destination chain, set on the source chain",
	}, []string{"src_chain", "dest_chain"})

	onChainConversionRate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xfeemngr",
		Name:      "on_chain_conversion_rate",
		Help:      "Dest-to-source conversion rate, set on the source chain",
	}, []string{"src_chain", "dest_chain", "src_token", "dest_token"})

	portalBalance = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xfeemngr",
		Name:      "portal_balance",
		Help:      "Balance of the portal contract",
	}, []string{"chain", "address"})
)
