package app

import (
	"context"
	"fmt"

	log "github.com/omni-network/omni/lib/log"
)

type Config struct {
}

func Run(ctx context.Context, conf Config) (err error) {
	log.Info(ctx, "starting explorer-api")

	fmt.Println("hello world")

	return nil
}
