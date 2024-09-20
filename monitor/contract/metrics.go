package contract

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	contractBalance = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "contract",
		Name:      "balance_ether",
		Help:      "The balance of the contract on a specific chain in ether. Alert if low.",
	}, []string{"chain", "name"})

	contractBalanceLow = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "contract",
		Name:      "balance_low",
		Help:      "Constant gauge indicating whether the contract balance is below the minimum threshold (1=true,0=false)",
	}, []string{"chain", "bame"})
)
