package headerdb

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

// Set exports the set method for testing.
func (db *DB) Set(ctx context.Context, h *types.Header) error {
	return db.set(ctx, h)
}

// DeleteFrom exports the deleteFrom method for testing.
func (db *DB) DeleteFrom(ctx context.Context, height uint64) (int, error) {
	return db.deleteFrom(ctx, height)
}
