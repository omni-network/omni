// Package defines list of targets supported by Omni's v1 solver. Targets are
// restricted to reduce attack surface area, and keep order flow predictable.
// Targets restriction will be removed / lessened in future versions.
package targets

import (
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

var (
	SymbioticSepoliaWSTETHVault1 = common.HexToAddress("0x77F170Dcd0439c0057055a6D7e5A1Eb9c48cCD2a")
	SymbioticSepoliaWSTETHVault2 = common.HexToAddress("0x1BAe55e4774372F6181DaAaB4Ca197A8D9CC06Dd")
	SymbioticSepoliaWSTETHVault3 = common.HexToAddress("0x6415D3B5fc615D4a00C71f4044dEc24C141EBFf8")
	SymbioticHoleskyWSTETHVault1 = common.HexToAddress("0xd88dDf98fE4d161a66FB836bee4Ca469eb0E4a75")
	SymbioticHoleskyWSTETHVault2 = common.HexToAddress("0xa4c81649c79f8378a4409178E758B839F1d57a54")
	OmniStaking                  = common.HexToAddress(predeploys.Staking)

	// targetsRestricted maps each network to whether targets should be restricted to the allowed set.
	targetsRestricted = map[netconf.ID]bool{
		netconf.Staging: true,
		netconf.Omega:   true,
		netconf.Mainnet: true,
	}

	// allowedTargets maps chain id to a set of allowed target addresses.
	allowedTargets = map[uint64]map[common.Address]bool{
		evmchain.IDSepolia: {
			SymbioticSepoliaWSTETHVault1: true,
			SymbioticSepoliaWSTETHVault2: true,
			SymbioticSepoliaWSTETHVault3: true,
		},
		evmchain.IDHolesky: {
			SymbioticHoleskyWSTETHVault1: true,
			SymbioticHoleskyWSTETHVault2: true,
		},
		evmchain.IDOmniStaging: {
			OmniStaking: true,
		},
		evmchain.IDOmniOmega: {
			OmniStaking: true,
		},
		evmchain.IDOmniMainnet: {
			OmniStaking: true,
		},
	}
)

// IsRestricted returns true if the given network restricts targets.
func IsRestricted(network netconf.ID) bool {
	return targetsRestricted[network]
}

// IsAllowedTarget returns true if the target address is allowed on the given chain.
func IsAllowedTarget(chainID uint64, target common.Address) bool {
	return allowedTargets[chainID][target]
}
