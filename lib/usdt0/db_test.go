package usdt0_test

import (
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/layerzero"
	"github.com/omni-network/omni/lib/usdt0"

	"github.com/ethereum/go-ethereum/crypto"

	cosmosdb "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

// randMsg generates a random USDT0 message for testing.
func randMsg() usdt0.MsgSend {
	amount := big.NewInt(rand.Int63n(1000) + 1)
	amount.Mul(amount, big.NewInt(1e6))

	srcChainID := uint64(rand.Intn(10) + 1)
	destChainID := uint64(rand.Intn(10) + 1)
	for destChainID == srcChainID {
		destChainID = uint64(rand.Intn(10) + 1)
	}

	txHash := crypto.Keccak256Hash([]byte(time.Now().String()))

	return usdt0.MsgSend{
		TxHash:      txHash,
		BlockHeight: uint64(rand.Intn(1000) + 1),
		SrcChainID:  srcChainID,
		DestChainID: destChainID,
		Amount:      amount,
		Status:      layerzero.MsgStatusConfirming,
	}
}

func TestMsgDB(t *testing.T) {
	t.Parallel()

	memDB := cosmosdb.NewMemDB()
	db, err := usdt0.NewDB(memDB)
	require.NoError(t, err)

	ctx := t.Context()

	numMsgs := 10
	msgs := make([]usdt0.MsgSend, numMsgs)
	for i := 0; i < numMsgs; i++ {
		msgs[i] = randMsg()
	}

	// Test InsertMsg and verify CreatedAt
	for i, msg := range msgs {
		err := db.InsertMsg(ctx, msg)
		require.NoError(t, err)

		// Verify CreatedAt is set
		createdAt, err := db.GetMsgCreatedAt(ctx, msg.TxHash)
		require.NoError(t, err)
		require.NotZero(t, createdAt)

		gotMsgs, err := db.GetMsgs(ctx)
		require.NoError(t, err)
		require.NotNil(t, gotMsgs)
		require.Len(t, gotMsgs, i+1)
		require.Contains(t, gotMsgs, msg)
	}

	// Test SetMsgStatus
	for _, msg := range msgs {
		err := db.SetMsgStatus(ctx, msg.TxHash, layerzero.MsgStatusDelivered)
		require.NoError(t, err)

		gotMsgs, err := db.GetMsgs(ctx)
		require.NoError(t, err)
		for _, gotMsg := range gotMsgs {
			if gotMsg.TxHash == msg.TxHash {
				require.Equal(t, layerzero.MsgStatusDelivered, gotMsg.Status)
				break
			}
		}
	}

	// Filter by src
	srcChainID := msgs[0].SrcChainID
	gotMsgs, err := db.GetMsgs(ctx, usdt0.FilterMsgBySrc(srcChainID))
	require.NoError(t, err)
	for _, msg := range gotMsgs {
		require.Equal(t, srcChainID, msg.SrcChainID)
	}

	// Filter by dest
	destChainID := msgs[0].DestChainID
	gotMsgs, err = db.GetMsgs(ctx, usdt0.FilterMsgByDest(destChainID))
	require.NoError(t, err)
	for _, msg := range gotMsgs {
		require.Equal(t, destChainID, msg.DestChainID)
	}

	// Filter by status
	gotMsgs, err = db.GetMsgs(ctx, usdt0.FilterMsgByStatus(layerzero.MsgStatusDelivered))
	require.NoError(t, err)
	for _, msg := range gotMsgs {
		require.Equal(t, layerzero.MsgStatusDelivered, msg.Status)
	}

	// Multiple filters
	srcChainID = msgs[0].SrcChainID
	destChainID = msgs[0].DestChainID
	gotMsgs, err = db.GetMsgs(ctx,
		usdt0.FilterMsgBySrc(srcChainID),
		usdt0.FilterMsgByDest(destChainID),
		usdt0.FilterMsgByStatus(layerzero.MsgStatusDelivered))
	require.NoError(t, err)
	for _, msg := range gotMsgs {
		require.Equal(t, srcChainID, msg.SrcChainID)
		require.Equal(t, destChainID, msg.DestChainID)
		require.Equal(t, layerzero.MsgStatusDelivered, msg.Status)
	}

	// Delete messages and verify
	for i, msg := range msgs {
		err := db.DeleteMsg(ctx, msg.TxHash)
		require.NoError(t, err)
		gotMsgs, err := db.GetMsgs(ctx)
		require.NoError(t, err)
		require.NotNil(t, gotMsgs)
		require.Len(t, gotMsgs, numMsgs-i-1)
		require.NotContains(t, gotMsgs, msg)
	}

	gotMsgs, err = db.GetMsgs(ctx)
	require.NoError(t, err)
	require.NotNil(t, gotMsgs)
	require.Empty(t, gotMsgs)
}
