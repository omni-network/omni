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

	const (
		errs  = 2
		total = 10

		chainID    = uint64(999)
		fromHeight = uint64(100)
	)

	backoff := new(testBackOff)
	fetcher := &testFetcher{errs: errs}

	p := provider.NewProviderForT(t, fetcher.Fetch, nil, nil, backoff.BackOff)

	var actual []xchain.Attestation
	p.Subscribe(ctx, chainID, fromHeight, "test", func(ctx context.Context, approved xchain.Attestation) error {
		actual = append(actual, approved)
		if len(actual) == total {
			cancel()
		}

		return nil
	})

	<-ctx.Done()

	// Test fetcher returns <errs> errors, then 0,1,2,3,4,... attestations per fetch.
	var expectFetched, expectCount int
	for i := 0; expectFetched <= total; i++ {
		expectCount++
		expectFetched += i
	}

	require.Empty(t, fetcher.Errs()) // All errors returned
	require.Equal(t, expectFetched, fetcher.Fetched())
	require.Equal(t, expectCount, fetcher.Count())
	require.Equal(t, errs+1, backoff.Count()) // 2 errors + 1 empty fetch

	require.Len(t, actual, total)
	for i, attestation := range actual {
		require.Equal(t, chainID, attestation.SourceChainID)
		require.Equal(t, fromHeight+uint64(i), attestation.BlockHeight)
	}
}

// testFetcher implements fetchFunc.
// It first returns errs errors.
// Then it returns 0,1,2,3,4,5... attestations up to max.
type testFetcher struct {
	mu      sync.Mutex
	errs    int
	count   int
	fetched int
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

func (f *testFetcher) Fetch(_ context.Context, chainID uint64, fromHeight uint64,
) ([]xchain.Attestation, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.errs > 0 {
		f.errs--
		return nil, errors.New("test error")
	}

	toReturn := f.count
	f.count++

	var resp []xchain.Attestation
	for i := 0; i < toReturn; i++ {
		resp = append(resp, xchain.Attestation{
			BlockHeader: xchain.BlockHeader{
				SourceChainID: chainID,
				BlockHeight:   fromHeight + uint64(i),
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
