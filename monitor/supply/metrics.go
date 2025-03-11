package supply

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	cChainSupply = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "supply",
		Name:      "cchain_total_ether",
		Help:      "Token supply on the consensus chain in ether",
	})

	l1Erc20Supply = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "supply",
		Name:      "l1_erc20_total_ether",
		Help:      "OMNI ERC20 token supply on Ethereum chain in ether",
	})

	bridgeBalance = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "supply",
		Name:      "l1_bridge_balance_ether",
		Help:      "OMNI ERC20 token balance of the bridge contract on Ethereum in ether",
	})
)
