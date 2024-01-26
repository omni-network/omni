package xchain

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/merkle"
)

// BlockTree is a merkle tree of a cross chain block.
// It is attested to by the consensus chain validators.
// It's proofs are used to submit messages to destination chains.
type BlockTree [][32]byte

func (t BlockTree) Root() [32]byte {
	return t[0]
}

// Proof returns the merkle multi proof for the provided header and messages.
func (t BlockTree) Proof(header BlockHeader, msgs []Msg) (merkle.MultiProof, error) {
	// Get the indices to prove
	indices := make([]int, 0, len(msgs)+1)

	headerLeaf, err := blockHeaderLeaf(header)
	if err != nil {
		return merkle.MultiProof{}, err
	}
	headerIndex, err := t.leafIndex(headerLeaf)
	if err != nil {
		return merkle.MultiProof{}, errors.Wrap(err, "header index")
	}
	indices = append(indices, headerIndex)

	for _, msg := range msgs {
		msgLeaf, err := msgLeaf(msg)
		if err != nil {
			return merkle.MultiProof{}, err
		}
		msgIndex, err := t.leafIndex(msgLeaf)
		if err != nil {
			return merkle.MultiProof{}, errors.Wrap(err, "msg index")
		}
		indices = append(indices, msgIndex)
	}

	return merkle.GetMultiProof(t, indices...)
}

func (t BlockTree) leafIndex(leaf [32]byte) (int, error) {
	// Linear search for the leaf (probably ok since trees are small; < 1000)
	for i, l := range t {
		if l == leaf {
			return i, nil
		}
	}

	return 0, errors.New("leaf not in tree")
}

// NewBlockTree returns the merkle root of the provided block
// to be attested to.
func NewBlockTree(block Block) (BlockTree, error) {
	leafs := make([][32]byte, 0, 1+len(block.Msgs))

	// First leaf is the block header
	headerLeaf, err := blockHeaderLeaf(block.BlockHeader)
	if err != nil {
		return BlockTree{}, err
	}
	leafs = append(leafs, headerLeaf)

	// Next leafs are the messages
	for _, msg := range block.Msgs {
		msgLeaf, err := msgLeaf(msg)
		if err != nil {
			return BlockTree{}, err
		}
		leafs = append(leafs, msgLeaf)
	}

	return merkle.MakeTree(leafs)
}

func msgLeaf(msg Msg) ([32]byte, error) {
	bz, err := encodeMsg(msg)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "encode message")
	}

	return merkle.StdLeafHash(bz), nil
}

func blockHeaderLeaf(header BlockHeader) ([32]byte, error) {
	bz, err := encodeHeader(header)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "encode block header")
	}

	return merkle.StdLeafHash(bz), nil
}
