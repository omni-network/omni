package app

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/ethclient"

	"golang.org/x/sync/errgroup"
)

func Run(ctx context.Context, cfg Config) error {
	httpServer, err := newHTTPServer(ctx, cfg)
	if err != nil {
		return errors.Wrap(err, "new http server")
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log.Info(ctx, "Starting fbproxy server", "address", cfg.ListenAddr)

	var eg errgroup.Group
	eg.Go(func() error {
		defer cancel() // Cancel the app context if serving fails.

		// ListenAndServe always returns an error.
		return errors.Wrap(httpServer.ListenAndServe(), "serve")
	})

	eg.Go(func() error {
		<-ctx.Done()
		log.Info(ctx, "Shutdown detected, stopping fbproxy server")

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

func Start(ctx context.Context, cfg Config) (string, error) {
	httpServer, err := newHTTPServer(ctx, cfg)
	if err != nil {
		return "", errors.Wrap(err, "new http server")
	}

	var eg errgroup.Group
	addrCh := make(chan string, 1)

	eg.Go(func() error {
		listener, err := net.Listen("tcp", cfg.ListenAddr)
		if err != nil {
			return errors.Wrap(err, "listen")
		}

		addrCh <- listener.Addr().String()

		// ListenAndServe always returns an error.
		return errors.Wrap(httpServer.Serve(listener), "serve")
	})

	addr := <-addrCh

	log.Info(ctx, "Started fbproxy server", "addr", addr)

	eg.Go(func() error {
		<-ctx.Done()
		log.Info(ctx, "Shutdown detected, stopping fbproxy server", "addr", addr)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil { //nolint:contextcheck // Explicit new shutdown context.
			return errors.Wrap(err, "server shutdown")
		}

		return nil
	})

	return addr, nil
}

func newHTTPServer(ctx context.Context, cfg Config) (*http.Server, error) {
	chainID, err := getChainID(ctx, cfg.BaseRPC)
	if err != nil {
		return nil, errors.Wrap(err, "get chain ID")
	}

	fireCl, err := newFireblocks(cfg, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "new fireblocks")
	}

	proxy, err := newProxy(cfg.BaseRPC, NewSendTxMiddleware(fireCl, chainID))
	if err != nil {
		return nil, errors.Wrap(err, "new proxy")
	}

	return &http.Server{
		ReadHeaderTimeout: 30 * time.Second,
		IdleTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		Handler:           http.HandlerFunc(proxy.Proxy),
	}, nil
}

func getChainID(ctx context.Context, rpc string) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if rpc == "" {
		return 0, errors.New("base RPC URL is required")
	}

	client, err := ethclient.DialContext(ctx, rpc)
	if err != nil {
		return 0, errors.Wrap(err, "dial base rpc")
	}

	id, err := client.ChainID(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get chain ID")
	}

	return id.Uint64(), nil
}
