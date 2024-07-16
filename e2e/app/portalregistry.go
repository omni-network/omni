package app

import (
	"context"
	"fmt"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"

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

	admin := eoa.MustAddress(def.Testnet.Network, eoa.RoleAdmin)
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
	chains, err := evmChains(def)
	if err != nil {
		return nil, err
	}

	return toPortalDepls(def, chains)
}

// evmChains returns the EVM chains from the definition.
func evmChains(def Definition) ([]types.EVMChain, error) {
	if len(def.Testnet.OmniEVMs) == 0 {
		return nil, errors.New("missing omni evm")
	}

	var chains []types.EVMChain

	for _, c := range def.Testnet.PublicChains {
		chains = append(chains, c.Chain())
	}

	for _, c := range def.Testnet.AnvilChains {
		chains = append(chains, c.Chain)
	}

	chains = append(chains, def.Testnet.OmniEVMs[0].Chain)

	return chains, nil
}

// toPortalDepls converts EVM chains to portal registry deployments.
func toPortalDepls(def Definition, chains []types.EVMChain) (map[uint64]bindings.PortalRegistryDeployment, error) {
	infos := def.DeployInfos()
	deps := make(map[uint64]bindings.PortalRegistryDeployment)

	for _, chain := range chains {
		info, ok := infos[chain.ChainID][types.ContractPortal]
		if !ok {
			return nil, errors.New("missing info", "chain", chain.ChainID)
		}

		deps[chain.ChainID] = bindings.PortalRegistryDeployment{
			Name:           chain.Name,
			ChainId:        chain.ChainID,
			Addr:           info.Address,
			BlockPeriod:    uint64(chain.BlockPeriod),
			AttestInterval: chain.AttestInterval(def.Testnet.Network),
			DeployHeight:   info.Height,
			Shards:         chain.ShardsUint64(),
		}
	}

	return deps, nil
}

func startAddingMockPortals(ctx context.Context, def Definition) func() error {
	quit := make(chan struct{})

	errChan := make(chan error, 1)
	returnErr := func(err error) {
		if err != nil {
			log.Error(ctx, "Adding mock portal failed", err)
		}
		select {
		case errChan <- err:
		default:
			log.Error(ctx, "Error channel full, dropping error", err)
		}
	}

	go func() {
		mngr, err := newRegistryMngr(ctx, def)
		if err != nil {
			returnErr(err)
			return
		}

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		chainID := uint64(999000)
		for {
			select {
			case <-ctx.Done():
				returnErr(nil)
				return
			case <-quit:
				returnErr(nil)
				return
			case <-ticker.C:
			}

			portal := bindings.PortalRegistryDeployment{
				ChainId:        chainID,
				Addr:           tutil.RandomAddress(),
				DeployHeight:   chainID, // does not matter
				AttestInterval: 60,      // 60 blocks,
				BlockPeriod:    1000,    // 1 second
				Shards:         []uint64{uint64(xchain.ShardFinalized0)},
				Name:           fmt.Sprintf("mock-portal-%d", chainID),
			}

			log.Debug(ctx, "Adding mock portal", "chain", chainID)
			if tx, err := mngr.contract.Register(mngr.txOpts, portal); err != nil {
				returnErr(err)
				return
			} else if _, err := mngr.backend.WaitMined(ctx, tx); err != nil {
				returnErr(err)
				return
			}
			chainID++
		}
	}()

	return func() error {
		close(quit)
		return <-errChan
	}
}
