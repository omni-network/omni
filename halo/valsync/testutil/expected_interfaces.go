//go:generate mockgen -source ./expected_interfaces.go -package testutil -destination ./mock_interfaces.go
package testutil

import (
	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/valsync/types"
)

type StakingKeeper interface {
	types.StakingKeeper
}

type AttestKeeper interface {
	atypes.AttestKeeper
}

type Subscriber interface {
	types.ValSetSubscriber
}
