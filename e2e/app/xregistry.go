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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// initXRegistries initializes the XRegistry and PortalRegistr Omni EVM predeploys.
func initXRegistries(ctx context.Context, def Definition) error {
	mngr, err := newRegistryMngr(def)
	if err != nil {
		return errors.Wrap(err, "new registry mngr")
	}

	if err := mngr.setXRegistryPortal(ctx); err != nil {
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
	admin      common.Address
	backend    *ethbackend.Backend
	portals    map[uint64]bindings.PortalRegistryDeployment
	def        Definition
	chainNamer func(uint64) string
}

// newRegistryMngr creates a new registry manager. A registry manager is used to
// initialize the XRegistry and PortalRegistry predeploys.
func newRegistryMngr(def Definition) (registryMngr, error) {
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

	portals, err := makePortalDeps(def)
	if err != nil {
		return registryMngr{}, err
	}

	return registryMngr{
		admin:      admin,
		xreg:       xregistry,
		preg:       portalRegistry,
		backend:    backend,
		portals:    portals,
		def:        def,
		chainNamer: netconf.ChainNamer(def.Testnet.Network),
	}, nil
}

// registerPortals registers each portal with the PortalRegistry.
func (m registryMngr) registerPortals(ctx context.Context) error {
	for chainID := range m.portals {
		if err := m.registerPortal(ctx, chainID); err != nil {
			return errors.Wrap(err, "register portal", "chain", chainID)
		}
	}

	return nil
}

var pregsitryABI = mustGetABI(bindings.PortalRegistryMetaData)

func mustGetABI(m *bind.MetaData) *abi.ABI {
	abi, err := m.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
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

	txOpts, err := m.backend.BindOpts(ctx, m.admin)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}
	txOpts.Value = fee

	input, err := pregsitryABI.Pack("register", p)
	if err != nil {
		return err
	}

	addr := common.HexToAddress(predeploys.PortalRegistry)

	msg := ethereum.CallMsg{
		From:  txOpts.From,
		To:    &addr,
		Data:  input,
		Value: fee,
	}

	gasLimit, err := m.backend.EstimateGasAt(ctx, msg, "pending")
	if err != nil {
		return errors.Wrap(err, "estimate gas")
	}

	gasLimitLatest, err := m.backend.EstimateGasAt(ctx, msg, "latest")
	if err != nil {
		return errors.Wrap(err, "estimate gas")
	}

	current, err := m.preg.List(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "list portals")
	}

	for _, p := range current {
		log.Debug(ctx, "Existing portal", "chain", p.ChainId, "addr", p.Addr.Hex())
	}

	txOpts.GasLimit = gasLimit

	tx, err := m.preg.Register(txOpts, p)
	if err != nil {
		return errors.Wrap(err, "register portal")
	}

	rec, err := m.backend.WaitMined(ctx, tx)
	if err != nil {
		if rec != nil {
			log.Debug(ctx, "Register portal failed", "block", rec.BlockNumber, "estimated_gas", gasLimit, "estimated_gas_latest", gasLimitLatest, "gas_used", rec.GasUsed)
		}

		return errors.Wrap(err, "wait mined")
	}

	log.Debug(ctx, "Register portal succeeded", "block", rec.BlockNumber, "estimated_gas", gasLimit, "estimated_gas_latest", gasLimitLatest, "gas_used", rec.GasUsed)

	return nil
}

// setXRegistryPortal sets the portal address for the XRegistry.
func (m registryMngr) setXRegistryPortal(ctx context.Context) error {
	if len(m.def.Testnet.OmniEVMs) == 0 {
		return errors.New("missing omni evm")
	}

	omniEVM := m.def.Testnet.OmniEVMs[0].Chain

	portal, ok := m.portals[omniEVM.ChainID]
	if !ok {
		return errors.New("missing portal", "chain", omniEVM.ChainID)
	}

	txOpts, err := m.backend.BindOpts(ctx, m.admin)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err := m.xreg.SetPortal(txOpts, portal.Addr)
	if err != nil {
		return errors.Wrap(err, "set portal")
	}

	_, err = m.backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
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

	txOpts, err := m.backend.BindOpts(ctx, m.admin)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err := m.xreg.SetReplica(txOpts, chainID, replica)
	if err != nil {
		return errors.Wrap(err, "set replica")
	}

	_, err = m.backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
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
			ChainId:           chain.ChainID,
			Addr:              info.Address,
			DeployHeight:      info.Height,
			FinalizationStrat: chain.FinalizationStrat.String(),
		}
	}

	for _, c := range tnet.AnvilChains {
		chain := c.Chain

		info, ok := infos[chain.ChainID][types.ContractPortal]
		if !ok {
			return nil, errors.New("missing info", "chain", chain.ChainID)
		}

		deps[chain.ChainID] = bindings.PortalRegistryDeployment{
			ChainId:           chain.ChainID,
			Addr:              info.Address,
			DeployHeight:      info.Height,
			FinalizationStrat: chain.FinalizationStrat.String(),
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
		ChainId:           chain.ChainID,
		Addr:              info.Address,
		DeployHeight:      info.Height,
		FinalizationStrat: chain.FinalizationStrat.String(),
	}

	return deps, nil
}
