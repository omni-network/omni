package contracts

import (
	"context"
	"sync"
	"time"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	devnetVersion  = "devnet"
	omegaVersion   = "v0.1.0"
	mainnetVersion = "v1.0.0"

	// solver net versions can progress independently of core.
	solvernetOmegaVersion   = "v0.1.2"
	solvernetMainnetVersion = "v1.0.0"

	NameAVS               = "avs"
	NameCreate3Factory    = "create3-factory"
	NameGasPump           = "gas-pump"
	NameGasStation        = "gas-station"
	NameL1Bridge          = "l1-bridge"
	NamePortal            = "portal"
	NameOmniToken         = "token"
	NameNomToken          = "nom-token"
	NameSolverNetInbox    = "solvernet-inbox"
	NameSolverNetOutbox   = "solvernet-outbox"
	NameSovlerNetExecutor = "solvernet-executor"
	NameFeeOracleV2       = "fee-oracle-v2"
)

type Versions struct {
	Core      string
	SolverNet string
}

// version returns the salt version for a given network. Staging version is block 1 hash.
func versions(ctx context.Context, network netconf.ID) (Versions, error) {
	switch network {
	case netconf.Devnet:
		// same for both on devnet
		return Versions{Core: devnetVersion, SolverNet: devnetVersion}, nil
	case netconf.Staging:
		for range 3 {
			v, err := StagingID(ctx)
			if err != nil {
				log.Warn(ctx, "Failed fetching staging ID (will retry)", err)
				time.Sleep(time.Second * 3)

				continue
			}

			// same for both on staging
			return Versions{Core: v, SolverNet: v}, nil
		}

		return Versions{}, errors.New("failed to fetch staging ID after 3 attempts")
	case netconf.Omega:
		return Versions{Core: omegaVersion, SolverNet: solvernetOmegaVersion}, nil
	case netconf.Mainnet:
		return Versions{Core: mainnetVersion, SolverNet: solvernetMainnetVersion}, nil
	default:
		return Versions{}, errors.New("unsupported network", "network", network)
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

// StagingID returns id for a staing instance (hash of block 1).
func StagingID(ctx context.Context) (string, error) {
	// Cache the staging version.
	if stagingVersion != "" {
		return stagingVersion, nil
	}

	rpc := netconf.Staging.Static().ExecutionRPC()
	if stagingOmniRPC != "" {
		rpc = stagingOmniRPC
	}

	client, err := ethclient.DialContext(ctx, "omni_evm", rpc)
	if err != nil {
		return "", errors.Wrap(err, "dial omni")
	}

	block1, err := client.BlockByNumber(ctx, bi.One())
	if err != nil {
		return "", errors.Wrap(err, "get block 1")
	}

	stagingVersion = block1.Hash().Hex()

	return stagingVersion, nil
}

type Addresses struct {
	AVS               common.Address
	Create3Factory    common.Address
	GasPump           common.Address
	GasStation        common.Address
	L1Bridge          common.Address
	Portal            common.Address
	Token             common.Address
	SolverNetInbox    common.Address
	SolverNetOutbox   common.Address
	SolverNetExecutor common.Address
	FeeOracleV2       common.Address
}

type Salts struct {
	AVS               string
	GasPump           string
	GasStation        string
	L1Bridge          string
	Portal            string
	Token             string
	SolverNetInbox    string
	SolverNetOutbox   string
	SolverNetExecutor string
	FeeOracleV2       string
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
	v, err := versions(ctx, network)
	if err != nil {
		return Addresses{}, err
	}

	addrsCache.mu.Lock()
	defer addrsCache.mu.Unlock()

	addrs, ok := addrsCache.cache[network]
	if ok {
		return addrs, nil
	}

	s := func(name string) string { return salt(network, name, v) }

	addrs = Addresses{
		Create3Factory:    Create3Factory(network),
		AVS:               Avs(network),
		Token:             TokenAddr(network),
		Portal:            addr(network, s(NamePortal)),
		L1Bridge:          addr(network, s(NameL1Bridge)),
		GasPump:           addr(network, s(NameGasPump)),
		GasStation:        addr(network, s(NameGasStation)),
		SolverNetInbox:    addr(network, s(NameSolverNetInbox)),
		SolverNetOutbox:   addr(network, s(NameSolverNetOutbox)),
		SolverNetExecutor: addr(network, s(NameSovlerNetExecutor)),
		FeeOracleV2:       addr(network, s(NameFeeOracleV2)),
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

	v, err := versions(ctx, network)
	if err != nil {
		return Salts{}, err
	}

	s := func(name string) string { return salt(network, name, v) }

	salts = Salts{
		AVS:               s(NameAVS),
		Portal:            s(NamePortal),
		L1Bridge:          s(NameL1Bridge),
		Token:             s(NameOmniToken),
		GasPump:           s(NameGasPump),
		GasStation:        s(NameGasStation),
		SolverNetInbox:    s(NameSolverNetInbox),
		SolverNetOutbox:   s(NameSolverNetOutbox),
		SolverNetExecutor: s(NameSovlerNetExecutor),
		FeeOracleV2:       s(NameFeeOracleV2),
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

	return addr(network, prefixNetwork(network, NameAVS))
}

// TokenAddr returns the Omni ERC20 token contract address for the given network.
func TokenAddr(network netconf.ID) common.Address {
	if network == netconf.Mainnet {
		return common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4")
	} else if network == netconf.Omega {
		return common.HexToAddress("0xD036C60f46FF51dd7Fbf6a819b5B171c8A076b07")
	}

	return addr(network, prefixNetwork(network, NameOmniToken))
}

func NomAddr(network netconf.ID) common.Address {
	if network == netconf.Mainnet || network == netconf.Omega {
		return common.HexToAddress("0x6e6f6d696e610000000000000000000000000000")
	}

	return addr(network, prefixNetwork(network, NameNomToken))
}

// Create3Factory returns the Create3 factory address for the given network.
func Create3Factory(network netconf.ID) common.Address {
	return crypto.CreateAddress(eoa.MustAddress(network, eoa.RoleCreate3Deployer), 0)
}

// Create3Address returns the Create3 address for the given network and salt.
func Create3Address(network netconf.ID, salt string) common.Address {
	return addr(network, salt)
}

// salt returns the salt string for a contract on a network / version.
func salt(network netconf.ID, name string, versions Versions) string {
	if !isVersioned(name) {
		return prefixNetwork(network, name)
	}

	if isSolvernet(name) {
		return prefixNetwork(network, suffixVersion(name, versions.SolverNet))
	}

	return prefixNetwork(network, suffixVersion(name, versions.Core))
}

func isVersioned(contract string) bool {
	// AVS not versioned, as requiring re-registration per each version is too cumbersome.
	// Token salt is static, as Omni ERC20 contract does not change.
	not := contract == NameAVS || contract == NameOmniToken || contract == NameNomToken

	return !not
}

func isSolvernet(contract string) bool {
	return (contract == NameSolverNetInbox ||
		contract == NameSolverNetOutbox ||
		contract == NameSovlerNetExecutor)
}

func suffixVersion(contract string, version string) string {
	return contract + "-" + version
}

func prefixNetwork(network netconf.ID, suffix string) string {
	return string(network) + "-" + suffix
}

func addr(network netconf.ID, salt string) common.Address {
	return create3.Address(Create3Factory(network), salt, eoa.MustAddress(network, eoa.RoleDeployer))
}
