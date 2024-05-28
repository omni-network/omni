package keeper

import "github.com/omni-network/omni/lib/xchain"

func (a *Attestation) XChainConfLevel() xchain.ConfLevel {
	return xchain.ConfLevel(a.GetConfLevel())
}

func (a *Attestation) XChainVersion() xchain.ChainVersion {
	return xchain.ChainVersion{
		ID:        a.GetChainId(),
		ConfLevel: a.XChainConfLevel(),
	}
}

func (a *Attestation) XChainConfLevelStr() string {
	return xchain.ConfLevel(a.GetConfLevel()).String()
}
