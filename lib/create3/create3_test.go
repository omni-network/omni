package create3_test

import (
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/netconf"

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

func TestAddress(t *testing.T) {
	t.Parallel()

	// test case is Devent proxy admin deployment

	factory := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	expected := common.HexToAddress("0xd8dc3f2817F4d87200443FBaEdE5ab8D5d742465")
	require.Equal(t, expected, create3.Address(factory, "devnet-proxy-admin", eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer)))
}
