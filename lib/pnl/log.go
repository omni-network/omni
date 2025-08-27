package pnl

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/omni-network/omni/lib/log"
)

type (
	Type     string
	Currency string
)

const (
	Expense Type = "expense"
	Income  Type = "income"

	USD Currency = "USD"
	ETH Currency = "ETH"
	NOM Currency = "NOM"
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
			"delta_gwei", deltaFstr(p.AmountGwei, p.Type),
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

// deltaFstr returns a string repr of a float with an optional negative sign for expenses.
func deltaFstr(f float64, t Type) string {
	if t == Expense {
		return "-" + fstr(f)
	}

	return fstr(f)
}

// mdstr returns a key-ordered string representation of metadata.
func mdstr(m map[string]any) string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var kvs []string
	for _, k := range keys {
		kvs = append(kvs, fmt.Sprintf("%s=%v", k, m[k]))
	}

	return strings.Join(kvs, "|")
}
