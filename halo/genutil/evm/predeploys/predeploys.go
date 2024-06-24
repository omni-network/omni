package predeploys

import (
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
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
	// OmniNamespace is namespace of for omni specific predeploys.
	OmniNamepsace = "0x121E240000000000000000000000000000000000"

	// OctaneNamespace is namespace of for octane specific predeploys.
	OctaneNamespace = "0xcccccc0000000000000000000000000000000000"

	// NamespaceSize is the number of proxies to deploy per namespace.
	NamespaceSize = 2048

	// ProxyAdmin for all namespaces.
	ProxyAdmin = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	// Omni Predeploys.
	PortalRegistry   = "0x121E240000000000000000000000000000000001"
	OmniBridgeNative = "0x121E240000000000000000000000000000000002"
	WOmni            = "0x121E240000000000000000000000000000000003"
	EthStakeInbox    = "0x121E240000000000000000000000000000000004"

	// Octane Predeploys.
	Staking  = "0xcccccc0000000000000000000000000000000001"
	Slashing = "0xcccccc0000000000000000000000000000000002"

	// TransparentUpgradeableProxy storage slots.
	ProxyImplementationSlot = "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"
	ProxyAdminSlot          = "0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103"
)

var (
	// Namespace big.Ints.
	omniNamespace   = common.HexToAddress(OmniNamepsace).Big()
	octaneNamespace = common.HexToAddress(OctaneNamespace).Big()

	// Predeploy addresses.
	proxyAdmin     = common.HexToAddress(ProxyAdmin)
	portalRegistry = common.HexToAddress(PortalRegistry)
	womni          = common.HexToAddress(WOmni)
	staking        = common.HexToAddress(Staking)
	slashing       = common.HexToAddress(Slashing)
	omniBridge     = common.HexToAddress(OmniBridgeNative)

	// Predeploy bytecodes.
	proxyCode          = hexutil.MustDecode(bindings.TransparentUpgradeableProxyDeployedBytecode)
	proxyAdminCode     = hexutil.MustDecode(bindings.ProxyAdminDeployedBytecode)
	portalRegistryCode = hexutil.MustDecode(bindings.PortalRegistryDeployedBytecode)
	womniCode          = hexutil.MustDecode(bindings.WOmniDeployedBytecode)
	stakingCode        = hexutil.MustDecode(bindings.StakingDeployedBytecode)
	slashingCode       = hexutil.MustDecode(bindings.SlashingDeployedBytecode)
	omniBridgeCode     = hexutil.MustDecode(bindings.OmniBridgeNativeDeployedBytecode)
)

// Alloc returns the genesis allocs for the predeployed contracts, initializing code and storage.
func Alloc(network netconf.ID, admin common.Address) (types.GenesisAlloc, error) {
	emptyGenesis := &core.Genesis{Alloc: types.GenesisAlloc{}}

	db := state.NewMemDB(emptyGenesis)

	setProxies(db)

	if err := setProxyAdmin(db, admin); err != nil {
		return nil, errors.Wrap(err, "set proxy admin")
	}

	if err := setPortalRegistry(db, admin); err != nil {
		return nil, errors.Wrap(err, "set portal registry")
	}

	if err := setOmniBridge(db, admin); err != nil {
		return nil, errors.Wrap(err, "set omni bridge")
	}

	if err := setWOmni(db); err != nil {
		return nil, errors.Wrap(err, "set womni")
	}

	if err := setStaking(db, admin, network.IsProtected()); err != nil {
		return nil, errors.Wrap(err, "set staking")
	}

	if err := setSlashing(db); err != nil {
		return nil, errors.Wrap(err, "set slashing")
	}

	return db.Genesis().Alloc, nil
}

// setProxies deployes TransparentUpgradeableProxy contracts for the first NumProxies addresses in the proxy namespace.
func setProxies(db *state.MemDB) {
	for _, namspace := range []*big.Int{omniNamespace, octaneNamespace} {
		for i := 0; i < NamespaceSize; i++ {
			// add one, so that we don't set the namespace zero addr
			addr := address(namspace, i+1)

			db.SetCode(addr, proxyCode)
			db.SetState(addr, common.HexToHash(ProxyAdminSlot), common.HexToHash(proxyAdmin.Hex()))
		}
	}
}

// setProxyAdmin sets the proxy admin predeploy.
func setProxyAdmin(db *state.MemDB, owner common.Address) error {
	storage := state.StorageValues{"_owner": owner}

	db.SetCode(proxyAdmin, proxyAdminCode)

	return setStrorage(db, proxyAdmin, bindings.ProxyAdminStorageLayout, storage)
}

// setPortalRegistry sets the PortalRegistry predeploy.
func setPortalRegistry(db *state.MemDB, owner common.Address) error {
	storage := state.StorageValues{
		"_initialized": uint8(1), // disable initializer
		"_owner":       owner,
	}

	return setPredeploy(db, portalRegistry, portalRegistryCode, bindings.PortalRegistryStorageLayout, storage)
}

// setOmniBridge sets the OmniBridgeNative predeploy.
func setOmniBridge(db *state.MemDB, owner common.Address) error {
	storage := state.StorageValues{
		"_initialized": uint8(1), // disable initializer
		"_owner":       owner,
	}

	// 100M total supply
	db.SetBalance(omniBridge, new(big.Int).Mul(big.NewInt(100e6), big.NewInt(1e18)))

	return setPredeploy(db, omniBridge, omniBridgeCode, bindings.OmniBridgeNativeStorageLayout, storage)
}

// setWOmni sets the WOmni predeploy.
func setWOmni(db *state.MemDB) error {
	storage := state.StorageValues{}

	return setPredeploy(db, womni, womniCode, bindings.WOmniStorageLayout, storage)
}

// setStaking sets the Staking predeploy.
func setStaking(db *state.MemDB, owner common.Address, isAllowlistEnabled bool) error {
	storage := state.StorageValues{
		"_initialized":       uint8(1), // disable initializer
		"_owner":             owner,
		"isAllowlistEnabled": isAllowlistEnabled,
	}

	return setPredeploy(db, staking, stakingCode, bindings.StakingStorageLayout, storage)
}

// setSlashing sets the Slashing predeploy.
func setSlashing(db *state.MemDB) error {
	storage := state.StorageValues{}

	return setPredeploy(db, slashing, slashingCode, bindings.SlashingStorageLayout, storage)
}

// setPredeploy sets the implementation code and proxy storage for the given predeploy.
func setPredeploy(db *state.MemDB, proxy common.Address, code []byte, layout solc.StorageLayout, storage state.StorageValues) error {
	impl := impl(proxy)
	setProxyImplementation(db, proxy, impl)

	// disable impl initializers, if needed
	if _, ok := solc.SlotOf(layout, "_initialized"); ok {
		slot, err := state.EncodeStorageEntry(layout, "_initialized", uint8(255)) // max uint8, disables all initializers
		if err != nil {
			return errors.Wrap(err, "encode impl storage", "addr", impl)
		}

		db.SetState(impl, slot.Key, slot.Value)
	}

	db.SetCode(impl, code)

	return setStrorage(db, proxy, layout, storage)
}

// setStrorage sets the code and storage for the given predeploy.
func setStrorage(db *state.MemDB, addr common.Address, layout solc.StorageLayout, storage state.StorageValues) error {
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

// impl returns the implementation address for the given proxy address.
func impl(addr common.Address) common.Address {
	// To get a unique implementation per each proxy address, we subtract the address from the max address.
	// Max address is odd, so the result will be unique.
	maxAddr := common.HexToAddress("0xffffffffffffffffffffffffffffffffffffffff").Big()
	return common.BigToAddress(new(big.Int).Sub(maxAddr, addr.Big()))
}

// address returns the address at the given index in the namespace.
func address(namespace *big.Int, i int) common.Address {
	return common.BigToAddress(new(big.Int).Add(namespace, big.NewInt(int64(i))))
}
