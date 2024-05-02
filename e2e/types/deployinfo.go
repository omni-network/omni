package types

import (
	"encoding/json"
	"os"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

type ContractName string

const (
	ContractPortal ContractName = "portal"
)

// DeployInfos contains the addresses of deployed xdapps and contracts by chainID.
//
// Note that ContractPortal is duplicated both here and in netconf.Chain. That is because netconf
// is used by halo and relayer and other production apps.
//
// This DeployInfos is similar to netconf, but it is only used by e2e deployment and tests.
// As soon as any production app needs to use any of these addresses, they should be moved to netconf.
// TODO(corver): Remove since this is stored in on-chain PortalRegistry.
type DeployInfos map[uint64]map[ContractName]DeployInfo

// DeployInfo contains the address and deploy height of a deployed contract.
type DeployInfo struct {
	Address common.Address
	Height  uint64
}

func (i DeployInfos) Addr(chainID uint64, contract ContractName) (common.Address, bool) {
	info, ok := i[chainID][contract]
	return info.Address, ok
}

func (i DeployInfos) Set(chainID uint64, contract ContractName, addr common.Address, height uint64) {
	if i[chainID] == nil {
		i[chainID] = make(map[ContractName]DeployInfo)
	}

	i[chainID][contract] = DeployInfo{
		Address: addr,
		Height:  height,
	}
}

func (i DeployInfos) Save(file string) error {
	bz, err := json.Marshal(i)
	if err != nil {
		return errors.Wrap(err, "marshal deploy info")
	}

	err = os.WriteFile(file, bz, 0644)
	if err != nil {
		return errors.Wrap(err, "write deploy info")
	}

	return nil
}

func LoadDeployInfos(file string) (DeployInfos, error) {
	bz, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "read deploy info")
	}

	var i DeployInfos
	err = json.Unmarshal(bz, &i)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal deploy info")
	}

	return i, nil
}
