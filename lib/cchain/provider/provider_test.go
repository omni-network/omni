package provider_test

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func TestProvider(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
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
	fetcher := newTestFetcher(errs, expectCount)

	p := provider.NewProviderForT(t, fetcher.Fetch, nil, nil, backoff.BackOff)

	var actual []xchain.Attestation
	var wg sync.WaitGroup
	wg.Add(1)

	chainVer := xchain.ChainVersion{ID: chainID, ConfLevel: conf}
	p.Subscribe(ctx, chainVer, fromHeight, "test", func(ctx context.Context, approved xchain.Attestation) error {
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

func newTestFetcher(errs, maxCount int) *testFetcher {
	return &testFetcher{
		errs:     errs,
		maxCount: maxCount,
	}
}

// testFetcher implements fetchFunc.
// It first returns errs errors.
// Then it returns 0,1,2,3,4,5... attestations up to max.
type testFetcher struct {
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

func (f *testFetcher) Fetch(ctx context.Context, chainVer xchain.ChainVersion, fromHeight uint64,
) ([]xchain.Attestation, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.errs > 0 {
		f.errs--
		return nil, errors.New("test error")
	} else if f.count >= f.maxCount {
		// Block and wait for test to cancel the context
		// This is required for deterministic fetch assertions since it is done async wrt callbacks.
		<-ctx.Done()
		return nil, ctx.Err()
	}

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

	return resp, nil
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
