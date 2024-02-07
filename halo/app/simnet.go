package app

import (
	"context"

	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	relayer "github.com/omni-network/omni/relayer/app"

	"github.com/cometbft/cometbft/node"
	rpclocal "github.com/cometbft/cometbft/rpc/client/local"
)

// maybeSetupSimnetRelayer sets up the simnet relayer if the network is simnet.
func maybeSetupSimnetRelayer(ctx context.Context, network netconf.Network, cmtNode *node.Node,
	xprovider xchain.Provider,
) error {
	if network.Name != netconf.Simnet {
		return nil
	}

	cprov := cprovider.NewABCIProvider(rpclocal.New(cmtNode))

	return relayer.StartRelayer(ctx, cprov, network, xprovider, relayer.CreateSubmissions,
		simnetSender{}.SendTransaction)
}

var _ relayer.SendFunc = simnetSender{}.SendTransaction

// simnetSender implements relayer.Sender for simnet by just logging.
type simnetSender struct{}

func (simnetSender) SendTransaction(ctx context.Context, submission xchain.Submission) error {
	var destChainID, offset uint64
	for _, msg := range submission.Msgs {
		destChainID = msg.DestChainID
		offset = msg.StreamOffset

		break
	}

	log.Info(ctx, "Simnet relayer sending transaction",
		"dest_chain", destChainID,
		"msgs", len(submission.Msgs),
		"first_offset", offset,
	)

	return nil
}
