package netconf

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/common"

	"github.com/google/uuid"

	_ "embed"
)

const consensusIDPrefix = "omni-"
const consensusIDOffset = 1_000_000
const maxValidators = 10

// Static defines static config and data for a network.
type Static struct {
	Version              string
	OmniExecutionChainID uint64
	AVSContractAddress   common.Address
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

// OmniConsensusChainIDUint64 returns the chain ID uint64 for the Omni consensus chain.
// It is calculated as 1_000_000 + OmniExecutionChainID.
func (s Static) OmniConsensusChainIDUint64() uint64 {
	return consensusIDOffset + s.OmniExecutionChainID
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

// Use random runid for staging version.
//
//nolint:gochecknoglobals // Static ID
var runid = uuid.New().String()

//nolint:gochecknoglobals // Static addresses
var (
	// Address matches lib/contracts.TestnetPortal() and lib/contracts.TestnetAVS().
	// We do not import to avoid cylic dependencies.
	testnetPortal = common.HexToAddress("0xFf22F3532C19a6f890c52c4CfcDB94007aA471Dc")
	testnetAVS    = common.HexToAddress("0xa7b2e7830C51728832D33421670DbBE30299fD92")

	// This address DOES NOT match lib/contracts.MainnetAVS().
	// This mainnet AVS was deployed outside of the e2e deployment flow, without Create3.
	mainnetAVS = common.HexToAddress("0xed2f4d90b073128ae6769a9A8D51547B1Df766C8")

	//go:embed testnet/consensus-genesis.json
	testnetConsensusGenesisJSON []byte

	//go:embed testnet/consensus-seeds.txt
	testnetConsensusSeedsTXT []byte

	//go:embed testnet/execution-genesis.json
	testnetExecutionGenesisJSON []byte

	//go:embed testnet/execution-seeds.txt
	testnetExecutionSeedsTXT []byte
)

//nolint:gochecknoglobals // Static mappings.
var statics = map[ID]Static{
	Simnet: {
		Version:              "simnet",
		OmniExecutionChainID: evmchain.IDOmniEphemeral,
		MaxValidators:        maxValidators,
	},
	Devnet: {
		Version:              "devnet",
		OmniExecutionChainID: evmchain.IDOmniEphemeral,
		MaxValidators:        maxValidators,
	},
	Staging: {
		Version:              runid,
		OmniExecutionChainID: evmchain.IDOmniEphemeral,
		MaxValidators:        maxValidators,
	},
	Testnet: {
		Version:              "v0.0.2",
		AVSContractAddress:   testnetAVS,
		OmniExecutionChainID: evmchain.IDOmniTestnet,
		MaxValidators:        maxValidators,
		Portals: []Deployment{
			{
				ChainID:      evmchain.IDHolesky,
				Address:      testnetPortal,
				DeployHeight: 1357819,
			},
			{
				ChainID:      evmchain.IDOpSepolia,
				Address:      testnetPortal,
				DeployHeight: 10731455,
			},
			{
				ChainID:      evmchain.IDArbSepolia,
				Address:      testnetPortal,
				DeployHeight: 34237972,
			},
		},
		ConsensusGenesisJSON: testnetConsensusGenesisJSON,
		ConsensusSeedTXT:     testnetConsensusSeedsTXT,
		ExecutionGenesisJSON: testnetExecutionGenesisJSON,
		ExecutionSeedTXT:     testnetExecutionSeedsTXT,
	},
	Mainnet: {
		Version:            "v0.0.1",
		AVSContractAddress: mainnetAVS,
		MaxValidators:      maxValidators,
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
