package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// initPortalRegistry initializes the PortalRegistry predeploy.
func initPortalRegistry(ctx context.Context, def Definition) error {
	mngr, err := newRegistryMngr(ctx, def)
	if err != nil {
		return errors.Wrap(err, "new registry mngr")
	}

	if err := mngr.registerPortals(ctx); err != nil {
		return errors.Wrap(err, "register portals")
	}

	return nil
}

type registryMngr struct {
	contract   *bindings.PortalRegistry
	txOpts     *bind.TransactOpts
	backend    *ethbackend.Backend
	portals    map[uint64]bindings.PortalRegistryDeployment
	def        Definition
	chainNamer func(uint64) string
}

// newRegistryMngr creates a new portal registry manager, used to
// register portal deployments with the PortalRegistry predeploy.
func newRegistryMngr(ctx context.Context, def Definition) (registryMngr, error) {
	if !def.Testnet.HasOmniEVM() {
		return registryMngr{}, errors.New("missing omni evm")
	}

	omniEVM := def.Testnet.OmniEVMs[0].Chain

	backend, err := def.Backends().Backend(omniEVM.ChainID)
	if err != nil {
		return registryMngr{}, err
	}

	contract, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), backend)
	if err != nil {
		return registryMngr{}, errors.Wrap(err, "new portal registry")
	}

	admin, err := eoa.Admin(def.Testnet.Network)
	if err != nil {
		return registryMngr{}, err
	}

	txOpts, err := backend.BindOpts(ctx, admin)
	if err != nil {
		return registryMngr{}, err
	}

	portals, err := makePortalDeps(def)
	if err != nil {
		return registryMngr{}, err
	}

	return registryMngr{
		contract:   contract,
		backend:    backend,
		txOpts:     txOpts,
		portals:    portals,
		def:        def,
		chainNamer: netconf.ChainNamer(def.Testnet.Network),
	}, nil
}

// registerPortals registers each portal with the PortalRegistry.
func (m registryMngr) registerPortals(ctx context.Context) error {
	log.Info(ctx, "Registering portals")

	var deps []bindings.PortalRegistryDeployment
	for _, dep := range m.portals {
		deps = append(deps, dep)
	}

	tx, err := m.contract.BulkRegister(m.txOpts, deps)
	if err != nil {
		return errors.Wrap(err, "register tx")
	}

	_, err = m.backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}

// makePortalDeps creates a map of portal deployments by chain id.
func makePortalDeps(def Definition) (map[uint64]bindings.PortalRegistryDeployment, error) {
	tnet := def.Testnet
	infos := def.DeployInfos()

	deps := make(map[uint64]bindings.PortalRegistryDeployment)

	for _, c := range tnet.PublicChains {
		chain := c.Chain()

		info, ok := infos[chain.ChainID][types.ContractPortal]
		if !ok {
			return nil, errors.New("missing info", "chain", chain.ChainID)
		}

		deps[chain.ChainID] = bindings.PortalRegistryDeployment{
			ChainId:      chain.ChainID,
			Addr:         info.Address,
			DeployHeight: info.Height,
			Shards:       chain.ShardsUint64(),
		}
	}

	for _, c := range tnet.AnvilChains {
		chain := c.Chain

		info, ok := infos[chain.ChainID][types.ContractPortal]
		if !ok {
			return nil, errors.New("missing info", "chain", chain.ChainID)
		}

		deps[chain.ChainID] = bindings.PortalRegistryDeployment{
			ChainId:      chain.ChainID,
			Addr:         info.Address,
			DeployHeight: info.Height,
			Shards:       chain.ShardsUint64(),
		}
	}

	if len(tnet.OmniEVMs) == 0 {
		return nil, errors.New("missing omni evm")
	}

	chain := tnet.OmniEVMs[0].Chain

	info, ok := infos[chain.ChainID][types.ContractPortal]
	if !ok {
		return nil, errors.New("missing info", "chain", chain.ChainID)
	}

	deps[chain.ChainID] = bindings.PortalRegistryDeployment{
		ChainId:      chain.ChainID,
		Addr:         info.Address,
		DeployHeight: info.Height,
		Shards:       chain.ShardsUint64(),
	}

	return deps, nil
}
