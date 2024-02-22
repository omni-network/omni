package types

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

func (a *Vote) Verify() error {
	if a == nil {
		return errors.New("nil attestation")
	}

	if a.BlockHeader == nil {
		return errors.New("nil block header")
	}

	if a.Signature == nil {
		return errors.New("nil signature")
	}

	if len(a.BlockRoot) != len(common.Hash{}) {
		return errors.New("invalid block root length")
	}

	if len(a.BlockHeader.Hash) != len(common.Hash{}) {
		return errors.New("invalid block header hash length")
	}

	if len(a.Signature.Signature) != len(xchain.Signature65{}) {
		return errors.New("invalid signature length")
	}

	if len(a.Signature.ValidatorAddress) != len(common.Address{}) {
		return errors.New("invalid validator address length")
	}

	ok, err := k1util.Verify(
		common.Address(a.Signature.ValidatorAddress),
		common.Hash(a.BlockRoot),
		xchain.Signature65(a.Signature.Signature),
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

func (m *SigTuple) Verify() error {
	if m == nil {
		return errors.New("nil sig tuple")
	}

	if len(m.ValidatorAddress) != len(common.Address{}) {
		return errors.New("invalid validator address length")
	}

	if len(m.Signature) != len(xchain.Signature65{}) {
		return errors.New("invalid signature length")
	}

	return nil
}

func (m *SigTuple) ToXChain() xchain.SigTuple {
	return xchain.SigTuple{
		ValidatorAddress: common.Address(m.ValidatorAddress),
		Signature:        xchain.Signature65(m.Signature),
	}
}

func (a *AggVote) Verify() error {
	if a == nil {
		return errors.New("nil aggregate attestation")
	}

	if err := a.BlockHeader.Verify(); err != nil {
		return errors.Wrap(err, "block header")
	}

	if len(a.BlockRoot) != len(common.Hash{}) {
		return errors.New("invalid block root length")
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
		return errors.New("nil aggregate attestation")
	}

	if err := a.BlockHeader.Verify(); err != nil {
		return errors.Wrap(err, "block header")
	}

	if len(a.BlockRoot) != len(common.Hash{}) {
		return errors.New("invalid block root length")
	}

	if len(a.ValidatorsHash) != len(common.Hash{}) {
		return errors.New("invalid validator set hash length")
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

func (a *Attestation) ToXChain() xchain.AggAttestation {
	sigs := make([]xchain.SigTuple, 0, len(a.Signatures))
	for _, sig := range a.Signatures {
		sigs = append(sigs, sig.ToXChain())
	}

	return xchain.AggAttestation{
		BlockHeader:      a.BlockHeader.ToXChain(),
		ValidatorSetHash: common.Hash(a.ValidatorsHash),
		BlockRoot:        common.Hash(a.BlockRoot),
		Signatures:       sigs,
	}
}

func (a *Vote) ToXChain() xchain.Attestation {
	return xchain.Attestation{
		BlockHeader: a.BlockHeader.ToXChain(),
		BlockRoot:   common.Hash(a.BlockRoot),
		Signature:   a.Signature.ToXChain(),
	}
}
