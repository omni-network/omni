package solve

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/common"
)

func salt(networkID netconf.ID) string {
	return networkID.String() + "::mock::vault"
}

func MockVaultAddress(networkID netconf.ID) common.Address {
	return create3.Address(
		contracts.Create3Factory(networkID),
		salt(networkID),
		eoa.MustAddress(networkID, eoa.RoleDeployer),
	)
}

// maybeDeployMockVault deploys a wstETH mock vault to the MockL2 chain.
func maybeDeployMockVault(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	if network.ID != netconf.Devnet {
		return nil
	}

	chain, ok := network.Chain(evmchain.IDMockL2)
	if !ok {
		return errors.New("chain not found")
	}

	backend, err := backends.Backend(chain.ID)
	if err != nil {
		return errors.Wrap(err, "get backend", "chain", chain.Name)
	}

	deployer := eoa.MustAddress(network.ID, eoa.RoleDeployer)

	txOpts, err := backend.BindOpts(ctx, deployer)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	factory, err := bindings.NewCreate3(contracts.Create3Factory(network.ID), backend)
	if err != nil {
		return errors.Wrap(err, "new create3")
	}

	salt := salt(network.ID)

	addr, deployed, err := isDeployed(ctx, backend, factory, deployer, salt)
	if err != nil {
		return errors.Wrap(err, "is deployed")
	}

	if deployed {
		log.Info(ctx, "MockVault already deployed", "addr", addr, "salt", salt)
		return nil
	}

	abi, err := bindings.MockVaultMetaData.GetAbi()
	if err != nil {
		return errors.Wrap(err, "get abi")
	}

	wstETHOnMockL2, err := Find(evmchain.IDMockL2, tokens.WSTETH.Symbol)
	if err != nil {
		return err
	}

	initCode, err := contracts.PackInitCode(abi, bindings.MockVaultMetaData.Bin, wstETHOnMockL2)
	if err != nil {
		return errors.Wrap(err, "pack init code")
	}

	tx, err := factory.DeployWithRetry(txOpts, create3.HashSalt(salt), initCode) //nolint:contextcheck // Context is txOpts
	if err != nil {
		return errors.Wrap(err, "deploy")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "MockVault deployed", "addr", addr, "salt", salt)

	return nil
}
