package genesis

import (
	"context"

	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

// MaybeDownload downloads the genesis files via cprovider the network if they are not already set.
func MaybeDownload(ctx context.Context, network netconf.ID) error {
	if network.IsProtected() {
		return nil // No need to download genesis for protected networks
	}

	if network.Static().ConsensusGenesisJSON != nil && network.Static().ExecutionGenesisJSON != nil {
		return nil // Already set
	}

	rpcServer := network.Static().ConsensusRPC()
	rpcCl, err := rpchttp.New(rpcServer, "/websocket")
	if err != nil {
		return errors.Wrap(err, "create rpc client")
	}
	stubNamer := func(xchain.ChainVersion) string { return "" }
	cprov := cprovider.NewABCIProvider(rpcCl, network, stubNamer)

	execution, consensus, err := cprov.GenesisFiles(ctx)
	if err != nil {
		return errors.Wrap(err, "fetching genesis files")
	} else if len(execution) == 0 {
		return errors.New("empty execution genesis file downloaded", "server", rpcServer)
	}

	log.Info(ctx, "Downloaded genesis files", "execution", len(execution), "consensus", len(consensus), "rpc", rpcServer)

	return netconf.SetEphemeralGenesisBz(network, execution, consensus)
}

// Init downloads the network's genesis files to static, if necessary.
func Init(ctx context.Context, network netconf.ID) error {
	if network != netconf.Staging {
		return nil // Only staging needs initialization
	}

	err := MaybeDownload(ctx, network)
	if err != nil {
		return errors.Wrap(err, "download genesis", "network", network)
	}

	return nil
}
