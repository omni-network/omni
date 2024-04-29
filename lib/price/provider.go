package price

import (
	"context"

	"github.com/omni-network/omni/lib/tokens"
)

// Provider is the price provider interface.
type Provider interface {
	// Price returns the price of each provided token in USD.
	Price(ctx context.Context, tokens ...tokens.Token) (map[tokens.Token]float64, error)
}
