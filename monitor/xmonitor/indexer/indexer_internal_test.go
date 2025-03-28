package indexer

import (
	"context"
	"fmt"
	"testing"

	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	dbm "github.com/cosmos/cosmos-db"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

//go:generate go test . -count=1000 -race

func TestIndexer(t *testing.T) {
	t.Parallel()

	f := fuzz.New().NilChance(0).NumElements(0, 10)
	ctx := t.Context()
	db := dbm.NewMemDB()

	streamNamer := func(s xchain.StreamID) string { return fmt.Sprint(s) }

	indexer, err := newIndexer(db, mockXProvider{}, streamNamer, nil)
	require.NoError(t, err)
	var samples []sample
	indexer.sampleFunc = func(s sample) {
		samples = append(samples, s)
	}

	var stream xchain.StreamID
	f.Fuzz(&stream)

	const total = 4

	var msgs []xchain.Msg
	var receipts []xchain.Receipt
	for i := 0; i < total; i++ {
		var msg xchain.Msg
		f.Fuzz(&msg)
		msgs = append(msgs, msg)

		var receipt xchain.Receipt
		f.Fuzz(&receipt)
		receipt.MsgID = msg.MsgID
		receipts = append(receipts, receipt)
	}

	blocks := []xchain.Block{
		fuzzBlock(f, msgs[0:2], nil),
		fuzzBlock(f, msgs[2:4], receipts[0:2]),
		fuzzBlock(f, nil, receipts[2:4]),
	}

	samplesPerBlock := map[int][]sample{
		0: nil, // No receipts yet
		1: {
			makeSample(blocks, receipts, msgs, 0),
			makeSample(blocks, receipts, msgs, 1),
		},
		2: {
			makeSample(blocks, receipts, msgs, 2),
			makeSample(blocks, receipts, msgs, 3),
		},
	}

	deletedPerBlock := map[int][]xchain.BlockHeader{
		0: nil,
		1: {blocks[0].BlockHeader},
		2: {blocks[1].BlockHeader, blocks[2].BlockHeader},
	}

	c, err := indexer.cursors(ctx)
	require.NoError(t, err)
	require.Empty(t, c)

	cursors := make(map[xchain.ChainVersion]uint64)
	for i, block := range blocks {
		// Insert the same block twice for idempotency
		for range 2 {
			err := indexer.index(ctx, block)
			tutil.RequireNoError(t, err)
			require.Equal(t, samplesPerBlock[i], samples)
			samples = nil
		}
		cursors[xchain.NewChainVersion(block.ChainID, xchain.ConfFinalized)] = block.BlockHeight

		deleted, err := indexer.delete(ctx)
		tutil.RequireNoError(t, err)
		require.Equal(t, deletedPerBlock[i], deleted)
	}

	c, err = indexer.cursors(ctx)
	require.NoError(t, err)
	require.Equal(t, cursors, c)

	// Ensure both blocks and msgLink tables are empty.
	biter, err := indexer.blockTable.List(ctx, BlockPrimaryKey{})
	require.NoError(t, err)
	defer biter.Close()
	require.False(t, biter.Next())
	liter, err := indexer.msgLinkTable.List(ctx, MsgLinkPrimaryKey{})
	require.NoError(t, err)
	defer liter.Close()
	require.False(t, liter.Next())

	// Ensure empty blocks also update cursors every emptyBlockCursorUpdate
	for i := range uint64(2) {
		emptyBlock := fuzzBlock(f, nil, nil)
		emptyBlock.BlockHeight -= emptyBlock.BlockHeight % emptyBlockCursorUpdate // Truncate
		emptyBlock.BlockHeight += i                                               // Maybe offset
		err = indexer.index(ctx, emptyBlock)
		require.NoError(t, err)
		require.Empty(t, samples)

		cursors, err := indexer.cursors(ctx)
		require.NoError(t, err)

		chainVer := xchain.NewChainVersion(emptyBlock.ChainID, xchain.ConfFinalized)
		if i == 0 {
			require.Equal(t, emptyBlock.BlockHeight, cursors[chainVer])
		} else {
			require.NotEqual(t, emptyBlock.BlockHeight, cursors[chainVer])
		}
	}
}

func makeSample(blocks []xchain.Block, receipts []xchain.Receipt, msgs []xchain.Msg, idx int) sample {
	return sample{
		Stream:        fmt.Sprint(receipts[idx].StreamID),
		XDApp:         unknown,
		SrcChain:      unknown,
		FeeToken:      unknown,
		FeeAmount:     msgs[idx].Fees,
		Latency:       getReceiptBlock(blocks, receipts[idx].MsgID).Timestamp.Sub(getMsgBlock(blocks, msgs[idx].MsgID).Timestamp),
		ExcessGas:     umath.SubtractOrZero(msgs[idx].DestGasLimit, receipts[idx].GasUsed),
		Success:       receipts[idx].Success,
		FuzzyOverride: expectFuzzyOverride(receipts[idx]),
	}
}

func expectFuzzyOverride(receipt xchain.Receipt) bool {
	if !receipt.ConfLevel().IsFuzzy() {
		return false
	}

	return !mockConfLevel(receipt.TxHash).IsFuzzy()
}

func getReceiptBlock(blocks []xchain.Block, msgID xchain.MsgID) xchain.Block {
	for _, block := range blocks {
		_, err := block.ReceiptByID(msgID)
		if err == nil {
			return block
		}
	}

	panic("msgID not in block")
}

func getMsgBlock(blocks []xchain.Block, msgID xchain.MsgID) xchain.Block {
	for _, block := range blocks {
		_, err := block.MsgByID(msgID)
		if err == nil {
			return block
		}
	}

	panic("msgID not in block")
}

func fuzzBlock(f *fuzz.Fuzzer, msgs []xchain.Msg, receipts []xchain.Receipt) xchain.Block {
	var resp xchain.Block
	f.Fuzz(&resp)
	resp.Msgs = msgs
	resp.Receipts = receipts

	return resp
}

type mockXProvider struct {
	xchain.Provider
}

func (m mockXProvider) GetSubmission(_ context.Context, chainID uint64, txHash common.Hash) (xchain.Submission, error) {
	return xchain.Submission{
		AttHeader: xchain.AttestHeader{
			ChainVersion: xchain.ChainVersion{
				ID:        chainID,
				ConfLevel: mockConfLevel(txHash),
			},
		},
	}, nil
}

func mockConfLevel(txHash common.Hash) xchain.ConfLevel {
	resp := xchain.ConfFinalized
	if txHash[0]%2 == 0 {
		resp = xchain.ConfLatest
	}

	return resp
}
