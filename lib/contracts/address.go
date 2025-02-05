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
	AVS             common.Address
	Create3Factory  common.Address
	GasPump         common.Address
	GasStation      common.Address
	L1Bridge        common.Address
	Portal          common.Address
	Token           common.Address
	SolverNetInbox  common.Address
	SolverNetOutbox common.Address
	FeeOracleV2     common.Address
}

type Salts struct {
	AVS             string
	GasPump         string
	GasStation      string
	L1Bridge        string
	Portal          string
	Token           string
	SolverNetInbox  string
	SolverNetOutbox string
	FeeOracleV2     string
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

	addrs = Addresses{
		Create3Factory:  Create3Factory(network),
		AVS:             avs(network),
		Portal:          portal(network, ver),
		L1Bridge:        l1Bridge(network, ver),
		Token:           TokenAddr(network),
		GasPump:         gasPump(network, ver),
		GasStation:      gasStation(network, ver),
		SolverNetInbox:  solverNetInbox(network, ver),
		SolverNetOutbox: solverNetOutbox(network, ver),
		FeeOracleV2:     feeOracleV2(network, ver),
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

	salts = Salts{
		AVS:             avsSalt(network),
		Portal:          portalSalt(network, ver),
		L1Bridge:        l1BridgeSalt(network, ver),
		Token:           tokenSalt(network),
		GasPump:         gasPumpSalt(network, ver),
		GasStation:      gasStationSalt(network, ver),
		SolverNetInbox:  solverNetInboxSalt(network, ver),
		SolverNetOutbox: solverNetOutboxSalt(network, ver),
		FeeOracleV2:     feeOracleV2Salt(network, ver),
	}

	saltsCache.cache[network] = salts

	return salts, nil
}

// avs returns the AVS contract address for the given network.
func avs(network netconf.ID) common.Address {
	if network == netconf.Mainnet {
		return common.HexToAddress("0xed2f4d90b073128ae6769a9A8D51547B1Df766C8")
	} else if network == netconf.Omega {
		return common.HexToAddress("0xa7b2e7830C51728832D33421670DbBE30299fD92")
	}

	return create3.Address(Create3Factory(network), avsSalt(network), eoa.MustAddress(network, eoa.RoleDeployer))
}

// Create3Factory returns the Create3 factory address for the given network.
func Create3Factory(network netconf.ID) common.Address {
	return crypto.CreateAddress(eoa.MustAddress(network, eoa.RoleCreate3Deployer), 0)
}

func Create3Address(ctx context.Context, network netconf.ID, saltPrefix string) (common.Address, error) {
	ver, err := version(ctx, network)
	if err != nil {
		return common.Address{}, err
	}

	return create3.Address(Create3Factory(network), salt(network, saltPrefix+"-"+ver), eoa.MustAddress(network, eoa.RoleDeployer)), nil
}

func Create3Salt(ctx context.Context, network netconf.ID, saltPrefix string) (string, error) {
	ver, err := version(ctx, network)
	if err != nil {
		return "", err
	}

	return salt(network, saltPrefix+"-"+ver), nil
}

// portal returns the Portal contract address for the given network.
func portal(network netconf.ID, saltVersion string) common.Address {
	return create3.Address(Create3Factory(network), portalSalt(network, saltVersion), eoa.MustAddress(network, eoa.RoleDeployer))
}

// l1Bridge returns the L1Bridge contract address for the given network.
func l1Bridge(network netconf.ID, version string) common.Address {
	return create3.Address(Create3Factory(network), l1BridgeSalt(network, version), eoa.MustAddress(network, eoa.RoleDeployer))
}

// TokenAddr returns the Omni ERC20 token contract address for the given network.
func TokenAddr(network netconf.ID) common.Address {
	if network == netconf.Mainnet {
		return common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4")
	} else if network == netconf.Omega {
		return common.HexToAddress("0xD036C60f46FF51dd7Fbf6a819b5B171c8A076b07")
	}

	return create3.Address(Create3Factory(network), tokenSalt(network), eoa.MustAddress(network, eoa.RoleDeployer))
}

// gasPump returns the GasPump contract address for the given network.
func gasPump(network netconf.ID, version string) common.Address {
	return create3.Address(Create3Factory(network), gasPumpSalt(network, version), eoa.MustAddress(network, eoa.RoleDeployer))
}

// gasStation returns the GasStation contract address for the given network.
func gasStation(network netconf.ID, version string) common.Address {
	return create3.Address(Create3Factory(network), gasStationSalt(network, version), eoa.MustAddress(network, eoa.RoleDeployer))
}

func solverNetInbox(network netconf.ID, version string) common.Address {
	return create3.Address(Create3Factory(network), solverNetInboxSalt(network, version), eoa.MustAddress(network, eoa.RoleDeployer))
}

func solverNetOutbox(network netconf.ID, version string) common.Address {
	return create3.Address(Create3Factory(network), solverNetOutboxSalt(network, version), eoa.MustAddress(network, eoa.RoleDeployer))
}

func feeOracleV2(network netconf.ID, version string) common.Address {
	return create3.Address(Create3Factory(network), feeOracleV2Salt(network, version), eoa.MustAddress(network, eoa.RoleDeployer))
}

//
// Salts.
//

func portalSalt(network netconf.ID, version string) string {
	return salt(network, "portal-"+version)
}

func avsSalt(network netconf.ID) string {
	// AVS not versioned, as requiring re-registration per each version is too cumbersome.
	return salt(network, "avs")
}

func l1BridgeSalt(network netconf.ID, version string) string {
	return salt(network, "l1-bridge-"+version)
}

// token salt is static, as Omni ERC20 contract does not change.
func tokenSalt(network netconf.ID) string {
	return salt(network, "token")
}

func gasPumpSalt(network netconf.ID, version string) string {
	return salt(network, "gas-pump-"+version)
}

func gasStationSalt(network netconf.ID, version string) string {
	return salt(network, "gas-station-"+version)
}

func solverNetInboxSalt(network netconf.ID, version string) string {
	return salt(network, "solvernet-inbox-"+version)
}

func solverNetOutboxSalt(network netconf.ID, version string) string {
	return salt(network, "solvernet-outbox-"+version)
}

func feeOracleV2Salt(network netconf.ID, version string) string {
	return salt(network, "fee-oracle-v2-"+version)
}

//
// Utils.
//

// salt generates a salt for a contract deployment. For ephemeral networks,
// the salt includes a random per-run suffix. For persistent networks, the
// sale is static.
func salt(network netconf.ID, contract string) string {
	return string(network) + "-" + contract
}
