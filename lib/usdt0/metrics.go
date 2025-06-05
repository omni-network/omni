package usdt0

import (
	"github.com/omni-network/omni/lib/promutil"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	usdt0Pending = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "usdt0",
		Name:      "pending_usdt0",
		Help:      "The amount of USDT0 in flight to / from specific chains",
	}, []string{"src_chain", "dst_chain"})

	msgsPending = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "usdt0",
		Name:      "pending_messages",
		Help:      "The number of messages in flight to / from specific chains",
	}, []string{"src_chain", "dst_chain"})

	oldestMsg = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "usdt0",
		Name:      "oldest_msg",
		Help:      "Oldest msg by status",
	}, []string{"status"})
)
