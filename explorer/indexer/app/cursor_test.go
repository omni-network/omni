package app_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/xprovidercursor"
	"github.com/omni-network/omni/explorer/indexer/app"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestCursor(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name   string
		chain  netconf.Chain
		blocks uint64
	}{
		{
			name: "zero_index_cursor",
			chain: netconf.Chain{
				ID:                100,
				Name:              "mock_l1",
				RPCURL:            "http://mock_l1:8545",
				PortalAddress:     common.Address([]byte("0x268bb5F3d4301b591288390E76b97BE8E8B1Ca82")),
				DeployHeight:      0,
				IsEthereum:        true,
				BlockPeriod:       time.Duration(1) * time.Second,
				FinalizationStrat: "latest",
			},
			blocks: 5,
		},
		{
			name: "public_chain_non_zero_cursor",
			chain: netconf.Chain{
				ID:                1016561,
				Name:              "omni_consensus",
				RPCURL:            "http://mock_arb:8545",
				PortalAddress:     common.Address([]byte("0x268bb5F3d4301b591288390E76b97BE8E8B1Ca82")),
				DeployHeight:      10687126,
				IsOmniConsensus:   true,
				BlockPeriod:       time.Duration(2) * time.Second,
				FinalizationStrat: "latest",
			},
			blocks: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			entClient := setupDB(t)

			// Get our initial deploy height
			deployHeight, err := app.InitChainCursor(ctx, entClient, tt.chain)
			require.NoError(t, err)

			// Insert TestBlocks
			for i := uint64(0); i < tt.blocks; i++ {
				insertBlock(t, ctx, entClient, tt.chain.ID, deployHeight+i)
			}

			// Gets our final cursor location
			cursor, err := entClient.XProviderCursor.Query().Where(xprovidercursor.ChainID(tt.chain.ID)).Only(ctx)
			require.NoError(t, err)

			// Check our results
			require.Equal(t, tt.chain.DeployHeight, deployHeight)

			// our cursoe should equal the height of the last block we inserted
			// we subtract one because our cursor starts at deployHeight - 1
			require.Equal(t, tt.chain.DeployHeight+tt.blocks-1, cursor.Height)
		})
	}
}

func insertBlock(t *testing.T, ctx context.Context, client *ent.Client, sourceChainID uint64, blockHeight uint64) {
	t.Helper()
	var msgTxHash, receiptTxHash common.Hash
	fuzz.New().NilChance(0).Fuzz(&msgTxHash)
	fuzz.New().NilChance(0).Fuzz(&receiptTxHash)

	var blockHash common.Hash
	fuzz.New().NilChance(0).Fuzz(&blockHash)

	var sourceMessageSender, destAddress, relayerAddress [20]byte
	fuzz.New().NilChance(0).Fuzz(&sourceMessageSender)
	fuzz.New().NilChance(0).Fuzz(&destAddress)
	fuzz.New().NilChance(0).Fuzz(&relayerAddress)

	var msgData []byte
	fuzz.New().NilChance(0).Fuzz(&msgData)

	destChainID := uint64(2)

	gasLimit := uint64(1000)
	streamOffset := uint64(0)
	gasUsed := uint64(100)

	tx, err := client.BeginTx(ctx, nil)
	require.NoError(t, err)

	err = app.InsertBlockTX(ctx, tx, xchain.Block{
		BlockHeader: xchain.BlockHeader{
			SourceChainID: sourceChainID,
			BlockHeight:   blockHeight,
			BlockHash:     blockHash,
		},
		Msgs: []xchain.Msg{
			{
				MsgID: xchain.MsgID{
					StreamID: xchain.StreamID{
						SourceChainID: sourceChainID,
						DestChainID:   destChainID,
					},
					StreamOffset: streamOffset,
				},
				SourceMsgSender: sourceMessageSender,
				DestAddress:     destAddress,
				Data:            msgData,
				DestGasLimit:    gasLimit,
				TxHash:          msgTxHash,
			},
		},
		Receipts: []xchain.Receipt{
			{
				MsgID: xchain.MsgID{
					StreamID: xchain.StreamID{
						SourceChainID: sourceChainID,
						DestChainID:   destChainID,
					},
					StreamOffset: streamOffset,
				},
				GasUsed:        gasUsed,
				Success:        true,
				RelayerAddress: common.Address(relayerAddress[:]),
				TxHash:         receiptTxHash,
			},
		},
		Timestamp: time.Now(),
	})
	require.NoError(t, err)
}
