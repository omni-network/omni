package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/scripts/trade/config"
	"github.com/omni-network/omni/scripts/trade/rpc"
	usersdb "github.com/omni-network/omni/scripts/trade/users/db"
	usersservice "github.com/omni-network/omni/scripts/trade/users/service"

	"github.com/jackc/pgx/v5"
)

// Start starts the trade service and returns:
// - a channel for async errors.
// - a function to stop the service.
// - an error if the service could not be started.
func Start(ctx context.Context, cfg config.Config) (<-chan error, func(context.Context) error, error) {
	log.Info(ctx, "Trade service starting up")

	conn, err := pgx.Connect(ctx, cfg.DBConn)
	if err != nil {
		return nil, nil, errors.Wrap(err, "connect to db")
	}

	userSvc := usersservice.New(usersdb.New(conn))

	asyncAbort := make(chan error, 1)
	stopRPC := rpc.Serve(cfg.RPCListen, asyncAbort,
		userSvc.RPCHandlers,
	)

	log.Info(ctx, "Trade service started", "rpc_listen", cfg.RPCListen)

	return asyncAbort,
		func(ctx context.Context) error {
			if err := stopRPC(ctx); err != nil {
				return errors.Wrap(err, "stop rpc server")
			}

			if err := conn.Close(ctx); err != nil {
				return errors.Wrap(err, "close db connection")
			}

			return nil
		},
		nil
}

// Run runs the trade app until the context is canceled.
func Run(ctx context.Context, cfg config.Config) error {
	async, stopFunc, err := Start(ctx, cfg)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		log.Info(ctx, "Shutdown detected, stopping...")
	case err := <-async:
		return err
	}

	// Use a fresh context for stopping (only allow 5 seconds).
	stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return stopFunc(stopCtx)
}
