package users

import (
	"context"

	"github.com/omni-network/omni/lib/uni"

	"github.com/google/uuid"
)

// Service defines the user service.
type Service interface {
	Create(ctx context.Context, req RequestCreate) (User, error)
	GetByID(ctx context.Context, uuid uuid.UUID) (User, error)
	GetByPrivyID(ctx context.Context, privyID string) (User, error)
	GetByAddress(ctx context.Context, address uni.Address) (User, error)
	ListAll(ctx context.Context) ([]User, error)
}

type User struct {
	ID      uuid.UUID
	PrivyID string
	Address uni.Address
	// TODO: Add created_at, updated_at, etc.
}

type RequestCreate struct {
	ID      uuid.UUID
	PrivyID string
	Address uni.Address
}
