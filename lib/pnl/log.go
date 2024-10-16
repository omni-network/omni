package pnl

import (
	"context"
	"fmt"

	"github.com/omni-network/omni/lib/log"
)

type (
	Type     string
	Currency string
	Unit     string
)

const (
	Expense Type = "expense"
	Income  Type = "income"

	USD  Currency = "USD"
	ETH  Currency = "ETH"
	OMNI Currency = "OMNI"
)

type LogP struct {
	Type        Type
	AmountGwei  float64
	Currency    Currency
	Category    string
	Subcategory string
	Chain       string
	ID          string
	Metadata    map[string]any
}

// Log logs a pnl event.
func Log(ctx context.Context, ps ...LogP) {
	for _, p := range ps {
		log.Info(ctx,
			"PnL",
			"type", p.Type,
			"amt_gwei", fstr(p.AmountGwei),
			"currency", p.Currency,
			"category", p.Category,
			"subcategory", p.Subcategory,
			"chain", p.Chain,
			"id", p.ID,
			"metadata", mdstr(p.Metadata),
		)
	}
}

// fstr returns a string repr of a float.
func fstr(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

// mdstr returns a string repr of metadat.
func mdstr(m map[string]any) string {
	var s string

	for k, v := range m {
		s += fmt.Sprintf("%s:%v|", k, v)
	}

	s = s[:len(s)-1] // remove last pipe

	return s
}
