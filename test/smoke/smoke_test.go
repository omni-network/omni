package smoke_test

import (
	"context"
	"math/big"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/halo/consensus"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/scripts/gethdevnet"

	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto"
	rpctest "github.com/cometbft/cometbft/rpc/test"

	fuzz "github.com/google/gofuzz"
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

// TestSmoke starts a genesis geth node and a halo core application ensuring that blocks are built.
// TODO(corver): improve this a lot.
func TestSmoke(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Use logproxy=true to debug engine API errors
	ethCl, cleanup, err := gethdevnet.StartGenesisGeth(ctx, gethDevNetPath, false)
	require.NoError(t, err)
	defer cleanup()

	const attestations = 10

	attSvc := &testAttSvc{
		totals:     attestations,
		pubKeyChan: make(chan crypto.PubKey, 1),
		fuzzer:     fuzz.New().NilChance(0),
	}

	state, err := consensus.LoadOrGenState(t.TempDir(), 1)
	require.NoError(t, err)

	snapshots, err := consensus.NewSnapshotStore(t.TempDir())
	require.NoError(t, err)

	core := consensus.NewCore(ethCl, attSvc, state, snapshots, 1)

	conf := rpctest.GetConfig(true)
	writeFiles(t, conf)

	node := rpctest.StartTendermint(core)
	defer rpctest.StopTendermint(node)

	pubKey, err := node.PrivValidator().GetPubKey()
	require.NoError(t, err)
	attSvc.pubKeyChan <- pubKey

	var lastHeight int64
	for i := 0; i < 3; i++ {
		env, err := node.ConfigureRPC()
		require.NoError(t, err)

		cHeight := env.BlockStore.Height()
		lastHeight = cHeight
		cblock := env.BlockStore.LoadBlock(cHeight)
		var cHash string
		if cblock != nil {
			cHash = cblock.Hash().String()
		}
		t.Logf("🔥!! Consensus Height=%v Hash=%v\n", cHeight, cHash)

		latest, err := ethCl.BlockNumber(ctx)
		require.NoError(t, err)

		eblock, err := ethCl.BlockByNumber(ctx, big.NewInt(int64(latest)))
		require.NoError(t, err)
		t.Logf("🔥!! Execution Height=%v Hash=%v\n", latest, eblock.Hash())

		time.Sleep(1 * time.Second)
	}

	// Assert chain made progress.
	require.NotEmpty(t, lastHeight, "Stuck at height 1")

	// Assert all attestations used and approved.
	require.Empty(t, attSvc.totals, "Not all attestations used")
	require.Lenf(t, core.ApprovedAggregates(), attestations, "Not all attestations approved")
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

type testAttSvc struct {
	mu         sync.Mutex
	totals     int
	pubKeyChan chan crypto.PubKey
	pubKey     crypto.PubKey
	fuzzer     *fuzz.Fuzzer
}

func (s *testAttSvc) decTotal() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.totals == 0 {
		return false
	}

	s.totals--

	return true
}

func (s *testAttSvc) getPubKey() crypto.PubKey { //nolint:ireturn  // Returning interface is fine.
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.pubKey == nil {
		s.pubKey = <-s.pubKeyChan
	}

	return s.pubKey
}

func (s *testAttSvc) newAttestation() xchain.Attestation {
	var att xchain.Attestation
	s.fuzzer.Fuzz(&att)
	copy(att.Signature.ValidatorPubKey[:], s.getPubKey().Bytes())

	return att
}

func (s *testAttSvc) GetAvailable() []xchain.Attestation {
	if !s.decTotal() {
		return nil
	}

	return []xchain.Attestation{s.newAttestation()}
}

func (s *testAttSvc) SetProposed([]xchain.BlockHeader) {}

func (s *testAttSvc) SetCommitted([]xchain.BlockHeader) {}

func (s *testAttSvc) LocalPubKey() [33]byte {
	return [33]byte(s.getPubKey().Bytes())
}

var _ attest.Service = (*testAttSvc)(nil)
