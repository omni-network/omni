package types_test

import (
	"testing"

	"github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/testutil"
	"github.com/omni-network/omni/lib/cctp/types"

	"github.com/stretchr/testify/require"
)

// TestMsgStatus asserts that MsgStatus type matches proto enum.
func TestMsgStatus(t *testing.T) {
	t.Parallel()

	require.Equal(t, int32(db.MsgStatus_MSG_STATUS_UNKNOWN), int32(types.MsgStatusUnknown))
	require.Equal(t, int32(db.MsgStatus_MSG_STATUS_SUBMITTED), int32(types.MsgStatusSubmitted))
	require.Equal(t, int32(db.MsgStatus_MSG_STATUS_MINTED), int32(types.MsgStatusMinted))
	require.Len(t, db.MsgStatus_name, 3)
}

func TestMsgEquals(t *testing.T) {
	t.Parallel()

	msg := testutil.RandMsg()
	same := msg
	diff := testutil.RandMsg()

	require.True(t, msg.Equals(same))
	require.False(t, msg.Equals(diff))
}
