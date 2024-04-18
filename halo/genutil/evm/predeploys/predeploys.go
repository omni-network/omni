package predeploys

import (
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm/state"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/solc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	// Namespace - namespace of predeploys users will interact with.
	Namespace     = "0x121E240000000000000000000000000000000000"
	NamespaceSize = 2048

	// Implementation namespaces.
	ImplNamespace = "0x121E241111111111111111111111111111000000"

	// Predeploy addresses.
	ProxyAdmin = "0x121E240000000000000000000000000000000001"
	OmniStake  = "0x121E240000000000000000000000000000000002"

	// TransparentUpgradeableProxy storage slots.
	ProxyImplementationSlot = "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"
	ProxyAdminSlot          = "0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103"
)

var (
	// Namespace big.Ints.
	mainNamespace = addrToBig(common.HexToAddress(Namespace))
	implNamespace = addrToBig(common.HexToAddress(ImplNamespace))

	// Predeploy addresses.
	proxyAdmin = common.HexToAddress(ProxyAdmin)
	omniStake  = common.HexToAddress(OmniStake)

	// Predeploy bytecodes.
	proxyCode      = mustDecodeHex(bindings.TransparentUpgradeableProxyDeployedBytecode)
	proxyAdminCode = mustDecodeHex(bindings.ProxyAdminDeployedBytecode)
	omniStakeCode  = mustDecodeHex(bindings.OmniStakeDeployedBytecode)
)

// Alloc returns the genesis allocs for the predeployed contracts, initializing code and storage.
func Alloc(network netconf.ID) (types.GenesisAlloc, error) {
	emptyGenesis := &core.Genesis{Alloc: types.GenesisAlloc{}}

	db := state.NewMemDB(emptyGenesis)

	setProxies(db)

	if err := setProxyAdmin(db, network); err != nil {
		return nil, errors.Wrap(err, "set proxy admin")
	}

	if err := setOmniStake(db); err != nil {
		return nil, errors.Wrap(err, "set omni stake")
	}

	return db.Genesis().Alloc, nil
}

// setProxies deployes TransparentUpgradeableProxy contracts for the first NumProxies addresses in the proxy namespace.
func setProxies(db *state.MemDB) {
	for i := 0; i < NamespaceSize; i++ {
		// add one, so that we don't set the namespace zero addr
		addr := namespaceAddr(mainNamespace, i+1)

		// do not deploy a proxy for the proxy admin
		if addr == proxyAdmin {
			continue
		}

		db.SetCode(addr, proxyCode)
		db.SetState(addr, common.HexToHash(ProxyAdminSlot), common.HexToHash(proxyAdmin.Hex()))
	}
}

// setOmniStake sets the omniStake predeploy.
func setOmniStake(db *state.MemDB) error {
	// OmniStake storage is empty
	storage := state.StorageValues{}

	impl := implementation(omniStake)
	setProxyImplementation(db, omniStake, impl)

	return setPredeploy(db, impl, omniStakeCode, bindings.OmniStakeStorageLayout, storage)
}

// setProxyAdmin sets the proxy admin predeploy.
func setProxyAdmin(db *state.MemDB, network netconf.ID) error {
	owner, err := eoa.Admin(network)
	if err != nil {
		return errors.Wrap(err, "network admin")
	}

	storage := state.StorageValues{
		"_owner": owner,
	}

	return setPredeploy(db, proxyAdmin, proxyAdminCode, bindings.ProxyAdminStorageLayout, storage)
}

// setPredeploy sets the code and storage for the given predeploy.
func setPredeploy(db *state.MemDB, addr common.Address, code []byte, layout solc.StorageLayout, storage state.StorageValues) error {
	db.SetCode(addr, code)

	slots, err := state.EncodeStorage(layout, storage)
	if err != nil {
		return errors.Wrap(err, "encode storage", "addr", addr)
	}

	for _, slot := range slots {
		db.SetState(addr, slot.Key, slot.Value)
	}

	return nil
}

// setProxyImplementation sets the implementation address for the given proxy address.
func setProxyImplementation(db *state.MemDB, proxy, impl common.Address) {
	db.SetState(proxy, common.HexToHash(ProxyImplementationSlot), common.HexToHash(impl.Hex()))
}

// implementation returns the implementation address for the given proxy address.
func implementation(addr common.Address) common.Address {
	return namespaceAddr(implNamespace, namespaceIdx(mainNamespace, addr))
}

// namespaceAddr returns the address at the given index in the namespace.
func namespaceAddr(namespace *big.Int, i int) common.Address {
	return common.BigToAddress(new(big.Int).Add(namespace, big.NewInt(int64(i))))
}

// namespaceIdx returns the index of the address in the namespace.
func namespaceIdx(namespace *big.Int, addr common.Address) int {
	return int(new(big.Int).Sub(addr.Big(), namespace).Int64())
}

// addrToBig converts an address to a big.Int.
func addrToBig(addr common.Address) *big.Int {
	return new(big.Int).SetBytes(addr.Bytes())
}

// mustDecodeHex decodes the given hex string. It panics on error.
func mustDecodeHex(hex string) []byte {
	b, err := hexutil.Decode(hex)
	if err != nil {
		panic(err)
	}

	return b
}
