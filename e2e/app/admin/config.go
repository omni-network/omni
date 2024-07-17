package admin

import (
	"github.com/omni-network/omni/lib/errors"
)

const chainAll = "all"

type PortalAdminConfig struct {
	Chain string // Name of chain to run admin command on, use "all" to run on all chains
}

func DefaultPortalAdminConfig() PortalAdminConfig {
	return PortalAdminConfig{
		Chain: "",
	}
}

func (cfg PortalAdminConfig) Validate() error {
	if cfg.Chain == "" {
		return errors.New("chain must be set")
	}

	return nil
}
