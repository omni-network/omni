package relayer

import (
	"context"
	"sync"
	"testing"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

var errSentAsync = errors.New("sent async")

func TestSendAsync(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	fuzzer := fuzz.New().NilChance(0).NumElements(1, 8)
	mockGasEstimator := func(destChain uint64, msgs []xchain.Msg) uint64 {
		return 0
	}
	txMgr := newMockTxMgr()
	sender := Sender{
		network:      netconf.Simnet,
		txMgr:        txMgr,
		gasEstimator: mockGasEstimator,
		abi:          nil,
		chain:        netconf.Chain{},
		gasToken:     "",
		chainNames:   nil,
		ethCl:        nil,
	}

	const total = 10

	// Enqueue a bunch of random submissions
	var resps []<-chan error
	for range total {
		var sub xchain.Submission
		fuzzer.Fuzz(&sub)
		sub.DestChainID = 0 // Ensure destChain matches zero netconf chain above

		resp := sender.SendAsync(ctx, sub)

		// Ensure it didn't error synchronously.
		select {
		case err := <-resp:
			require.Fail(t, "unexpected synchronous error", "error: %v", err)
		default:
			// Nothing in the channel yet, expected.
		}

		resps = append(resps, resp)
	}

	// Ensure nonces were reserved sequentially.
	require.EqualValues(t, total, txMgr.ReservedNonces())

	// Trigger each send and ensure expected result (errSentAsync)
	for _, resp := range resps {
		txMgr.MineNext()
		require.ErrorIs(t, <-resp, errSentAsync)
	}
}

var _ txmgr.TxManager = (*mockTxMgr)(nil)

func newMockTxMgr() *mockTxMgr {
	return &mockTxMgr{
		sends: make(chan txmgr.TxCandidate, 1),
	}
}

type mockTxMgr struct {
	sync.Mutex
	nonces uint64
	sends  chan txmgr.TxCandidate
}

func (m *mockTxMgr) Send(_ context.Context, candidate txmgr.TxCandidate) (*types.Transaction, *types.Receipt, error) {
	// Blocks until MineNext is called.
	m.sends <- candidate

	return nil, nil, errSentAsync
}

func (m *mockTxMgr) MineNext() txmgr.TxCandidate {
	return <-m.sends
}

func (m *mockTxMgr) From() common.Address {
	panic("implement me")
}

func (m *mockTxMgr) ReservedNonces() uint64 {
	m.Lock()
	defer m.Unlock()

	return m.nonces
}

func (m *mockTxMgr) ReserveNextNonce(context.Context) (uint64, error) {
	m.Lock()
	defer m.Unlock()

	m.nonces++

	return m.nonces, nil
}
