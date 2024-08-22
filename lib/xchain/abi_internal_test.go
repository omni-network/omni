package xchain

import (
	"encoding/hex"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestEncodeMsg(t *testing.T) {
	t.Parallel()

	msg := Msg{
		MsgID: MsgID{
			StreamID: StreamID{
				SourceChainID: 1,
				DestChainID:   2,
				ShardID:       1,
			},
			StreamOffset: 100,
		},
		SourceMsgSender: common.HexToAddress("0xcbbc5Da52ea2728279560Dca8f4ec08d5F829985"),
		DestAddress:     common.HexToAddress("0x9CC971e84FE5d09d0967f15AE05dfd553C5A1FA6"),
		Data:            common.Hex2Bytes("d09de08a"),
		DestGasLimit:    250_000,
		TxHash:          common.Hash{},
	}

	packed, err := encodeMsg(msg)
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, []byte(hex.EncodeToString(packed)))
}

func TestEncodeHeader(t *testing.T) {
	t.Parallel()

	aHeader := AttestHeader{
		ConsensusChainID: 1,
		ChainVersion:     ChainVersion{ID: 1, ConfLevel: ConfLatest},
		AttestOffset:     100,
	}
	bHeader := BlockHeader{
		ChainID:     1,
		BlockHeight: 99,
		BlockHash:   common.HexToHash("0x412d62a6a3115ab5a0e0cae9d63082ff8dfb002a98cc889d06dc986a9461586b"),
	}

	packed, err := encodeSubmissionHeader(aHeader, bHeader)
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, []byte(hex.EncodeToString(packed)))
}

func TestSubmissionToFromBinding(t *testing.T) {
	t.Parallel()
	var sub Submission
	fuzz.New().NilChance(0).Fuzz(&sub)
	sub.AttHeader.ChainVersion.ID = sub.BlockHeader.ChainID // Align headers

	xsub := SubmissionToBinding(sub)
	reversedSub := SubmissionFromBinding(xsub, sub.DestChainID)

	// Zero TxHash and ChainID for comparison since they aren't translated.
	for i := range sub.Msgs {
		sub.Msgs[i].TxHash = common.Hash{}
		sub.Msgs[i].SourceChainID = 0
	}

	// Zero BlockHeight as we only submit AttestOffset
	sub.BlockHeader.BlockHeight = 0

	require.Equal(t, sub, reversedSub)
}

func TestXSubmitEncodeDecode(t *testing.T) {
	t.Parallel()
	var sub bindings.XSubmission
	fuzz.New().NilChance(0).Fuzz(&sub)

	calldata, err := EncodeXSubmit(sub)
	require.NoError(t, err)

	decoded, err := DecodeXSubmit(calldata)
	require.NoError(t, err)
	require.Equal(t, sub, decoded)
}
