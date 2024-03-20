package avs

import (
	"context"

	"github.com/omni-network/omni/contracts/bindingsv1"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

// Monitor starts monitoring the AVS contract.
func Monitor(ctx context.Context, network netconf.Network) error {
	if network.Name != netconf.Testnet && network.Name != netconf.Mainnet {
		// only monitor in Testned and Mainnet
		return nil
	}

	ch, ok := network.EthereumChain()
	if !ok {
		return errors.New("no avs chain found")
	}

	log.Info(ctx, "Starting AVS monitor")

	client, err := ethclient.Dial(ch.Name, ch.RPCURL)
	if err != nil {
		return errors.Wrap(err, "dialing", "chain", ch.Name)
	}

	avs, err := newAVS(client, ch.AVSContractAddr)
	if err != nil {
		return err
	}

	startMonitoring(ctx, avs)

	return nil
}

// newAVS returns a new AVS contract instance.
func newAVS(client ethclient.Client, address common.Address) (*bindingsv1.OmniAVS, error) {
	avs, err := bindingsv1.NewOmniAVS(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "new AVS")
	}

	return avs, nil
}
