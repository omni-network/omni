package xchain

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/merkle"

	"github.com/ethereum/go-ethereum/common"
)

const (
	DSTUnknown merkle.DomainSeparationTag = 0
	DSTHeader  merkle.DomainSeparationTag = 1
	DSTMessage merkle.DomainSeparationTag = 2
)

// MsgTree is a merkle tree of all the messages in a cross-chain block.
// It is used as a leaf when calculating the AttestationRoot.
// Its proofs are used to submit messages to destination chains.
type MsgTree struct {
	tree    [][32]byte       // Merkle tree with all nodes, incl leaves
	indices map[[32]byte]int // Reverse lookup table for node indices
}

func (t MsgTree) MsgRoot() [32]byte {
	return t.tree[0]
}

// Proof returns the merkle multi proof for the provided submissionHeader and messages.
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

	return merkle.GetMultiProof(t.tree, indices...)
}

func (t MsgTree) leafIndex(leaf [32]byte) (int, error) {
	index, ok := t.indices[leaf]
	if !ok {
		return 0, errors.New("leaf not in tree")
	}

	return index, nil
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

	tree, err := merkle.MakeTree(leafs)
	if err != nil {
		return MsgTree{}, err
	}

	// Create a lookup table for node indices
	indices := make(map[[32]byte]int)
	for i, node := range tree {
		indices[node] = i
	}

	return MsgTree{
		tree:    tree,
		indices: indices,
	}, nil
}

func msgLeaf(msg Msg) ([32]byte, error) {
	bz, err := encodeMsg(msg)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "encode message")
	}

	return merkle.StdLeafHash(DSTMessage, bz), nil
}

func submissionHeaderLeaf(attHeader AttestHeader, blockHeader BlockHeader) ([32]byte, error) {
	bz, err := encodeSubmissionHeader(attHeader, blockHeader)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "encode block submissionHeader")
	}

	return merkle.StdLeafHash(DSTHeader, bz), nil
}

// AttestationRoot returns the attestation root of the provided block submissionHeader and message root.
func AttestationRoot(attHeader AttestHeader, blockHeader BlockHeader, msgRoot common.Hash) (common.Hash, error) {
	headerLeaf, err := submissionHeaderLeaf(attHeader, blockHeader)
	if err != nil {
		return [32]byte{}, err
	}

	tree, err := merkle.MakeTree([][32]byte{msgRoot, headerLeaf})
	if err != nil {
		return [32]byte{}, err
	}

	return tree[0], nil
}
