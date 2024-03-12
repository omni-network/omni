package avs

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

// use hardcoded avs address for now
// TODO: add avs address to network config.
const (
	goerliAVSAddr = "0x848BE3DBcd054c17EbC712E0d29D15C2e638aBCe"
	goerliRPC     = "https://ethereum-goerli-rpc.publicnode.com"
)

// Monitor starts monitoring the AVS contract.
func Monitor(ctx context.Context, network netconf.Network) error {
	if network.Name != netconf.Staging {
		// only monitor in staging for now
		return nil
	}

	log.Info(ctx, "Starting AVS monitor")

	client, err := ethclient.Dial("goerli", goerliRPC)
	if err != nil {
		return errors.Wrap(err, "dialing goerli")
	}

	addr := common.HexToAddress(goerliAVSAddr)
	avs, err := newAVS(client, addr)
	if err != nil {
		return err
	}

	startMonitoring(ctx, avs)

	return nil
}

// newAVS returns a new AVS contract instance.
func newAVS(client ethclient.Client, address common.Address) (*bindings.OmniAVS, error) {
	avs, err := bindings.NewOmniAVS(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "new AVS")
	}

	return avs, nil
}
