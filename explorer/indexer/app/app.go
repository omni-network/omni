package app

import (
	"context"

	"github.com/omni-network/omni/explorer/db"
	"github.com/omni-network/omni/lib/log"
)

func Run(c context.Context, conf ExplorerIndexerConfig) error {
	log.Info(c, "Config: %v", conf)
	ctx, cancel := context.WithCancel(c)

	go func() {
		var err error

		// create ent client
		client := db.NewClient()

		entClient, err := client.CreateNewEntClientWithSchema(ctx, conf.DBUrl)

		if err != nil {
			log.Error(ctx, "Ent client %v", err, entClient)
			return
		}

		cancel()
	}()

	<-ctx.Done()

	return nil
}
