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
	common.HexToAddress("0xb86DDe58C05dF3C09a8eB7476152793138D113C9"), // Chorusone
	common.HexToAddress("0xcf8EB4Ee3cb3C9c14a4b290bD902dC06d2926ec1"), // P-OPS
	common.HexToAddress("0x27eA917d14d637797FDeb3f9A9824395e7744941"), // DAIC
	common.HexToAddress("0x44Fb4c265E551139e4D3956Aba6fe2DEa27AE4De"), // Finoa
}

// mainnetOperators are the operators that are allowed to be validators on the Mainnet network.
var mainnetOperators = []common.Address{
	common.HexToAddress("0x85bD563C4636a6006BCFB3fd367757676e2Dd269"), // RHINO
	common.HexToAddress("0xc6A510df0F6d2D8CEbD6112FB3132aba4bAc23d1"), // A41
}

// AllowOperators ensures that all operators hard-coded in this package is allowed as validators.
// Note it only adds any of the operators that are missing, it doesn't remove any ever.
func AllowOperators(ctx context.Context, def app.Definition, cfg Config) error {
	network := def.Testnet.Network
	if network.Static().Network != netconf.Omega || network.Static().Network != netconf.Mainnet {
		return errors.New("allow operator only supported on omega or mainnet", "network", network.Static().Network.String())
	}

	backend, err := def.Backends().Backend(network.Static().OmniExecutionChainID)
	if err != nil {
		return err
	}

	contract, err := bindings.NewStaking(common.HexToAddress(predeploys.Staking), backend)
	if err != nil {
		return errors.Wrap(err, "new staking contract")
	}

	var operatorsRange []common.Address
	if network.Static().Network == netconf.Omega {
		operatorsRange = omegaOperators
	} else {
		// Mainnet
		operatorsRange = mainnetOperators
	}
	
	var toAllow []common.Address
	for _, operator := range operatorsRange {
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
		log.Info(ctx, "All operators already allowed to be validators", "count", len(operatorsRange))
		return nil
	}

	txOpts, err := backend.BindOpts(ctx, eoa.MustAddress(network, eoa.RoleManager))
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

	var link string
	if network.Static().Network == netconf.Omega {
		link = fmt.Sprintf("https://%s.omniscan.network/tx/%s", network, tx.Hash().Hex())
	} else {
		// Mainnet
		link = fmt.Sprintf("https://omniscan.network/tx/%s", tx.Hash().Hex())
	}
	
	log.Info(ctx, "🎉 Successfully allowed operators as validators",
		"count", len(toAllow),
		"link", link,
		"network", network,
	)

	return nil
}
