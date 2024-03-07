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
	devnetAVSAddr = "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853"
	goerliAVSAddr = "0x848BE3DBcd054c17EbC712E0d29D15C2e638aBCe"
)

// Monitor starts monitoring the AVS contract.
func Monitor(ctx context.Context, network netconf.Network) error {
	log.Info(ctx, "Starting AVS monitor")

	l1Client, err := newL1Client(network)
	if err != nil {
		return err
	}

	addr := common.HexToAddress(devnetAVSAddr)
	// monitor goerli avs in staging, for now
	if network.Name == netconf.Staging {
		addr = common.HexToAddress(goerliAVSAddr)
	}

	avs, err := newAVS(l1Client, addr)
	if err != nil {
		return err
	}

	startMonitoring(ctx, avs)

	return nil
}

// newL1Client returns a new ethclient.Client for the chain marked `IsEthereum` in the network config.
func newL1Client(network netconf.Network) (ethclient.Client, error) {
	for _, chain := range network.Chains {
		if chain.IsEthereum {
			client, err := ethclient.Dial(chain.Name, chain.RPCURL)
			if err != nil {
				return nil, errors.Wrap(err, "dial eth client")
			}

			return client, nil
		}
	}

	return nil, errors.New("no ethereum chain found")
}

// newAVS returns a new AVS contract instance.
func newAVS(client ethclient.Client, address common.Address) (*bindings.OmniAVS, error) {
	avs, err := bindings.NewOmniAVS(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "new AVS")
	}

	return avs, nil
}
