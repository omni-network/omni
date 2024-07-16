package admin

import (
	"github.com/omni-network/omni/lib/errors"
)

type PortalAdminConfig struct {
	Chain string // Name of chain to run admin command on
}

func DefaultPortalAdminConfig() PortalAdminConfig {
	return PortalAdminConfig{
		Chain: "",
	}
}

func (cfg PortalAdminConfig) Validate() error {
	if cfg.Chain == "" {
		return errors.New("chain is required")
	}

	return nil
}
