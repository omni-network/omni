package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
)

type TokenDeployment struct {
	WstETH common.Address
	RETH   common.Address
	METH   common.Address
}

type SymbioticVaultDeployment struct {
	WstETHVault common.Address
	RETHVault   common.Address
	METHVault   common.Address
}

func DeployMockSymbiotic(ctx context.Context, def Definition) (map[uint64]TokenDeployment, map[uint64]SymbioticVaultDeployment, error) {
	deployedTokens, err := deployTokens(ctx, def)
	if err != nil {
		return nil, nil, errors.Wrap(err, "deploy tokens")
	}

	deployedVaults, err := deploySymbioticVaults(ctx, def, deployedTokens)
	if err != nil {
		return nil, nil, errors.Wrap(err, "deploy symbiotic vaults")
	}

	return deployedTokens, deployedVaults, nil
}

func deployTokens(ctx context.Context, def Definition) (map[uint64]TokenDeployment, error) {
	tokens := []struct {
		name   string
		symbol string
	}{
		{"Wrapped Staked ETH", "wstETH"},
		{"Rocket Pool ETH", "rETH"},
		{"Mantle ETH", "mETH"},
	}

	// Track deployed token addresses per chain
	deployments := make(map[uint64]TokenDeployment)

	for _, chain := range def.Testnet.EVMChains() {
		backend, err := def.Backends().Backend(chain.ChainID)
		if err != nil {
			return nil, errors.Wrap(err, "backend", "chain", chain.Name)
		}

		deployer := eoa.MustAddress(def.Testnet.Network, eoa.RoleDeployer)
		auth, err := backend.BindOpts(ctx, deployer)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get TransactionOpts for bindings")
		}

		var deployment TokenDeployment
		// Deploy each token
		for i, token := range tokens {
			addr, _ /* tx */, mockToken, err := bindings.DeployMockToken(auth, backend, token.name, token.symbol)
			// receipt, _ := backend.WaitMined(ctx, tx)
			if err != nil {
				return nil, errors.Wrap(
					err,
					"failed to deploy mock token",
					"chain", chain.Name,
					"token", token.symbol,
				)
			}

			// Store address in deployment struct
			switch i {
			case 0:
				deployment.WstETH = addr
			case 1:
				deployment.RETH = addr
			case 2:
				deployment.METH = addr
			}

			log.Info(ctx, "Token deployed",
				"token", mockToken.Name,
				"symbol", token.symbol,
				"chain", chain.Name,
				"address", addr.Hex(),
				// "txid", receipt.TxHash.Hex()
			)
		}

		deployments[chain.ChainID] = deployment
	}

	return deployments, nil
}

func deploySymbioticVaults(ctx context.Context, def Definition, deployedTokens map[uint64]TokenDeployment) (map[uint64]SymbioticVaultDeployment, error) {
	vaults := []string{
		"wstETH Vault",
		"rETH Vault",
		"mETH Vault",
	}

	// Track deployed vault addresses per chain
	deployments := make(map[uint64]SymbioticVaultDeployment)

	for _, chain := range def.Testnet.EVMChains() {
		tokens := deployedTokens[chain.ChainID]

		backend, err := def.Backends().Backend(chain.ChainID)
		if err != nil {
			return nil, errors.Wrap(err, "backend", "chain", chain.Name)
		}

		deployer := eoa.MustAddress(def.Testnet.Network, eoa.RoleDeployer)
		auth, err := backend.BindOpts(ctx, deployer)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get TransactionOpts for bindings")
		}

		var deployment SymbioticVaultDeployment
		// Deploy each vault
		for i, vault := range vaults {
			var tokenAddr common.Address
			switch i {
			case 0:
				tokenAddr = tokens.WstETH
			case 1:
				tokenAddr = tokens.RETH
			case 2:
				tokenAddr = tokens.METH
			}

			addr, _ /* transactionType */, _ /* mockVault */, err := bindings.DeployMockSymbioticVault(auth, backend, tokenAddr)
			if err != nil {
				return nil, errors.Wrap(
					err,
					"failed to deploy mock symbiotic vault",
					"chain", chain.Name,
					"vault", vault,
				)
			}

			// Store address in deployment struct
			switch i {
			case 0:
				deployment.WstETHVault = addr
			case 1:
				deployment.RETHVault = addr
			case 2:
				deployment.METHVault = addr
			}

			log.Info(ctx, "Symbiotic vault deployed",
				"vault", vault,
				"chain", chain.Name,
				"address", addr.Hex())
		}

		deployments[chain.ChainID] = deployment
	}

	return deployments, nil
}
