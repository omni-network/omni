package app

import (
	"fmt"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/fireblocks"
	"github.com/omni-network/omni/lib/netconf"
)

func newFireblocks(cfg Config, chainID *big.Int) (fireblocks.Client, error) {
	key, err := fireblocks.LoadKey(cfg.FireKeyPath)
	if err != nil {
		return fireblocks.Client{}, errors.Wrap(err, "load fireblocks key")
	}

	opts := []fireblocks.Option{
		fireblocks.WithSignNote(fmt.Sprintf("fireblocks proxy to chainID=%d, network=%s", chainID.Int64(), cfg.Network)),
		fireblocks.WithQueryInterval(5 * time.Second),
	}

	network := netconf.ID(cfg.Network)
	if err := network.Verify(); err != nil {
		return fireblocks.Client{}, errors.Wrap(err, "invalid network", "network", cfg.Network)
	}

	fireCl, err := fireblocks.New(network, cfg.FireAPIKey, key, opts...)
	if err != nil {
		return fireblocks.Client{}, errors.Wrap(err, "new fireblocks")
	}

	return fireCl, nil
}
