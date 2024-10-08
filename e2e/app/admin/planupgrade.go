package admin

import (
	"context"
	"fmt"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/eoa"
	uluwatu1 "github.com/omni-network/omni/halo/app/upgrades/uluwatu"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

var upgradePlans = map[netconf.ID]bindings.UpgradePlan{
	netconf.Staging: {
		Name:   uluwatu1.UpgradeName,
		Height: 0, // Dynamically calculated for ephemeral networks
	},
}

// PlanUpgrade plans the above configured network upgrade.
func PlanUpgrade(ctx context.Context, def app.Definition, cfg Config) error {
	network := def.Manifest.Network
	plan, ok := upgradePlans[network]
	if !ok {
		return errors.New("no network upgrade configured for network", "network", network)
	}

	backend, err := def.Backends().Backend(network.Static().OmniExecutionChainID)
	if err != nil {
		return err
	}

	// Override height with latest+5 for ephemeral networks.
	if network.IsEphemeral() {
		latest, err := backend.BlockNumber(ctx)
		if err != nil {
			return err
		}

		const delay = 100 // Upgrades must be planned in the future, add a buffer of few minutes
		plan.Height = latest + delay
	}

	contract, err := bindings.NewUpgrade(common.HexToAddress(predeploys.Upgrade), backend)
	if err != nil {
		return errors.Wrap(err, "new staking contract")
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

	link := fmt.Sprintf("https://%s.omniscan.network/tx/%s", network, tx.Hash().Hex())
	log.Info(ctx, "🎉 Successfully planned network upgrade",
		"upgrade", plan.Name,
		"height", plan.Height,
		"network", network,
		"link", link,
	)

	return nil
}
