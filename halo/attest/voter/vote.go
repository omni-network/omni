package voter

import (
	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"
)

// CreateVote creates an vote for the given block.
func CreateVote(privKey crypto.PrivKey, block xchain.Block) (*types.Vote, error) {
	pubkey := privKey.PubKey().Bytes()
	if len(pubkey) != 33 {
		return nil, errors.New("invalid pubkey length", "length", len(pubkey))
	}

	tree, err := xchain.NewBlockTree(block)
	if err != nil {
		return nil, err
	}
	root := tree.Root()

	sig, err := k1util.Sign(privKey, root)
	if err != nil {
		return nil, errors.Wrap(err, "sign attestation")
	}

	address, err := k1util.PubKeyToAddress(privKey.PubKey())
	if err != nil {
		return nil, err
	}

	return &types.Vote{
		BlockHeader: &types.BlockHeader{
			ChainId: block.SourceChainID,
			Height:  block.BlockHeight,
			Hash:    block.BlockHash[:],
		},
		BlockRoot: root[:],
		Signature: &types.SigTuple{
			ValidatorAddress: address[:],
			Signature:        sig[:],
		},
		MsgOffsets:     msgOffsets(block),
		ReceiptOffsets: receiptOffsets(block),
	}, nil
}

func msgOffsets(block xchain.Block) []*types.MsgOffset {
	maxes := make(map[uint64]uint64)
	for _, msg := range block.Msgs {
		if msg.StreamOffset > maxes[msg.DestChainID] {
			maxes[msg.DestChainID] = msg.StreamOffset
		}
	}

	offsets := make([]*types.MsgOffset, 0, len(maxes))
	for destChainIO, offset := range maxes {
		offsets = append(offsets, &types.MsgOffset{
			DestChainId:  destChainIO,
			StreamOffset: offset,
		})
	}

	return offsets
}

func receiptOffsets(block xchain.Block) []*types.ReceiptOffset {
	maxes := make(map[uint64]uint64)
	for _, receipt := range block.Receipts {
		if receipt.StreamOffset > maxes[receipt.SourceChainID] {
			maxes[receipt.SourceChainID] = receipt.StreamOffset
		}
	}

	offsets := make([]*types.ReceiptOffset, 0, len(maxes))
	for srcChainID, offset := range maxes {
		offsets = append(offsets, &types.ReceiptOffset{
			SourceChainId: srcChainID,
			StreamOffset:  offset,
		})
	}

	return offsets
}
