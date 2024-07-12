package main

import (
	"context"
	"fmt"

	"github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/client/http"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	xprovider "github.com/omni-network/omni/lib/xchain/provider"
	"github.com/omni-network/omni/monitor/xmonitor/emitcache"

	dbm "github.com/cosmos/cosmos-db"
)

// Lets us test our emit cursor cache - this is codes only for OP Sepolia

func main() {
	ctx := context.Background()

	halo_endpoint := "localhost:9999"       // Needs to be a ssh tunnel to a halo instance
	op_sepolia_endpoint := "localhost:8545" // Needs to be an ssh tunnel to an op sepolia instance

	network := netconf.Network{
		ID:     "staging",
		Chains: []netconf.Chain{{ID: 11155420, Name: "op sepolia", Shards: []xchain.ShardID{xchain.ShardFinalized0, xchain.ShardLatest0}}}}

	tmClient, err := newClient(halo_endpoint)
	if err != nil {
		fmt.Printf("halo endpoint error: err=%v\n", err)
	}

	cProvider := cprovider.NewABCIProvider(tmClient, network.ID, netconf.ChainVersionNamer("staging"))

	ethClients := make(map[uint64]ethclient.Client)
	ethCl, err := ethclient.Dial("op_sepolia", op_sepolia_endpoint)
	if err != nil {
		fmt.Printf("ethclient error: err=%v\n", err)
	}

	ethClients[11155420] = ethCl

	xProvider := xprovider.New(network, ethClients, cProvider)

	db := dbm.NewMemDB()

	cache, err := emitcache.Start(ctx, network, xProvider, db)
	if err != nil {
		fmt.Printf("emit cache start error: err=%v\n", err)
	}

	println(cache)

}

func newClient(tmNodeAddr string) (client.Client, error) {
	c, err := http.New("tcp://"+tmNodeAddr, "/websocket")
	if err != nil {
		return nil, errors.Wrap(err, "new tendermint client")
	}

	return c, nil
}
