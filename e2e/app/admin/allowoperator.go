package admin

import (
	"context"

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

var allowedOperators = map[netconf.ID][]common.Address{
	netconf.Omega: {
		common.HexToAddress("0xdc5754Fb79163A65753D2CAF23dDA2398cC1f277"), // A41
		common.HexToAddress("0x446924c33A33F413B773d952E7054504788E4c08"), // BlockDaemon
		common.HexToAddress("0xb3E5246B42BC6a12033d5758Dc1752d43807B1dC"), // RHINO
		common.HexToAddress("0x1B6881C66fFd311eE7b0C9b925EB7fd612E1C7B9"), // Kingnodes
		common.HexToAddress("0x641F5938E0d093988d7Cf99509C3152FC7922B88"), // Galaxy
		common.HexToAddress("0xb86DDe58C05dF3C09a8eB7476152793138D113C9"), // Chorusone
		common.HexToAddress("0xcf8EB4Ee3cb3C9c14a4b290bD902dC06d2926ec1"), // P-OPS
		common.HexToAddress("0x27eA917d14d637797FDeb3f9A9824395e7744941"), // DAIC
		common.HexToAddress("0x44Fb4c265E551139e4D3956Aba6fe2DEa27AE4De"), // Finoa
		common.HexToAddress("0xCE624ce5C5717b63CED36AfB76857183E0a8a6eb"), // validator01
		common.HexToAddress("0x98Eb13371c095905985cddE937018881d4D7f229"), // validator02
		common.HexToAddress("0xe96cF9Ad91cD6dc911817603Dfb3c65d5f532B95"), // validator03
		common.HexToAddress("0xdBd26a685DB4475b6c58ADEC0DE06c6eE387EAa8"), // validator04
	},
	netconf.Mainnet: {
		common.HexToAddress("0xc6A510df0F6d2D8CEbD6112FB3132aba4bAc23d1"), // A41
		common.HexToAddress("0xb278D1dd8a92c0537f3801Acc7242a5e223cA2a9"), // BlockDaemon
		common.HexToAddress("0x85bD563C4636a6006BCFB3fd367757676e2Dd269"), // RHINO
		common.HexToAddress("0x7E634AC3186E9ef6fd36C1207bb87445Fd5CAA7d"), // Kingnodes
		common.HexToAddress("0x39C4Ce600Ef5450a3C06461E4C0bC1E2F70dD4eF"), // Galaxy
		common.HexToAddress("0x3848BEa0cd7D6b521E8CBF74310C8Ab797003537"), // Chorusone
		common.HexToAddress("0xd456Db9730a346211797D42de2272f2406936194"), // P-OPS
		common.HexToAddress("0xd2eC4b16426CaE96a1720d20Bd141dd0DA0b6C9f"), // DAIC
		common.HexToAddress("0x231a7A2872208A3e3EF75Ae6580a16308D60706b"), // Finoa
		common.HexToAddress("0x58D2A4e3880635B7682A1BB7Ed8a43F5ac6cFD3d"), // validator01
		common.HexToAddress("0x19a4Cb685af95A96BEd67C764b6dB137978a5B17"), // validator02
		common.HexToAddress("0xD5f9e687c1EA2b0Da7C06bEbe80ddAb03B33C075"), // validator03
		common.HexToAddress("0x8be1aBb26435fc1AF39Fc88DF9499f626094f9AF"), // validator04
	},
}

// AllowOperators ensures that all operators hard-coded in this package is allowed as validators.
// Note it only adds any of the operators that are missing, it doesn't remove any ever.
func AllowOperators(ctx context.Context, def app.Definition, cfg Config) error {
	network := def.Testnet.Network
	if !network.IsProtected() {
		return errors.New("allow operator only supported on protected networks", "network", def.Testnet.Network)
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
	for _, operator := range allowedOperators[network] {
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
		log.Info(ctx, "All operators already allowed to be validators", "count", len(allowedOperators[network]))
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

	log.Info(ctx, "🎉 Successfully allowed operators as validators",
		"count", len(toAllow),
		"link", network.Static().OmniScanTXURL(tx.Hash()),
		"network", network,
	)

	return nil
}
