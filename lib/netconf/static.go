package netconf

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	"github.com/google/uuid"

	_ "embed"
)

const consensusIDPrefix = "omni-"
const consensusIDOffset = 1_000_000
const maxValidators = 30

// Static defines static config and data for a network.
type Static struct {
	Network              ID
	Version              string
	OmniExecutionChainID uint64
	AVSContractAddress   common.Address
	TokenAddress         common.Address
	L1BridgeAddress      common.Address
	Portals              []Deployment
	MaxValidators        uint32
	ConsensusGenesisJSON []byte
	ConsensusSeedTXT     []byte
	ExecutionGenesisJSON []byte
	ExecutionSeedTXT     []byte
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

// Use random runid for staging version.
//
//nolint:gochecknoglobals // Static ID
var runid = uuid.New().String()

//nolint:gochecknoglobals // Static addresses
var (
	omegaAVS     = common.HexToAddress("0xa7b2e7830C51728832D33421670DbBE30299fD92")
	mainnetAVS   = common.HexToAddress("0xed2f4d90b073128ae6769a9A8D51547B1Df766C8")
	omegaPortal  = common.HexToAddress("0x0eEaA62754A0628C062F45bd7605E57a9d3e3813")
	mainnetToken = common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4")

	//go:embed omega/consensus-genesis.json
	omegaConsensusGenesisJSON []byte

	//go:embed omega/consensus-seeds.txt
	omegaConsensusSeedsTXT []byte

	//go:embed omega/execution-genesis.json
	omegaExecutionGenesisJSON []byte

	//go:embed omega/execution-seeds.txt
	omegaExecutionSeedsTXT []byte

	//go:embed staging/consensus-seeds.txt
	stagingConsensusSeedsTXT []byte

	//go:embed staging/execution-seeds.txt
	stagingExecutionSeedsTXT []byte
)

//nolint:gochecknoglobals // Static mappings.
var statics = map[ID]Static{
	Simnet: {
		Network:              Simnet,
		Version:              "simnet",
		OmniExecutionChainID: evmchain.IDOmniEphemeral,
		MaxValidators:        maxValidators,
	},
	Devnet: {
		Network:              Devnet,
		Version:              "devnet",
		OmniExecutionChainID: evmchain.IDOmniEphemeral,
		MaxValidators:        maxValidators,
	},
	Staging: {
		Network:              Staging,
		Version:              runid,
		OmniExecutionChainID: evmchain.IDOmniEphemeral,
		MaxValidators:        maxValidators,
		ConsensusSeedTXT:     stagingConsensusSeedsTXT,
		ExecutionSeedTXT:     stagingExecutionSeedsTXT,
	},
	Omega: {
		Network:              Omega,
		Version:              "v0.0.8",
		AVSContractAddress:   omegaAVS,
		OmniExecutionChainID: evmchain.IDOmniOmega,
		MaxValidators:        maxValidators,
		Portals: []Deployment{
			{ChainID: evmchain.IDArbSepolia, Address: omegaPortal, DeployHeight: 64990937},
			{ChainID: evmchain.IDBaseSepolia, Address: omegaPortal, DeployHeight: 12808134},
			{ChainID: evmchain.IDHolesky, Address: omegaPortal, DeployHeight: 1958423},
			{ChainID: evmchain.IDOpSepolia, Address: omegaPortal, DeployHeight: 14791006},
		},
		ConsensusGenesisJSON: omegaConsensusGenesisJSON,
		ConsensusSeedTXT:     omegaConsensusSeedsTXT,
		ExecutionGenesisJSON: omegaExecutionGenesisJSON,
		ExecutionSeedTXT:     omegaExecutionSeedsTXT,
	},
	Mainnet: {
		Network:              Mainnet,
		Version:              "v0.0.1",
		AVSContractAddress:   mainnetAVS,
		OmniExecutionChainID: evmchain.IDOmniMainnet,
		MaxValidators:        maxValidators,
		TokenAddress:         mainnetToken,
	},
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
			mustSimnetChain(evmchain.IDMockL1Fast, xchain.ShardFinalized0, xchain.ShardLatest0),
			mustSimnetChain(evmchain.IDMockL2, xchain.ShardFinalized0),
			Simnet.Static().OmniConsensusChain(),
		},
	}
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

// SetEphemeralGenesis sets the ephemeral genesis files.
func SetEphemeralGenesis(network ID, execution, consensus []byte) error {
	if network.IsProtected() {
		return errors.New("cannot set ephemeral genesis for protected network", "network", network)
	}

	static := statics[network]
	static.ExecutionGenesisJSON = execution
	static.ConsensusGenesisJSON = consensus
	statics[network] = static

	return nil
}
