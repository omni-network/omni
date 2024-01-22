package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/omni-network/omni/explorer/api/docs"
	svr "github.com/omni-network/omni/explorer/api/server"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

func Run(ctx context.Context, conf ExplorerAPIConfig) error {
	log.Info(ctx, "Config: %v", conf)
	ctx, cancelCtx := context.WithCancel(ctx)

	go func() {
		client := svr.NewClient(conf.Port)
		server, err := client.CreateServer(ctx)
		if err != nil {
			log.Error(ctx, "Failed to create rest server %v", err)
		}
		mux := http.NewServeMux()

		// api with prefix /api/v1
		mux.Handle("/api/v1/", http.StripPrefix("/api/v1", server))

		// static files
		mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

		// hosting our documentation
		mux.HandleFunc("/docs", docs.GetHandler())

		httpServer := &http.Server{
			Addr:              fmt.Sprintf(":%v", conf.Port),
			ReadHeaderTimeout: 30 * time.Second,
			IdleTimeout:       30 * time.Second,
			WriteTimeout:      30 * time.Second,
			Handler:           mux,
		}

		log.Info(ctx, "Starting to serve API on port: %v", httpServer.Addr)

		err = httpServer.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Info(ctx, "Closed rest server @%v", httpServer.Addr)
		} else {
			log.Error(ctx, "Error listening for %v:", err, conf.Port)
		}
		cancelCtx()
	}()

	<-ctx.Done()

	return nil
}
