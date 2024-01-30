package relayer

import (
	"context"
	"fmt"

	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/gitinfo"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	"github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/client/http"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting relayer")

	commit, timestamp := gitinfo.Get()
	log.Info(ctx, "Version info", "git_commit", commit, "git_timestamp", timestamp)

	network, err := netconf.Load(cfg.NetworkFile)
	if err != nil {
		return err
	}

	rpcClientPerChain, err := initializeRPCClients(network.Chains)
	if err != nil {
		return err
	}

	privateKey, err := ethcrypto.LoadECDSA(cfg.PrivateKey)
	if err != nil {
		return errors.Wrap(err, "failed to load private key")
	}

	sender, err := NewSimpleSender(network.Chains, rpcClientPerChain, *privateKey)
	if err != nil {
		return err
	}

	tmClient, err := newClient(cfg.HaloURL)
	if err != nil {
		return err
	}

	err = StartRelayer(ctx,
		cprovider.NewABCIProvider(tmClient),
		network,
		xprovider.New(network, rpcClientPerChain),
		CreateSubmissions,
		sender)

	if err != nil {
		return err
	}

	<-ctx.Done()
	log.Info(ctx, "Shutdown detected, stopping...")

	return nil
}

func newClient(tmNodeAddr string) (client.Client, error) {
	c, err := http.New(fmt.Sprintf("tcp://%s", tmNodeAddr), "/websocket")
	if err != nil {
		return nil, errors.Wrap(err, "new tendermint client")
	}

	return c, nil
}
