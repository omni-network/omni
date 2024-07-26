package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"
)

func (a *Attestation) XChainVersion() xchain.ChainVersion {
	return xchain.ChainVersion{
		ID:        a.GetChainId(),
		ConfLevel: xchain.ConfLevel(a.GetConfLevel()),
	}
}

func (a *Attestation) IsFuzzy() bool {
	return xchain.ConfLevel(a.GetConfLevel()).IsFuzzy()
}

func AttestationFromDB(att *Attestation, consensusChainID uint64, sigs []*Signature) *types.Attestation {
	return &types.Attestation{
		AttestHeader: &types.AttestHeader{
			ConsensusChainId: consensusChainID,
			SourceChainId:    att.GetChainId(),
			ConfLevel:        att.GetConfLevel(),
			AttestOffset:     att.GetAttestOffset(),
		},
		BlockHeader: &types.BlockHeader{
			ChainId:     att.GetChainId(),
			BlockHeight: att.GetBlockHeight(),
			BlockHash:   att.GetBlockHash(),
		},
		ValidatorSetId: att.GetValidatorSetId(),
		MsgRoot:        att.GetMsgRoot(),
		Signatures:     sigsFromDB(sigs),
	}
}

func sigsFromDB(sigs []*Signature) []*types.SigTuple {
	resp := make([]*types.SigTuple, 0, len(sigs))
	for _, sig := range sigs {
		resp = append(resp, &types.SigTuple{
			ValidatorAddress: sig.GetValidatorAddress(),
			Signature:        sig.GetSignature(),
		})
	}

	return resp
}

// lookupFunc returns the earliest or latest approved attestation offset for the given chain version.
// It returns false if none exist.
type lookupFunc func(context.Context, xchain.ChainVersion) (uint64, bool, error)

func newLatestLookupCache(k *Keeper) lookupFunc {
	return newLookupCache(k.latestAttestation)
}

func newEarliestLookupCache(k *Keeper) lookupFunc {
	return newLookupCache(k.earliestAttestation)
}

func newLookupCache(lookup func(context.Context, xchain.ChainVersion) (*Attestation, bool, error)) lookupFunc {
	cache := make(map[xchain.ChainVersion]uint64)

	return func(ctx context.Context, chainVer xchain.ChainVersion) (uint64, bool, error) {
		if offset, ok := cache[chainVer]; ok {
			return offset, offset > 0, nil
		}

		latest, ok, err := lookup(ctx, chainVer)
		if err != nil {
			return 0, false, err
		} else if !ok {
			cache[chainVer] = 0 // Populate miss
			return 0, false, nil
		} else if latest.GetAttestOffset() == 0 { // Offsets start at 1
			return 0, false, errors.New("invalid zero attestation offset [BUG]")
		}

		// Populate hit
		offset := latest.GetAttestOffset()
		cache[chainVer] = offset

		return offset, true, nil
	}
}

// fuzzyDependents returns all the fuzzy chain versions that depend on the given finalized chain version.
func fuzzyDependents(chainVer xchain.ChainVersion, confLevels map[uint64][]xchain.ConfLevel) []xchain.ChainVersion {
	if chainVer.ConfLevel.IsFuzzy() {
		// Fuzzy chain versions don't have other fuzzy chain dependents.
		return nil
	}

	var resp []xchain.ChainVersion
	for _, confLevel := range confLevels[chainVer.ID] {
		if !confLevel.IsFuzzy() {
			continue // Skip non-fuzzy chain versions
		}

		resp = append(resp, xchain.ChainVersion{
			ID:        chainVer.ID,
			ConfLevel: confLevel,
		})
	}

	return resp
}
