package predeploys

import (
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	// ProxyAdmin = "0x121E240000000000000000000000000000000001".
	OmniStake = "0x121E240000000000000000000000000000000002"
	// EthStakeInbox = "0x121E240000000000000000000000000000000003".
	OmniXChainRegistry = "0x121E240000000000000000000000000000000004"
)

func mustDecodeHex(hex string) []byte {
	b, err := hexutil.Decode(hex)
	if err != nil {
		panic(err)
	}

	return b
}

// Alloc returns the genesis allocs for the predeployed contracts, initializing code and storage.
func Alloc() types.GenesisAlloc {
	// TODO: Use TransparentUpgradeableProxy for all predeploys
	return types.GenesisAlloc{
		// NOTE:
		// - OmniStake has note has no immutables, and no storage so we can
		// 	 use the deployed bytecode as is, and do not need to initialize storage
		// - balance is encoded to null when not set, so we set to 1 (same as precompilesAlloc in evm.go)
		common.HexToAddress(OmniStake): {Balance: big.NewInt(1), Code: mustDecodeHex(bindings.OmniStakeDeployedBytecode)},
		// NOTE:
		// OmniXCahinRegistry has immutables, but does require initialized storage (to set the owner).
		// For now, we set it right after deploying the chain, permissionlessly.
		// We need to set storage at genesis.
		common.HexToAddress(OmniXChainRegistry): {Balance: big.NewInt(1), Code: mustDecodeHex(bindings.OmniXChainRegistryDeployedBytecode)},
	}
}
