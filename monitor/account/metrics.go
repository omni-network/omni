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
		Help:      "The native balance of the account on a specific chain in ether. Alert if low.",
	}, []string{"chain", "role"})

	accountNonce = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "account",
		Name:      "nonce",
		Help:      "The nonce of the account on a specific chain",
	}, []string{"chain", "role"})

	accountBalanceLow = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "account",
		Name:      "balance_low",
		Help:      "Constant gauge indicating whether the account balance is below the minimum threshold (1=true,0=false)",
	}, []string{"chain", "role"})

	tokenBalance = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "account",
		Name:      "token_balance_ether",
		Help:      "The token balance of the account on a specific chain in ether",
	}, []string{"chain", "role", "token"})

	tokenBalanceLow = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "account",
		Name:      "token_balance_low",
		Help:      "Constant gauge indicating whether the token balance is below the minimum threshold (1=true,0=false)",
	}, []string{"chain", "role", "token"})
)
