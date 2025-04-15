package db_test

import (
	crand "crypto/rand"
	"math/big"
	mrand "math/rand"
	"testing"

	"github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	cosmosdb "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

func TestCursorDB(t *testing.T) {
	t.Parallel()

	// Create a new database
	memDB := cosmosdb.NewMemDB()
	db, err := db.New(memDB)
	require.NoError(t, err)

	ctx := t.Context()

	chainID := uint64(1)
	height := uint64(100)

	// Test setting cursor
	err = db.SetCursor(ctx, chainID, height)
	require.NoError(t, err)

	// Test getting cursor
	got, ok, err := db.GetCursor(ctx, chainID)
	require.NoError(t, err)
	require.True(t, ok)
	require.Equal(t, height, got)

	// Test getting non-existent cursor
	_, ok, err = db.GetCursor(ctx, 999)
	require.NoError(t, err)
	require.False(t, ok)
}

func TestMsgDB(t *testing.T) {
	t.Parallel()

	// Create a new database
	memDB := cosmosdb.NewMemDB()
	db, err := db.New(memDB)
	require.NoError(t, err)

	ctx := t.Context()

	numMsgs := 10
	msgs := make([]types.MsgSendUSDC, numMsgs)
	for i := 0; i < numMsgs; i++ {
		msgBz := mustRandBytes(100)
		msgHash := crypto.Keccak256Hash(msgBz)

		msgs[i] = types.MsgSendUSDC{
			TxHash:       common.BytesToHash(mustRandBytes(32)),
			SrcChainID:   uint64(mrand.Uint32()), // cctp uses uint32 domain ids
			DestChainID:  uint64(mrand.Uint32()),
			Amount:       big.NewInt(mrand.Int63()),
			MessageBytes: msgBz,
			MessageHash:  msgHash,
			Recipient:    common.BytesToAddress(mustRandBytes(20)),
		}
	}

	// Test InsertMsg and GetMsg
	for i, msg := range msgs {
		// Insert message
		err := db.InsertMsg(ctx, msg)
		require.NoError(t, err)

		// Assert message exists
		ok, err := db.HasMsg(ctx, msg.TxHash)
		require.NoError(t, err)
		require.True(t, ok)

		// Assert message can be retrieved
		gotMsg, ok, err := db.GetMsg(ctx, msg.TxHash)
		require.NoError(t, err)
		require.True(t, ok)
		require.Equal(t, msg, gotMsg)

		// List and verify all messages added so far
		listMsgs, err := db.ListMsgs(ctx)
		require.NoError(t, err)
		require.NotNil(t, listMsgs)
		require.Len(t, listMsgs, i+1)
		require.Contains(t, listMsgs, msg)
	}

	// Test SetMsg
	for _, msg := range msgs {
		// Modify the message hash / bytes (simulates reorg)
		updated := msg
		updated.MessageBytes = mustRandBytes(200)
		updated.MessageHash = crypto.Keccak256Hash(updated.MessageBytes)

		// Update the message
		err := db.SetMsg(ctx, updated)
		require.NoError(t, err)

		// Assert message was updated
		gotMsg, ok, err := db.GetMsg(ctx, msg.TxHash)
		require.NoError(t, err)
		require.True(t, ok)
		require.Equal(t, updated, gotMsg)
		require.NotEqual(t, msg, gotMsg)
	}

	// Test DeleteMsg
	for i, msg := range msgs {
		// Delete message
		err := db.DeleteMsg(ctx, msg.TxHash)
		require.NoError(t, err)

		// Assert message no longer exists
		ok, err := db.HasMsg(ctx, msg.TxHash)
		require.NoError(t, err)
		require.False(t, ok)

		// Assert message cannot be retrieved
		_, ok, err = db.GetMsg(ctx, msg.TxHash)
		require.NoError(t, err)
		require.False(t, ok)

		// List and verify remaining
		listMsgs, err := db.ListMsgs(ctx)
		require.NoError(t, err)
		require.NotNil(t, listMsgs)
		require.Len(t, listMsgs, numMsgs-i-1)
		require.NotContains(t, listMsgs, msg)
	}

	// Assert no messages left
	listMsgs, err := db.ListMsgs(ctx)
	require.NoError(t, err)
	require.NotNil(t, listMsgs)
	require.Empty(t, listMsgs)
}

func mustRandBytes(n int) []byte {
	b := make([]byte, n)

	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}

	return b
}
