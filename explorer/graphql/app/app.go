package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/omni-network/omni/explorer/graphql/data"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

func Run(ctx context.Context, conf ExplorerAPIConfig) error {
	log.Info(ctx, "Config: %v", conf)
	ctx, cancel := context.WithCancel(ctx)
	provider := data.Provider{}

	go func() {
		mux := http.NewServeMux()

		mux.HandleFunc("/", home)
		mux.Handle("/query", GraphQL(provider))

		httpServer := &http.Server{
			Addr:              fmt.Sprintf(":%v", conf.Port),
			ReadHeaderTimeout: 30 * time.Second,
			IdleTimeout:       30 * time.Second,
			WriteTimeout:      30 * time.Second,
			Handler:           mux,
		}

		log.Info(ctx, "Starting to serve API on port: %v", httpServer.Addr)

		err := httpServer.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Info(ctx, "Closed http server @%v", httpServer.Addr)
		} else {
			log.Error(ctx, "Error listening for %v:", err, conf.Port)
		}
		cancel()
	}()

	<-ctx.Done()

	return nil
}
