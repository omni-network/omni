package contracts

import (
	"context"
	"math/big"
	"sync"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	devnetVersion  = "devnet"
	omegaVersion   = "v0.1.0"
	mainnetVersion = "v1.0.0"

	NameAVS                = "avs"
	NameCreate3Factory     = "create3-factory"
	NameGasPump            = "gas-pump"
	NameGasStation         = "gas-station"
	NameL1Bridge           = "l1-bridge"
	NamePortal             = "portal"
	NameToken              = "token"
	NameSolverNetInbox     = "solvernet-inbox"
	NameSolverNetOutbox    = "solvernet-outbox"
	NameSolverNetMiddleman = "solvernet-middleman"
	NameFeeOracleV2        = "fee-oracle-v2"
)

// version returns the salt version for a given network. Staging version is block 1 hash.
func version(ctx context.Context, network netconf.ID) (string, error) {
	switch network {
	case netconf.Devnet:
		return devnetVersion, nil
	case netconf.Staging:
		return getStagingVersion(ctx)
	case netconf.Omega:
		return omegaVersion, nil
	case netconf.Mainnet:
		return mainnetVersion, nil
	default:
		return "", errors.New("unsupported network", "network", network)
	}
}

var (
	// Cached staging version.
	stagingVersion string

	// Overrides default https://staging.omni.network
	stagingOmniRPC string
)

// UseStagingOmniRPC overrides the default staging Omni EVM RPC URL.
func UseStagingOmniRPC(rpc string) {
	stagingOmniRPC = rpc
}

func getStagingVersion(ctx context.Context) (string, error) {
	// Cache the staging version.
	if stagingVersion != "" {
		return stagingVersion, nil
	}

	rpc := netconf.Staging.Static().ExecutionRPC()
	if stagingOmniRPC != "" {
		rpc = stagingOmniRPC
	}

	client, err := ethclient.Dial("omni_evm", rpc)
	if err != nil {
		return "", errors.Wrap(err, "dial omni")
	}

	block1, err := client.BlockByNumber(ctx, big.NewInt(1))
	if err != nil {
		return "", errors.Wrap(err, "get block 1")
	}

	stagingVersion = block1.Hash().Hex()

	return stagingVersion, nil
}

type Addresses struct {
	AVS                common.Address
	Create3Factory     common.Address
	GasPump            common.Address
	GasStation         common.Address
	L1Bridge           common.Address
	Portal             common.Address
	Token              common.Address
	SolverNetInbox     common.Address
	SolverNetOutbox    common.Address
	SolverNetMiddleman common.Address
	FeeOracleV2        common.Address
}

type Salts struct {
	AVS                string
	GasPump            string
	GasStation         string
	L1Bridge           string
	Portal             string
	Token              string
	SolverNetInbox     string
	SolverNetOutbox    string
	SolverNetMiddleman string
	FeeOracleV2        string
}

type cache[T any] struct {
	mu    sync.Mutex
	cache T
}

var (
	// cached addresses by network.
	addrsCache = cache[map[netconf.ID]Addresses]{
		cache: map[netconf.ID]Addresses{},
	}

	// cached salts by network.
	saltsCache = cache[map[netconf.ID]Salts]{
		cache: map[netconf.ID]Salts{},
	}
)

// GetAddresses returns the contract addresses for the given network.
func GetAddresses(ctx context.Context, network netconf.ID) (Addresses, error) {
	ver, err := version(ctx, network)
	if err != nil {
		return Addresses{}, err
	}

	addrsCache.mu.Lock()
	defer addrsCache.mu.Unlock()

	addrs, ok := addrsCache.cache[network]
	if ok {
		return addrs, nil
	}

	s := func(name string) string {
		if !isVersioned(name) {
			return salt(network, name)
		}

		return salt(network, versioned(name, ver))
	}

	addrs = Addresses{
		Create3Factory:     Create3Factory(network),
		AVS:                Avs(network),
		Token:              TokenAddr(network),
		Portal:             addr(network, s(NamePortal)),
		L1Bridge:           addr(network, s(NameL1Bridge)),
		GasPump:            addr(network, s(NameGasPump)),
		GasStation:         addr(network, s(NameGasStation)),
		SolverNetInbox:     addr(network, s(NameSolverNetInbox)),
		SolverNetOutbox:    addr(network, s(NameSolverNetOutbox)),
		SolverNetMiddleman: addr(network, s(NameSolverNetMiddleman)),
		FeeOracleV2:        addr(network, s(NameFeeOracleV2)),
	}

	addrsCache.cache[network] = addrs

	return addrs, nil
}

// GetSalts returns the contract salts for the given network.
func GetSalts(ctx context.Context, network netconf.ID) (Salts, error) {
	saltsCache.mu.Lock()
	defer saltsCache.mu.Unlock()

	salts, ok := saltsCache.cache[network]
	if ok {
		return salts, nil
	}

	ver, err := version(ctx, network)
	if err != nil {
		return Salts{}, err
	}

	s := func(name string) string {
		if !isVersioned(name) {
			return salt(network, name)
		}

		return salt(network, versioned(name, ver))
	}

	salts = Salts{
		AVS:                s(NameAVS),
		Portal:             s(NamePortal),
		L1Bridge:           s(NameL1Bridge),
		Token:              s(NameToken),
		GasPump:            s(NameGasPump),
		GasStation:         s(NameGasStation),
		SolverNetInbox:     s(NameSolverNetInbox),
		SolverNetOutbox:    s(NameSolverNetOutbox),
		SolverNetMiddleman: s(NameSolverNetMiddleman),
		FeeOracleV2:        s(NameFeeOracleV2),
	}

	saltsCache.cache[network] = salts

	return salts, nil
}

// Avs returns the AVS contract address for the given network.
func Avs(network netconf.ID) common.Address {
	if network == netconf.Mainnet {
		return common.HexToAddress("0xed2f4d90b073128ae6769a9A8D51547B1Df766C8")
	} else if network == netconf.Omega {
		return common.HexToAddress("0xa7b2e7830C51728832D33421670DbBE30299fD92")
	}

	return addr(network, salt(network, NameAVS))
}

// TokenAddr returns the Omni ERC20 token contract address for the given network.
func TokenAddr(network netconf.ID) common.Address {
	if network == netconf.Mainnet {
		return common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4")
	} else if network == netconf.Omega {
		return common.HexToAddress("0xD036C60f46FF51dd7Fbf6a819b5B171c8A076b07")
	}

	return addr(network, salt(network, NameToken))
}

// Create3Factory returns the Create3 factory address for the given network.
func Create3Factory(network netconf.ID) common.Address {
	return crypto.CreateAddress(eoa.MustAddress(network, eoa.RoleCreate3Deployer), 0)
}

// Create3Address returns the Create3 address for the given network and salt id.
func Create3Address(ctx context.Context, network netconf.ID, saltID string) (common.Address, error) {
	s, err := Create3Salt(ctx, network, saltID)
	if err != nil {
		return common.Address{}, err
	}

	return addr(network, s), nil
}

// Create3Salt returns the Create3 salt for the given network and salt id.
func Create3Salt(ctx context.Context, network netconf.ID, saltID string) (string, error) {
	ver, err := version(ctx, network)
	if err != nil {
		return "", err
	}

	return salt(network, versioned(saltID, ver)), nil
}

func isVersioned(contract string) bool {
	// AVS not versioned, as requiring re-registration per each version is too cumbersome.
	// Token salt is static, as Omni ERC20 contract does not change.
	not := contract == NameAVS || contract == NameToken

	return !not
}

func versioned(contract string, version string) string {
	return contract + "-" + version
}

func salt(network netconf.ID, suffix string) string {
	return string(network) + "-" + suffix
}

func addr(network netconf.ID, salt string) common.Address {
	return create3.Address(Create3Factory(network), salt, eoa.MustAddress(network, eoa.RoleDeployer))
}
