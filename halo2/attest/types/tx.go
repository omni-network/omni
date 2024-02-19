package types

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

func (a *Attestation) Verify() error {
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
