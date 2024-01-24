package app

import (
	"context"

	"github.com/omni-network/omni/halo/comet"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/lib/xchain/provider"
	relayer "github.com/omni-network/omni/relayer/app"
)

// maybeSetupSimnetRelayer sets up the simnet relayer if the network is simnet.
func maybeSetupSimnetRelayer(ctx context.Context, network netconf.Network, app *comet.App, xprovider xchain.Provider,
) error {
	if network.Name != netconf.Simnet {
		return nil // Skip if not simnet.
	}

	fetchFunc := func(ctx context.Context, chainID uint64, fromHeight uint64,
	) ([]xchain.AggAttestation, error) {
		return app.ApprovedFrom(chainID, fromHeight), nil
	}

	backoffFunc := func(ctx context.Context) (func(), func()) {
		return expbackoff.NewWithReset(ctx, expbackoff.WithFastConfig())
	}

	cprov := cprovider.NewProviderForT(nil, fetchFunc, backoffFunc)

	mockXPriv, ok := xprovider.(*provider.Mock)
	if !ok {
		return errors.New("xchain provider is not a mock")
	}

	return relayer.StartRelayer(ctx, cprov, network.ChainIDs(), mockXPriv, relayer.CreateSubmissions, simnetSender{})
}

var _ relayer.Sender = simnetSender{}

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
