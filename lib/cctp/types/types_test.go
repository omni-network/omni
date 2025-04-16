package types_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/testutil"
	"github.com/omni-network/omni/lib/cctp/types"

	"github.com/ethereum/go-ethereum/common"

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

func TestMsgDiff(t *testing.T) {
	t.Parallel()

	base := testutil.RandMsg()

	tests := []struct {
		name     string
		modify   func(msg types.MsgSendUSDC) types.MsgSendUSDC
		expected func(base, modified types.MsgSendUSDC) map[string]string
	}{
		{
			name: "no differences",
			modify: func(msg types.MsgSendUSDC) types.MsgSendUSDC {
				return msg
			},
			expected: func(_, _ types.MsgSendUSDC) map[string]string {
				return map[string]string{}
			},
		},
		{
			name: "different tx hash",
			modify: func(msg types.MsgSendUSDC) types.MsgSendUSDC {
				msg.TxHash = common.BytesToHash(testutil.RandBytes(32))
				return msg
			},
			expected: func(base, modified types.MsgSendUSDC) map[string]string {
				return map[string]string{
					"tx_hash": fmt.Sprintf("%v != %v", base.TxHash, modified.TxHash),
				}
			},
		},
		{
			name: "different block height",
			modify: func(msg types.MsgSendUSDC) types.MsgSendUSDC {
				msg.BlockHeight++
				return msg
			},
			expected: func(base, modified types.MsgSendUSDC) map[string]string {
				return map[string]string{
					"block_height": fmt.Sprintf("%v != %v", base.BlockHeight, modified.BlockHeight),
				}
			},
		},
		{
			name: "different message hash",
			modify: func(msg types.MsgSendUSDC) types.MsgSendUSDC {
				msg.MessageHash = common.BytesToHash(testutil.RandBytes(32))
				return msg
			},
			expected: func(base, modified types.MsgSendUSDC) map[string]string {
				return map[string]string{
					"message_hash": fmt.Sprintf("%v != %v", base.MessageHash, modified.MessageHash),
				}
			},
		},
		{
			name: "different source chain",
			modify: func(msg types.MsgSendUSDC) types.MsgSendUSDC {
				msg.SrcChainID++
				return msg
			},
			expected: func(base, modified types.MsgSendUSDC) map[string]string {
				return map[string]string{
					"src_chain_id": fmt.Sprintf("%v != %v", base.SrcChainID, modified.SrcChainID),
				}
			},
		},
		{
			name: "different destination chain",
			modify: func(msg types.MsgSendUSDC) types.MsgSendUSDC {
				msg.DestChainID++
				return msg
			},
			expected: func(base, modified types.MsgSendUSDC) map[string]string {
				return map[string]string{
					"dest_chain_id": fmt.Sprintf("%v != %v", base.DestChainID, modified.DestChainID),
				}
			},
		},
		{
			name: "different amount",
			modify: func(msg types.MsgSendUSDC) types.MsgSendUSDC {
				msg.Amount = big.NewInt(100)
				return msg
			},
			expected: func(base, modified types.MsgSendUSDC) map[string]string {
				return map[string]string{
					"amount": fmt.Sprintf("%v != %v", base.Amount, modified.Amount),
				}
			},
		},
		{
			name: "different recipient",
			modify: func(msg types.MsgSendUSDC) types.MsgSendUSDC {
				msg.Recipient = common.BytesToAddress(testutil.RandBytes(20))
				return msg
			},
			expected: func(base, modified types.MsgSendUSDC) map[string]string {
				return map[string]string{
					"recipient": fmt.Sprintf("%v != %v", base.Recipient, modified.Recipient),
				}
			},
		},
		{
			name: "different status",
			modify: func(msg types.MsgSendUSDC) types.MsgSendUSDC {
				msg.Status = types.MsgStatusMinted
				return msg
			},
			expected: func(base, modified types.MsgSendUSDC) map[string]string {
				return map[string]string{
					"status": fmt.Sprintf("%v != %v", base.Status, modified.Status),
				}
			},
		},
		{
			name: "different message bytes",
			modify: func(msg types.MsgSendUSDC) types.MsgSendUSDC {
				msg.MessageBytes = testutil.RandBytes(32)
				return msg
			},
			expected: func(base, modified types.MsgSendUSDC) map[string]string {
				return map[string]string{
					"message_bytes": fmt.Sprintf("%v != %v", abbrevBz(base.MessageBytes), abbrevBz(modified.MessageBytes)),
				}
			},
		},
		{
			name: "multiple differences",
			modify: func(msg types.MsgSendUSDC) types.MsgSendUSDC {
				msg.TxHash = common.BytesToHash(testutil.RandBytes(32))
				msg.BlockHeight++
				msg.Status = types.MsgStatusMinted

				return msg
			},
			expected: func(base, modified types.MsgSendUSDC) map[string]string {
				return map[string]string{
					"tx_hash":      fmt.Sprintf("%v != %v", base.TxHash, modified.TxHash),
					"block_height": fmt.Sprintf("%v != %v", base.BlockHeight, modified.BlockHeight),
					"status":       fmt.Sprintf("%v != %v", base.Status, modified.Status),
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			modified := tt.modify(base)
			diff := base.Diff(modified)
			expected := tt.expected(base, modified)
			require.Equal(t, expected, diff)
		})
	}
}

func abbrevBz(b []byte) string {
	if len(b) <= 8 {
		return "0x" + common.Bytes2Hex(b)
	}

	return "0x" + common.Bytes2Hex(b[:4]) + "..." + common.Bytes2Hex(b[len(b)-4:])
}
