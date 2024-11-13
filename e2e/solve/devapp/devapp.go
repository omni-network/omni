package devapp

import (
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type App struct {
	L1Vault common.Address
	L1Token common.Address
	L2Token common.Address
}

var (
	vaultABI     = mustGetABI(bindings.MockVaultMetaData)
	tokenABI     = mustGetABI(bindings.MockTokenMetaData)
	vaultDeposit = mustGetMethod(vaultABI, "deposit")

	// static is the static devnt app instance.
	static = App{
		L1Vault: create3.Address(create3Factory, l1VaultSalt, deployer),
		L1Token: create3.Address(create3Factory, l1TokenSalt, deployer),
		L2Token: create3.Address(create3Factory, l2TokenSalt, deployer),
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
