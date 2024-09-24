package netconf

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"

	cosmgen "github.com/cosmos/cosmos-sdk/x/genutil/types"

	_ "embed"
)

const consensusIDPrefix = "omni-"
const consensusIDOffset = 1_000_000
const maxValidators = 30

// Static defines static config and data for a network.
type Static struct {
	Network              ID
	OmniExecutionChainID uint64
	AVSContractAddress   common.Address
	TokenAddress         common.Address
	L1BridgeAddress      common.Address
	Portals              []Deployment
	MaxValidators        uint32
	ConsensusGenesisJSON []byte
	ConsensusSeedTXT     []byte
	ConsensusArchiveTXT  []byte
	ExecutionGenesisJSON []byte
	ExecutionSeedTXT     []byte

	version string
}

type Deployment struct {
	ChainID      uint64
	Address      common.Address
	DeployHeight uint64
}

// OmniConsensusChainIDStr returns the chain ID string for the Omni consensus chain.
// It is calculated as "omni-<OmniConsensusChainIDUint64>".
func (s Static) OmniConsensusChainIDStr() string {
	return fmt.Sprintf("%s%d", consensusIDPrefix, s.OmniConsensusChainIDUint64())
}

// OmniExecutionChainName returns the name of the Omni execution chain.
func (s Static) OmniExecutionChainName() string {
	meta, _ := evmchain.MetadataByID(s.OmniExecutionChainID)
	return meta.Name
}

// OmniConsensusChainIDUint64 returns the chain ID uint64 for the Omni consensus chain.
// It is calculated as 1_000_000 + OmniExecutionChainID.
func (s Static) OmniConsensusChainIDUint64() uint64 {
	return consensusIDOffset + s.OmniExecutionChainID
}

// OmniConsensusChain returns the omni consensus Chain struct.
func (s Static) OmniConsensusChain() Chain {
	return Chain{
		ID:             s.OmniConsensusChainIDUint64(),
		Name:           "omni_consensus",
		BlockPeriod:    time.Second * 2,
		Shards:         []xchain.ShardID{xchain.ShardBroadcast0}, // Consensus chain only supports broadcast shard.
		DeployHeight:   1,                                        // Emit portal blocks start at 1, not 0.
		AttestInterval: 0,                                        // Emit portal blocks are never empty, so this isn't required.
	}
}

// PortalDeployment returns the portal deployment for the given chainID.
// If there is none, it returns an empty deployment.
func (s Static) PortalDeployment(chainID uint64) (Deployment, bool) {
	for _, d := range s.Portals {
		if d.ChainID == chainID {
			return d, true
		}
	}

	return Deployment{}, false
}

func (s Static) ConsensusSeeds() []string {
	var resp []string
	for _, seed := range strings.Split(string(s.ConsensusSeedTXT), "\n") {
		if seed = strings.TrimSpace(seed); seed != "" {
			resp = append(resp, seed)
		}
	}

	return resp
}

func (s Static) ConsensusArchives() []string {
	var resp []string
	for _, archive := range strings.Split(string(s.ConsensusArchiveTXT), "\n") {
		if archive = strings.TrimSpace(archive); archive != "" {
			resp = append(resp, archive)
		}
	}

	return resp
}

func (s Static) ExecutionSeeds() []string {
	var resp []string
	for _, seed := range strings.Split(string(s.ExecutionSeedTXT), "\n") {
		if seed = strings.TrimSpace(seed); seed != "" {
			resp = append(resp, seed)
		}
	}

	return resp
}

func (s Static) ExecutionRPC() string {
	if s.Network == Devnet {
		return "http://localhost:8001"
	}

	return fmt.Sprintf("https://%s.omni.network", s.Network)
}

func (s Static) ConsensusRPC() string {
	if s.Network == Devnet {
		return "http://localhost:5701"
	}

	return fmt.Sprintf("https://consensus.%s.omni.network", s.Network)
}

//nolint:gochecknoglobals // Static addresses
var (
	omegaAVS     = common.HexToAddress("0xa7b2e7830C51728832D33421670DbBE30299fD92")
	omegaPortal  = common.HexToAddress("0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3")
	mainnetAVS   = common.HexToAddress("0xed2f4d90b073128ae6769a9A8D51547B1Df766C8")
	mainnetToken = common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4")

	//go:embed omega/consensus-genesis.json
	omegaConsensusGenesisJSON []byte

	//go:embed omega/consensus-seeds.txt
	omegaConsensusSeedsTXT []byte

	//go:embed omega/consensus-archives.txt
	omegaConsensusArchivesTXT []byte

	//go:embed omega/execution-genesis.json
	omegaExecutionGenesisJSON []byte

	//go:embed omega/execution-seeds.txt
	omegaExecutionSeedsTXT []byte

	//go:embed staging/consensus-seeds.txt
	stagingConsensusSeedsTXT []byte

	//go:embed staging/consensus-archives.txt
	stagingConsensusArchivesTXT []byte

	//go:embed staging/execution-seeds.txt
	stagingExecutionSeedsTXT []byte
)

//nolint:gochecknoglobals // Static mappings.
var statics = map[ID]Static{
	Simnet: {
		Network:              Simnet,
		version:              "simnet",
		OmniExecutionChainID: evmchain.IDOmniDevnet,
		MaxValidators:        maxValidators,
	},
	Devnet: {
		Network:              Devnet,
		version:              "devnet",
		OmniExecutionChainID: evmchain.IDOmniDevnet,
		MaxValidators:        maxValidators,
	},
	Staging: {
		Network:              Staging,
		OmniExecutionChainID: evmchain.IDOmniStaging,
		MaxValidators:        maxValidators,
		ConsensusSeedTXT:     stagingConsensusSeedsTXT,
		ConsensusArchiveTXT:  stagingConsensusArchivesTXT,
		ExecutionSeedTXT:     stagingExecutionSeedsTXT,
	},
	Omega: {
		Network:              Omega,
		version:              "v0.1.0",
		AVSContractAddress:   omegaAVS,
		OmniExecutionChainID: evmchain.IDOmniOmega,
		MaxValidators:        maxValidators,
		Portals: []Deployment{
			{ChainID: evmchain.IDArbSepolia, Address: omegaPortal, DeployHeight: 71015563},
			{ChainID: evmchain.IDBaseSepolia, Address: omegaPortal, DeployHeight: 13932203},
			{ChainID: evmchain.IDHolesky, Address: omegaPortal, DeployHeight: 2130892},
			{ChainID: evmchain.IDOpSepolia, Address: omegaPortal, DeployHeight: 15915062},
		},
		ConsensusGenesisJSON: omegaConsensusGenesisJSON,
		ConsensusSeedTXT:     omegaConsensusSeedsTXT,
		ConsensusArchiveTXT:  omegaConsensusArchivesTXT,
		ExecutionGenesisJSON: omegaExecutionGenesisJSON,
		ExecutionSeedTXT:     omegaExecutionSeedsTXT,
	},
	Mainnet: {
		Network:              Mainnet,
		version:              "v0.0.1",
		AVSContractAddress:   mainnetAVS,
		OmniExecutionChainID: evmchain.IDOmniMainnet,
		MaxValidators:        maxValidators,
		TokenAddress:         mainnetToken,
	},
}

// Version returns the version of the network.
func (s Static) Version() string {
	if s.version == "" {
		log.Warn(context.Background(), "Using unset network version. Must call netcong/genesis::Init()", nil, "network", s.Network)
		// Or panic? panic(errors.New("using unset network version - must call genesis.Init()", "network", s.Network))
	}

	return s.version
}

// ConsensusChainIDStr2Uint64 parses the uint suffix from the provided a consensus chain ID string.
func ConsensusChainIDStr2Uint64(id string) (uint64, error) {
	if !strings.HasPrefix(id, consensusIDPrefix) {
		return 0, errors.New("invalid consensus chain ID", "id", id)
	}

	suffix := strings.TrimPrefix(id, consensusIDPrefix)

	resp, err := strconv.ParseUint(suffix, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "parse consensus chain ID", "id", id)
	}

	return resp, nil
}

// IsOmniConsensus returns true if provided chainID is the omni consensus chain for the network.
func IsOmniConsensus(network ID, chainID uint64) bool {
	return network.Static().OmniConsensusChainIDUint64() == chainID
}

// IsOmniExecution returns true if provided chainID is the omni execution chain for the network.
func IsOmniExecution(network ID, chainID uint64) bool {
	return network.Static().OmniExecutionChainID == chainID
}

// SimnetNetwork defines the simnet network configuration.
// Simnet is the only statically defined network, all others are dynamically defined in manifests.
func SimnetNetwork() Network {
	return Network{
		ID: Simnet,
		Chains: []Chain{
			mustSimnetChain(Simnet.Static().OmniExecutionChainID, xchain.ShardFinalized0),
			mustSimnetChain(evmchain.IDMockL1, xchain.ShardFinalized0, xchain.ShardLatest0),
			mustSimnetChain(evmchain.IDMockL2, xchain.ShardFinalized0),
			Simnet.Static().OmniConsensusChain(),
		},
	}
}

// GetGenesis returns the genesis files for the network if they are set.
// If they are not set, or set but invalid, it returns an err.
func Genesis(network ID) (cosmgen.AppGenesis, ethcore.Genesis, error) {
	var consensus cosmgen.AppGenesis
	var execution ethcore.Genesis

	static := statics[network]

	if static.ExecutionGenesisJSON == nil || static.ConsensusGenesisJSON == nil {
		return consensus, execution, errors.New("genesis not set", "network", network)
	}

	err := json.Unmarshal(static.ConsensusGenesisJSON, &consensus)
	if err != nil {
		return consensus, execution, errors.Wrap(err, "unmarshal consensus genesis")
	}

	err = json.Unmarshal(static.ExecutionGenesisJSON, &execution)
	if err != nil {
		return consensus, execution, errors.Wrap(err, "unmarshal execution genesis")
	}

	return consensus, execution, nil
}

// SetEphemeralGenesis sets the genesis files for the network.
func SetEphemeralGenesis(network ID, consensus cosmgen.AppGenesis, execution ethcore.Genesis) error {
	if network.IsProtected() {
		return errors.New("cannot set genesis for protected network", "network", network)
	}

	consensusBz, err := json.Marshal(consensus)
	if err != nil {
		return errors.Wrap(err, "marshal consensus genesis")
	}

	executionBz, err := json.Marshal(execution)
	if err != nil {
		return errors.Wrap(err, "marshal execution genesis")
	}

	return SetEphemeralGenesisBz(network, executionBz, consensusBz)
}

func mustSimnetChain(id uint64, shards ...xchain.ShardID) Chain {
	meta, ok := evmchain.MetadataByID(id)
	if !ok {
		panic("missing chain metadata")
	}

	return Chain{
		ID:          meta.ChainID,
		Name:        meta.Name,
		BlockPeriod: time.Millisecond * 500, // Speed up block times for testing
		Shards:      shards,
	}
}

// SetEphemeralGenesisBz sets the genesis files (in bytes) for the network.
func SetEphemeralGenesisBz(network ID, executionBz, consensusBz []byte) error {
	if network.IsProtected() {
		return errors.New("cannot set ephemeral genesis for protected network", "network", network)
	}

	static := statics[network]
	static.ExecutionGenesisJSON = executionBz
	static.ConsensusGenesisJSON = consensusBz

	// set staging version to hash of consensus genesis. could also use timestamp.
	if network == Staging {
		static.version = crypto.Keccak256Hash(consensusBz).Hex()
	}

	statics[network] = static

	return nil
}
