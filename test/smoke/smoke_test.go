//nolint:ireturn  // Returning interface below is fine.
package smoke_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/halo/comet"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"
	relayer "github.com/omni-network/omni/relayer/app"
	"github.com/omni-network/omni/scripts/gethdevnet"
	"github.com/omni-network/omni/test/tutil"

	"github.com/cometbft/cometbft/privval"
	rpclocal "github.com/cometbft/cometbft/rpc/client/local"
	rpctest "github.com/cometbft/cometbft/rpc/test"
	"github.com/cometbft/cometbft/types"

	"github.com/stretchr/testify/require"

	_ "embed"
)

const gethDevNetPath = "../../scripts/gethdevnet"

// TestSmoke run a cobbled-together instance of halo and relayer ensuring that blocks are built
// and that the cross chain message flow works.
//
// It has two Engine API variants:
// - geth: uses a real geth devnet (integration test)
// - mock: uses a mock Engine API (unit test)
//
// Each variant includes:
// - Mock XProvider generates periodic xblocks for 1 src chain incl messages to 2 dest chains.
// - Uses real cometBFT with single validator
// - Uses real halo implementations of: app, attestation service, app state, snapshot store
// - Uses relayer code with mocked creator and sender
// - Integrate relayer using cprovider directly connected to app
// - Assert that stream updates are generated for all xblocks.
func TestSmoke(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		ethClFunc func(t *testing.T) engine.API
	}{
		{
			name: "geth",
			ethClFunc: func(t *testing.T) engine.API {
				t.Helper()
				// Use logproxy=true to debug engine API errors
				ethCl, cleanup, err := gethdevnet.StartGenesisGeth(context.Background(), gethDevNetPath, false)
				require.NoError(t, err)
				t.Cleanup(cleanup)

				return ethCl
			},
		},
		{
			name: "mock",
			ethClFunc: func(t *testing.T) engine.API {
				t.Helper()
				mock, err := engine.NewMock()
				require.NoError(t, err)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testSmoke(t, tt.ethClFunc(t))
		})
	}
}

func testSmoke(t *testing.T, ethCl engine.API) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const (
		srcChainBlockPeriod = 1 * time.Millisecond * 100
		srcChainID          = 1
	)

	// Write genesis and priv validator files to temp dir.
	conf := tutil.PrepRPCTestConfig(t)

	// Load the private validator
	privVal := privval.LoadFilePV(conf.PrivValidatorKeyFile(), conf.PrivValidatorStateFile())

	// Start a mock xprovider (this is the source of xblocks)
	xprov := xprovider.NewMock(srcChainBlockPeriod)

	// Create the attestation service.
	path := filepath.Join(t.TempDir(), "state.json")
	err := attest.GenEmptyStateFile(path)
	require.NoError(t, err)
	attSvc, err := attest.LoadAttester(ctx, privVal.Key.PrivKey, path, xprov, []uint64{srcChainID})
	require.NoError(t, err)

	// Create application state
	state, err := comet.LoadOrGenState(t.TempDir(), 1)
	require.NoError(t, err)

	// Create snapshot store.
	snapshots, err := comet.NewSnapshotStore(t.TempDir())
	require.NoError(t, err)

	// Create the comet application.
	app := comet.NewApp(ethCl, attSvc, state, snapshots, 1)

	// Start cometbft
	node := rpctest.StartTendermint(app)
	defer rpctest.StopTendermint(node)

	// Start the relayer, collecting all updates.
	updates := make(chan relayer.StreamUpdate)
	err = relayer.StartRelayer(ctx,
		cprovider.NewABCIProvider(rpclocal.New(node)),
		[]uint64{srcChainID},
		xprov,
		func(update relayer.StreamUpdate) ([]xchain.Submission, error) {
			updates <- update
			return nil, nil
		},
		panicSender{},
	)
	require.NoError(t, err)

	// Subscribe cometbft blocks
	blocksSub, err := node.EventBus().Subscribe(ctx, "", types.EventQueryNewBlock)
	require.NoError(t, err)

	// Wait for 10 stream updates.
	stopAfter := 10
	offsets := make(map[xchain.StreamID]uint64)
	for {
		select {
		case event := <-blocksSub.Out():
			blockEvent, ok := event.Data().(types.EventDataNewBlock)
			require.True(t, ok)
			t.Logf("🔥!! produced block=%d\n", blockEvent.Block.Height)

		case update := <-updates:
			t.Logf("🔥!! stream update: destChain=%v msgs=%v\n", update.DestChainID, len(update.Msgs))

			// Assert the update is good
			require.EqualValues(t, srcChainID, update.SourceChainID)
			require.NotEmpty(t, update.Msgs)
			require.Len(t, update.AggAttestation.Signatures, 1)
			require.EqualValues(t, privVal.Key.PubKey.Bytes(), update.AggAttestation.Signatures[0].ValidatorPubKey)

			// Assert offsets are sequential
			for _, msg := range update.Msgs {
				offsets[update.StreamID]++
				require.EqualValues(t, offsets[update.StreamID], msg.StreamOffset)
			}

			// Stop when we have received enough updates
			stopAfter--
			if stopAfter == 0 {
				cancel()
				return
			}

		case <-time.After(5 * time.Second):
			t.Fatal("timed out waiting for the node to produce a block")
		}
	}
}

var _ relayer.Sender = panicSender{}

type panicSender struct{}

func (panicSender) SendTransaction(context.Context, xchain.Submission) error {
	panic("this should never be called")
}
