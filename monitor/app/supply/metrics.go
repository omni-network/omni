package supply

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	cChainSupply = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "supply",
		Name:      "cchain_total_gwei",
		Help:      "Token supply on the consensus chain in gwei",
	})

	eChainSupply = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "supply",
		Name:      "l1_erc20_total_gwei",
		Help:      "ERC20 token supply on Ethereum chain in gwei",
	})

	bridgeBalance = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "supply",
		Name:      "l1_bridge_balance_gwei",
		Help:      "ERC20 token balance of the bridge contract on Ethereum in gwei",
	})
)
