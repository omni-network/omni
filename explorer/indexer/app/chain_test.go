package app_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/explorer/db/ent/chain"
	"github.com/omni-network/omni/explorer/db/ent/xprovidercursor"
	"github.com/omni-network/omni/explorer/indexer/app"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestChain(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type want struct {
		cursorHeight uint64
		deployHeight uint64
	}

	tests := []struct {
		name          string
		chain         netconf.Chain
		prerequisites prerequisite // These functions create entries on our db before the evaluation
		want          want
	}{
		{
			name: "insert_chain_height_zero",
			chain: netconf.Chain{
				ID:            100,
				Name:          "mock_l1",
				PortalAddress: common.Address([]byte("0x268bb5F3d4301b591288390E76b97BE8E8B1Ca82")),
				DeployHeight:  0,
				BlockPeriod:   1 * time.Second,
				Shards:        []xchain.ShardID{xchain.ShardLatest0},
			},
			want: want{
				cursorHeight: 0,
				deployHeight: 0,
			},
		},
		{
			name: "insert_chain_height_non_zero",
			chain: netconf.Chain{
				ID:            1016561,
				Name:          "omni_consensus",
				PortalAddress: common.Address([]byte("0x268bb5F3d4301b591288390E76b97BE8E8B1Ca82")),
				DeployHeight:  10687126,
				BlockPeriod:   2 * time.Second,
				Shards:        []xchain.ShardID{xchain.ShardLatest0},
			},
			want: want{
				cursorHeight: 10687125,
				deployHeight: 10687126,
			},
		},
	}

	for _, tt := range tests {
		// this returns the deploy height
		entClient := setupDB(t)
		// TODO (pavel): verify correct offset
		height, _, err := app.InitChainCursor(ctx, entClient, tt.chain)
		require.NoError(t, err)
		require.Equal(t, tt.want.deployHeight, height)

		// this should return the chain height for our starting queries
		cursor, err := entClient.XProviderCursor.Query().Where(xprovidercursor.ChainID(tt.chain.ID)).Only(ctx)
		require.NoError(t, err)
		require.Equal(t, tt.want.cursorHeight, cursor.Height)

		chain, err := entClient.Chain.Query().Where(chain.ChainID(tt.chain.ID)).Only(ctx)
		require.NoError(t, err)
		require.Equal(t, tt.chain.Name, chain.Name)
	}
}
