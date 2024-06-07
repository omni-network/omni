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

// initXRegistries initializes the XRegistry and PortalRegistr Omni EVM predeploys.
func initXRegistries(ctx context.Context, def Definition) error {
	mngr, err := newRegistryMngr(ctx, def)
	if err != nil {
		return errors.Wrap(err, "new registry mngr")
	}

	if err := mngr.setXRegisryPortal(ctx); err != nil {
		return errors.Wrap(err, "set xregistry portal")
	}

	if err := mngr.setReplicas(ctx); err != nil {
		return errors.Wrap(err, "set replicas")
	}

	if err := mngr.registerPortals(ctx); err != nil {
		return errors.Wrap(err, "register portals")
	}

	return nil
}

type registryMngr struct {
	xreg       *bindings.XRegistry
	preg       *bindings.PortalRegistry
	txOpts     *bind.TransactOpts
	backend    *ethbackend.Backend
	portals    map[uint64]bindings.PortalRegistryDeployment
	def        Definition
	chainNamer func(uint64) string
}

// newRegistryMngr creates a new registry manager. A registry manager is used to
// initialize the XRegistry and PortalRegistry predeploys.
func newRegistryMngr(ctx context.Context, def Definition) (registryMngr, error) {
	if !def.Testnet.HasOmniEVM() {
		return registryMngr{}, errors.New("missing omni evm")
	}

	omniEVM := def.Testnet.OmniEVMs[0].Chain

	backend, err := def.Backends().Backend(omniEVM.ChainID)
	if err != nil {
		return registryMngr{}, err
	}

	xregistry, err := bindings.NewXRegistry(common.HexToAddress(predeploys.XRegistry), backend)
	if err != nil {
		return registryMngr{}, errors.Wrap(err, "new xregistry")
	}

	portalRegistry, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), backend)
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
		xreg:       xregistry,
		preg:       portalRegistry,
		backend:    backend,
		txOpts:     txOpts,
		portals:    portals,
		def:        def,
		chainNamer: netconf.ChainNamer(def.Testnet.Network),
	}, nil
}

// registerPortals registers each portal with the PortalRegistry.
func (m registryMngr) registerPortals(ctx context.Context) error {
	omniEVM := m.def.Testnet.OmniEVMs[0].Chain

	// register omni portal first, required to initialize omni shards
	if err := m.registerPortal(ctx, omniEVM.ChainID); err != nil {
		return errors.Wrap(err, "register omni portal")
	}

	for chainID := range m.portals {
		if chainID == omniEVM.ChainID {
			continue
		}

		if err := m.registerPortal(ctx, chainID); err != nil {
			return errors.Wrap(err, "register portal", chainID)
		}
	}

	return nil
}

// registerPortal registers a portal with the PortalRegistry.
func (m registryMngr) registerPortal(ctx context.Context, chainID uint64) error {
	p, ok := m.portals[chainID]
	if !ok {
		return errors.New("missing portal", "chain", chainID)
	}

	fee, err := m.preg.RegistrationFee(&bind.CallOpts{Context: ctx}, p)
	if err != nil {
		return errors.Wrap(err, "registration fee")
	}

	log.Info(ctx, "Registering portal", "chain", m.chainNamer(chainID))

	m.txOpts.Value = fee
	tx, err := m.preg.Register(m.txOpts, p)
	m.txOpts.Value = nil

	if err != nil {
		return errors.Wrap(err, "register tx")
	}

	receipt, err := m.backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	} else if receipt.Status != 1 {
		return errors.New("tx failed", "tx", tx.Hash().Hex())
	}

	return nil
}

// setXRegisryPortal sets the portal address for the XRegistry.
func (m registryMngr) setXRegisryPortal(ctx context.Context) error {
	if len(m.def.Testnet.OmniEVMs) == 0 {
		return errors.New("missing omni evm")
	}

	omniEVM := m.def.Testnet.OmniEVMs[0].Chain

	portal, ok := m.portals[omniEVM.ChainID]
	if !ok {
		return errors.New("missing portal", "chain", omniEVM.ChainID)
	}

	tx, err := m.xreg.SetPortal(m.txOpts, portal.Addr)
	if err != nil {
		return errors.Wrap(err, "set portal")
	}

	receipt, err := m.backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	} else if receipt.Status != 1 {
		return errors.New("tx failed", "tx", tx.Hash().Hex())
	}

	return nil
}

// setReplicas sets the replica for each chain in the XRegistry.
func (m registryMngr) setReplicas(ctx context.Context) error {
	for chainID := range m.portals {
		if err := m.setReplica(ctx, chainID); err != nil {
			return errors.Wrap(err, "set replica", "chain", chainID)
		}
	}

	return nil
}

// setReplica sets the replica for a chain in the XRegistry.
func (m registryMngr) setReplica(ctx context.Context, chainID uint64) error {
	replica, err := m.getReplicaAddr(ctx, chainID)
	if err != nil {
		return errors.Wrap(err, "get replica", "chain", chainID)
	}

	tx, err := m.xreg.SetReplica(m.txOpts, chainID, replica)
	if err != nil {
		return errors.Wrap(err, "set replica")
	}

	receipt, err := m.backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	} else if receipt.Status != 1 {
		return errors.New("tx failed", "tx", tx.Hash().Hex())
	}

	return nil
}

// getReplicaAddr gets the xregistry replica address for a chain.
func (m registryMngr) getReplicaAddr(ctx context.Context, chainID uint64) (common.Address, error) {
	backend, err := m.def.Backends().Backend(chainID)
	if err != nil {
		return common.Address{}, err
	}

	p, ok := m.portals[chainID]
	if !ok {
		return common.Address{}, errors.New("missing portal", "chain", chainID)
	}

	portal, err := bindings.NewOmniPortal(p.Addr, backend)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "new portal")
	}

	replica, err := portal.Xregistry(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, errors.Wrap(err, "portal xregistry")
	}

	return replica, nil
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
