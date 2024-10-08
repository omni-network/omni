package cast_test

import (
	"crypto/rand"
	"testing"

	"github.com/omni-network/omni/lib/cast"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestCast(t *testing.T) {
	t.Parallel()

	slice := make([]byte, 100)
	var resp common.Address
	_, _ = rand.Read(resp[:])

	// 8
	_, err := cast.Array8(slice)
	require.Error(t, err)
	_, err = cast.Array8(slice[:1])
	require.Error(t, err)
	resp8, err := cast.Array8(slice[:8])
	require.NoError(t, err)
	require.Equal(t, slice[:8], resp8[:])

	// 20
	_, err = cast.Array20(slice)
	require.Error(t, err)
	_, err = cast.Array20(slice[:1])
	require.Error(t, err)
	resp20, err := cast.Array20(slice[:20])
	require.NoError(t, err)
	require.Equal(t, slice[:20], resp20[:])
	addr, err := cast.EthAddress(slice[:20])
	require.NoError(t, err)
	require.EqualValues(t, resp20, addr)

	// 32
	_, err = cast.Array32(slice)
	require.Error(t, err)
	_, err = cast.Array32(slice[:1])
	require.Error(t, err)
	resp32, err := cast.Array32(slice[:32])
	require.NoError(t, err)
	require.Equal(t, slice[:32], resp32[:])
	hash, err := cast.EthHash(slice[:32])
	require.NoError(t, err)
	require.EqualValues(t, resp32, hash)

	// 65
	_, err = cast.Array65(slice)
	require.Error(t, err)
	_, err = cast.Array65(slice[:1])
	require.Error(t, err)
	resp65, err := cast.Array65(slice[:65])
	require.NoError(t, err)
	require.Equal(t, slice[:65], resp65[:])
}
