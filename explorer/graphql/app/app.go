package app

import (
	"context"
	"net/http"
	"time"

	"github.com/omni-network/omni/explorer/db"
	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/graphql/data"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
)

func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting Explorer GraphQL server")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	buildinfo.Instrument(ctx)

	// create ent client
	entCl, err := db.NewPostgressClient(cfg.ExplorerDBConn)
	if err != nil {
		return errors.Wrap(err, "create db client")
	}

	defer func(entCl *ent.Client) {
		err := entCl.Close()
		if err != nil {
			log.Error(ctx, "Failed to close ent client", err)
		}
	}(entCl)

	provider := data.Provider{
		EntClient: entCl,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.Handle("/query", GraphQL(provider))
	handler := cors.Default().Handler(mux)

	httpServer := &http.Server{
		Addr:              cfg.ListenAddr,
		ReadHeaderTimeout: 30 * time.Second,
		IdleTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		Handler:           handler,
	}

	log.Info(ctx, "Starting server", "address", httpServer.Addr)

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
	eg.Go(func() error {
		return serveMonitoring(cfg.MonitoringAddr)
	})

	if err := eg.Wait(); errors.Is(err, http.ErrServerClosed) {
		return nil // No error on shutdown.
	} else if err != nil {
		return errors.Wrap(err, "server")
	}

	return nil
}

// serveMonitoring serves the monitoring API.
func serveMonitoring(address string) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:              address,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		Handler:           mux,
	}

	return errors.Wrap(srv.ListenAndServe(), "serve monitoring")
}
