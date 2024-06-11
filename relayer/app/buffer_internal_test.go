package relayer

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/xchain"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_activeBuffer_AddInput(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	limit := int64(5)
	sender := &mockBufSender{}
	buffer := newActiveBuffer("test", limit, sender.Send)

	// Have a reader ready as we are unbuffered and blocking
	go func() {
		for range buffer.buffer {
		}
	}()

	err := buffer.AddInput(ctx, xchain.Submission{})
	require.NoError(t, err)

	select {
	case <-ctx.Done():
		t.Errorf("AddInput is blocking and should not have")
	default:
	}
}

type mockBufSender struct {
	sendChan chan xchain.Submission
}

func newMockSender() *mockBufSender {
	return &mockBufSender{
		sendChan: make(chan xchain.Submission),
	}
}

func (m *mockBufSender) Send(_ context.Context, sub xchain.Submission) error {
	m.sendChan <- sub
	return nil
}

func (m *mockBufSender) Next() xchain.Submission {
	return <-m.sendChan
}

// Test_activeBuffer_Run tests that the buffer is blocking when the number of submissions is greater
// than the mempoolLimit.
func Test_activeBuffer_Run(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//
	const (
		memLimit = int64(5) // mempoolLimit
		size     = 10
	)

	sender := newMockSender()
	buffer := newActiveBuffer("test", memLimit, sender.Send)

	var input []xchain.Submission
	fuzz.New().NilChance(0).NumElements(size, size).Fuzz(&input)

	go func() {
		err := buffer.Run(ctx)
		assert.ErrorIs(t, err, context.Canceled)
	}()

	counter := new(atomic.Int64)
	go func() {
		for _, sub := range input {
			err := buffer.AddInput(ctx, sub)
			assert.NoError(t, err)
			counter.Add(1)
		}
	}()

	require.Eventuallyf(t, func() bool {
		return counter.Load() == memLimit+1
	},
		time.Second, time.Millisecond, "expected %d", memLimit+1)

	// assert again that buf is blocking
	require.Equal(t, memLimit+1, counter.Load())

	// Retrieve output submissions
	var output []xchain.Submission
	for len(input) != len(output) {
		output = append(output, sender.Next())
	}

	require.Eventuallyf(t, func() bool {
		return counter.Load() == int64(size)
	}, time.Second, time.Millisecond, "expected %d", size)

	// Assert equality of input and output submissions
	require.Len(t, input, len(output))
}
