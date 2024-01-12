package app

import (
	"context"
	"fmt"

	log "github.com/omni-network/omni/lib/log"
)

type Config struct {
}

func Run(ctx context.Context, conf Config) (err error) {
	log.Info(ctx, "starting collector")

	fmt.Println("hello world")

	fmt.Printf("%v", conf)

	return nil
}
