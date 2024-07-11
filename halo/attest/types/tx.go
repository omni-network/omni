package types

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

// VoteSource identifies the source block being voted on.
// Multiple votes to the same source by the same validator qualifies as double signing as is slashable.
type VoteSource struct {
	ChainID     uint64
	ConfLevel   uint32
	BlockOffset uint64
}

// VoteSource returns the source block being voted on.
func (a *AggVote) VoteSource() VoteSource {
	return VoteSource{
		ChainID:     a.BlockHeader.GetChainId(),
		ConfLevel:   a.BlockHeader.GetConfLevel(),
		BlockOffset: a.BlockHeader.GetOffset(),
	}
}

// VoteSource returns the source block being voted on.
func (v *Vote) VoteSource() VoteSource {
	return VoteSource{
		ChainID:     v.BlockHeader.GetChainId(),
		ConfLevel:   v.BlockHeader.GetConfLevel(),
		BlockOffset: v.BlockHeader.GetOffset(),
	}
}

func (v *Vote) AttestationRoot() (common.Hash, error) {
	return xchain.AttestationRoot(v.BlockHeader.ToXChain(), common.Hash(v.MsgRoot))
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

	if len(v.MsgRoot) != len(common.Hash{}) {
		return errors.New("invalid message root length")
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

// BlockRoot returns a unique key for the block header.
func (h *BlockHeader) BlockRoot() ([32]byte, error) {
	return xchain.BlockHeaderLeaf(h.ToXChain())
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
	return xchain.AttestationRoot(a.BlockHeader.ToXChain(), common.Hash(a.MsgRoot))
}

func (a *Attestation) Verify() error {
	if a == nil {
		return errors.New("nil attestation")
	}

	if err := a.BlockHeader.Verify(); err != nil {
		return errors.Wrap(err, "block header")
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
	return xchain.AttestationRoot(a.BlockHeader.ToXChain(), common.Hash(a.MsgRoot))
}

func (v *Vote) ToXChain() xchain.Vote {
	return xchain.Vote{
		BlockHeader: v.BlockHeader.ToXChain(),
		MsgRoot:     common.Hash(v.MsgRoot),
		Signature:   v.Signature.ToXChain(),
	}
}
