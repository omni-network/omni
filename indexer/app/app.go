package app

import (
	"context"

	"github.com/omni-network/omni/lib/log"
)

type Config struct{}

func Run(ctx context.Context, conf Config) error {
	log.Info(ctx, "Starting Collector")
	log.Info(ctx, "Config: %v", conf)

	return nil
}
