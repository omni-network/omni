//go:generate mockgen -source ./expected_interfaces.go -package testutil -destination ./mock_interfaces.go
package testutil

import (
	"github.com/omni-network/omni/halo/attest/types"
	rtypes "github.com/omni-network/omni/halo/registry/types"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/xchain"

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
	ChainName(chainVer xchain.ChainVersion) string
}

type Registry interface {
	rtypes.PortalRegistry
}
