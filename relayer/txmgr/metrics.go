package txmgr

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//nolint:gochecknoglobals // Promauto metrics are global by nature
var (
	resendTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "relayer",
		Subsystem: "txmgr",
		Name:      "resend_total",
		Help:      "The total number of transaction resends to a destination chain",
	}, []string{"chain"})

	txL1GasFee = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "txmgr",
		Name:      "tx_fee_gwei",
		Help:      "L1 gas fee for transactions in GWEI",
	}, []string{"chain"})

	txFees = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "relayer",
		Subsystem: "txmgr",
		Name:      "tx_fee_gwei_total",
		Help:      "Sum of fees spent for all transactions in GWEI",
	}, []string{"chain"})

	txConfirmationLatency = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "txmgr",
		Name:      "tx_confirmation_latency",
		Help:      "Latency between transaction submission and confirmation in seconds",
	}, []string{"chain"})
)
