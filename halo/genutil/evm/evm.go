package evm

import (
	"math/big"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

var eth1k = math.NewInt(1000).MulRaw(params.Ether).BigInt()
var eth1m = math.NewInt(1000000).MulRaw(params.Ether).BigInt()

func newUint64(val uint64) *uint64 { return &val }

// MakeGenesis returns a genesis block for a development chain.
// See geth reference: https://github.com/ethereum/go-ethereum/blob/master/core/genesis.go#L564
func MakeGenesis(network netconf.ID) (core.Genesis, error) {
	predeps, err := predeploys.Alloc(network)
	if err != nil {
		return core.Genesis{}, errors.Wrap(err, "predeploys")
	}

	allocs := mergeAllocs(precompilesAlloc(), predeps)

	if network.IsEphemeral() {
		allocs = mergeAllocs(allocs, stagingPrefundAlloc())
	} else if network == netconf.Omega {
		allocs = mergeAllocs(allocs, testnetPrefundAlloc())
	} else {
		return core.Genesis{}, errors.New("unsupported network", "network", network.String())
	}

	return core.Genesis{
		Config:     defaultChainConfig(network),
		GasLimit:   miner.DefaultConfig.GasCeil,
		BaseFee:    big.NewInt(params.InitialBaseFee),
		Difficulty: big.NewInt(0),
		Alloc:      allocs,
	}, nil
}

// defaultChainConfig returns the default chain config for a network.
// See geth reference: https://github.com/ethereum/go-ethereum/blob/master/params/config.go#L65
func defaultChainConfig(network netconf.ID) *params.ChainConfig {
	return &params.ChainConfig{
		ChainID:                       big.NewInt(int64(network.Static().OmniExecutionChainID)),
		HomesteadBlock:                big.NewInt(0),
		EIP150Block:                   big.NewInt(0),
		EIP155Block:                   big.NewInt(0),
		EIP158Block:                   big.NewInt(0),
		ByzantiumBlock:                big.NewInt(0),
		ConstantinopleBlock:           big.NewInt(0),
		PetersburgBlock:               big.NewInt(0),
		IstanbulBlock:                 big.NewInt(0),
		MuirGlacierBlock:              big.NewInt(0),
		BerlinBlock:                   big.NewInt(0),
		LondonBlock:                   big.NewInt(0),
		ArrowGlacierBlock:             big.NewInt(0),
		GrayGlacierBlock:              big.NewInt(0),
		ShanghaiTime:                  newUint64(0),
		CancunTime:                    newUint64(0),
		TerminalTotalDifficulty:       big.NewInt(0),
		TerminalTotalDifficultyPassed: true,
	}
}

// precompilesAlloc returns allocs for precompiled contracts
// precompile balances are set to 1 as a performance optimization, as done in geth.
func precompilesAlloc() types.GenesisAlloc {
	return types.GenesisAlloc{
		common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
		common.BytesToAddress([]byte{9}): {Balance: big.NewInt(1)}, // BLAKE2b
	}
}

// devPrefundAlloc returns allocs for pre-funded geth dev accounts.
func stagingPrefundAlloc() types.GenesisAlloc {
	return types.GenesisAlloc{
		// anvil pre-funded accounts
		anvil.DevAccount0(): {Balance: eth1m},
		anvil.DevAccount1(): {Balance: eth1m},
		anvil.DevAccount2(): {Balance: eth1m},
		anvil.DevAccount3(): {Balance: eth1m},
		anvil.DevAccount4(): {Balance: eth1m},
		anvil.DevAccount5(): {Balance: eth1m},
		anvil.DevAccount6(): {Balance: eth1m},
		anvil.DevAccount7(): {Balance: eth1m},
		anvil.DevAccount8(): {Balance: eth1m},
		anvil.DevAccount9(): {Balance: eth1m},

		// Relayer and Monitor EOAs
		eoa.MustAddress(netconf.Staging, eoa.RoleMonitor): {Balance: eth1m},
		eoa.MustAddress(netconf.Staging, eoa.RoleRelayer): {Balance: eth1m},

		// team ops accounts
		common.HexToAddress("0xfE921e06Ed0a22c035b4aCFF0A5D3a434A330c96"): {Balance: eth1m}, // dev relayer (local)
		common.HexToAddress("0xfC9D554D69DdCfC0A731b2DC64550177b0723bE5"): {Balance: eth1m}, // dev deployer (local)
		common.HexToAddress("0x7a6cF389082dc698285474976d7C75CAdE08ab7e"): {Balance: eth1m}, // fb: dev
		common.HexToAddress("0xC8103859Ac7CB547d70307EdeF1A2319FC305fdC"): {Balance: eth1m}, // fb: create3-deployer
		common.HexToAddress("0x274c4B3e5d27A65196d63964532366872F81D261"): {Balance: eth1m}, // fb: deployer
		common.HexToAddress("0x4891925c4f13A34FC26453FD168Db80aF3273014"): {Balance: eth1m}, // fb: owner
	}
}

func testnetPrefundAlloc() types.GenesisAlloc {
	return types.GenesisAlloc{
		// team ops accounts
		common.HexToAddress("0x7a6cF389082dc698285474976d7C75CAdE08ab7e"): {Balance: eth1m}, // fb: dev
		common.HexToAddress("0xeC5134556da0797A5C5cD51DD622b689Cac97Fe9"): {Balance: eth1k}, // fb: create3-deployer
		common.HexToAddress("0x0CdCc644158b7D03f40197f55454dc7a11Bd92c1"): {Balance: eth1k}, // fb: deployer
		common.HexToAddress("0xEAD625eB2011394cdD739E91Bf9D51A7169C22F5"): {Balance: eth1k}, // fb: owner
		common.HexToAddress("0xf63316AA39fEc9D2109AB0D9c7B1eE3a6F60AEA4"): {Balance: eth1k}, // fb: funder
		common.HexToAddress("0xfE921e06Ed0a22c035b4aCFF0A5D3a434A330c96"): {Balance: eth1k}, // dev: relayer

		// Faucet
		common.HexToAddress("0xAd5c6124305AB7F0101eC2e3BC9C2A5ec06E3594"): {Balance: eth1m},
		common.HexToAddress("0xC4C6ABeDf3C585C5efD755CB6591c92aa6E66079"): {Balance: eth1m},
		common.HexToAddress("0x7aAA75265052BFCe2A4910a4f50B68939E157fBD"): {Balance: eth1m},
		common.HexToAddress("0x815BA064f72667B11da226201991A2106C6f4ae0"): {Balance: eth1m},
		common.HexToAddress("0x78b223A22000f65033E1422A623cC3d6B566c3E8"): {Balance: eth1m},
		common.HexToAddress("0xAE2927bFEBF272a74340FC99D23d002035e5a0EA"): {Balance: eth1m},
		common.HexToAddress("0x64BC5a64713d4642c38a4e6521b0Fa6F293Ed85b"): {Balance: eth1m},
		common.HexToAddress("0x4A5EB30AFED3ED1A5eAB5bb95D5e6769fF1bc44E"): {Balance: eth1m},
		common.HexToAddress("0x24D2f0e25e1a321B1dFe2fAB5936907047CEfc97"): {Balance: eth1m},
		common.HexToAddress("0x371A5561c150C1ac051F3C190e4072f56B251cE6"): {Balance: eth1m},

		// Relayer and Monitor EOAs
		eoa.MustAddress(netconf.Omega, eoa.RoleMonitor): {Balance: eth1m},
		eoa.MustAddress(netconf.Omega, eoa.RoleRelayer): {Balance: eth1m},

		// team dev accounts
		common.HexToAddress("0xFbD05C5dD1c09970D30Ad8BBf29BC34283E84E20"): {Balance: eth1m}, // corver
		common.HexToAddress("0xe3481474b23f88a8917DbcB4cBC55Efcf0f68CC7"): {Balance: eth1m}, // kevin
		common.HexToAddress("0x38E2a3FC1923767F74d2308a529a353e91763EBF"): {Balance: eth1m}, // tt
		common.HexToAddress("0x5816EBCe5421c85B952F1f193980b96462B296De"): {Balance: eth1m}, // fab
		common.HexToAddress("0xDEdDf2DA39E0E39469a28F5A0392DcFbe40323de"): {Balance: eth1m}, // chase
		common.HexToAddress("0xb95512856C7044431E300C9b72C90297529B53DC"): {Balance: eth1m}, // austin
		common.HexToAddress("0x12444cDeD3BC994684D4Dc109240a22B8AC64f7C"): {Balance: eth1m}, // graham
		common.HexToAddress("0x29f26d43B2639aa8C6E99478C55a8645aD466766"): {Balance: eth1m}, // mark
		common.HexToAddress("0xEa64ab3af247d241E5389D1eE1aAB46262753906"): {Balance: eth1m}, // aayush

		// TODO: add validators, relayer
	}
}

// mergeAllocs merges multiple allocs into one.
func mergeAllocs(allocs ...types.GenesisAlloc) types.GenesisAlloc {
	merged := make(types.GenesisAlloc)
	for _, alloc := range allocs {
		for addr, account := range alloc {
			merged[addr] = account
		}
	}

	return merged
}
