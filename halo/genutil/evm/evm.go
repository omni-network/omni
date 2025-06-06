package evm

import (
	"encoding/json"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/params"
)

var (
	eth1k = bi.Ether(1_000)
	eth1m = bi.Ether(1_000_000)
)

func newUint64(val uint64) *uint64 { return &val }

// MakeGenesis returns a genesis block for a development chain.
// See geth reference: https://github.com/ethereum/go-ethereum/blob/master/core/genesis.go#L564
func MakeGenesis(network netconf.ID) (core.Genesis, error) {
	predeps, err := predeploys.Alloc(network)
	if err != nil {
		return core.Genesis{}, errors.Wrap(err, "predeploys")
	}

	prefunds, err := PrefundAlloc(network)
	if err != nil {
		return core.Genesis{}, errors.Wrap(err, "prefund alloc")
	}

	return core.Genesis{
		Config:     DefaultChainConfig(network),
		GasLimit:   miner.DefaultConfig.GasCeil,
		BaseFee:    bi.N(params.InitialBaseFee),
		Difficulty: bi.Zero(),
		Alloc:      mergeAllocs(PrecompilesAlloc(), predeps, prefunds),
	}, nil
}

// MarshallBackwardsCompatible marshals a genesis into a backwards compatible JSON format.
func MarshallBackwardsCompatible(genesis core.Genesis) ([]byte, error) {
	// Geth version 1.14.12 removed TerminalTotalDifficultyPassed,
	// but this is still required by older v1.14.* versions.
	backwardsConfig := struct {
		*params.ChainConfig                // Extend latest chain config
		TerminalTotalDifficultyPassed bool `json:"terminalTotalDifficultyPassed"`
	}{
		ChainConfig:                   genesis.Config,
		TerminalTotalDifficultyPassed: true,
	}
	configBz, err := json.MarshalIndent(backwardsConfig, "", "  ")
	if err != nil {
		return nil, errors.Wrap(err, "marshal backwards compatible config")
	}

	// Marshal new genesis
	bz, err := json.MarshalIndent(genesis, "", "  ")
	if err != nil {
		return nil, errors.Wrap(err, "marshal genesis")
	}

	// Unmarshal into generic map
	temp1 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(bz, &temp1); err != nil {
		return nil, errors.Wrap(err, "unmarshal genesis")
	}

	// Replace config with backwards compatible config
	temp1["config"] = configBz

	// Marshal backwards compatible genesis
	bz, err = json.MarshalIndent(temp1, "", "  ")
	if err != nil {
		return nil, errors.Wrap(err, "marshal backwards compatible genesis")
	}

	return bz, nil
}

// DefaultChainConfig returns the default chain config for a network.
// See geth reference: https://github.com/ethereum/go-ethereum/blob/master/params/config.go#L65
func DefaultChainConfig(network netconf.ID) *params.ChainConfig {
	return &params.ChainConfig{
		ChainID:                 bi.N(network.Static().OmniExecutionChainID),
		HomesteadBlock:          bi.Zero(),
		EIP150Block:             bi.Zero(),
		EIP155Block:             bi.Zero(),
		EIP158Block:             bi.Zero(),
		ByzantiumBlock:          bi.Zero(),
		ConstantinopleBlock:     bi.Zero(),
		PetersburgBlock:         bi.Zero(),
		IstanbulBlock:           bi.Zero(),
		MuirGlacierBlock:        bi.Zero(),
		BerlinBlock:             bi.Zero(),
		LondonBlock:             bi.Zero(),
		ArrowGlacierBlock:       bi.Zero(),
		GrayGlacierBlock:        bi.Zero(),
		ShanghaiTime:            newUint64(0),
		CancunTime:              newUint64(0),
		TerminalTotalDifficulty: bi.Zero(),
		BlobScheduleConfig:      params.DefaultBlobSchedule,
	}
}

// precompilesAlloc returns allocs for precompiled contracts
// precompile balances are set to 1 as a performance optimization, as done in geth.
//
//nolint:forbidigo // Explicitly use BytesToAddress with left padding.
func PrecompilesAlloc() types.GenesisAlloc {
	return types.GenesisAlloc{
		common.BytesToAddress([]byte{1}): {Balance: bi.One()}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: bi.One()}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: bi.One()}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: bi.One()}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: bi.One()}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: bi.One()}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: bi.One()}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: bi.One()}, // ECPairing
		common.BytesToAddress([]byte{9}): {Balance: bi.One()}, // BLAKE2b
	}
}

func PrefundAlloc(network netconf.ID) (types.GenesisAlloc, error) {
	if network.IsEphemeral() {
		return ephemeralPrefundAlloc(network), nil
	} else if network == netconf.Omega {
		return omegaPrefundAlloc(), nil
	} else if network == netconf.Mainnet {
		return mainnetPrefundAllocs(), nil
	}

	return nil, errors.New("unsupported network", "network", network.String())
}

func ephemeralPrefundAlloc(network netconf.ID) types.GenesisAlloc {
	allocs := anvilPrefundAlloc()

	for _, role := range eoa.AllRoles() {
		allocs[eoa.MustAddress(network, role)] = types.Account{Balance: eth1m}
	}

	return allocs
}

func anvilPrefundAlloc() types.GenesisAlloc {
	resp := make(types.GenesisAlloc)
	for _, addr := range anvil.DevAccounts() { // anvil pre-funded accounts
		resp[addr] = types.Account{Balance: eth1m}
	}

	for addr := range anvil.ExternalAccounts { // external pre-funded accounts
		resp[addr] = types.Account{Balance: eth1m}
	}

	return resp
}

func omegaPrefundAlloc() types.GenesisAlloc {
	return types.GenesisAlloc{
		// team ops accounts
		common.HexToAddress("0x7a6cF389082dc698285474976d7C75CAdE08ab7e"): {Balance: eth1m}, // fb: dev
		common.HexToAddress("0xeC5134556da0797A5C5cD51DD622b689Cac97Fe9"): {Balance: eth1k}, // fb: create3-deployer
		common.HexToAddress("0x0CdCc644158b7D03f40197f55454dc7a11Bd92c1"): {Balance: eth1k}, // fb: deployer
		common.HexToAddress("0xEAD625eB2011394cdD739E91Bf9D51A7169C22F5"): {Balance: eth1k}, // fb: owner
		common.HexToAddress("0xf63316AA39fEc9D2109AB0D9c7B1eE3a6F60AEA4"): {Balance: eth1k}, // fb: hot
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
	}
}

// mainnetPrefundAllocs returns allocs for prefunded geth accounts on mainnet.
//
// mainnet prefunds are kept minimal, and are accounted for w.r.t OMNI's total
// supply via "genesis bridging". That is, decrementing the pre-funded amount
// from the native bridge's prefund balance, and transferring the same amount
// to the bridge contract on L1.
func mainnetPrefundAllocs() types.GenesisAlloc {
	allocs := make(types.GenesisAlloc)

	for _, role := range eoa.AllRoles() {
		fund, ok := eoa.GetFundThresholds(tokens.OMNI, netconf.Mainnet, role)
		if !ok {
			continue
		}

		acc := types.Account{Balance: fund.TargetBalance()}
		allocs[eoa.MustAddress(netconf.Mainnet, role)] = acc
	}

	return allocs
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
