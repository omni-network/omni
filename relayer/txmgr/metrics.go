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
)
