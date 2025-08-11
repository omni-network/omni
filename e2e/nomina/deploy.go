package nomina

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
)

func DeployNomina(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	err := deployNomTokenIfNeeded(ctx, network, backends)
	if err != nil {
		return errors.Wrap(err, "deploy nom token")
	}

	// TODO(zodomo): Deploy WNOM after EVM is upgraded

	return nil
}
