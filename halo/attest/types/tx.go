package types

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

func (v *Vote) AttestationRoot() (common.Hash, error) {
	return xchain.AttestationRoot(v.AttestHeader.ToXChain(), v.BlockHeader.ToXChain(), common.Hash(v.MsgRoot))
}

func (v *Vote) Verify() error {
	if v == nil {
		return errors.New("nil attestation")
	}
	if v.Signature == nil {
		return errors.New("nil signature")
	}

	if len(v.MsgRoot) != len(common.Hash{}) {
		return errors.New("invalid message root length")
	}

	if err := v.BlockHeader.Verify(); err != nil {
		return errors.Wrap(err, "verify block header")
	}

	if err := v.AttestHeader.Verify(); err != nil {
		return errors.Wrap(err, "verify attestation header")
	}

	if len(v.Signature.Signature) != len(xchain.Signature65{}) {
		return errors.New("invalid signature length")
	}

	if len(v.Signature.ValidatorAddress) != len(common.Address{}) {
		return errors.New("invalid validator address length")
	}

	attRoot, err := v.AttestationRoot()
	if err != nil {
		return err
	}

	ok, err := k1util.Verify(
		common.Address(v.Signature.ValidatorAddress),
		attRoot,
		xchain.Signature65(v.Signature.Signature),
	)
	if err != nil {
		return err
	} else if !ok {
		return errors.New("invalid attestation signature")
	}

	return nil
}

func (h *AttestHeader) Verify() error {
	if h == nil {
		return errors.New("nil attestation header")
	}

	if h.SourceChainId == 0 {
		return errors.New("empty source chain id")
	}

	if conf := xchain.ConfLevel(byte(h.GetConfLevel())); !conf.Valid() {
		return errors.New("invalid conf level", "conf_level", conf.String())
	}

	if h.ConsensusChainId == 0 {
		return errors.New("zero consensus chain ID")
	}

	if h.AttestOffset == 0 {
		return errors.New("zero attestation offset")
	}

	return nil
}

func (h *AttestHeader) ToXChain() xchain.AttestHeader {
	return AttestHeaderFromProto(h)
}

func (h *BlockHeader) Verify() error {
	if h == nil {
		return errors.New("nil block header")
	}

	if len(h.BlockHash) != len(common.Hash{}) {
		return errors.New("invalid block header hash length")
	}

	return nil
}

func (h *BlockHeader) ToXChain() xchain.BlockHeader {
	return BlockHeaderFromProto(h)
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
		return errors.Wrap(err, "verify block header")
	}

	if err := a.AttestHeader.Verify(); err != nil {
		return errors.Wrap(err, "verify block header")
	}

	if len(a.MsgRoot) != len(common.Hash{}) {
		return errors.New("invalid message root length")
	}

	if len(a.Signatures) == 0 {
		return errors.New("empty signatures")
	}

	attRoot, err := a.AttestationRoot()
	if err != nil {
		return err
	}

	duplicateVals := make(map[common.Address]bool)
	for _, sig := range a.Signatures {
		if err := sig.Verify(); err != nil {
			return errors.Wrap(err, "signature")
		}

		addr := common.BytesToAddress(sig.ValidatorAddress)

		ok, err := k1util.Verify(
			addr,
			attRoot,
			xchain.Signature65(sig.Signature),
		)
		if err != nil {
			return err
		} else if !ok {
			return errors.New("invalid attestation signature")
		}

		if duplicateVals[addr] {
			return errors.New("duplicate validator signature", "validator", addr)
		}
		duplicateVals[addr] = true
	}

	return nil
}

func (a *AggVote) AttestationRoot() (common.Hash, error) {
	return xchain.AttestationRoot(a.AttestHeader.ToXChain(), a.BlockHeader.ToXChain(), common.Hash(a.MsgRoot))
}

func (a *Attestation) Verify() error {
	if a == nil {
		return errors.New("nil attestation")
	}

	if err := a.BlockHeader.Verify(); err != nil {
		return errors.Wrap(err, "verify block header")
	}

	if err := a.AttestHeader.Verify(); err != nil {
		return errors.Wrap(err, "verify attest header")
	}

	if len(a.MsgRoot) != len(common.Hash{}) {
		return errors.New("invalid message root length")
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
		BlockHeader:    a.BlockHeader.ToXChain(),
		ValidatorSetID: a.ValidatorSetId,
		MsgRoot:        common.Hash(a.MsgRoot),
		Signatures:     sigs,
	}
}

func (a *Attestation) AttestationRoot() (common.Hash, error) {
	return xchain.AttestationRoot(a.AttestHeader.ToXChain(), a.BlockHeader.ToXChain(), common.Hash(a.MsgRoot))
}

func (v *Vote) ToXChain() xchain.Vote {
	return xchain.Vote{
		BlockHeader: v.BlockHeader.ToXChain(),
		MsgRoot:     common.Hash(v.MsgRoot),
		Signature:   v.Signature.ToXChain(),
	}
}
