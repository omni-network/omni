package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	err := run()
	if err != nil {
		fmt.Printf("netjon failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("usage: go run netjson <network-id>")
	}

	// So that only json is printed
	ctx := log.WithNoopLogger(context.Background())

	netID := netconf.ID(os.Args[1])
	if err := netID.Verify(); err != nil {
		return errors.Wrap(err, "invalid network ID")
	}

	portalReg, err := makePortalRegistry(netID)
	if err != nil {
		return errors.Wrap(err, "portals")
	}

	network, err := netconf.AwaitOnChain(ctx, netID, portalReg, nil)
	if err != nil {
		return errors.Wrap(err, "await on chain")
	}

	out, err := json.MarshalIndent(network, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal")
	}

	fmt.Print(string(out))

	return nil
}

func makePortalRegistry(netID netconf.ID) (*bindings.PortalRegistry, error) {
	omniEVMID := netID.Static().OmniExecutionChainID
	omniEVMMetadata, ok := evmchain.MetadataByID(omniEVMID)
	if !ok {
		return nil, errors.New("no omni evm")
	}

	ethCl, err := ethclient.Dial(omniEVMMetadata.Name, netID.Static().ExecutionRPC())
	if err != nil {
		return nil, err
	}

	resp, err := bindings.NewPortalRegistry(common.HexToAddress(predeploys.PortalRegistry), ethCl)
	if err != nil {
		return nil, errors.Wrap(err, "create portal registry")
	}

	return resp, nil
}
