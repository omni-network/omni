package app

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"time"

	"github.com/omni-network/omni/explorer/db"
	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

func Run(c context.Context, conf ExplorerIndexerConfig) error {
	log.Info(c, "Config: %v", conf)
	ctx, cancel := context.WithCancel(c)

	go func() {
		var err error

		// create ent client
		client := db.NewClient()

		_, err = client.CreateNewEntClientWithSchema(ctx, conf.DBUrl)

		if err != nil {
			log.Error(ctx, "Failed to open ent client", err)
			return
		}

		mux := http.NewServeMux()

		mux.HandleFunc("/", hello)

		httpServer := &http.Server{
			Addr:              fmt.Sprintf(":%v", conf.Port),
			ReadHeaderTimeout: 30 * time.Second,
			IdleTimeout:       30 * time.Second,
			WriteTimeout:      30 * time.Second,
			Handler:           mux,
		}

		log.Info(ctx, "Starting to serve GraphQL - API on port: %v", httpServer.Addr)

		err = httpServer.ListenAndServe()
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

func hello(w http.ResponseWriter, r *http.Request) {
	hash := make([]byte, 32)
	_, err := rand.Read(hash)
	if err != nil {
		log.Error(r.Context(), "error generating random hash", err)
	}

	var entClient *ent.Client
	// create ent client
	client := db.NewClient()

	entClient, err = client.CreateNewEntClient(r.Context(), DBURL)
	if err != nil {
		log.Error(r.Context(), "failed to create ent client", err)
		panic(err)
	}

	height, err := entClient.Block.Query().Count(r.Context())
	if err != nil {
		log.Error(r.Context(), "failed to get block height", err)
	}

	block, err := entClient.Block.Create().
		SetBlockHash(hash).
		SetSourceChainID(1234).
		SetBlockHeight(uint64(height)).
		Save(r.Context())

	if err != nil {
		log.Error(r.Context(), "failed to create ent client", err)
		panic(err)
	}

	log.Debug(r.Context(), "%v", block)

	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	_, err = w.Write([]byte(fmt.Sprintf("successfully created block %v\n", height)))
	if err != nil {
		log.Error(r.Context(), "graphql home err", err)
	}
}
