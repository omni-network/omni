package db_test

import (
	"testing"
	"time"

	"github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/testutil"
	"github.com/omni-network/omni/lib/cctp/types"

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
		msgs[i] = testutil.RandMsg()
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
		gotMsgs, err := db.GetMsgs(ctx)
		require.NoError(t, err)
		require.NotNil(t, gotMsgs)
		require.Len(t, gotMsgs, i+1)
		require.Contains(t, gotMsgs, msg)
	}

	// Asset all timestamps non-zero, save to check they remain unchanged
	timestamps := make([]time.Time, numMsgs)
	for i, msg := range msgs {
		createdAt, err := db.GetMsgCreatedAt(ctx, msg.TxHash)
		require.NoError(t, err)
		require.NotZero(t, createdAt)
		timestamps[i] = createdAt
	}

	// Test SetMsg
	for i, msg := range msgs {
		// Modify the message hash / bytes (simulates reorg)
		updated := msg
		updated.MessageBytes = testutil.RandBytes(200)
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

		// Asset timestamp remains unchanged
		createdAt, err := db.GetMsgCreatedAt(ctx, msg.TxHash)
		require.NoError(t, err)
		require.Equal(t, timestamps[i], createdAt)
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
		gotMsgs, err := db.GetMsgs(ctx)
		require.NoError(t, err)
		require.NotNil(t, gotMsgs)
		require.Len(t, gotMsgs, numMsgs-i-1)
		require.NotContains(t, gotMsgs, msg)
	}

	// Assert no messages left
	gotMsgs, err := db.GetMsgs(ctx)
	require.NoError(t, err)
	require.NotNil(t, gotMsgs)
	require.Empty(t, gotMsgs)
}
