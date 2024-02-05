package types

import (
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

// InfrastructureData wraps e2e.InfrastructureData with additional omni-specific fields.
// TODO(corver): Maybe remove this type if not used.
type InfrastructureData struct {
	e2e.InfrastructureData
}
