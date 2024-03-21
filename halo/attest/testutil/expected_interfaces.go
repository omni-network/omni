//go:generate mockgen -source ./expected_interfaces.go -package testutil -destination ./mock_interfaces.go
package testutil

import (
	"github.com/omni-network/omni/halo/attest/types"
	vtypes "github.com/omni-network/omni/halo/valsync/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
)

type StakingKeeper interface {
	baseapp.ValidatorStore
}

type Voter interface {
	types.Voter
}

type ValProvider interface {
	vtypes.ValidatorProvider
}

type ChainNamer interface {
	ChainName(chainID uint64) string
}
