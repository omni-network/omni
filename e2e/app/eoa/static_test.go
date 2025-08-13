package eoa_test

import (
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/manifests"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -run=TestThresholdReference -golden

var (
	fundTokens = []tokens.Asset{tokens.ETH, tokens.OMNI, tokens.MNT, tokens.HYPE}
)

func TestThresholdReference(t *testing.T) {
	t.Parallel()

	resp := make(map[netconf.ID]map[string]map[eoa.Role]map[string]string)
	for _, network := range []netconf.ID{netconf.Staging, netconf.Omega, netconf.Mainnet} {
		resp[network] = make(map[string]map[eoa.Role]map[string]string)
		for _, token := range fundTokens {
			resp[network][token.Symbol] = make(map[eoa.Role]map[string]string)
			for _, role := range eoa.AllRoles() {
				resp[network][token.Symbol][role] = make(map[string]string)
				thresholds, _ := eoa.GetFundThresholds(token, network, role)

				resp[network][token.Symbol][role]["target"] = token.FormatAmt(thresholds.TargetBalance())
				resp[network][token.Symbol][role]["min"] = token.FormatAmt(thresholds.MinBalance())
			}
		}
	}

	tutil.RequireGoldenJSON(t, resp, tutil.WithFilename("threshold_reference.json"))
}

func TestStatic(t *testing.T) {
	t.Parallel()
	for _, network := range []netconf.ID{netconf.Devnet, netconf.Staging, netconf.Omega, netconf.Mainnet} {
		chains, err := manifests.EVMChains(network)
		require.NoError(t, err)

		for _, chain := range chains {
			for _, role := range eoa.AllRoles() {
				if !shouldExist(role, network) {
					continue
				}

				acc, ok := eoa.AccountForRole(network, role)
				require.True(t, ok, "account not found: %s %s", network, role)
				require.NotZero(t, acc.Address)
				require.True(t, common.IsHexAddress(acc.Address.Hex()))

				thresholds, _ := eoa.GetFundThresholds(chain.NativeToken, network, acc.Role)

				require.NotPanics(t, func() {
					mini := thresholds.MinBalance()
					target := thresholds.TargetBalance()
					t.Logf("Thresholds: network=%s, role=%s, min=%s, target=%s",
						network, acc.Role, chain.NativeToken.FormatAmt(mini), chain.NativeToken.FormatAmt(target))
				})
			}
		}
	}
}

// TestMainnet makes sure the initial accounts used for deployment won't change in the future.
func TestMainnet(t *testing.T) {
	t.Parallel()
	fixed := map[eoa.Role]string{
		eoa.RoleManager:         "0xd09DD1126385877352d24B669Fd68f462200756E",
		eoa.RoleUpgrader:        "0xF8740c09f25E2cbF5C9b34Ef142ED7B343B42360",
		eoa.RoleCreate3Deployer: "0x992b9de7D42981B90A75C523842C01e27875b65B",
		eoa.RoleDeployer:        "0x9496Bf1Bd2Fa5BCba72062cC781cC97eA6930A13",
		eoa.RoleHot:             "0x8F609f4d58355539c48C98464E1e54ab2709aCfe",
		eoa.RoleCold:            "0x8b6b217572582C57616262F9cE02A951A1D1b951",
		eoa.RoleNomAuthority:    "0x9016516dcf1Bc94d24DA583Dfc31FEB87f852922",
	}

	n := netconf.Mainnet
	for role := range fixed {
		acc, ok := eoa.AccountForRole(n, role)
		require.True(t, ok, "account not found: %s %s", n, role)
		require.Equal(t, fixed[acc.Role], acc.Address.Hex())
	}
}

func shouldExist(role eoa.Role, id netconf.ID) bool {
	switch {
	case role == eoa.RoleTester && id == netconf.Mainnet: // RoleTester not supported on mainnet
		return false
	case role == eoa.RoleRedenomizer && id != netconf.Devnet: // Redenomizer only supported on devnet
		return false
	default:
		return true
	}
}
