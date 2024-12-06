package app

import (
	"context"

	"github.com/omni-network/omni/lib/contracts/avs"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// DeployOmniAVSImpl deploys the OmniAVS implementation contract to Ethereum mainnet without upgrading the proxy
// The proxy is managed by a manual multisig external from e2e, and will be upgraded manually.
func DeployOmniAVSImpl(ctx context.Context, def Definition) error {
	ethMainnet, ok := def.Testnet.EthereumChain()
	if !ok {
		return errors.New("no ethereum mainnet chain")
	}

	backend, err := def.Backends().Backend(ethMainnet.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend", "chain", ethMainnet.Name)
	}

	addr, receipt, err := avs.DeployImpl(ctx, def.Testnet.Network, backend)
	if err != nil {
		return errors.Wrap(err, "deploy", "chain", ethMainnet.Name, "tx", maybeTxHash(receipt))
	}

	log.Info(ctx, "OmniAVS implementation deployed", "chain", ethMainnet.Name, "address", addr.Hex(), "tx", maybeTxHash(receipt))

	return nil
}
