package app

import (
	"context"

	"github.com/omni-network/omni/lib/log"
)

func Run(ctx context.Context, cfg Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	log.Info(ctx, "Loadgen config:", cfg.String())
	// @TODO: add load generation logic
	return nil
}
