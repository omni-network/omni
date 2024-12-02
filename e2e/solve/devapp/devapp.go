package devapp

import (
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type App struct {
	L1Vault common.Address
	L1Token common.Address
	L2Token common.Address
	L1      evmchain.Metadata
	L2      evmchain.Metadata
}

var (
	vaultABI     = mustGetABI(bindings.MockVaultMetaData)
	tokenABI     = mustGetABI(bindings.MockTokenMetaData)
	vaultDeposit = mustGetMethod(vaultABI, "deposit")
	mockL1       = mustChainMeta(evmchain.IDMockL1)
	mockL2       = mustChainMeta(evmchain.IDMockL2)

	create3Factory = contracts.Create3Factory(netconf.Devnet)
	deployer       = eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer)
	manager        = eoa.MustAddress(netconf.Devnet, eoa.RoleManager)

	static = App{
		L1Vault: create3.Address(create3Factory, l1VaultSalt, deployer),
		L1Token: create3.Address(create3Factory, l1TokenSalt, deployer),
		L2Token: create3.Address(create3Factory, l2TokenSalt, deployer),
		L1:      mockL1,
		L2:      mockL2,
	}
)

func GetApp() App {
	return static
}

func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

func mustGetMethod(abi *abi.ABI, name string) abi.Method {
	method, ok := abi.Methods[name]
	if !ok {
		panic(errors.New("missing method", "name", name))
	}

	return method
}

func mustChainMeta(chainID uint64) evmchain.Metadata {
	meta, ok := evmchain.MetadataByID(chainID)
	if !ok {
		panic(errors.New("missing chain meta", "chain_id", chainID))
	}

	return meta
}
