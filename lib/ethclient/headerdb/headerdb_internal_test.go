package headerdb

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

// Set exports the set method for testing.
func (db *DB) Set(ctx context.Context, h *types.Header) error {
	return db.set(ctx, h)
}
