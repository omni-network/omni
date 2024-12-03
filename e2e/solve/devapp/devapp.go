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

	mockL1      = mustChainMeta(evmchain.IDMockL1)
	mockL2      = mustChainMeta(evmchain.IDMockL2)
	holesky     = mustChainMeta(evmchain.IDHolesky)
	baseSepolia = mustChainMeta(evmchain.IDBaseSepolia)
)

func GetApp(network netconf.ID) (App, error) {
	if !network.IsEphemeral() {
		return App{}, errors.New("only ephemeral networks")
	}

	deployer := eoa.MustAddress(network, eoa.RoleDeployer)
	create3Factory := contracts.Create3Factory(network)

	app := App{
		L1Vault: create3.Address(create3Factory, l1VaultSalt, deployer),
		L1Token: create3.Address(create3Factory, l1TokenSalt, deployer),
		L2Token: create3.Address(create3Factory, l2TokenSalt, deployer),
	}

	if network == netconf.Devnet {
		app.L1 = mockL1
		app.L2 = mockL2
	}

	if network == netconf.Staging {
		app.L1 = holesky
		app.L2 = baseSepolia
	}

	return app, nil
}

func MustGetApp(network netconf.ID) App {
	app, err := GetApp(network)
	if err != nil {
		panic(err)
	}

	return app
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
