//nolint:ireturn  // Returning interface below is fine.
package smoke_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/halo/consensus"
	cprovider "github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"
	relayer "github.com/omni-network/omni/relayer/app"
	"github.com/omni-network/omni/scripts/gethdevnet"

	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/privval"
	rpctest "github.com/cometbft/cometbft/rpc/test"
	"github.com/cometbft/cometbft/types"

	"github.com/stretchr/testify/require"

	_ "embed"
)

const gethDevNetPath = "../../scripts/gethdevnet"

var (
	//go:embed testdata/genesis.json
	genesisJSON []byte

	//go:embed testdata/priv_validator_key.json
	privValKeyJSON []byte

	//go:embed testdata/priv_validator_state.json
	privValStateJSON []byte
)

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

// TestSmoke starts a genesis geth node and a halo core application ensuring that blocks are built.
// TODO(corver): improve this a lot.
func testSmoke(t *testing.T, ethCl engine.API) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const (
		srcChainBlockPeriod = 1 * time.Millisecond * 100
		srcChainID          = 1
	)

	// Write genesis and priv validator files to temp dir.
	conf := rpctest.GetConfig(true)
	writeFiles(t, conf)

	// Load the private validator
	privVal := privval.LoadFilePV(conf.PrivValidatorKeyFile(), conf.PrivValidatorStateFile())

	// Create the attestation service.
	attSvc := attest.NewAttesterForT(t, privVal.Key.PrivKey)

	// Create application state
	state, err := consensus.LoadOrGenState(t.TempDir(), 1)
	require.NoError(t, err)

	// Create snapshot store.
	snapshots, err := consensus.NewSnapshotStore(t.TempDir())
	require.NoError(t, err)

	// Create the core application.
	core := consensus.NewCore(ethCl, attSvc, state, snapshots, 1)

	// Start a mock xprovider (this is the source of xblocks)
	xprov := xprovider.NewMock(srcChainBlockPeriod)

	// Subscribe the attestation service to the mock xprovider.
	err = xprov.Subscribe(ctx, srcChainID, 0, attSvc.Attest)
	require.NoError(t, err)

	// Setup a cprovider that reads directly from app state.
	cprov := cprovider.NewProviderForT(t, adaptFetcher(core), 99, noopBackoff)

	// Start the relayer, collecting all updates.
	updates := make(chan relayer.StreamUpdate)
	relayer.StartRelayer(ctx, cprov, []uint64{srcChainID}, xprov,
		func(update relayer.StreamUpdate) ([]xchain.Submission, error) {
			updates <- update
			return nil, nil
		},
		panicSender{},
	)

	// Start cometbft
	node := rpctest.StartTendermint(core)
	defer rpctest.StopTendermint(node)

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

// adaptFetcher adapts the core application to implement the cprovider.FetchFunc.
func adaptFetcher(core *consensus.Core) cprovider.FetchFunc {
	return func(ctx context.Context, chainID uint64, fromHeight uint64, max uint64) ([]xchain.AggAttestation, error) {
		return core.ApprovedFrom(chainID, fromHeight, max), nil
	}
}

func writeFiles(t *testing.T, conf *config.Config) {
	t.Helper()

	err := os.WriteFile(conf.GenesisFile(), genesisJSON, 0o644)
	require.NoError(t, err)

	err = os.WriteFile(conf.PrivValidatorKeyFile(), privValKeyJSON, 0o644)
	require.NoError(t, err)

	err = os.WriteFile(conf.PrivValidatorStateFile(), privValStateJSON, 0o644)
	require.NoError(t, err)
}

func noopBackoff(context.Context) (func(), func()) {
	return func() {}, func() {}
}

var _ relayer.Sender = panicSender{}

type panicSender struct{}

func (panicSender) SendTransaction(context.Context, xchain.Submission) error {
	panic("this should never be called")
}
