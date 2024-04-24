package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"

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
	xreg    *bindings.XRegistry
	preg    *bindings.PortalRegistry
	txOpts  *bind.TransactOpts
	backend *ethbackend.Backend
	portals map[uint64]bindings.PortalRegistryDeployment
	def     Definition
}

// newRegistryMngr creates a new registry manager. A registry manager is used to
// initialize the XRegistry and PortalRegistry predeploys.
func newRegistryMngr(ctx context.Context, def Definition) (registryMngr, error) {
	if !def.Testnet.HasOmniEVM() {
		return registryMngr{}, errors.New("missing omni evm")
	}

	omniEVM := def.Testnet.OmniEVMs[0].Chain

	backend, err := def.Backends().Backend(omniEVM.ID)
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
		xreg:    xregistry,
		preg:    portalRegistry,
		backend: backend,
		txOpts:  txOpts,
		portals: portals,
		def:     def,
	}, nil
}

// registerPortals registers each portal with the PortalRegistry.
func (mngr registryMngr) registerPortals(ctx context.Context) error {
	for chainID := range mngr.portals {
		if err := mngr.registerPortal(ctx, chainID); err != nil {
			return errors.Wrap(err, "register portal", "chain", chainID)
		}
	}

	return nil
}

// registerPortal registers a portal with the PortalRegistry.
func (mngr registryMngr) registerPortal(ctx context.Context, chainID uint64) error {
	p, ok := mngr.portals[chainID]
	if !ok {
		return errors.New("missing portal", "chain", chainID)
	}

	fee, err := mngr.preg.RegistrationFee(&bind.CallOpts{Context: ctx}, p)
	if err != nil {
		return errors.Wrap(err, "registration fee")
	}

	mngr.txOpts.Value = fee
	tx, err := mngr.preg.Register(mngr.txOpts, p)
	mngr.txOpts.Value = nil

	if err != nil {
		return errors.Wrap(err, "register portal")
	}

	receipt, err := mngr.backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	} else if receipt.Status != 1 {
		return errors.New("tx failed", "tx", tx.Hash().Hex())
	}

	return nil
}

// setXRegisryPortal sets the portal address for the XRegistry.
func (mngr registryMngr) setXRegisryPortal(ctx context.Context) error {
	if len(mngr.def.Testnet.OmniEVMs) == 0 {
		return errors.New("missing omni evm")
	}

	omniEVM := mngr.def.Testnet.OmniEVMs[0].Chain

	portal, ok := mngr.portals[omniEVM.ID]
	if !ok {
		return errors.New("missing portal", "chain", omniEVM.ID)
	}

	tx, err := mngr.xreg.SetPortal(mngr.txOpts, portal.Addr)
	if err != nil {
		return errors.Wrap(err, "set portal")
	}

	receipt, err := mngr.backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	} else if receipt.Status != 1 {
		return errors.New("tx failed", "tx", tx.Hash().Hex())
	}

	return nil
}

// setReplicas sets the replica for each chain in the XRegistry.
func (mngr registryMngr) setReplicas(ctx context.Context) error {
	for chainID := range mngr.portals {
		if err := mngr.setReplica(ctx, chainID); err != nil {
			return errors.Wrap(err, "set replica", "chain", chainID)
		}
	}

	return nil
}

// setReplica sets the replica for a chain in the XRegistry.
func (mngr registryMngr) setReplica(ctx context.Context, chainID uint64) error {
	replica, err := mngr.getReplicaAddr(ctx, chainID)
	if err != nil {
		return errors.Wrap(err, "get replica", "chain", chainID)
	}

	tx, err := mngr.xreg.SetReplica(mngr.txOpts, chainID, replica)
	if err != nil {
		return errors.Wrap(err, "set replica")
	}

	receipt, err := mngr.backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	} else if receipt.Status != 1 {
		return errors.New("tx failed", "tx", tx.Hash().Hex())
	}

	return nil
}

// getReplicaAddr gets the xregistry replica address for a chain.
func (mngr registryMngr) getReplicaAddr(ctx context.Context, chainID uint64) (common.Address, error) {
	backend, err := mngr.def.Backends().Backend(chainID)
	if err != nil {
		return common.Address{}, err
	}

	p, ok := mngr.portals[chainID]
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

		info, ok := infos[chain.ID][types.ContractPortal]
		if !ok {
			return nil, errors.New("missing info", "chain", chain.ID)
		}

		deps[chain.ID] = bindings.PortalRegistryDeployment{
			ChainId:           chain.ID,
			Addr:              info.Address,
			DeployHeight:      info.Height,
			FinalizationStrat: chain.FinalizationStrat.String(),
		}
	}

	for _, c := range tnet.AnvilChains {
		chain := c.Chain

		info, ok := infos[chain.ID][types.ContractPortal]
		if !ok {
			return nil, errors.New("missing info", "chain", chain.ID)
		}

		deps[chain.ID] = bindings.PortalRegistryDeployment{
			ChainId:           chain.ID,
			Addr:              info.Address,
			DeployHeight:      info.Height,
			FinalizationStrat: chain.FinalizationStrat.String(),
		}
	}

	if len(tnet.OmniEVMs) == 0 {
		return nil, errors.New("missing omni evm")
	}

	chain := tnet.OmniEVMs[0].Chain

	info, ok := infos[chain.ID][types.ContractPortal]
	if !ok {
		return nil, errors.New("missing info", "chain", chain.ID)
	}

	deps[chain.ID] = bindings.PortalRegistryDeployment{
		ChainId:           chain.ID,
		Addr:              info.Address,
		DeployHeight:      info.Height,
		FinalizationStrat: chain.FinalizationStrat.String(),
	}

	return deps, nil
}
