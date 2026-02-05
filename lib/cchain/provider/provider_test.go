package provider_test

import (
	"context"
	"errors"
	"flag"
	"sync"
	"testing"

	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "run integration tests")

func TestUpgradeQueries(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()

	cprov, err := provider.Dial(netconf.Staging)
	require.NoError(t, err)

	_, ok, err := cprov.AppliedPlan(ctx, "not an upgrade")
	require.NoError(t, err)
	require.False(t, ok)
}

func TestExecutionHead(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()

	cprov, err := provider.Dial(netconf.Omega)
	require.NoError(t, err)

	head, err := cprov.ExecutionHead(ctx)
	require.NoError(t, err)
	require.NotZero(t, head.BlockNumber, "block number should not be zero")
	require.NotEqual(t, "0x0000000000000000000000000000000000000000000000000000000000000000", head.BlockHash.Hex(), "block hash should not be zero")

	log.Info(ctx, "ExecutionHead retrieved",
		"block_number", head.BlockNumber,
		"block_hash", head.BlockHash.Hex(),
	)
}

func TestSigningInfos(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()

	cprov, err := provider.Dial(netconf.Omega)
	require.NoError(t, err)

	infos, err := cprov.SDKSigningInfos(ctx)
	require.NoError(t, err)

	for _, info := range infos {
		consCmtAddr, err := info.ConsensusCmtAddr()
		require.NoError(t, err)

		log.Info(ctx, "Validator Signing Info",
			"consensus_addr", consCmtAddr,
			"jailed", info.Jailed(),
			"uptime", info.Uptime,
			"tombstoned", info.Tombstoned,
			"expected_blocks", info.IndexOffset,
			"missed_blocks", info.MissedBlocksCounter,
		)
	}
}
func TestSDKValidator(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()

	cprov, err := provider.Dial(netconf.Omega)
	require.NoError(t, err)

	vals, err := cprov.SDKValidators(ctx)
	require.NoError(t, err)

	for _, val := range vals {
		opAddr, err := val.OperatorEthAddr()
		require.NoError(t, err)

		consEthAddr, err := val.ConsensusEthAddr()
		require.NoError(t, err)

		consCmtAddr, err := val.ConsensusCmtAddr()
		require.NoError(t, err)

		power, err := val.Power()
		require.NoError(t, err)

		log.Info(ctx, "Validator",
			"operator", opAddr,
			"consensus_eth_addr", consEthAddr,
			"consensus_addr", consCmtAddr,
			"power", power,
			"is_jailed", val.IsJailed(),
			"is_bonded", val.IsBonded(),
		)
	}
}

func TestProvider(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel() // Ensure all paths cancel the context to avoid context leak

	const (
		errs  = 2
		total = 10

		chainID    = uint64(999)
		conf       = xchain.ConfFinalized
		fromHeight = uint64(100)
	)

	// Test fetcher returns <errs> errors, then 0,1,2,3,4,... attestations per fetch.
	var expectFetched, expectCount int
	for i := 0; expectFetched < total; i++ {
		expectCount++
		expectFetched += i
	}

	backoff := new(testBackOff)
	fetcher := newTestFetcher(t, errs, expectCount)

	p := provider.NewProviderForT(t, fetcher.Fetch, nil, nil, backoff.BackOff)

	var actual []xchain.Attestation
	var wg sync.WaitGroup
	wg.Add(1)

	chainVer := xchain.ChainVersion{ID: chainID, ConfLevel: conf}
	p.StreamAsync(ctx, chainVer, fromHeight, "test", func(ctx context.Context, approved xchain.Attestation) error {
		actual = append(actual, approved)
		if len(actual) == total {
			cancel()  // Cancel the context to stop further fetch operations
			wg.Done() // Signal that we have fetched enough attestations

			return nil
		}

		return nil
	})

	wg.Wait() // Wait for all fetching to complete before proceeding with assertions

	require.Empty(t, fetcher.Errs()) // All errors returned
	require.Equal(t, expectFetched, fetcher.Fetched())
	require.Equal(t, expectCount, fetcher.Count())
	require.Equal(t, errs+1, backoff.Count()) // 2 errors + 1 empty fetch

	require.Len(t, actual, total)
	for i, attestation := range actual {
		require.Equal(t, chainID, attestation.ChainID)
		require.Equal(t, conf, attestation.ChainVersion.ConfLevel)
		require.Equal(t, fromHeight+uint64(i), attestation.AttestOffset)
	}
}

func newTestFetcher(t *testing.T, errs, maxCount int) *testFetcher {
	t.Helper()
	return &testFetcher{
		t:        t,
		errs:     errs,
		maxCount: maxCount,
	}
}

// testFetcher implements fetchFunc.
// It first returns errs errors.
// Then it returns 0,1,2,3,4,5... attestations up to max.
type testFetcher struct {
	t        *testing.T
	mu       sync.Mutex
	errs     int
	maxCount int
	count    int
	fetched  int
}

func (f *testFetcher) Count() int {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.count
}

func (f *testFetcher) Fetched() int {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.fetched
}

func (f *testFetcher) Errs() int {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.errs
}

func (f *testFetcher) Fetch(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	fromHeight uint64,
	cursor uint64,
) ([]xchain.Attestation, uint64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.errs > 0 {
		f.errs--
		return nil, 0, errors.New("test error")
	} else if f.count >= f.maxCount {
		// Block and wait for test to cancel the context
		// This is required for deterministic fetch assertions since it is done async wrt callbacks.
		<-ctx.Done()
		return nil, cursor, ctx.Err()
	}

	// we use count as consensus block height
	// assert.Equal(f.t, uint64(f.count), cursor, "search start height invalid")
	require.Empty(f.t, cursor, "cursor not disabled")

	toReturn := f.count
	f.count++

	var resp []xchain.Attestation
	for i := 0; i < toReturn; i++ {
		resp = append(resp, xchain.Attestation{
			AttestHeader: xchain.AttestHeader{
				ChainVersion: chainVer,
				AttestOffset: fromHeight + uint64(i),
			},
			BlockHeader: xchain.BlockHeader{
				ChainID: chainVer.ID,
			},
		})
	}

	f.fetched += len(resp)

	return resp, uint64(f.count), nil
}

type testBackOff struct {
	mu      sync.Mutex
	backoff int
}

func (b *testBackOff) Count() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.backoff
}

func (b *testBackOff) BackOff(context.Context) func() {
	return func() {
		b.mu.Lock()
		defer b.mu.Unlock()
		b.backoff++
	}
}
