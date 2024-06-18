package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	xconnect "github.com/omni-network/omni/lib/xchain/connect"
)

func main() {
	ctx := context.Background()
	netID := netconf.Staging
	const rpc_endpoint string = "https://sepolia.optimism.io"
	const go_back_mins = 5
	const seconds_per_block = 2

	endpoints := xchain.RPCEndpoints{
		fmt.Sprint(evmchain.IDOpSepolia): rpc_endpoint,
	}
	xconn, err := xconnect.New(ctx, netID, endpoints)
	if err != nil {
		fmt.Printf("Error: err=%v\n", err)
	}

	for i := 1; i < go_back_mins+1; i++ {

		fmt.Printf("Go back %d min(s)\n", i)

		block_num, _ := blockHeight(rpc_endpoint, seconds_per_block, i)

		fmt.Println("Block number: ", block_num)

		for _, stream := range xconn.Network.StreamsFrom(evmchain.IDOpSepolia) {

			ref := xchain.EmitRef{
				Height: &block_num,
			}

			xchain.ConfEmitRef(xchain.ConfFinalized)

			_, _, err := xconn.XProvider.GetEmittedCursor(ctx, ref, stream)
			if err != nil {
				fmt.Printf("💀\n")
			} else {
				fmt.Printf("😀\n")
			}

			time.Sleep(1 * time.Second)
		}
	}
}

func blockHeight(endpoint string, secondsPerBlock int, offsetMins int) (uint64, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return 0, err
	}
	defer client.Close()

	block, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}

	var offset_block uint64 = block - (uint64(offsetMins)*60)/uint64(secondsPerBlock)
	return offset_block, nil
}
