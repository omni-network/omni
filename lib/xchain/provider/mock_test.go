package provider_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/lib/xchain/provider"

	"github.com/stretchr/testify/require"
)

//nolint:testifylint // This makes it less readable.
func TestMock(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())

	const (
		chainID    = 123
		fromHeight = 456
		total      = 5
	)

	var mock provider.Mock

	var blocks []xchain.Block
	err := mock.Subscribe(ctx, chainID, fromHeight, func(ctx context.Context, block *xchain.Block) error {
		blocks = append(blocks, *block)
		if len(blocks) == total {
			cancel()
		}

		return nil
	})
	require.NoError(t, err)

	<-ctx.Done()
	require.Len(t, blocks, total)

	// Just some very basic sanity checks
	require.Len(t, blocks[0].Msgs, 0)
	require.Len(t, blocks[1].Msgs, 1)
	require.Len(t, blocks[2].Msgs, 1)
	require.Len(t, blocks[3].Msgs, 2)
	require.Len(t, blocks[4].Msgs, 0)
}
