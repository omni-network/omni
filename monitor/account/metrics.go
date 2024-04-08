package account

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	accountBalance = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "account",
		Name:      "balance_ether",
		Help:      "The balance of the account on a specific chain in ether. Alert if low.",
	}, []string{"chain", "type"})

	accountNonce = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "account",
		Name:      "nonce",
		Help:      "The nonce of the account on a specific chain",
	}, []string{"chain", "type"})
)
