package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	xconnect "github.com/omni-network/omni/lib/xchain/connect"
)

// Will go and check from a specific block e.g. the standard Geth full node cut of, of 128 blocks for the ability to get an emit cursor
// This proves or disproves whether we can access and use state at that point.

func main() {
	ctx := context.Background()
	netID := netconf.Staging
	const rpc_endpoint string = "http://localhost:8545" // Ensure you have an SSH tunnel first e.g. ssh -L 8545:localhost:8545 op-sepolia-1
	go_back_blocks := 128

	endpoints := xchain.RPCEndpoints{
		fmt.Sprint(evmchain.IDOpSepolia): rpc_endpoint,
	}
	xconn, err := xconnect.New(ctx, netID, endpoints)
	if err != nil {
		fmt.Printf("xconnect Error: err=%v\n", err)
	}

	head_block, err := blockHeight(rpc_endpoint)
	if err != nil {
		fmt.Printf("Error getting block height: err=%v\n", err)
	}
	fmt.Printf("Current block height %d\n", head_block)
	fmt.Printf("Checking back %d block(s)\n", go_back_blocks)

	for i := go_back_blocks; i > 0; i-- {

		height := head_block - uint64(i)
		fmt.Printf("Checking for cursor at block %d", height)

		for _, stream := range xconn.Network.StreamsFrom(evmchain.IDOpSepolia) {

			ref := xchain.EmitRef{
				Height: &height,
			}

			_, _, err := xconn.XProvider.GetEmittedCursor(ctx, ref, stream)
			if err != nil {
				fmt.Printf("💀\n")
			} else {
				fmt.Printf("😀\n")
			}

		}
	}
}

func blockHeight(endpoint string) (uint64, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return 0, err
	}
	defer client.Close()

	block, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}

	return block, nil
}
