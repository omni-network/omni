package netconf

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/omni-network/omni/lib/chainids"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"

	"github.com/google/uuid"
)

const consensusIDPrefix = "omni-"
const consensusIDOffset = 1_000_000

// Static defines static config and data for a network.
type Static struct {
	Version              string
	OmniExecutionChainID uint64
	AVSContractAddress   common.Address
	Portals              []Deployment
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

// Use random runid for version in ephemeral networks.
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
)

//nolint:gochecknoglobals // Static mappings.
var statics = map[ID]Static{
	Simnet: {
		Version:              runid,
		OmniExecutionChainID: chainids.OmniDevnet,
	},
	Devnet: {
		Version:              runid,
		OmniExecutionChainID: chainids.OmniDevnet,
	},
	Staging: {
		Version:              runid,
		OmniExecutionChainID: chainids.OmniDevnet,
	},
	Testnet: {
		Version:              "v0.0.2",
		AVSContractAddress:   testnetAVS,
		OmniExecutionChainID: chainids.OmniTestnet,
		Portals: []Deployment{
			{
				ChainID:      chainids.Holesky,
				Address:      testnetPortal,
				DeployHeight: 1357819,
			},
			{
				ChainID:      chainids.OpSepolia,
				Address:      testnetPortal,
				DeployHeight: 10731455,
			},
			{
				ChainID:      chainids.ArbSepolia,
				Address:      testnetPortal,
				DeployHeight: 34237972,
			},
		},
	},
	Mainnet: {
		Version:            "v0.0.1",
		AVSContractAddress: mainnetAVS,
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
