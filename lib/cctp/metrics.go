package cctp

import (
	"github.com/omni-network/omni/lib/promutil"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	usdcInFlight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "cctp",
		Name:      "inflight_usdc",
		Help:      "The amount of USDC in flight to / from specific chains",
	}, []string{"src_chain", "dst_chain"})

	msgsInFlight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "cctp",
		Name:      "inflight_messages",
		Help:      "The number of messages in flight to / from specific chains",
	}, []string{"src_chain", "dst_chain"})

	auditHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "cctp",
		Name:      "audit_height",
		Help:      "The height of the last audited block for a chain",
	}, []string{"chain", "recipient"})

	auditCorrectionsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "cctp",
		Name:      "corrected_msgs_total",
		Help:      "The total number of corrected messages by audit",
	}, []string{"chain", "recipient"})

	auditInsertsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "cctp",
		Name:      "missed_msgs_total",
		Help:      "The total number of messages inserted by audit",
	}, []string{"chain", "recipient"})

	oldestMsg = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "cctp",
		Name:      "oldest_msg",
		Help:      "Oldest msg by status (submitted | minted)",
	}, []string{"status"})
)
