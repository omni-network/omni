package account

import (
	"context"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

type accountType string

const (
	deployer        accountType = "deployer"
	create3Deployer accountType = "create3-deployer"
	devFireblocks   accountType = "dev-fireblocks"
)

type account struct {
	addr        common.Address
	addressType accountType
}

// Monitor starts monitoring account balances.
func Monitor(ctx context.Context, network netconf.Network) error {
	rpcClientPerChain := make(map[uint64]ethclient.Client)
	for _, chain := range network.Chains {
		if chain.IsOmniConsensus {
			continue // skip non-EVM chains
		}
		c, err := ethclient.Dial(chain.Name, chain.RPCURL)
		if err != nil {
			return errors.Wrap(err, "dial rpc", "chain_id", chain.ID, "rpc_url", chain.RPCURL)
		}
		rpcClientPerChain[chain.ID] = c
	}

	accounts := map[netconf.ID][]account{
		netconf.Testnet: {
			{eoa.MustAddress(netconf.Testnet, eoa.RoleCreate3Deployer), create3Deployer},
			{eoa.MustAddress(netconf.Testnet, eoa.RoleDeployer), deployer},
		},
		netconf.Staging: {
			{eoa.MustAddress(netconf.Staging, eoa.RoleCreate3Deployer), create3Deployer},
			{eoa.MustAddress(netconf.Staging, eoa.RoleDeployer), deployer},
			{common.HexToAddress("0x7a6cF389082dc698285474976d7C75CAdE08ab7e"), devFireblocks}, // fb: dev
		},
	}

	startMonitoring(ctx, network, accounts[network.ID], rpcClientPerChain)

	return nil
}
