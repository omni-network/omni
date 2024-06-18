package xchain

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/merkle"

	"github.com/ethereum/go-ethereum/common"
)

// MsgTree is a merkle tree of all the messages in a cross chain block.
// It is used a leaf when calculating the AttestationRoot.
// It's proofs are used to submit messages to destination chains.
type MsgTree [][32]byte

func (t MsgTree) MsgRoot() [32]byte {
	return t[0]
}

// Proof returns the merkle multi proof for the provided header and messages.
func (t MsgTree) Proof(msgs []Msg) (merkle.MultiProof, error) {
	// Get the indices to prove
	indices := make([]int, 0, len(msgs))
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

func (t MsgTree) leafIndex(leaf [32]byte) (int, error) {
	// Linear search for the leaf (probably ok since trees are small; < 1000)
	for i, l := range t {
		if l == leaf {
			return i, nil
		}
	}

	return 0, errors.New("leaf not in tree")
}

// NewMsgTree returns the merkle root of the provided messages
// to be submitted.
func NewMsgTree(msgs []Msg) (MsgTree, error) {
	leafs := make([][32]byte, 0, len(msgs))
	for _, msg := range msgs {
		msgLeaf, err := msgLeaf(msg)
		if err != nil {
			return MsgTree{}, err
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

func BlockHeaderLeaf(header BlockHeader) ([32]byte, error) {
	bz, err := encodeHeader(header)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "encode block header")
	}

	return merkle.StdLeafHash(bz), nil
}

// AttestationRoot returns the attestation root of the provided block header and message root.
func AttestationRoot(header BlockHeader, msgRoot common.Hash) (common.Hash, error) {
	headerLeaf, err := BlockHeaderLeaf(header)
	if err != nil {
		return [32]byte{}, err
	}

	tree, err := merkle.MakeTree([][32]byte{msgRoot, headerLeaf})
	if err != nil {
		return [32]byte{}, err
	}

	return tree[0], nil
}
