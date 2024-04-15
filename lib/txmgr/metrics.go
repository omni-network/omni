package txmgr

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	resendTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "relayer",
		Subsystem: "txmgr",
		Name:      "resend_total",
		Help:      "The total number of transaction resends to a destination chain",
	}, []string{"chain"})

	txEffectiveGasPrice = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "txmgr",
		Name:      "tx_fee_gwei",
		Help:      "Effective gas price for transactions in GWEI",
	}, []string{"chain"})

	txConfirmationLatency = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "txmgr",
		Name:      "tx_confirmation_latency",
		Help:      "Latency between transaction submission and confirmation in seconds",
	}, []string{"chain"})

	txGasUsed = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "relayer",
		Subsystem: "txmgr",
		Name:      "tx_gas_used",
		Help:      "Gas used by transactions",
		Buckets:   prometheus.ExponentialBucketsRange(21_000, 10_000_000, 8),
	}, []string{"chain"})
)
