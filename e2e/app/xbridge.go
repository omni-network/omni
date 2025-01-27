package app

import (
	"context"

	"github.com/omni-network/omni/e2e/xbridge"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

func DeployEphemeralXBridge(ctx context.Context, def Definition) error {
	// Restrict deployment to ephemeral networks
	if !def.Testnet.Network.IsEphemeral() {
		return nil
	}

	// Skip deployment if no Ethereum-labeled chain is available
	_, ok := def.Testnet.EthereumChain()
	if !ok {
		return nil
	}

	if err := deployTokens(ctx, def); err != nil {
		return errors.Wrap(err, "deploy token")
	}

	if err := deployLockbox(ctx, def); err != nil {
		return errors.Wrap(err, "deploy lockbox")
	}

	if err := deployBridge(ctx, def); err != nil {
		return errors.Wrap(err, "deploy xbridge")
	}

	log.Debug(ctx, "XBridge deployment complete")

	return nil
}

func deployTokens(ctx context.Context, def Definition) error {
	omniEVM, ok := def.Testnet.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	ethMainnet, ok := def.Testnet.EthereumChain()
	if !ok {
		return errors.New("no ethereum chain")
	}

	for _, chain := range def.Testnet.EVMChains() {
		// Token contracts are not deployed on OmniEVM
		if chain.ChainID == omniEVM.ChainID {
			continue
		}

		backend, err := def.Backends().Backend(chain.ChainID)
		if err != nil {
			return errors.Wrap(err, "backend", "chain", chain.Name)
		}

		if chain.ChainID == ethMainnet.ChainID {
			addr, receipt, err := xbridge.DeployTokenIfNeeded(ctx, def.Testnet.Network, backend, false)
			if err != nil {
				return errors.Wrap(err, "deploy", "type", "RLUSD", "chain", chain.Name, "tx", maybeTxHash(receipt))
			}

			log.Info(ctx, "RLUSD deployed", "chain", chain.Name, "address", addr.Hex(), "tx", maybeTxHash(receipt))
		}

		addr, receipt, err := xbridge.DeployTokenIfNeeded(ctx, def.Testnet.Network, backend, true)
		if err != nil {
			return errors.Wrap(err, "deploy", "type", "RLUSDe", "chain", chain.Name, "tx", maybeTxHash(receipt))
		}

		log.Info(ctx, "RLUSDe deployed", "chain", chain.Name, "address", addr.Hex(), "tx", maybeTxHash(receipt))
	}

	return nil
}

func deployLockbox(ctx context.Context, def Definition) error {
	// Lockbox is only deployed to Ethereum chains where RLUSD is deployed
	ethMainnet, ok := def.Testnet.EthereumChain()
	if !ok {
		return errors.New("no ethereum chain")
	}

	backend, err := def.Backends().Backend(ethMainnet.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend", "chain", ethMainnet.Name)
	}

	addr, receipt, err := xbridge.DeployLockboxIfNeeded(ctx, def.Testnet.Network, backend)
	if err != nil {
		return errors.Wrap(err, "deploy", "chain", ethMainnet.Name, "tx", maybeTxHash(receipt))
	}

	log.Info(ctx, "Lockbox deployed", "chain", ethMainnet.Name, "address", addr.Hex(), "tx", maybeTxHash(receipt))

	return nil
}

func deployBridge(ctx context.Context, def Definition) error {
	omniEVM, ok := def.Testnet.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	ethMainnet, ok := def.Testnet.EthereumChain()
	if !ok {
		return errors.New("no ethereum chain")
	}

	for _, chain := range def.Testnet.EVMChains() {
		// Bridge contracts are not deployed on OmniEVM
		if chain.ChainID == omniEVM.ChainID {
			continue
		}

		backend, err := def.Backends().Backend(chain.ChainID)
		if err != nil {
			return errors.Wrap(err, "backend", "chain", chain.Name)
		}

		// Lockbox is only deployed to Ethereum chains where RLUSD is deployed
		lockbox := false
		if chain.ChainID == ethMainnet.ChainID {
			lockbox = true
		}

		addr, receipt, err := xbridge.DeployBridgeIfNeeded(ctx, def.Testnet.Network, backend, lockbox)
		if err != nil {
			return errors.Wrap(err, "deploy", "chain", chain.Name, "tx", maybeTxHash(receipt))
		}

		log.Info(ctx, "XBridge deployed", "chain", chain.Name, "address", addr.Hex(), "tx", maybeTxHash(receipt))
	}

	return nil
}
