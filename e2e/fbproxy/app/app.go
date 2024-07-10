package app

import (
	"context"
	"math/big"
	"net/http"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/ethclient"

	"golang.org/x/sync/errgroup"
)

func Run(ctx context.Context, cfg Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	chainID, err := getChainID(ctx, cfg.BaseRPC)
	if err != nil {
		return errors.Wrap(err, "get chain ID")
	}

	fireCl, err := newFireblocks(cfg, chainID)
	if err != nil {
		return errors.Wrap(err, "new fireblocks")
	}

	proxy, err := newProxy(cfg.BaseRPC, NewSendTxMiddleware(fireCl, chainID))
	if err != nil {
		return errors.Wrap(err, "new proxy")
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(proxy.Proxy))

	httpServer := &http.Server{
		Addr:              cfg.ListenAddr,
		ReadHeaderTimeout: 30 * time.Second,
		IdleTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		Handler:           mux,
	}

	log.Info(ctx, "Starting fbproxy server", "address", cfg.ListenAddr)

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

		return nil
	})

	if err := eg.Wait(); errors.Is(err, http.ErrServerClosed) {
		return nil // No error on shutdown.
	} else if err != nil {
		return errors.Wrap(err, "run server")
	}

	return nil
}

func getChainID(ctx context.Context, rpc string) (*big.Int, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpc)
	if err != nil {
		return nil, errors.Wrap(err, "dial base rpc")
	}

	return client.ChainID(ctx)
}
