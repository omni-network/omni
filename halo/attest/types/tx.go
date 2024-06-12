package types

import (
	"crypto/sha256"
	"encoding/binary"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

// UniqueKey returns a unique key for the vote
// It panics if the vote is invalid. Ensure Verify() is called before UniqueKey().
func (v *Vote) UniqueKey() [32]byte {
	return uniqueVoteHash(v.BlockHeader, v.AttestationRoot)
}

func (v *Vote) Verify() error {
	if v == nil {
		return errors.New("nil attestation")
	}

	if err := v.BlockHeader.Verify(); err != nil {
		return err
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

	if conf := xchain.ConfLevel(byte(h.GetConfLevel())); !conf.Valid() {
		return errors.New("invalid conf level", "conf_level", conf.String())
	}

	return nil
}

func (h *BlockHeader) ToXChain() xchain.BlockHeader {
	return BlockHeaderFromProto(h)
}

// UniqueKey returns a unique key for the block header.
// It panics if the vote is invalid. Ensure Verify() is called before UniqueKey().
func (h *BlockHeader) UniqueKey() [32]byte {
	return uniqueVoteHash(h, nil)
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

// UniqueKey returns a unique key for the agg vote
// It panics if the vote is invalid. Ensure Verify() is called before UniqueKey().
func (a *AggVote) UniqueKey() [32]byte {
	return uniqueVoteHash(a.BlockHeader, a.AttestationRoot)
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

// UniqueKey returns a unique key for the agg vote
// It panics if the vote is invalid. Ensure Verify() is called before UniqueKey().
func (a *Attestation) UniqueKey() [32]byte {
	return uniqueVoteHash(a.BlockHeader, a.AttestationRoot)
}

func (v *Vote) ToXChain() xchain.Vote {
	return xchain.Vote{
		BlockHeader:     v.BlockHeader.ToXChain(),
		AttestationRoot: common.Hash(v.AttestationRoot),
		Signature:       v.Signature.ToXChain(),
	}
}

// uniqueVoteHash returns a hash that uniquely identifies the vote attestation.
//
// The unique key hash is required since the attestation root cannot be verified on-chain.
// So instead, a unique key hash is calculated from on-chain metadata to uniquely identify the attestation and group votes.
func uniqueVoteHash(header *BlockHeader, attRoot []byte) [32]byte {
	h := sha256.New()

	_ = binary.Write(h, binary.BigEndian, header.ChainId)
	_ = binary.Write(h, binary.BigEndian, header.Height)
	_ = binary.Write(h, binary.BigEndian, header.Offset)
	_ = binary.Write(h, binary.BigEndian, header.ConfLevel)
	_, _ = h.Write(header.Hash)

	_, _ = h.Write(attRoot)

	return [32]byte(h.Sum(nil))
}
