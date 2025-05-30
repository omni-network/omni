package relayer

import (
	"context"
	"sync"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/relayer/app/cursor"

	db "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

func TestWorker_Run(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(t.Context())

	const (
		srcChain         = 1
		destChainA       = 2
		destChainB       = 3
		destChainACursor = 10 // ChainA is lagging
		destChainBCursor = 20 // ChainB is ahead

	)
	expectChainA := destChainBCursor - destChainACursor + 1
	expectChainB := 1
	// totalMsgs := expectChainA + expectChainB

	streamA := xchain.StreamID{
		SourceChainID: srcChain,
		DestChainID:   destChainA,
		ShardID:       xchain.ShardFinalized0,
	}
	streamB := xchain.StreamID{
		SourceChainID: srcChain,
		DestChainID:   destChainB,
		ShardID:       xchain.ShardLatest0,
	}
	submittedCursors := map[xchain.StreamID]xchain.SubmitCursor{
		streamA: {StreamID: streamA, MsgOffset: destChainACursor, AttestOffset: destChainACursor},
		streamB: {StreamID: streamB, MsgOffset: destChainBCursor, AttestOffset: destChainBCursor},
	}

	// Return mock blocks (with a single msg per dest chain).
	mockXClient := &mockXChainClient{
		GetBlockFn: func(ctx context.Context, req xchain.ProviderRequest) (xchain.Block, bool, error) {
			require.EqualValues(t, srcChain, req.ChainID) // Only fetch blocks for source chains.

			// Each block has two messages, one for each stream.
			return xchain.Block{
				BlockHeader: xchain.BlockHeader{
					ChainID:     req.ChainID,
					BlockHeight: req.Height,
				},
				Msgs: []xchain.Msg{
					{MsgID: xchain.MsgID{StreamID: streamA, StreamOffset: req.Height}, LogIndex: 0},
					{MsgID: xchain.MsgID{StreamID: streamB, StreamOffset: req.Height}, LogIndex: 1},
				},
			}, true, nil
		},
		GetSubmittedCursorFn: func(_ context.Context, ref xchain.Ref, stream xchain.StreamID) (xchain.SubmitCursor, bool, error) {
			resp, ok := submittedCursors[stream]
			return resp, ok, nil
		},
	}
	done := make(chan struct{})
	// Collect all stream updates via the creator, stop as soon as we get one msg from for streamB.
	var submissions []xchain.Submission
	var mutex sync.Mutex
	submissionsChan := make(chan xchain.Submission)

	mockCreateFunc := func(streamUpdate StreamUpdate) ([]xchain.Submission, error) {
		mutex.Lock()
		defer mutex.Unlock()
		subs, err := CreateSubmissions(streamUpdate)
		if err != nil {
			return nil, err
		}

		for _, sub := range subs {
			submissionsChan <- sub
		}

		return nil, nil
	}

	// mockSender should never be called, since we return empty slices from the creator.
	mockSender := &mockSender{
		SendTransactionFn: func(ctx context.Context, submission xchain.Submission) <-chan error {
			require.Fail(t, "should not be called")
			return nil
		},
	}

	// Provider mock attestations as requested until context canceled.
	mockProvider := &mockProvider{
		SubscribeFn: func(ctx context.Context, chainVer xchain.ChainVersion, attestOffset uint64, callback cchain.ProviderCallback) {
			if chainVer.ID != srcChain {
				return // Only subscribe to source chain.
			}
			if attestOffset != destChainACursor && attestOffset != destChainBCursor {
				return
			}

			offset := attestOffset
			nextAtt := func() xchain.Attestation {
				defer func() { offset++ }()

				// Calculate the attestation root
				block, _, _ := mockXClient.GetBlock(ctx, xchain.ProviderRequest{
					ChainID:   chainVer.ID,
					ConfLevel: chainVer.ConfLevel,
					// We treat the offset as the requested height for the test.
					Height: offset,
				})
				tree, err := xchain.NewMsgTree(block.Msgs)
				require.NoError(t, err)

				return xchain.Attestation{
					MsgRoot: tree.MsgRoot(),
					AttestHeader: xchain.AttestHeader{
						ChainVersion: chainVer,
						AttestOffset: offset,
					},
					BlockHeader: xchain.BlockHeader{
						ChainID:     chainVer.ID,
						BlockHeight: offset,
					},
					ValidatorSetID: mockValSetID,
				}
			}

			for ctx.Err() == nil {
				err := callback(ctx, nextAtt())
				if ctx.Err() != nil {
					return
				}
				require.NoError(t, err)
			}
		},
	}

	network := netconf.Network{Chains: []netconf.Chain{
		{ID: srcChain, Name: "source", Shards: []xchain.ShardID{xchain.ShardFinalized0, xchain.ShardLatest0}, HasEmitPortal: true},
		{ID: destChainA, Name: "mock_l1", HasSubmitPortal: true},
		{ID: destChainB, Name: "mock_l2", HasSubmitPortal: true},
	}}

	noAwait := func(context.Context, uint64) error { return nil }

	cursors, err := cursor.New(db.NewMemDB(), mockXClient.GetSubmittedCursor, network)
	require.NoError(t, err)

	for _, chain := range network.Chains {
		w := NewWorker(
			chain,
			network,
			mockProvider,
			mockXClient,
			mockCreateFunc,
			func() (SendAsync, error) { return mockSender.SendTransaction, nil },
			noAwait,
			cursors)
		go w.Run(ctx)
	}

	go func() {
		for sub := range submissionsChan {
			chainBMsgs := 0
			chainAMsgs := 0
			for _, msg := range submissions {
				switch msg.DestChainID {
				case destChainA:
					chainAMsgs++
				case destChainB:
					chainBMsgs++
				}
			}

			for _, msg := range sub.Msgs {
				if chainAMsgs < expectChainA && msg.DestChainID == destChainA {
					submissions = append(submissions, sub)
					chainAMsgs++
				} else if chainBMsgs < expectChainB && msg.DestChainID == destChainB {
					submissions = append(submissions, sub)
					chainBMsgs++
				} else if chainAMsgs == expectChainA && chainBMsgs == expectChainB {
					close(done)
					cancel()

					return
				}
			}
		}
	}()

	// Wait for messages to come in.
	<-done

	// Verify responses
	require.Len(t, submissions, expectChainA+expectChainB)

	// Ensure msgs are delivered in sequence
	var actualChainA, actualChainB int
	prevChainA, prevChainB := destChainACursor, destChainBCursor
	for _, submission := range submissions {
		require.Len(t, submission.Msgs, 1)
		next := submission.Msgs[0].StreamOffset
		if submission.DestChainID == destChainA {
			actualChainA++
			prevChainA++
			require.EqualValues(t, prevChainA, next)
		} else {
			actualChainB++
			prevChainB++
			require.EqualValues(t, prevChainB, next)
		}
	}

	// Ensure totals.
	require.Equal(t, expectChainA, actualChainA)
	require.Equal(t, expectChainB, actualChainB)
}
