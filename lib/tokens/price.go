package tokens

import (
	"context"
)

// Pricer is the token price provider interface.
type Pricer interface {
	// Price returns the price of each provided token in USD.
	Price(ctx context.Context, tokens ...Token) (map[Token]float64, error)
}
