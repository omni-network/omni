package admin

import (
	"context"
	"fmt"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// omegaOperators are the operators that are allowed to be validators on the Omega network.
var omegaOperators = []common.Address{
	common.HexToAddress("0xdc5754Fb79163A65753D2CAF23dDA2398cC1f277"), // A41
	common.HexToAddress("0x446924c33A33F413B773d952E7054504788E4c08"), // BlockDaemon
	common.HexToAddress("0xb3E5246B42BC6a12033d5758Dc1752d43807B1dC"), // RHINO
	common.HexToAddress("0x1B6881C66fFd311eE7b0C9b925EB7fd612E1C7B9"), // Kingnodes
	common.HexToAddress("0x641F5938E0d093988d7Cf99509C3152FC7922B88"), // Galaxy
}

// AllowOperators ensures that all operators hard-coded in this package is allowed as validators.
// Note it only adds any of the operators that are missing, it doesn't remove any ever.
func AllowOperators(ctx context.Context, def app.Definition, cfg Config) error {
	network := def.Testnet.Network
	if network.Static().Network != netconf.Omega {
		return errors.New("allow operator only supported on omega", "network", network.Static().Network.String())
	}

	backend, err := def.Backends().Backend(network.Static().OmniExecutionChainID)
	if err != nil {
		return err
	}

	contract, err := bindings.NewStaking(common.HexToAddress(predeploys.Staking), backend)
	if err != nil {
		return errors.Wrap(err, "new staking contract")
	}

	var toAllow []common.Address
	for _, operator := range omegaOperators {
		if ok, err := contract.IsAllowedValidator(&bind.CallOpts{}, operator); err != nil {
			return errors.Wrap(err, "call is allowed validator")
		} else if ok {
			log.Info(ctx, "Operator already allowed as validator, skipping", "operator", operator)
			continue
		}

		toAllow = append(toAllow, operator)
		log.Info(ctx, "Operator not allowed yet, adding to transaction", "operator", operator)
	}

	if len(toAllow) == 0 {
		log.Info(ctx, "All operators already allowed to be validators", "count", len(omegaOperators))
		return nil
	}

	txOpts, err := backend.BindOpts(ctx, eoa.MustAddress(network, eoa.RoleOwner))
	if err != nil {
		return errors.Wrap(err, "bind tx opts")
	}

	if !cfg.Broadcast {
		log.Info(ctx, "Dry-run mode, skipping transaction broadcast", "count", len(toAllow))
		return nil
	}

	tx, err := contract.AllowValidators(txOpts, toAllow)
	if err != nil {
		return errors.Wrap(err, "allow validators")
	}

	if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait minded")
	}

	link := fmt.Sprintf("https://%s.omniscan.network/tx/%s", network, tx.Hash().Hex())
	log.Info(ctx, "🎉 Successfully allowed operators as validators",
		"count", len(toAllow),
		"link", link,
		"network", network,
	)

	return nil
}
