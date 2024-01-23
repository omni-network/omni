package relayer

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/gitinfo"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
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
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return errors.Wrap(err, "generate private key")
	}
	sender, err := NewSimpleSender(network.Chains, rpcClientPerChain, *privateKey)
	if err != nil {
		return err
	}

	var noopCProvider cchain.Provider = &NoopCProvider{}
	var noopXProvider XChainClient = &NoopXChainClient{}

	StartRelayer(ctx, noopCProvider, network.ChainIDs(), noopXProvider, CreateSubmissions, sender)

	<-ctx.Done()
	log.Info(ctx, "Shutdown detected, stopping...")

	return nil
}

var _ cchain.Provider = (*NoopCProvider)(nil)
var _ XChainClient = (*NoopXChainClient)(nil)

// NoopCProvider is a no-op implementation of the Provider interface.
type NoopCProvider struct{}

// Subscribe implements the Subscribe method of the Provider interface.
func (*NoopCProvider) Subscribe(ctx context.Context, sourceChainID uint64, sourceHeight uint64,
	_ cchain.ProviderCallback) {
	log.Info(ctx, "No-op: Subscribe called with:", "source_chain_id", sourceChainID, "source_height:", sourceHeight)
}

// NoopXChainClient is a no-op implementation of the XChainClient interface.
type NoopXChainClient struct{}

func (NoopXChainClient) GetBlock(_ context.Context, _ uint64, _ uint64) (xchain.Block, bool, error) {
	return xchain.Block{}, false, nil
}

func (NoopXChainClient) GetSubmittedCursors(_ context.Context, _ uint64) ([]xchain.StreamCursor, error) {
	return nil, nil
}
