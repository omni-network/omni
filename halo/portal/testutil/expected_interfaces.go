//go:generate mockgen -source ./expected_interfaces.go -package testutil -destination ./mock_interfaces.go
package testutil

import (
	"github.com/omni-network/omni/halo/portal/types"
)

type Portal interface {
	types.EmitPortal
}
