package fireblocks

import (
	"context"

	f "github.com/omni-network/omni/lib/fireblocks/api"
)

type FireBlocks interface {
	// CreateTransaction creates a new transaction on the FireBlocks API.
	// We use raw signing by default
	CreateTransaction(ctx context.Context, request f.CreateTransactionRequest) (string, error)

	// GetTransactionById retrieves a transaction by its ID.
	GetTransactionById(ctx context.Context, transactionID string) (*f.CreateTransactionRequest, error)
}
