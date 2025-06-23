package service

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/uni"
	"github.com/omni-network/omni/scripts/trade/users"
	"github.com/omni-network/omni/scripts/trade/users/db"

	"github.com/google/uuid"
)

var _ users.Service = Service{}

type Service struct {
	db *db.Queries
}

func New(db *db.Queries) Service {
	return Service{
		db: db,
	}
}

func (s Service) Create(ctx context.Context, req users.RequestCreate) (users.User, error) {
	if err := req.Validate(); err != nil {
		return users.User{}, errors.Wrap(err, "validate create request")
	}

	u, err := s.db.Insert(ctx, db.InsertParams{
		ID:      req.ID,
		PrivyID: req.PrivyID,
		Address: req.Address.String(),
	})
	if err != nil {
		return users.User{}, errors.Wrap(err, "insert user")
	}

	return userFromDB(u)
}

func (s Service) GetByID(ctx context.Context, uuid uuid.UUID) (users.User, error) {
	u, err := s.db.GetByID(ctx, uuid)
	if err != nil {
		return users.User{}, errors.Wrap(err, "get user by id")
	}

	return userFromDB(u)
}

func (s Service) GetByPrivyID(ctx context.Context, s2 string) (users.User, error) {
	u, err := s.db.GetByPrivyID(ctx, s2)
	if err != nil {
		return users.User{}, errors.Wrap(err, "get user by privy id")
	}

	return userFromDB(u)
}

func (s Service) GetByAddress(ctx context.Context, address uni.Address) (users.User, error) {
	u, err := s.db.GetByWalletAddress(ctx, address.String())
	if err != nil {
		return users.User{}, errors.Wrap(err, "get user by address")
	}

	return userFromDB(u)
}

func (s Service) ListAll(ctx context.Context) ([]users.User, error) {
	us, err := s.db.ListAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list all users")
	}

	return usersFromDB(us)
}
