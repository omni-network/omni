package predeploys_test

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/require"
)

var (
	// Namespace big.Ints.
	omniNamespace   = common.HexToAddress(predeploys.OmniNamepsace).Big()
	octaneNamespace = common.HexToAddress(predeploys.OctaneNamespace).Big()

	// Predeploy addresses.
	proxyAdmin     = common.HexToAddress(predeploys.ProxyAdmin)
	portalRegistry = common.HexToAddress(predeploys.PortalRegistry)
	omniBridge     = common.HexToAddress(predeploys.OmniBridgeNative)
	womni          = common.HexToAddress(predeploys.WOmni)
	staking        = common.HexToAddress(predeploys.Staking)
	slashing       = common.HexToAddress(predeploys.Slashing)

	// Predeploy bytecodes.
	proxyCode          = hexutil.MustDecode(bindings.TransparentUpgradeableProxyDeployedBytecode)
	proxyAdminCode     = hexutil.MustDecode(bindings.ProxyAdminDeployedBytecode)
	portalRegistryCode = hexutil.MustDecode(bindings.PortalRegistryDeployedBytecode)
	womniCode          = hexutil.MustDecode(bindings.WOmniDeployedBytecode)
	omniBridgeCode     = hexutil.MustDecode(bindings.OmniBridgeNativeDeployedBytecode)
	stakingCode        = hexutil.MustDecode(bindings.StakingDeployedBytecode)
	slashingCode       = hexutil.MustDecode(bindings.SlashingDeployedBytecode)
)

func TestAlloc(t *testing.T) {
	t.Parallel()

	admin, err := eoa.Admin(netconf.Staging)
	require.NoError(t, err)

	predeps, err := predeploys.Alloc(netconf.Staging, admin)
	require.NoError(t, err)

	// Check each namespace filled with proxies
	for _, namspace := range []*big.Int{omniNamespace, octaneNamespace} {
		for i := 0; i < predeploys.NamespaceSize; i++ {
			// add one, so that we don't set the namespace zero addr
			addr := address(namspace, i+1)

			alloc, ok := predeps[addr]
			require.Truef(t, ok, "proxy alloc")
			require.Equalf(t, hexutil.Encode(proxyCode), hexutil.Encode(alloc.Code), "proxy code mismatch")
			require.Equalf(t, common.HexToHash(predeploys.ProxyAdmin), alloc.Storage[common.HexToHash(predeploys.ProxyAdminSlot)], "proxy admin slot empty")
		}
	}

	// Check ProxyAdmin is predeployed
	alloc, ok := predeps[proxyAdmin]
	require.True(t, ok, "proxy admin alloc")
	require.Equal(t, proxyAdminCode, alloc.Code, "proxy admin code mismatch")

	// Check each predpeploy has implementation set
	tests := []struct {
		name      string
		predeploy common.Address
		code      []byte
	}{
		{"PortalRegistry", portalRegistry, portalRegistryCode},
		{"OmniBridge", omniBridge, omniBridgeCode},
		{"WOmni", womni, womniCode},
		{"Staking", staking, stakingCode},
		{"Slashing", slashing, slashingCode},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alloc, ok := predeps[tt.predeploy]
			require.True(t, ok, "proxy alloc")
			require.Equal(t, proxyCode, alloc.Code, "proxy code mismatch")

			impl, ok := alloc.Storage[common.HexToHash(predeploys.ProxyImplementationSlot)]
			require.True(t, ok, "impl addr")

			implAlloc, ok := predeps[common.BytesToAddress(impl.Bytes())]
			require.True(t, ok, "impl alloc")
			require.Equal(t, tt.code, implAlloc.Code, "impl code mismatch")
		})
	}
}

// address returns the address at the given index in the namespace.
func address(namespace *big.Int, i int) common.Address {
	return common.BigToAddress(new(big.Int).Add(namespace, big.NewInt(int64(i))))
}
