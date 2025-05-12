package anchorinbox_test

import (
	"testing"

	"github.com/omni-network/omni/anchor/anchorinbox"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
)

func TestFindOrderStateAddress(t *testing.T) {
	t.Parallel()

	var zeroPub solana.PublicKey

	addr, bump, err := anchorinbox.FindOrderStateAddress(zeroPub)
	require.NoError(t, err)
	require.Equal(t, uint8(255), bump)
	require.Equal(t, solana.MustPublicKeyFromBase58("rGHaZ7FoTFpXVbadryjN25EgdYd6dKC1ESKUCrWjbgy"), addr)
}

func TestFindInboxStateAddress(t *testing.T) {
	t.Parallel()

	addr, bump, err := anchorinbox.FindInboxStateAddress()
	require.NoError(t, err)
	require.Equal(t, uint8(254), bump)
	require.Equal(t, solana.MustPublicKeyFromBase58("AY6Uhn4bacArkUYqsD5LAgpUehQ6Kq7qRbT3zzSZGV6r"), addr)
}
