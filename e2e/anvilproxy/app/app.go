package app

import (
	"context"
	"net/http"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"golang.org/x/sync/errgroup"
)

func Run(ctx context.Context, cfg Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	portProvider := newPortProvider()
	rootCfg := anvilConfig{
		Config: cfg,
		Port:   portProvider(),
	}
	root, err := startAnvil(ctx, rootCfg)
	if err != nil {
		return errors.Wrap(err, "start root anvil")
	}

	proxy, err := newProxy(root)
	if err != nil {
		return errors.Wrap(err, "new proxy")
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(proxy.Proxy))
	mux.Handle("/fuzzy_enable", http.HandlerFunc(proxy.EnableFuzzyHead))
	mux.Handle("/fuzzy_disable", http.HandlerFunc(proxy.DisableFuzzyHead))

	httpServer := &http.Server{
		Addr:              cfg.ListenAddr,
		ReadHeaderTimeout: 30 * time.Second,
		IdleTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		Handler:           mux,
	}

	log.Info(ctx, "Starting anvilproxy server", "address", cfg.ListenAddr)

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

		if err := httpServer.Shutdown(shutdownCtx); err != nil { //nolint:contextcheck // Explicit new shutdown context.
			return errors.Wrap(err, "server shutdown")
		}

		proxy.getInstance().stop()

		return nil
	})

	if err := eg.Wait(); errors.Is(err, http.ErrServerClosed) {
		return nil // No error on shutdown.
	} else if err != nil {
		return errors.Wrap(err, "run server")
	}

	return nil
}
