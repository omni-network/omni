package service

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/uni"
	"github.com/omni-network/omni/scripts/trade/rpc"
	"github.com/omni-network/omni/scripts/trade/users"

	"github.com/google/uuid"
)

func (s Service) RPCHandlers() (string, []rpc.Handler) {
	return "users", []rpc.Handler{
		{
			Endpoint: "create",
			ZeroReq:  func() any { return new(users.RequestCreate) },
			HandleFunc: func(ctx context.Context, anyReq any) (any, error) {
				req, ok := anyReq.(*users.RequestCreate)
				if !ok {
					return nil, errors.New("invalid request type [BUG]")
				}

				return s.Create(ctx, *req)
			},
		},
		{
			Endpoint: "get_by_id",
			ZeroReq:  func() any { return new(requestGetByID) },
			HandleFunc: func(ctx context.Context, anyReq any) (any, error) {
				req, ok := anyReq.(*requestGetByID)
				if !ok {
					return nil, errors.New("invalid request type [BUG]")
				}

				return s.GetByID(ctx, req.ID)
			},
		},
		{
			Endpoint: "get_by_privy_id",
			ZeroReq:  func() any { return new(requestGetByPrivyID) },
			HandleFunc: func(ctx context.Context, anyReq any) (any, error) {
				req, ok := anyReq.(*requestGetByPrivyID)
				if !ok {
					return nil, errors.New("invalid request type [BUG]")
				}

				return s.GetByPrivyID(ctx, req.PrivyID)
			},
		},
		{
			Endpoint: "get_by_address",
			ZeroReq:  func() any { return new(requestGetByAddress) },
			HandleFunc: func(ctx context.Context, anyReq any) (any, error) {
				req, ok := anyReq.(*requestGetByAddress)
				if !ok {
					return nil, errors.New("invalid request type [BUG]")
				}

				return s.GetByAddress(ctx, req.Address)
			},
		},
		{
			Endpoint: "list_all",
			ZeroReq:  func() any { return nil }, // No request needed for listing all users
			HandleFunc: func(ctx context.Context, _ any) (any, error) {
				return s.ListAll(ctx)
			},
		},
	}
}

// requestGetByID wraps the single ID field in a json object for RPC requests.
type requestGetByID struct {
	ID uuid.UUID `json:"id"`
}

// requestGetByPrivyID wraps the single PrivyID field in a json object for RPC requests.
type requestGetByPrivyID struct {
	PrivyID string `json:"privy_id"`
}

// requestGetByAddress wraps the single Address field in a json object for RPC requests.
type requestGetByAddress struct {
	Address uni.Address `json:"address"`
}
