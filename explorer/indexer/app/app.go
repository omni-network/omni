package app

import (
	"context"

	"github.com/omni-network/omni/explorer/db"
	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain/provider"
)

func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting Explorer Indexer")

	buildinfo.Instrument(ctx)

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

// provider: dial rpc: dial: dial unix: missing address" chain_id=2 rpc_url="" chain=omni_consensus url="" stacktrace="[app.go:38 cmd.go:34 command.go:983 command.go:1115 command.go:1039 command.go:1032 cmd.go:34 main.go:10 proc.go:271]"
// startXProvider all of our providers and subscribes to the chains in the network config.
func startXProvider(ctx context.Context, network netconf.Network, entCl *ent.Client) error {
	rpcClientPerChain, err := initializeRPCClients(network.EVMChains())
	if err != nil {
		return err
	}

	xprovider := provider.New(network, rpcClientPerChain, nil)
	if xprovider == nil {
		return errors.New("failed to create xchain provider")
	}
	callback := newCallback(entCl)

	for _, chain := range network.EVMChains() {
		fromHeight, err := initChainCursor(ctx, entCl, chain)
		if err != nil {
			return errors.Wrap(err, "initialize chain cursor", "chain_id", chain.ID)
		}

		err = xprovider.StreamAsync(ctx, chain.ID, fromHeight, callback)
		if err != nil {
			return errors.Wrap(err, "subscribe", "chain_id", chain.ID)
		}
	}

	return nil
}

// initializeRPCClients initializes the rpc clients for all evm chains in the network.
func initializeRPCClients(chains []netconf.Chain) (map[uint64]ethclient.Client, error) {
	rpcClientPerChain := make(map[uint64]ethclient.Client)
	for _, chain := range chains {
		client, err := ethclient.Dial(chain.Name, chain.RPCURL)
		if err != nil {
			return nil, errors.Wrap(err, "dial rpc", "chain_id", chain.ID, "rpc_url", chain.RPCURL)
		}
		rpcClientPerChain[chain.ID] = client
	}

	return rpcClientPerChain, nil
}

// initChainCursor return the initial cursor height to start streaming from (inclusive).
// If a cursor exists, it returns the cursor height + 1.
// Else a new cursor is created with chain deploy height.
func initChainCursor(ctx context.Context, entCl *ent.Client, chain netconf.Chain) (uint64, error) {
	cursor, ok, err := getCursor(ctx, entCl.XProviderCursor, chain.ID)
	if err != nil {
		return 0, errors.Wrap(err, "get cursor")
	} else if ok {
		return cursor.Height + 1, nil
	}

	// Store the cursor at deploy height - 1, so first cursor update will be at deploy height.
	deployMinOne := chain.DeployHeight - 1
	if chain.DeployHeight == 0 { // Except for 0, we handle this explicitly.
		deployMinOne = 0
	}

	// cursor doesn't exist yet, create it
	_, err = entCl.XProviderCursor.Create().
		SetChainID(chain.ID).
		SetHeight(deployMinOne).
		Save(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "create cursor")
	}

	return chain.DeployHeight, nil
}
