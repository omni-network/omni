package app

import (
	"context"

	"github.com/omni-network/omni/lib/log"
)

type Config struct{}

func Run(ctx context.Context, conf Config) error {
	log.Info(ctx, "starting collector")
	log.Info(ctx, "config: %v", conf)

	return nil
}
