package xbridge

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
)

type tokenDescriptors struct {
	symbol string
	name   string
}

type xBridgeDeployment struct {
	token   tokenDescriptors
	wrapped tokenDescriptors
}

var xBridgeDeployments = []xBridgeDeployment{
	{
		token:   tokenDescriptors{symbol: "RLUSD", name: "Ripple USD"},
		wrapped: tokenDescriptors{symbol: "RLUSDe", name: "Bridged RLUSD (Omni)"},
	},
}

func DeployEphemeralXBridge(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	// Restrict deployment to ephemeral networks
	if !network.ID.IsEphemeral() {
		return nil
	}

	// Skip deployment if no Ethereum-labeled chain is available
	_, ok := network.EthereumChain()
	if !ok {
		return nil
	}

	for _, deployment := range xBridgeDeployments {
		if err := deployXBridgeTokens(ctx, network, backends, deployment); err != nil {
			return errors.Wrap(err, "deploy token")
		}

		if err := deployXBridgeLockbox(ctx, network, backends, deployment); err != nil {
			return errors.Wrap(err, "deploy lockbox")
		}

		if err := deployXBridge(ctx, network, backends, deployment); err != nil {
			return errors.Wrap(err, "deploy xbridge")
		}

		log.Debug(ctx, "XBridge deployment finished", "deployment", deployment)
	}

	log.Debug(ctx, "XBridge deployments completed")

	return nil
}

func deployXBridgeTokens(ctx context.Context, network netconf.Network, backends ethbackend.Backends, deployment xBridgeDeployment) error {
	omniEVM, ok := network.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	ethMainnet, ok := network.EthereumChain()
	if !ok {
		return errors.New("no ethereum chain")
	}

	for _, chain := range network.EVMChains() {
		// Token contracts are not deployed on OmniEVM
		if chain.ID == omniEVM.ID {
			continue
		}

		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "backend", "chain", chain.Name)
		}

		if chain.ID == ethMainnet.ID {
			addr, receipt, err := DeployTokenIfNeeded(ctx, network.ID, backend, false, deployment)
			if err != nil {
				return errors.Wrap(err, "deploy", "chain", chain.Name, "name", deployment.token.name, "symbol", deployment.token.symbol, "tx", maybeTxHash(receipt))
			}

			log.Info(ctx, "Token deployed", "chain", chain.Name, "name", deployment.token.name, "symbol", deployment.token.symbol, "address", addr.Hex(), "tx", maybeTxHash(receipt))
		}

		addr, receipt, err := DeployTokenIfNeeded(ctx, network.ID, backend, true, deployment)
		if err != nil {
			return errors.Wrap(err, "deploy", "chain", chain.Name, "name", deployment.wrapped.name, "symbol", deployment.wrapped.symbol, "tx", maybeTxHash(receipt))
		}

		log.Info(ctx, "Wrapper deployed", "chain", chain.Name, "name", deployment.wrapped.name, "symbol", deployment.wrapped.symbol, "address", addr.Hex(), "tx", maybeTxHash(receipt))
	}

	return nil
}

func deployXBridgeLockbox(ctx context.Context, network netconf.Network, backends ethbackend.Backends, deployment xBridgeDeployment) error {
	// Lockbox is only deployed to Ethereum chains where primary tokens are deployed
	ethMainnet, ok := network.EthereumChain()
	if !ok {
		return errors.New("no ethereum chain")
	}

	backend, err := backends.Backend(ethMainnet.ID)
	if err != nil {
		return errors.Wrap(err, "backend", "chain", ethMainnet.Name)
	}

	addr, receipt, err := DeployLockboxIfNeeded(ctx, network.ID, backend, deployment)
	if err != nil {
		return errors.Wrap(err, "deploy", "chain", ethMainnet.Name, "deployment", deployment, "tx", maybeTxHash(receipt))
	}

	log.Info(ctx, "Lockbox deployed", "chain", ethMainnet.Name, "deployment", deployment, "address", addr.Hex(), "tx", maybeTxHash(receipt))

	return nil
}

func deployXBridge(ctx context.Context, network netconf.Network, backends ethbackend.Backends, deployment xBridgeDeployment) error {
	omniEVM, ok := network.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	ethMainnet, ok := network.EthereumChain()
	if !ok {
		return errors.New("no ethereum chain")
	}

	for _, chain := range network.EVMChains() {
		// Bridge contracts are not deployed on OmniEVM
		if chain.ID == omniEVM.ID {
			continue
		}

		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "backend", "chain", chain.Name)
		}

		// Lockbox is only deployed to Ethereum chains where primary tokens are deployed
		lockbox := false
		if chain.ID == ethMainnet.ID {
			lockbox = true
		}

		addr, receipt, err := DeployBridgeIfNeeded(ctx, network.ID, backend, lockbox, deployment)
		if err != nil {
			return errors.Wrap(err, "deploy", "chain", chain.Name, "deployment", deployment, "tx", maybeTxHash(receipt))
		}

		log.Info(ctx, "XBridge deployed", "chain", chain.Name, "deployment", deployment, "address", addr.Hex(), "tx", maybeTxHash(receipt))
	}

	return nil
}
