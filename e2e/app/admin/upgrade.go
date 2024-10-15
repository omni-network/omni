package admin

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// UpgradePortal upgrades the portal contracts on a network.
func UpgradePortal(ctx context.Context, def app.Definition, cfg Config) error {
	return setup(def, cfg).run(ctx, upgradePortal)
}

// UpgradeFeeOracleV1 upgrades the FeeOracleV1 contracts on a network.
func UpgradeFeeOracleV1(ctx context.Context, def app.Definition, cfg Config) error {
	return setup(def, cfg).run(ctx, upgradeFeeOracleV1)
}

// UpgradeGasStation upgrades the GasStation contracts on a network.
func UpgradeGasStation(ctx context.Context, def app.Definition, cfg Config) error {
	s := setup(def, cfg)

	c, err := setupChain(ctx, s, omniEVMName)
	if err != nil {
		return errors.Wrap(err, "setup chain")
	}

	return upgradeGasStation(ctx, s, c)
}

// UpgradeGasPump upgrades the OmniGasPump contracts on a network.
func UpgradeGasPump(ctx context.Context, def app.Definition, cfg Config) error {
	return setup(def, cfg).run(ctx, upgradeGasPump, withExclude(omniEVMName))
}

// UpgradeSlashing upgrades the Slashing predeploy.
func UpgradeSlashing(ctx context.Context, def app.Definition, cfg Config) error {
	s := setup(def, cfg)

	c, err := setupChain(ctx, s, omniEVMName)
	if err != nil {
		return errors.Wrap(err, "setup chain")
	}

	return ugpradeSlashing(ctx, s, c)
}

// UpgradeStaking upgrades the Staking predeploy.
func UpgradeStaking(ctx context.Context, def app.Definition, cfg Config) error {
	s := setup(def, cfg)

	c, err := setupChain(ctx, s, omniEVMName)
	if err != nil {
		return errors.Wrap(err, "setup chain")
	}

	return upgradeStaking(ctx, s, c)
}

// UpgradeBridgeNative upgrades the OmniBridgeNative predeploy.
func UpgradeBridgeNative(ctx context.Context, def app.Definition, cfg Config) error {
	s := setup(def, cfg)

	c, err := setupChain(ctx, s, omniEVMName)
	if err != nil {
		return errors.Wrap(err, "setup chain")
	}

	return upgradeBridgeNative(ctx, s, c)
}

// UpgradeBridgeL1 upgrades the OmniBridgeL1 contract.
func UpgradeBridgeL1(ctx context.Context, def app.Definition, cfg Config) error {
	s := setup(def, cfg)

	l1, ok := s.network.EthereumChain()
	if !ok {
		return errors.New("no l1 eth chain")
	}

	c, err := setupChain(ctx, s, l1.Name)
	if err != nil {
		return errors.Wrap(err, "setup chain")
	}

	return upgradeBridgeL1(ctx, s, c)
}

// UpgradePortalRegistry upgrades the PortalRegistry predeploy.
func UpgradePortalRegistry(ctx context.Context, def app.Definition, cfg Config) error {
	s := setup(def, cfg)

	c, err := setupChain(ctx, s, omniEVMName)
	if err != nil {
		return errors.Wrap(err, "setup chain")
	}

	return upgradePortalRegistry(ctx, s, c)
}

func upgradePortal(ctx context.Context, s shared, c chain) error {
	// TODO: replace if re-initialization is required
	initializer := []byte{}

	calldata, err := adminABI.Pack("upgradePortal", s.upgrader, s.deployer, c.PortalAddress, initializer)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.upgrader, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "Portal upgraded ✅", "chain", c.Name, "addr", c.PortalAddress, "out", out)

	return nil
}

func upgradeFeeOracleV1(ctx context.Context, s shared, c chain) error {
	// FeeOracleV1 contracts were not deployed via Create3
	// The address must be read from the portal

	client, err := ethclient.Dial(c.Name, c.rpc)
	if err != nil {
		return errors.Wrap(err, "dial rpc")
	}

	portal, err := bindings.NewOmniPortal(c.PortalAddress, client)
	if err != nil {
		return errors.Wrap(err, "new portal")
	}

	proxy, err := portal.FeeOracle(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "fee oracle")
	}

	// TODO: replace if re-initialization is required
	initializer := []byte{}

	calldata, err := adminABI.Pack("upgradeFeeOracleV1", s.upgrader, s.deployer, proxy, initializer)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.upgrader, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "FeeOracleV1 upgraded ✅", "chain", c.Name, "addr", proxy, "out", out)

	return nil
}

func upgradeGasStation(ctx context.Context, s shared, c chain) error {
	addrs, err := contracts.GetAddresses(ctx, s.network.ID)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	// TODO: replace if re-initialization is required
	initializer := []byte{}

	calldata, err := adminABI.Pack("upgradeGasStation", s.upgrader, s.deployer, addrs.GasStation, initializer)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.upgrader, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "GasStation upgraded ✅", "chain", c.Name, "addr", addrs.GasStation, "out", out)

	return nil
}

func upgradeGasPump(ctx context.Context, s shared, c chain) error {
	addrs, err := contracts.GetAddresses(ctx, s.network.ID)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	// TODO: replace if re-initialization is required
	initializer := []byte{}

	calldata, err := adminABI.Pack("upgradeGasPump", s.upgrader, s.deployer, addrs.GasPump, initializer)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.upgrader, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "GasPump upgraded ✅", "chain", c.Name, "addr", addrs.GasPump, "out", out)

	return nil
}

func ugpradeSlashing(ctx context.Context, s shared, c chain) error {
	// TODO: replace if re-initialization is required
	initializer := []byte{}

	calldata, err := adminABI.Pack("upgradeSlashing", s.upgrader, s.deployer, initializer)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.upgrader, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "Slashing upgraded ✅", "chain", c.Name, "out", out)

	return nil
}

func upgradeStaking(ctx context.Context, s shared, c chain) error {
	// TODO: replace if re-initialization is required
	initializer := []byte{}

	calldata, err := adminABI.Pack("upgradeStaking", s.upgrader, s.deployer, initializer)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.upgrader, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "Staking upgraded ✅", "chain", c.Name, "out", out)

	return nil
}

func upgradeBridgeNative(ctx context.Context, s shared, c chain) error {
	// TODO: replace if re-initialization is required
	initializer := []byte{}

	calldata, err := adminABI.Pack("upgradeBridgeNative", s.upgrader, s.deployer, initializer)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.upgrader, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "OmniBridgeNative upgraded ✅", "chain", c.Name, "out", out)

	return nil
}

func upgradeBridgeL1(ctx context.Context, s shared, c chain) error {
	addrs, err := contracts.GetAddresses(ctx, s.network.ID)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	// TODO: replace if re-initialization is required
	initializer := []byte{}

	calldata, err := adminABI.Pack("upgradeBridgeL1", s.upgrader, s.deployer, addrs.L1Bridge, initializer)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.upgrader, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "OmniBridgeL1 upgraded ✅", "chain", c.Name, "addr", addrs.L1Bridge, "out", out)

	return nil
}

func upgradePortalRegistry(ctx context.Context, s shared, c chain) error {
	// TODO: replace if re-initialization is required
	initializer := []byte{}

	calldata, err := adminABI.Pack("upgradePortalRegistry", s.upgrader, s.deployer, initializer)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := s.runForge(ctx, c.rpc, calldata, s.upgrader, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "PortalRegistry upgraded ✅", "chain", c.Name, "out", out)

	return nil
}
