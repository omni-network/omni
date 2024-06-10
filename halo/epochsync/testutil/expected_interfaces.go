//go:generate mockgen -source ./expected_interfaces.go -package testutil -destination ./mock_interfaces.go
package testutil

import (
	"github.com/omni-network/omni/halo/epochsync/types"
)

type StakingKeeper interface {
	types.StakingKeeper
}

type AttestKeeper interface {
	types.AttestKeeper
}

type Subscriber interface {
	types.ValSetSubscriber
}
