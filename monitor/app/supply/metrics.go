package supply

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	cChainSupply = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "supply",
		Name:      "supply_cchain_gwei",
		Help:      "Token supply on the consensus chain in gwei",
	})

	eChainSupply = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "supply",
		Name:      "total_supply_erc20_gwei",
		Help:      "ERC20 token supply on Ethereum chain in gwei",
	})

	bridgeBalance = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "supply",
		Name:      "bridge_balance_gwei",
		Help:      "ERC20 token balance of the bridge contract on Ethereum in gwei",
	})
)
