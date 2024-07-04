package account

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

// StartMonitor starts monitoring account balances. It doesn't block and returns immediately.
func StartMonitor(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints) error {
	rpcClientPerChain := make(map[uint64]ethclient.Client)
	for _, chain := range network.EVMChains() {
		rpc, err := endpoints.ByNameOrID(chain.Name, chain.ID)
		if err != nil {
			return err
		}
		c, err := ethclient.Dial(chain.Name, rpc)
		if err != nil {
			return errors.Wrap(err, "dial rpc", "chain", chain.Name, "rpc_url", rpc)
		}
		rpcClientPerChain[chain.ID] = c
	}

	startMonitoring(ctx, network, rpcClientPerChain)

	return nil
}
