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

// PortalDeployments returns the portal deployment for the given chainID.
// If there is none, it returns an empty deployment.
func (Static) PortalDeployment(chainID uint64) Deployment {
	for _, d := range statics[Testnet].Portals {
		if d.ChainID == chainID {
			return d
		}
	}

	return Deployment{}
}

// HasPortalDeployment returns true if there is a portal deployment for the given chainID.
func (Static) HasPortalDeployment(chainID uint64) bool {
	for _, d := range statics[Testnet].Portals {
		if d.ChainID == chainID {
			return true
		}
	}

	return false
}

// Use random runid for version in ephemeral networks.
//
//nolint:gochecknoglobals // Static ID
var runid = uuid.New().String()

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
		Version:              "v0.0.1",
		AVSContractAddress:   common.HexToAddress("0xa7b2e7830C51728832D33421670DbBE30299fD92"),
		OmniExecutionChainID: chainids.OmniTestnet,
		Portals: []Deployment{
			{
				ChainID: chainids.Holesky,
				// Address matches lib/contracts.TestnetPortal()
				// We do not import to avoid cylic dependencies.
				Address:      common.HexToAddress("0x71d510f4dc4e7E7716D03209c603C76F4398cF53"),
				DeployHeight: 1280141,
			},
		},
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

// Version returns the version for the given network.
func Version(network ID) string {
	return statics[network].Version
}
