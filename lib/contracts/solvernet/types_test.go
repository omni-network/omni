package solvernet_test

import (
	"testing"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/require"
)

func TestCallBindings(t *testing.T) {
	t.Parallel()

	calls := []solvernet.Call{
		{
			Value:  umath.Ether,
			Target: common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			// selector + params
			Data: hexutil.MustDecode("0x70a08231000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc7"),
		},
		{
			Value:  umath.Ether,
			Target: common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			// just selector
			Data: hexutil.MustDecode("0x70a08231"),
		},
		{
			Value:  umath.Ether,
			Target: common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			// no calldata
			Data: nil,
		},
		{
			// nil value
			Value:  nil,
			Target: common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			Data:   nil,
		},
	}

	bindings := solvernet.CallsToBindings(calls)

	require.Len(t, bindings, 4)

	// full calldata
	require.Equal(t, calls[0].Value, bindings[0].Value)
	require.Equal(t, calls[0].Target, bindings[0].Target)
	require.Equal(t, [4]byte(calls[0].Data[:4]), bindings[0].Selector)
	require.Equal(t, calls[0].Data[4:], bindings[0].Params)

	// just selector
	require.Equal(t, calls[1].Value, bindings[1].Value)
	require.Equal(t, calls[1].Target, bindings[1].Target)
	require.Equal(t, [4]byte(calls[1].Data[:4]), bindings[1].Selector)
	require.Equal(t, []byte(nil), bindings[1].Params)

	// no calldata
	require.Equal(t, calls[2].Value, bindings[2].Value)
	require.Equal(t, calls[2].Target, bindings[2].Target)
	require.Equal(t, [4]byte{}, bindings[2].Selector)
	require.Equal(t, []byte(nil), bindings[2].Params)

	// nil value
	require.Equal(t, calls[3].Value, bindings[3].Value)
	require.Equal(t, calls[3].Target, bindings[3].Target)
	require.Equal(t, [4]byte{}, bindings[3].Selector)
	require.Equal(t, []byte(nil), bindings[3].Params)
}
