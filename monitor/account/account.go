package account

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
)

// Monitor starts monitoring account balances.
func Monitor(ctx context.Context, network netconf.Network) error {

	rpcClientPerChain := make(map[uint64]ethclient.Client)
	for _, chain := range network.Chains {
		if chain.IsOmniConsensus {
			continue // Below monitors only apply to EVM chains.
		}
		c, err := ethclient.Dial(chain.Name, chain.RPCURL)
		if err != nil {
			return errors.Wrap(err, "dial rpc", "chain_id", chain.ID, "rpc_url", chain.RPCURL)
		}
		rpcClientPerChain[chain.ID] = c
	}

	addresses := map[netconf.ID][]common.Address{
		netconf.Devnet: {
			contracts.TestnetCreate3Deployer(),
			contracts.TestnetDeployer(),
		},
		netconf.Staging: {
			contracts.StagingCreate3Deployer(),
			contracts.StagingDeployer(),
			common.HexToAddress("0x7a6cF389082dc698285474976d7C75CAdE08ab7e"), // fb: dev
		},
	}

	startMonitoring(ctx, network, addresses[network.ID], rpcClientPerChain)

	return nil
}
