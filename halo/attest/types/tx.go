package types

import (
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

const (
	StatusUnknown  uint32 = 0
	StatusPending  uint32 = 1
	StatusApproved uint32 = 2
)

func (v *Vote) AttestationRoot() (common.Hash, error) {
	if v == nil {
		return common.Hash{}, errors.New("nil vote")
	}

	msgRoot, err := cast.Array32(v.MsgRoot)
	if err != nil {
		return common.Hash{}, err
	}

	header, err := v.BlockHeader.ToXChain()
	if err != nil {
		return common.Hash{}, err
	}

	return xchain.AttestationRoot(v.AttestHeader.ToXChain(), header, msgRoot)
}

func (v *Vote) Verify() error {
	if v == nil {
		return errors.New("nil attestation")
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

	attRoot, err := v.AttestationRoot()
	if err != nil {
		return err
	}

	if err := v.Signature.Verify(); err != nil {
		return errors.Wrap(err, "verify signature")
	}

	sigTuple, err := v.Signature.ToXChain()
	if err != nil {
		return err
	}

	ok, err := k1util.Verify(
		sigTuple.ValidatorAddress,
		attRoot,
		sigTuple.Signature,
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

	if (h.ConfLevel >> 8) > 0 { // Ensure conf level is max 1 byte
		return errors.New("invalid conf level")
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

func (h *BlockHeader) ToXChain() (xchain.BlockHeader, error) {
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

func (s *SigTuple) ValidatorEthAddress() (common.Address, error) {
	if s == nil {
		return common.Address{}, errors.New("nil sig tuple")
	}

	return cast.EthAddress(s.ValidatorAddress)
}

func (s *SigTuple) ToXChain() (xchain.SigTuple, error) {
	if s == nil {
		return xchain.SigTuple{}, errors.New("nil sig tuple")
	}

	addr, err := cast.EthAddress(s.ValidatorAddress)
	if err != nil {
		return xchain.SigTuple{}, err
	}

	sig, err := cast.Array65(s.Signature)
	if err != nil {
		return xchain.SigTuple{}, err
	}

	return xchain.SigTuple{
		ValidatorAddress: addr,
		Signature:        sig,
	}, nil
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

		sigTup, err := sig.ToXChain()
		if err != nil {
			return err
		}

		ok, err := k1util.Verify(
			sigTup.ValidatorAddress,
			attRoot,
			sigTup.Signature,
		)
		if err != nil {
			return err
		} else if !ok {
			return errors.New("invalid attestation signature")
		}

		if duplicateVals[sigTup.ValidatorAddress] {
			return errors.New("duplicate validator signature", "validator", sigTup.ValidatorAddress)
		}
		duplicateVals[sigTup.ValidatorAddress] = true
	}

	return nil
}

func (a *AggVote) AttestationRoot() (common.Hash, error) {
	if a == nil {
		return common.Hash{}, errors.New("nil agg vote")
	}

	msgRoot, err := cast.Array32(a.MsgRoot)
	if err != nil {
		return common.Hash{}, err
	}

	blockHeader, err := a.BlockHeader.ToXChain()
	if err != nil {
		return common.Hash{}, err
	}

	return xchain.AttestationRoot(a.AttestHeader.ToXChain(), blockHeader, msgRoot)
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

	if a.BlockHeader.ChainId != a.AttestHeader.SourceChainId {
		return errors.New("chain ID mismatch")
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

	attRoot, err := a.AttestationRoot()
	if err != nil {
		return err
	}

	duplicateVals := make(map[common.Address]bool)
	duplicateSig := make(map[string]bool)
	for _, sig := range a.Signatures {
		valEthAddr, err := sig.ValidatorEthAddress()
		if err != nil {
			return err
		}

		if duplicateVals[valEthAddr] {
			return errors.New("duplicate validator signature")
		}

		if duplicateSig[sig.String()] {
			return errors.New("duplicate attestation signature")
		}

		if err := sig.Verify(); err != nil {
			return errors.Wrap(err, "signature")
		}

		sig65, err := cast.Array65(sig.Signature)
		if err != nil {
			return err
		}

		ok, err := k1util.Verify(
			valEthAddr,
			attRoot,
			sig65,
		)
		if err != nil {
			return err
		} else if !ok {
			return errors.New("invalid attestation signature")
		}

		duplicateVals[valEthAddr] = true
		duplicateSig[sig.String()] = true
	}

	return nil
}

func (a *Attestation) ToXChain() (xchain.Attestation, error) {
	if a == nil {
		return xchain.Attestation{}, errors.New("nil attestation")
	}

	sigs := make([]xchain.SigTuple, 0, len(a.Signatures))
	for _, sig := range a.Signatures {
		sigTuple, err := sig.ToXChain()
		if err != nil {
			return xchain.Attestation{}, err
		}

		sigs = append(sigs, sigTuple)
	}

	msgRoot, err := cast.Array32(a.MsgRoot)
	if err != nil {
		return xchain.Attestation{}, err
	}

	blockHeader, err := a.BlockHeader.ToXChain()
	if err != nil {
		return xchain.Attestation{}, err
	}

	return xchain.Attestation{
		AttestHeader:   a.AttestHeader.ToXChain(),
		BlockHeader:    blockHeader,
		ValidatorSetID: a.ValidatorSetId,
		MsgRoot:        msgRoot,
		Signatures:     sigs,
	}, nil
}

func (a *Attestation) AttestationRoot() (common.Hash, error) {
	if a == nil {
		return common.Hash{}, errors.New("nil attestation")
	}

	msgRoot, err := cast.Array32(a.MsgRoot)
	if err != nil {
		return common.Hash{}, err
	}

	blockHeader, err := a.BlockHeader.ToXChain()
	if err != nil {
		return common.Hash{}, err
	}

	return xchain.AttestationRoot(a.AttestHeader.ToXChain(), blockHeader, msgRoot)
}

func (v *Vote) ToXChain() (xchain.Vote, error) {
	if v == nil {
		return xchain.Vote{}, errors.New("nil vote")
	}

	msgRoot, err := cast.Array32(v.MsgRoot)
	if err != nil {
		return xchain.Vote{}, err
	}

	blockHeader, err := v.BlockHeader.ToXChain()
	if err != nil {
		return xchain.Vote{}, err
	}

	sigTuple, err := v.Signature.ToXChain()
	if err != nil {
		return xchain.Vote{}, err
	}

	return xchain.Vote{
		AttestHeader: v.AttestHeader.ToXChain(),
		BlockHeader:  blockHeader,
		MsgRoot:      msgRoot,
		Signature:    sigTuple,
	}, nil
}
