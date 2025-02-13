package xbridge

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/xbridge/rlusd"
	"github.com/omni-network/omni/e2e/xbridge/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

// Deploy idempotently deploys all xtokens, bridges and lockboxes for a given network.
func Deploy(ctx context.Context, networkID netconf.ID, backends ethbackend.Backends) error {
	portalReg, err := makePortalRegistry(networkID, backends)
	if err != nil {
		return err
	}

	network, err := netconf.AwaitOnExecutionChain(ctx, networkID, portalReg, chainNames(networkID, backends))
	if err != nil {
		return err
	}

	for _, tkn := range Tokens() {
		bridge, err := BridgeAddr(ctx, networkID, tkn)
		if err != nil {
			return errors.Wrap(err, "bridge addr")
		}

		lockbox, err := LockboxAddr(ctx, networkID, tkn)
		if err != nil {
			return errors.Wrap(err, "lockbox addr")
		}

		canon, err := tkn.Canonical(ctx, networkID)
		if err != nil {
			return errors.Wrap(err, "canonical addr")
		}

		if _, ok := network.Chain(canon.ChainID); !ok {
			log.Debug(ctx, "Skipping xbridge deployment", "token", tkn.Symbol())
			continue
		}

		if err := deployXToken(ctx, network, backends, tkn, bridge, lockbox); err != nil {
			return errors.Wrap(err, "deploy xtoken", "xtoken", tkn.Symbol())
		}

		if err := deployLockbox(ctx, network.ID, backends, tkn); err != nil {
			return errors.Wrap(err, "deploy lockbox", "xtoken", tkn.Symbol())
		}

		// locbock and token must be deployed before the bridge
		if err := deployBridges(ctx, network, backends, tkn); err != nil {
			return errors.Wrap(err, "deploy bridges", "xtoken", tkn.Symbol())
		}
	}

	return nil
}

func deployXToken(
	ctx context.Context,
	network netconf.Network,
	backends ethbackend.Backends,
	tkn types.XToken,
	bridge common.Address, lockbox common.Address) error {
	switch tkn.Symbol() {
	case rlusd.Symbol():
		return rlusd.Deploy(ctx, network, backends, bridge, lockbox)
	default:
		return errors.New("unknown xtoken", "xtoken", tkn.Symbol())
	}
}

func chainNames(networkID netconf.ID, backends ethbackend.Backends) []string {
	var names []string
	namer := netconf.ChainNamer(networkID)

	for chainID := range backends.All() {
		names = append(names, namer(chainID))
	}

	return names
}

func makePortalRegistry(network netconf.ID, backends ethbackend.Backends) (*bindings.PortalRegistry, error) {
	meta := netconf.MetadataByID(network, network.Static().OmniExecutionChainID)
	backend, err := backends.Backend(meta.ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "backend", "chain", meta.Name)
	}

	resp, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), backend)
	if err != nil {
		return nil, errors.Wrap(err, "create portal registry")
	}

	return resp, nil
}
