package avs

import (
	"encoding/json"

	"github.com/omni-network/omni/e2e/app/static"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

//nolint:gochecknoglobals // static deployment json
var devnetEigenDeployments = mustDevnetEigenDeployments()

type EigenDeployments struct {
	AVSDirectory      common.Address `json:"AVSDirectory"`
	DelegationManager common.Address `json:"DelegationManager"`
}

func (e EigenDeployments) Validate() error {
	if (e.AVSDirectory == common.Address{}) {
		return errors.New("avs directory is zero")
	}
	if (e.DelegationManager == common.Address{}) {
		return errors.New("delegation manager is zero")
	}

	return nil
}

func holeskyEigenDeployments() EigenDeployments {
	return EigenDeployments{
		DelegationManager: common.HexToAddress("0xA44151489861Fe9e3055d95adC98FbD462B948e7"),
		AVSDirectory:      common.HexToAddress("0x055733000064333CaDDbC92763c58BF0192fFeBf"),
	}
}

func mustDevnetEigenDeployments() EigenDeployments {
	var el EigenDeployments
	err := json.Unmarshal(static.GetDevnetElDeployments(), &el)
	if err != nil {
		panic(err)
	}

	return el
}
