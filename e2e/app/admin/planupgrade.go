package admin

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/eoa"
	earhart4 "github.com/omni-network/omni/halo/app/upgrades/earhart"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/cchain/queryutil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

var upgradePlans = map[netconf.ID]bindings.UpgradePlan{
	netconf.Omega: {
		Name:   earhart4.UpgradeName,
		Height: 20_128_000, // Tue 26 Aug 13pm UTC
	},
	// netconf.Mainnet: {
	//	Name:   drake3.UpgradeName,
	//	Height: 12_110_000, // Thu 22 May 12pm UTC
	// },
}

// PlanUpgrade plans the above configured network upgrade.
func PlanUpgrade(ctx context.Context, def app.Definition, cfg Config) error {
	network := def.Manifest.Network

	backend, err := def.Backends().Backend(network.Static().OmniExecutionChainID)
	if err != nil {
		return err
	}

	client, err := def.Testnet.BroadcastNode().Client()
	if err != nil {
		return errors.Wrap(err, "broadcast client")
	}
	cprov := provider.NewABCI(client, network)

	next, ok, err := queryutil.NextUpgrade(ctx, cprov)
	if err != nil {
		return err
	} else if !ok {
		return errors.New("network fully upgraded")
	}

	latest, err := backend.BlockNumber(ctx)
	if err != nil {
		return err
	}

	const delay = 100 // Upgrades must be planned in the future, add a buffer of few minutes
	plan := bindings.UpgradePlan{
		Name:   next,
		Height: latest + delay,
	}

	if network.IsProtected() {
		plan, ok = upgradePlans[network]
		if !ok || plan.Height == 0 || plan.Name == "" {
			return errors.New("no network upgrade configured for protected network", "network", network)
		} else if plan.Name != next {
			return errors.New("configured and next upgrade mismatch", "network", network, "conf", plan.Name, "next", next)
		} else if latest+delay > plan.Height {
			return errors.New("configured height not far in future", "network", network, "conf", plan.Height, "latest", latest)
		}
	}

	log.Debug(ctx, "Planning network upgrade",
		"network", network,
		"upgrade", plan.Name,
		"height", plan.Height,
	)

	contract, err := bindings.NewUpgrade(common.HexToAddress(predeploys.Upgrade), backend)
	if err != nil {
		return errors.Wrap(err, "new upgrade contract")
	}

	txOpts, err := backend.BindOpts(ctx, eoa.MustAddress(network, eoa.RoleUpgrader))
	if err != nil {
		return errors.Wrap(err, "bind tx opts")
	}

	if !cfg.Broadcast {
		log.Info(ctx, "Dry-run mode, skipping transaction broadcast")
		return nil
	}

	tx, err := contract.PlanUpgrade(txOpts, plan)
	if err != nil {
		return errors.Wrap(err, "allow validators")
	}

	if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait minded")
	}

	log.Info(ctx, "🎉 Successfully planned network upgrade",
		"upgrade", plan.Name,
		"height", plan.Height,
		"network", network,
		"link", network.Static().OmniScanTXURL(tx.Hash()),
	)

	return nil
}

// CancelPlannedUpgrade cancels the current planned upgrade.
func CancelPlannedUpgrade(ctx context.Context, def app.Definition, cfg Config) error {
	network := def.Manifest.Network

	backend, err := def.Backends().Backend(network.Static().OmniExecutionChainID)
	if err != nil {
		return err
	}

	client, err := def.Testnet.BroadcastNode().Client()
	if err != nil {
		return errors.Wrap(err, "broadcast client")
	}
	cprov := provider.NewABCI(client, network)

	plan, ok, err := cprov.CurrentPlannedPlan(ctx)
	if err != nil {
		return err
	} else if !ok {
		return errors.New("no current planned upgrade")
	}

	log.Debug(ctx, "Canceling planned network upgrade",
		"network", network,
		"upgrade", plan.Name,
		"height", plan.Height,
	)

	contract, err := bindings.NewUpgrade(common.HexToAddress(predeploys.Upgrade), backend)
	if err != nil {
		return errors.Wrap(err, "new upgrade contract")
	}

	txOpts, err := backend.BindOpts(ctx, eoa.MustAddress(network, eoa.RoleUpgrader))
	if err != nil {
		return errors.Wrap(err, "bind tx opts")
	}

	if !cfg.Broadcast {
		log.Info(ctx, "Dry-run mode, skipping transaction broadcast")
		return nil
	}

	tx, err := contract.CancelUpgrade(txOpts)
	if err != nil {
		return errors.Wrap(err, "allow validators")
	}

	if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait minded")
	}

	log.Info(ctx, "🎉 Successfully canceled network upgrade",
		"upgrade", plan.Name,
		"height", plan.Height,
		"network", network,
		"link", network.Static().OmniScanTXURL(tx.Hash()),
	)

	return nil
}
