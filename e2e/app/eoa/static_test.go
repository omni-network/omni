package eoa_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -run=TestDynamicThresholds -clean -golden

func TestDynamicThresholds(t *testing.T) {
	t.Parallel()

	type val struct {
		Target string
		Min    string
	}

	results := make(map[string]val)
	for _, network := range []netconf.ID{netconf.Staging, netconf.Omega, netconf.Mainnet} {
		for _, role := range []eoa.Role{eoa.RoleSafe, eoa.RoleHot} {
			for _, token := range []tokens.Token{tokens.ETH, tokens.OMNI} {
				thresholds, ok := eoa.GetFundThresholds(token, network, role)
				require.True(t, ok, "thresholds not found: %s %s %s", network, role, token)

				key := fmt.Sprintf("%s %s %s", network, role, token)
				results[key] = val{
					Target: etherStr(thresholds.TargetBalance()),
					Min:    etherStr(thresholds.MinBalance()),
				}
			}
		}
	}

	tutil.RequireGoldenJSON(t, results)
}

func TestStatic(t *testing.T) {
	t.Parallel()
	for _, chain := range evmchain.All() {
		for _, network := range []netconf.ID{netconf.Devnet, netconf.Staging, netconf.Omega, netconf.Mainnet} {
			for _, role := range eoa.AllRoles() {
				shouldExist := role != eoa.RoleTester || network != netconf.Mainnet // skip tester on mainnet

				acc, ok := eoa.AccountForRole(network, role)
				require.Equal(t, shouldExist, ok, "account not found: %s %s", network, role)
				if shouldExist {
					require.NotZero(t, acc.Address)
					require.True(t, common.IsHexAddress(acc.Address.Hex()))
				}

				thresholds, ok := eoa.GetFundThresholds(chain.NativeToken, network, acc.Role)
				require.Equal(t, shouldExist, ok, "thresholds not found")

				if shouldExist {
					require.NotPanics(t, func() {
						mini := thresholds.MinBalance()
						target := thresholds.TargetBalance()
						t.Logf("Thresholds: network=%s, role=%s, min=%s, target=%s",
							network, acc.Role, etherStr(mini), etherStr(target))
					})
				}
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
	}

	n := netconf.Mainnet
	for role := range fixed {
		acc, ok := eoa.AccountForRole(n, role)
		require.True(t, ok, "account not found: %s %s", n, role)
		require.Equal(t, fixed[acc.Role], acc.Address.Hex())
	}
}

func etherStr(amount *big.Int) string {
	b, _ := amount.Float64()
	b /= params.Ether

	return fmt.Sprintf("%.4f", b)
}
