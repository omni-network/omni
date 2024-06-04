package keeper

import "github.com/omni-network/omni/lib/xchain"

func (a *Attestation) XChainVersion() xchain.ChainVersion {
	return xchain.ChainVersion{
		ID:        a.GetChainId(),
		ConfLevel: xchain.ConfLevel(a.GetConfLevel()),
	}
}
