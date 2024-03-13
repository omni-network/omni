package evm

import (
	"math/big"

	"github.com/omni-network/omni/halo/genutil/evm/predeploys"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

func newUint64(val uint64) *uint64 { return &val }

func DefaultDevConfig() params.ChainConfig {
	return params.ChainConfig{
		ChainID:                       big.NewInt(1), // TODO: choose new dev chain id
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
		TerminalTotalDifficulty:       big.NewInt(0),
		TerminalTotalDifficultyPassed: true,
	}
}

// precompilesAlloc returns allocs for precompiled contracts
// TODO: this matches go-ethereum's precompiles, but we should understand why balances are set to 1.
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
func devPrefundAlloc() types.GenesisAlloc {
	amt := big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(params.Ether))
	return types.GenesisAlloc{
		common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"): {Balance: amt},
		common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"): {Balance: amt},
		common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"): {Balance: amt},
		common.HexToAddress("0x90F79bf6EB2c4f870365E785982E1f101E93b906"): {Balance: amt},
		common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"): {Balance: amt},
		common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc"): {Balance: amt},
		common.HexToAddress("0x976EA74026E726554dB657fA54763abd0C3a0aa9"): {Balance: amt},
		common.HexToAddress("0x14dC79964da2C08b23698B3D3cc7Ca32193d9955"): {Balance: amt},
		common.HexToAddress("0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f"): {Balance: amt},
		common.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"): {Balance: amt},
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

// MakeDevGenesis returns a genesis block for a development chain
// TODO: add staging / testnet / mainnet genesis.
func MakeDevGenesis() core.Genesis {
	config := DefaultDevConfig()

	return core.Genesis{
		Config:     &config,
		GasLimit:   params.GenesisGasLimit,
		BaseFee:    big.NewInt(params.InitialBaseFee),
		Difficulty: big.NewInt(1),
		Alloc: mergeAllocs(
			precompilesAlloc(),
			predeploys.Alloc(),
			devPrefundAlloc(),
		),
	}
}
