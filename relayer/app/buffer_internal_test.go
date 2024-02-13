package relayer

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func Test_activeBuffer_AddInput(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	size := 0 // Unbuffered
	limit := int64(5)
	sender := &mockSender{}
	buffer := newActiveBuffer("test", limit, size, sender.Send)

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

type mockSender struct{}

func (m *mockSender) Send(context.Context, xchain.Submission) error {
	return nil
}

func Test_activeBuffer_Run(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	limit := int64(5)
	size := 0 // unbuffered
	sender := &mockSender{}
	buffer := newActiveBuffer("test", limit, size, sender.Send)

	// Run the buffer in a separate goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			err := buffer.Run(ctx)
			require.NoError(t, err)
		}
	}()

	// Add an item to the buffer
	err := buffer.AddInput(ctx, xchain.Submission{})
	require.NoError(t, err)
	require.Emptyf(t, buffer.buffer, "buffer should be empty after AddInput")
}
