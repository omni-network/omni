package tokenprice

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	liveTokenPrice = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xfeemngr",
		Name:      "live_token_price",
		Help:      "Live gas price",
	}, []string{"token"})

	bufferedTokenPrice = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xfeemngr",
		Name:      "buffered_token_price",
		Help:      "Buffered gas price",
	}, []string{"token"})

	liveConversionRate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xfeemngr",
		Name:      "live_conversion_rate",
		Help:      "Live conversion rate between tokens",
	}, []string{"from_token", "to_token"})

	bufferedConversionRate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xfeemngr",
		Name:      "buffered_conversion_rate",
		Help:      "Buffered conversion rate between tokens",
	}, []string{"from_token", "to_token"})
)
