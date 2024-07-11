package app

import (
	"fmt"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/fireblocks"
)

func newFireblocks(cfg Config, chainID uint64) (fireblocks.Client, error) {
	key, err := fireblocks.LoadKey(cfg.FireKeyPath)
	if err != nil {
		return fireblocks.Client{}, errors.Wrap(err, "load fireblocks key")
	}

	network := cfg.Network

	opts := []fireblocks.Option{
		fireblocks.WithSignNote(fmt.Sprintf("fireblocks proxy to chainID=%d, network=%s", chainID, network)),
		fireblocks.WithQueryInterval(5 * time.Second),
	}

	if err := network.Verify(); err != nil {
		return fireblocks.Client{}, errors.Wrap(err, "invalid network", "network", network)
	}

	fireCl, err := fireblocks.New(network, cfg.FireAPIKey, key, opts...)
	if err != nil {
		return fireblocks.Client{}, errors.Wrap(err, "new fireblocks")
	}

	return fireCl, nil
}
