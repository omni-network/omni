package appv2

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	statusOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "solver_v2",
		Subsystem: "processor",
		Name:      "status_offset",
		Help:      "Last inbox offset processed by chain and status",
	}, []string{"chain", "target", "status"})

	processedEvents = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "solver_v2",
		Subsystem: "processor",
		Name:      "processed_events_total",
		Help:      "Total number of events processed by chain and status",
	}, []string{"chain", "target", "status"})

	rejectedOrders = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "solver_v2",
		Subsystem: "processor",
		Name:      "rejected_orders_total",
		Help:      "Total number of rejected orders by chain and reason",
	}, []string{"src_chain", "dest_chain", "target", "reason"})

	tokenBalance = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "solver_v2",
		Subsystem: "liquidity",
		Name:      "token_balance",
		Help:      "Token balance of solver",
	}, []string{"chain", "solver_addr", "token_addr", "token_symbol", "is_native"})
)

func sampleBalance(
	chain string,
	token Token,
	solver common.Address,
	amount *big.Int,
) {
	tokenBalance.WithLabelValues(
		chain,                                // chain
		solver.Hex(),                         // solver_addr
		token.address.Hex(),                  // token_addr
		token.symbol,                         // token_symbol
		strconv.FormatBool(token.isNative()), // is_native
	).Set(float64(amount.Uint64()))
}
