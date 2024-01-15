package provider_test

import (
	"context"
	"errors"
	"math/rand"
	"testing"

	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func TestProvider(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())

	const (
		errs      = 2
		batchSize = 5
		total     = 10
	)

	chainID := rand.Uint64()
	fromHeight := rand.Uint64()

	var backoff testBackOff
	fetcher := &testFetcher{errs: errs}

	p := provider.NewProviderForT(t, fetcher.Fetch, batchSize, backoff.BackOff)

	var actual []xchain.AggAttestation
	p.Subscribe(ctx, chainID, fromHeight, func(ctx context.Context, approved xchain.AggAttestation) error {
		actual = append(actual, approved)
		if len(actual) == total {
			cancel()
		}

		return nil
	})

	<-ctx.Done()

	require.Empty(t, fetcher.errs) // All errors returned
	require.Equal(t, total, fetcher.fetched)
	require.Equal(t, errs+1, backoff.backoff) // 2 errors + 1 empty fetch
	require.Equal(t, total+errs+1, backoff.reset)

	require.Len(t, actual, total)
	for i, attestation := range actual {
		require.Equal(t, chainID, attestation.SourceChainID)
		require.Equal(t, fromHeight+uint64(i), attestation.BlockHeight)
	}
}

// testFetcher implements FetchFunc.
// It first returns errs errors.
// Then it returns 0,1,2,3,4,5... attestations up to max.
type testFetcher struct {
	errs    int
	count   int
	fetched int
}

func (f *testFetcher) Fetch(_ context.Context, chainID uint64, fromHeight uint64, max uint64,
) ([]xchain.AggAttestation, error) {
	if f.errs > 0 {
		f.errs--
		return nil, errors.New("test error")
	}

	toReturn := f.count
	f.count++

	if toReturn > int(max) {
		toReturn = int(max)
	}

	var resp []xchain.AggAttestation
	for i := 0; i < toReturn; i++ {
		resp = append(resp, xchain.AggAttestation{
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
	backoff int
	reset   int
}

func (b *testBackOff) BackOff(context.Context) (func(), func()) {
	return func() { b.backoff++ }, func() { b.reset++ }
}
