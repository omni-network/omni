package withdraw

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var dustCounter = promauto.NewCounter(prometheus.CounterOpts{
	Namespace:   "halo",
	Subsystem:   "withdraw",
	Name:        "dust_total",
	Help:        "Total withdrawal creation dust in wei (dust is amounts less than 1 gwei)",
	ConstLabels: nil,
})
