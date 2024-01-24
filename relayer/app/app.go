package relayer

import (
	"context"
	"crypto/rand"
	"fmt"

	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/gitinfo"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/client/http"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

func Run(ctx context.Context, cfg Config) error {
	log.Info(ctx, "Starting relayer")

	commit, timestamp := gitinfo.Get()
	log.Info(ctx, "Version info", "git_commit", commit, "git_timestamp", timestamp)

	network, err := netconf.Load(cfg.NetworkFile())
	if err != nil {
		return err
	}

	rpcClientPerChain, err := initializeRPCClients(network.Chains)
	if err != nil {
		return err
	}

	// todo(lazar955): load from cfg
	privateKey, err := ecies.GenerateKey(rand.Reader, ethcrypto.S256(), nil)
	if err != nil {
		return errors.Wrap(err, "generate private key")
	}
	sender, err := NewSimpleSender(network.Chains, rpcClientPerChain, *privateKey.ExportECDSA())
	if err != nil {
		return err
	}

	var noopXProvider XChainClient = &NoopXChainClient{}

	tmClient, err := newClient(cfg.HaloURL)
	if err != nil {
		return err
	}

	err = StartRelayer(ctx,
		cprovider.NewABCIProvider(tmClient),
		network.ChainIDs(),
		noopXProvider,
		CreateSubmissions,
		sender)
	if err != nil {
		return err
	}

	<-ctx.Done()
	log.Info(ctx, "Shutdown detected, stopping...")

	return nil
}

var _ XChainClient = (*NoopXChainClient)(nil)

// NoopXChainClient is a no-op implementation of the XChainClient interface.
type NoopXChainClient struct{}

func (NoopXChainClient) GetSubmittedCursor(context.Context, uint64, uint64) (xchain.StreamCursor, error) {
	return xchain.StreamCursor{}, nil
}

func (NoopXChainClient) GetBlock(_ context.Context, _ uint64, _ uint64) (xchain.Block, bool, error) {
	return xchain.Block{}, false, nil
}

func newClient(tmNodeAddr string) (client.Client, error) {
	c, err := http.New(fmt.Sprintf("tcp://%s", tmNodeAddr), "/websocket")
	if err != nil {
		return nil, errors.Wrap(err, "new tendermint client")
	}

	return c, nil
}
