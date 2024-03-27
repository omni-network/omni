package relayer

import (
	"context"
	"sync"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func TestWorker_Run(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())

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
	}
	streamB := xchain.StreamID{
		SourceChainID: srcChain,
		DestChainID:   destChainB,
	}
	cursors := map[uint64]xchain.StreamCursor{
		destChainA: {StreamID: streamA, Offset: destChainACursor, SourceBlockHeight: destChainACursor},
		destChainB: {StreamID: streamB, Offset: destChainBCursor, SourceBlockHeight: destChainBCursor},
	}

	// Return mock blocks (with a single msg per dest chain).
	mockXClient := &mockXChainClient{
		GetBlockFn: func(ctx context.Context, chainID uint64, height uint64) (xchain.Block, bool, error) {
			require.EqualValues(t, srcChain, chainID) // Only fetch blocks for source chains.

			// Each block has two messages, one for each stream.
			return xchain.Block{
				BlockHeader: xchain.BlockHeader{SourceChainID: chainID, BlockHeight: height},
				Msgs: []xchain.Msg{
					{MsgID: xchain.MsgID{StreamID: streamA, StreamOffset: height}},
					{MsgID: xchain.MsgID{StreamID: streamB, StreamOffset: height}},
				},
			}, true, nil
		},
		GetSubmittedCursorFn: func(_ context.Context, srcChainID uint64, _ uint64) (xchain.StreamCursor, bool, error) {
			return cursors[srcChainID], true, nil
		},
		GetEmittedCursorFn: func(_ context.Context, _ uint64, destChainID uint64) (xchain.StreamCursor, bool, error) {
			return cursors[destChainID], true, nil
		},
	}
	done := make(chan struct{})
	// Collect all stream updates via the creator, stop as soon as we get one msg from for streamB.
	var submissions []xchain.Submission
	var mutex = &sync.Mutex{}
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
		SendTransactionFn: func(ctx context.Context, submission xchain.Submission) error {
			require.Fail(t, "should not be called")
			return nil
		},
	}

	// Provider mock attestations as requested until context canceled.
	mockProvider := &mockProvider{
		SubscribeFn: func(ctx context.Context, chainID uint64, fromHeight uint64, callback cchain.ProviderCallback) {
			if chainID != srcChain {
				return // Only subscribe to source chain.
			}
			// require.EqualValues(t, destChainACursor, fromHeight)
			if fromHeight != destChainACursor && fromHeight != destChainBCursor {
				return
			}

			height := fromHeight
			nextAtt := func() xchain.Attestation {
				defer func() { height++ }()
				return xchain.Attestation{
					BlockHeader: xchain.BlockHeader{SourceChainID: chainID, BlockHeight: height},
				}
			}

			for ctx.Err() == nil {
				err := callback(ctx, nextAtt())
				require.NoError(t, err)
			}
		},
	}

	network := netconf.Network{Chains: []netconf.Chain{
		{ID: srcChain, Name: "source"},
		{ID: destChainA, Name: "chain_a"},
		{ID: destChainB, Name: "chain_b"},
	}}

	state := NewEmptyState("/tmp/relayer-state.json")

	noAwait := func(context.Context, uint64) error { return nil }

	for _, chain := range network.Chains {
		w := NewWorker(chain, network, mockProvider, mockXClient, mockCreateFunc, func() (SendFunc, error) {
			return mockSender.SendTransaction, nil
		}, state, noAwait)
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
	require.EqualValues(t, expectChainA, actualChainA)
	require.EqualValues(t, expectChainB, actualChainB)
}
