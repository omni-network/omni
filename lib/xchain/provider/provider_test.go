package provider_test

import (
	"context"
	"math/big"
	"sync"
	"testing"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/mock"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/lib/xchain/provider"

	"github.com/ethereum/go-ethereum"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestProvider(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())

	var mu sync.Mutex
	const (
		errs    = 2
		total   = 20
		workers = 2

		chainID    = uint64(999)
		fromHeight = uint64(100)
	)

	// Setup test input and mocks.
	backoff := new(testBackOff)
	network := netconf.Network{
		ID: netconf.Simnet,
		Chains: []netconf.Chain{{
			ID:                chainID,
			FinalizationStrat: netconf.StratFinalized,
		}},
	}

	ctrl := gomock.NewController(t)
	mockEthCl := mock.NewMockClient(ctrl)

	// Return a few errors from HeaderByType, then return a very high number so it is cached and not queried again.
	var remainErrs = errs
	mockEthCl.EXPECT().HeaderByType(gomock.Any(), ethclient.HeadFinalized).AnyTimes().DoAndReturn(func(ctx context.Context, typ ethclient.HeadType) (*ethtypes.Header, error) {
		mu.Lock()
		defer mu.Unlock()
		if remainErrs > 0 {
			remainErrs--
			return nil, errors.New("test errpr")
		}

		return &ethtypes.Header{
			Number: big.NewInt(int64(fromHeight * 10)),
		}, nil
	})

	// Return simple headers when queried
	mockEthCl.EXPECT().HeaderByNumber(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(_ context.Context, number *big.Int) (*ethtypes.Header, error) {
		return &ethtypes.Header{
			Number: number,
		}, nil
	})

	// Return empty logs when queried
	mockEthCl.EXPECT().FilterLogs(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(func(context.Context, ethereum.FilterQuery) ([]ethtypes.Log, error) {
		return nil, nil
	})

	xprov := provider.NewForT(t, network, map[uint64]ethclient.Client{chainID: mockEthCl}, backoff.BackOff, workers)

	var actual []xchain.Block
	err := xprov.StreamBlocks(ctx, chainID, fromHeight, func(ctx context.Context, block xchain.Block) error {
		actual = append(actual, block)
		if len(actual) == total {
			cancel()
		}

		return nil
	})
	require.NoError(t, err)

	<-ctx.Done()

	require.Equal(t, errs, backoff.Count()) // 2 errors

	require.Len(t, actual, total)
	for i, attestation := range actual {
		require.Equal(t, chainID, attestation.SourceChainID)
		require.Equal(t, fromHeight+uint64(i), attestation.BlockHeight)
	}
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
