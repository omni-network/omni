package testutil

import (
	crand "crypto/rand"
	"math/big"
	mrand "math/rand"
	"testing"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/cctp/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

// RandMsg returns a random tpes.MsgStatus.
func RandMsg() types.MsgSendUSDC {
	msgBz := RandBytes(32 * 10)
	msgHash := crypto.Keccak256Hash(msgBz)

	return types.MsgSendUSDC{
		TxHash:       common.BytesToHash(RandBytes(32)),
		BlockHeight:  mrand.Uint64(),
		SrcChainID:   uint64(mrand.Uint32()), // cctp uses uint32 domain ids
		DestChainID:  uint64(mrand.Uint32()),
		Amount:       RandBigInt(),
		MessageBytes: msgBz,
		MessageHash:  msgHash,
		Recipient:    RandAddr(),
		Status:       types.MsgStatusSubmitted,
	}
}

// RandBytes generates random bytes array of lenth n.
func RandBytes(n int) []byte {
	bz := make([]byte, n)

	_, err := crand.Read(bz)
	if err != nil {
		panic(err)
	}

	return bz
}

// ABIEncodeBytes encodes a byte slice into an ABI-encoded byte array.
func ABIEncodeBytes(bz []byte) []byte {
	typ, err := abi.NewType("bytes", "", nil)
	if err != nil {
		panic(err)
	}

	packed, err := abi.Arguments{{Type: typ}}.Pack(bz)
	if err != nil {
		panic(err)
	}

	return packed
}

// RandAddr returnsa random ETH address.
func RandAddr() common.Address {
	return cast.MustEthAddress(RandBytes(20))
}

// RandBigInt returns a random big int.
func RandBigInt() *big.Int {
	return big.NewInt(mrand.Int63())
}

// AssertMsgsEqual asserts that the expected messages are the same as the actual messages.
func AssertMsgsEqual(t *testing.T, expected, got []types.MsgSendUSDC) {
	t.Helper()

	require.Len(t, got, len(expected), "expected %d messages, got %d", len(expected), len(got))

	for _, e := range expected {
		require.Containsf(t, got, e, "expected message (tx_hash=%s, msg_hash=%s)", e.TxHash, e.MessageHash)
	}
}
