package create3_test

import (
	"testing"

	"github.com/omni-network/omni/lib/contracts/create3"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestHashSalt(t *testing.T) {
	t.Parallel()

	hash := create3.HashSalt("eip1967.proxy.implementation")

	require.Equal(
		t,
		// keccak-256 hash of "eip1967.proxy.implementation"
		"360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd",
		common.Bytes2Hex(hash[:]),
	)
}
