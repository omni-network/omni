// Package targets defines list of targets supported by Omni's v1 solver. Targets are
// restricted to reduce attack surface area, and keep order flow predictable.
// Targets restriction will be removed / lessened in future versions.
package targets

import (
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

type Target struct {
	Name      string
	Addresses func(chainID uint64) map[common.Address]bool
}

var (
	// Symbiotic testnet.
	SymbioticSepoliaWSTETHVault1 = addr("0x77F170Dcd0439c0057055a6D7e5A1Eb9c48cCD2a")
	SymbioticSepoliaWSTETHVault2 = addr("0x1BAe55e4774372F6181DaAaB4Ca197A8D9CC06Dd")
	SymbioticSepoliaWSTETHVault3 = addr("0x6415D3B5fc615D4a00C71f4044dEc24C141EBFf8")
	SymbioticHoleskyWSTETHVault1 = addr("0xd88dDf98fE4d161a66FB836bee4Ca469eb0E4a75")
	SymbioticHoleskyWSTETHVault2 = addr("0xa4c81649c79f8378a4409178E758B839F1d57a54")

	// Eigen testnet.
	EigenHoleskyStrategyManager = addr("0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6")

	// Eigen mainnet.
	EigenMainnetStrategyManager = addr("0x858646372CC42E1A627fcE94aa7A7033e7CF075A")

	// targetsRestricted maps each network to whether targets should be restricted to the allowed set.
	targetsRestricted = map[netconf.ID]bool{
		netconf.Staging: true,
		netconf.Omega:   true,
		netconf.Mainnet: true,
	}

	targets = []Target{
		{
			Name: "Symbiotic",
			Addresses: networkChainAddrs(map[uint64]map[common.Address]bool{
				evmchain.IDSepolia: set(
					SymbioticSepoliaWSTETHVault1,
					SymbioticSepoliaWSTETHVault2,
					SymbioticSepoliaWSTETHVault3,
				),
				evmchain.IDHolesky: set(
					SymbioticHoleskyWSTETHVault1,
					SymbioticHoleskyWSTETHVault2,
				),
			}),
		},
		{
			Name: "Eigen",
			Addresses: networkChainAddrs(map[uint64]map[common.Address]bool{
				evmchain.IDHolesky:  set(EigenHoleskyStrategyManager),
				evmchain.IDEthereum: set(EigenMainnetStrategyManager),
			}),
		},
		{
			Name: "OmniStaking",
			Addresses: func(uint64) map[common.Address]bool {
				return map[common.Address]bool{
					common.HexToAddress(predeploys.Staking): true,
				}
			},
		},
	}
)

func networkChainAddrs(m map[uint64]map[common.Address]bool) func(uint64) map[common.Address]bool {
	return func(chainID uint64) map[common.Address]bool {
		return m[chainID]
	}
}

// IsRestricted returns true if the given network restricts targets.
func IsRestricted(network netconf.ID) bool {
	return targetsRestricted[network]
}

// Get returns the allowed target for the given chain and address.
func Get(chainID uint64, target common.Address) (Target, bool) {
	for _, t := range targets {
		if _, ok := t.Addresses(chainID)[target]; ok {
			return t, true
		}
	}

	return Target{}, false
}

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}

func set(addrs ...common.Address) map[common.Address]bool {
	s := make(map[common.Address]bool)
	for _, addr := range addrs {
		s[addr] = true
	}

	return s
}
