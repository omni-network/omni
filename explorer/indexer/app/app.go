package app

import (
	"context"

	"github.com/omni-network/omni/explorer/db"
	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/gitinfo"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain/provider"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting Explorer Indexer")

	commit, timestamp := gitinfo.Get()
	log.Info(ctx, "Version info", "git_commit", commit, "git_timestamp", timestamp)

	network, err := netconf.Load(cfg.NetworkFile)
	if err != nil {
		return errors.Wrap(err, "load network config")
	}

	entCl, err := db.NewPostgressClient(cfg.DBUrl)
	if err != nil {
		return errors.Wrap(err, "new db client")
	}

	if err := db.CreateSchema(ctx, entCl); err != nil {
		return errors.Wrap(err, "create schema")
	}

	err = startXProvider(ctx, network, entCl)
	if err != nil {
		return errors.Wrap(err, "provider")
	}

	<-ctx.Done()

	log.Info(ctx, "Shutdown detected, stopping indexer")

	return nil
}

// startXProvider all of our providers and subscribes to the chains in the network config.
func startXProvider(ctx context.Context, network netconf.Network, entCl *ent.Client) error {
	rpcClientPerChain, err := initializeRPCClients(network.Chains)
	if err != nil {
		return err
	}

	xprovider := provider.New(network, rpcClientPerChain)
	callback := newCallback(entCl)

	for _, chain := range network.Chains {
		err = xprovider.Subscribe(ctx, chain.ID, chain.DeployHeight, callback)
		if err != nil {
			return errors.Wrap(err, "subscribe", "chain_id", chain.ID)
		}
	}

	return nil
}

func initializeRPCClients(chains []netconf.Chain) (map[uint64]*ethclient.Client, error) {
	rpcClientPerChain := make(map[uint64]*ethclient.Client)
	for _, chain := range chains {
		client, err := ethclient.Dial(chain.RPCURL)
		if err != nil {
			return nil, errors.Wrap(err, "dial rpc", "chain_id", chain.ID, "rpc_url", chain.RPCURL)
		}
		rpcClientPerChain[chain.ID] = client
	}

	return rpcClientPerChain, nil
}
