package types

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

// UniqueKey uniquely identifies a vote/aggregate vote/attestation.
type UniqueKey struct {
	xchain.BlockHeader
	AttestationRoot common.Hash
}

// UniqueKey returns a unique key for the vote
// It panics if the vote is invalid. Ensure Verify() is called before UniqueKey().
func (v *Vote) UniqueKey() UniqueKey {
	return UniqueKey{
		BlockHeader:     v.BlockHeader.ToXChain(),
		AttestationRoot: common.Hash(v.AttestationRoot),
	}
}

func (v *Vote) Verify() error {
	if v == nil {
		return errors.New("nil attestation")
	}

	if v.BlockHeader == nil {
		return errors.New("nil block header")
	}

	if v.Signature == nil {
		return errors.New("nil signature")
	}

	if len(v.AttestationRoot) != len(common.Hash{}) {
		return errors.New("invalid attestation root length")
	}

	if len(v.BlockHeader.Hash) != len(common.Hash{}) {
		return errors.New("invalid block header hash length")
	}

	if len(v.Signature.Signature) != len(xchain.Signature65{}) {
		return errors.New("invalid signature length")
	}

	if len(v.Signature.ValidatorAddress) != len(common.Address{}) {
		return errors.New("invalid validator address length")
	}

	ok, err := k1util.Verify(
		common.Address(v.Signature.ValidatorAddress),
		common.Hash(v.AttestationRoot),
		xchain.Signature65(v.Signature.Signature),
	)
	if err != nil {
		return err
	} else if !ok {
		return errors.New("invalid attestation signature")
	}

	return nil
}

func (h *BlockHeader) Verify() error {
	if h == nil {
		return errors.New("nil block header")
	}

	if len(h.Hash) != len(common.Hash{}) {
		return errors.New("invalid block header hash length")
	}

	return nil
}

func (h *BlockHeader) ToXChain() xchain.BlockHeader {
	return xchain.BlockHeader{
		SourceChainID: h.ChainId,
		BlockHeight:   h.Height,
		BlockHash:     common.BytesToHash(h.Hash),
	}
}

func (s *SigTuple) Verify() error {
	if s == nil {
		return errors.New("nil sig tuple")
	}

	if len(s.ValidatorAddress) != len(common.Address{}) {
		return errors.New("invalid validator address length")
	}

	if len(s.Signature) != len(xchain.Signature65{}) {
		return errors.New("invalid signature length")
	}

	return nil
}

func (s *SigTuple) ToXChain() xchain.SigTuple {
	return xchain.SigTuple{
		ValidatorAddress: common.Address(s.ValidatorAddress),
		Signature:        xchain.Signature65(s.Signature),
	}
}

func (a *AggVote) Verify() error {
	if a == nil {
		return errors.New("nil aggregate vote")
	}

	if err := a.BlockHeader.Verify(); err != nil {
		return errors.Wrap(err, "block header")
	}

	if len(a.AttestationRoot) != len(common.Hash{}) {
		return errors.New("invalid attestation root length")
	}

	if len(a.Signatures) == 0 {
		return errors.New("empty signatures")
	}

	for _, sig := range a.Signatures {
		if err := sig.Verify(); err != nil {
			return errors.Wrap(err, "signature")
		}
	}

	return nil
}

func (a *Attestation) Verify() error {
	if a == nil {
		return errors.New("nil attestation")
	}

	if err := a.BlockHeader.Verify(); err != nil {
		return errors.Wrap(err, "block header")
	}

	if len(a.AttestationRoot) != len(common.Hash{}) {
		return errors.New("invalid attestation root length")
	}

	if a.ValidatorSetId == 0 {
		return errors.New("zero validator set ID")
	}

	if len(a.Signatures) == 0 {
		return errors.New("empty signatures")
	}

	for _, sig := range a.Signatures {
		if err := sig.Verify(); err != nil {
			return errors.Wrap(err, "signature")
		}
	}

	return nil
}

func (a *Attestation) ToXChain() xchain.Attestation {
	sigs := make([]xchain.SigTuple, 0, len(a.Signatures))
	for _, sig := range a.Signatures {
		sigs = append(sigs, sig.ToXChain())
	}

	return xchain.Attestation{
		BlockHeader:     a.BlockHeader.ToXChain(),
		ValidatorSetID:  a.ValidatorSetId,
		AttestationRoot: common.Hash(a.AttestationRoot),
		Signatures:      sigs,
	}
}

func (v *Vote) ToXChain() xchain.Vote {
	return xchain.Vote{
		BlockHeader:     v.BlockHeader.ToXChain(),
		AttestationRoot: common.Hash(v.AttestationRoot),
		Signature:       v.Signature.ToXChain(),
	}
}
