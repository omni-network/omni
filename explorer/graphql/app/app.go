package app

import (
	"context"
	"net/http"
	"time"

	"github.com/omni-network/omni/explorer/db"
	"github.com/omni-network/omni/explorer/graphql/data"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"golang.org/x/sync/errgroup"
)

func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting Explorer GraphQL server")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	buildinfo.Instrument(ctx)

	// create ent client
	entCl, err := db.NewPostgressClient(cfg.DBUrl)
	if err != nil {
		return errors.Wrap(err, "create db client")
	}
	defer entCl.Close()

	if err := db.CreateSchema(ctx, entCl); err != nil {
		return errors.Wrap(err, "create schema")
	}

	provider := data.Provider{
		EntClient: entCl,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.Handle("/query", GraphQL(provider))

	httpServer := &http.Server{
		Addr:              cfg.ListenAddress,
		ReadHeaderTimeout: 30 * time.Second,
		IdleTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		Handler:           mux,
	}

	var eg errgroup.Group
	eg.Go(func() error {
		defer cancel() // Cancel the app context if serving fails.

		// ListenAndServe always returns an error.
		return errors.Wrap(httpServer.ListenAndServe(), "serve")
	})
	eg.Go(func() error {
		<-ctx.Done()
		log.Info(ctx, "Shutdown detected, stopping server")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := httpServer.Shutdown(shutdownCtx) //nolint:contextcheck // Fresh context is used for shutdown.
		if err != nil {
			return errors.Wrap(err, "server shutdown")
		}

		return nil
	})

	if err := eg.Wait(); errors.Is(err, http.ErrServerClosed) {
		return nil // No error on shutdown.
	} else if err != nil {
		return errors.Wrap(err, "server")
	}

	return nil
}
