package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/contracts/staking"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var stakingAdminABI = mustGetABI(bindings.StakingAdminMetaData)

// UpgradeStaking upgrades the staking predeploy.
func UpgradeStaking(ctx context.Context, def Definition) error {
	s, err := setup(def)
	if err != nil {
		return err
	}

	c, err := setupChain(ctx, s, chainOmniEVM)
	if err != nil {
		return err
	}

	return upgradeStaking(ctx, s, c)
}

// ConfigureStaking ensures on chain Staking config matches config in lib/config/staking.
func ConfigureStaking(ctx context.Context, def Definition) error {
	s, err := setup(def)
	if err != nil {
		return err
	}

	c, err := setupChain(ctx, s, chainOmniEVM)
	if err != nil {
		return err
	}

	return configureStaking(ctx, s, c)
}

func upgradeStaking(ctx context.Context, s shared, c chain) error {
	calldata, err := stakingAdminABI.Pack("upgrade", s.admin, s.deployer)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := runForge(ctx, scriptStakingAdmin, c.rpc, calldata, s.admin, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "Staking predeploy upgraded", "out", out)

	return nil
}

func configureStaking(ctx context.Context, s shared, c chain) error {
	client, err := ethclient.Dial(c.rpc)
	if err != nil {
		return errors.Wrap(err, "dial rpc")
	}

	contract, err := bindings.NewStaking(common.HexToAddress(predeploys.Staking), client)
	if err != nil {
		return errors.Wrap(err, "new staking contract")
	}

	cfg, ok := staking.ConfigByNetwork(s.network.ID)
	if !ok {
		log.Info(ctx, "No static staking config, skipping", "network", s.network.ID)
		return nil
	}

	//
	// Check isAllowlistEnabled
	//

	isAllowlistEnabled, err := contract.IsAllowlistEnabled(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "is allowlist enabled")
	}

	if isAllowlistEnabled != cfg.AllowlistEnabled {
		log.Info(ctx, "Updating allowlist enabled", "from", isAllowlistEnabled, "to", cfg.AllowlistEnabled)

		err := setAllowlistEnabled(ctx, s, c, cfg.AllowlistEnabled)
		if err != nil {
			return errors.Wrap(err, "set allowlist enabled")
		}
	}

	//
	// Check allowlist
	//

	allowlist, err := contract.Allowlist(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "allowlist")
	}

	if !cmp(allowlist, cfg.Allowlist) {
		log.Info(ctx, "Updating allowlist", "from", allowlist, "to", cfg.Allowlist)

		err := setAllowlist(ctx, s, c, cfg.Allowlist)
		if err != nil {
			return errors.Wrap(err, "set allowlist")
		}
	}

	return nil
}

func setAllowlistEnabled(ctx context.Context, s shared, c chain, enabled bool) error {
	calldata, err := stakingAdminABI.Pack("setAllowlistEnabled", s.admin, enabled)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := runForge(ctx, scriptStakingAdmin, c.rpc, calldata, s.admin, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "Staking isAllowlistEnabled set", "enabled", enabled, "out", out)

	return nil
}

func setAllowlist(ctx context.Context, s shared, c chain, allowlist []common.Address) error {
	calldata, err := stakingAdminABI.Pack("setAllowlist", s.admin, allowlist)
	if err != nil {
		return errors.Wrap(err, "pack calldata")
	}

	out, err := runForge(ctx, scriptStakingAdmin, c.rpc, calldata, s.admin, s.deployer)
	if err != nil {
		return errors.Wrap(err, "run forge", "out", out)
	}

	log.Info(ctx, "Staking allowlist set", "allowlist", allowlist, "out", out)

	return nil
}

func cmp[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
