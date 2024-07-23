// Package merkle provides an API to generate merkle trees and proofs from 32 byte leaves.
// It is a port of the OpenZeppelin JS merkle-tree library.
// See https://github.com/OpenZeppelin/merkle-tree/tree/master.
package merkle

import (
	"bytes"
	"sort"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/crypto"
)

// Specifies the domain of a single leaf in the tree.
type DomainSeparationTag byte

// StdLeafHash returns the standard leaf hash of the given data.
// The data is hashed twice with keccak256 to prevent pre-image attacks.
func StdLeafHash(dst DomainSeparationTag, data []byte) [32]byte {
	h1 := hash(append([]byte{byte(dst)}, data...))
	h2 := hash(h1[:])

	return h2
}

// MakeTree returns a merkle tree given the leaves.
func MakeTree(leaves [][32]byte) ([][32]byte, error) {
	if len(leaves) == 0 {
		return nil, errors.New("no leaves provided")
	}

	treeLen := 2*len(leaves) - 1
	tree := make([][32]byte, treeLen)

	// Fill in leaves in reverse order.
	for i, leaf := range leaves {
		tree[treeLen-1-i] = leaf
	}

	// Fill in the intermediate nodes up to the root.
	for i := treeLen - 1 - len(leaves); i >= 0; i-- {
		tree[i] = hashPair(tree[leftChildIndex(i)], tree[rightChildIndex(i)])
	}

	// Ensure the tree always has an odd number of nodes.
	if treeLen%2 == 0 {
		return nil, errors.New("invalid even tree [BUG]")
	}

	return tree, nil
}

// MultiProof is a merkle-multi-proof for multiple leaves.
type MultiProof struct {
	Leaves     [][32]byte
	Proof      [][32]byte
	ProofFlags []bool
}

// GetMultiProof returns a merkle-multi-proof for the given leaf indices.
func GetMultiProof(tree [][32]byte, indices ...int) (MultiProof, error) {
	if len(indices) == 0 {
		return MultiProof{}, errors.New("no indices provided")
	} else if len(tree)%2 == 0 {
		return MultiProof{}, errors.New("invalid even tree [BUG]")
	}

	for _, i := range indices {
		if err := checkLeafNode(tree, i); err != nil {
			return MultiProof{}, err
		}
	}

	// Sort indices in reverse order.
	sort.Slice(indices, func(i, j int) bool {
		return indices[i] > indices[j]
	})

	// Check for duplicates.
	for i, j := range indices[1:] {
		if j == indices[i] {
			return MultiProof{}, errors.New("cannot prove duplicated index")
		}
	}

	stack := make([]int, len(indices))
	copy(stack, indices)
	var proof [][32]byte
	var proofFlags []bool

	for len(stack) > 0 && stack[0] > 0 {
		// Pop from the beginning.
		j := stack[0]
		stack = stack[1:]

		s := siblingIndex(j)
		p := parentIndex(j)

		if s >= len(tree) {
			return MultiProof{}, errors.New("invalid sibling index, invalid tree [BUG]")
		}

		if len(stack) > 0 && s == stack[0] {
			proofFlags = append(proofFlags, true)
			stack = stack[1:]
		} else {
			proofFlags = append(proofFlags, false)
			proof = append(proof, tree[s])
		}
		stack = append(stack, p) //nolint:makezero // Appending to non-zero initialized slice is ok
	}

	leaves := make([][32]byte, 0, len(indices))
	for _, i := range indices {
		leaves = append(leaves, tree[i])
	}

	return MultiProof{
		Leaves:     leaves,
		Proof:      proof,
		ProofFlags: proofFlags,
	}, nil
}

// isTreeNode returns true if the given index is a node in the tree.
func isTreeNode(tree [][32]byte, i int) bool {
	return i >= 0 && i < len(tree)
}

// isInternalNode returns true if the given index is an internal node (not a leaf) in the tree.
func isInternalNode(tree [][32]byte, i int) bool {
	return isTreeNode(tree, leftChildIndex(i))
}

// isLeafNode returns true if the given index is a leaf node in the tree.
func isLeafNode(tree [][32]byte, i int) bool {
	return isTreeNode(tree, i) && !isInternalNode(tree, i)
}

// checkLeafNode returns an error if the given index is not a leaf node.
func checkLeafNode(tree [][32]byte, i int) error {
	if !isLeafNode(tree, i) {
		return errors.New("index is not a leaf")
	}

	return nil
}

// leftChildIndex returns the index of the left child of the node at the given index.
func leftChildIndex(i int) int {
	return 2*i + 1
}

// rightChildIndex returns the index of the right child of the node at the given index.
func rightChildIndex(i int) int {
	return 2*i + 2
}

// parentIndex returns the index of the parent of the node at the given index.
// It panics if the given index is 0.
func parentIndex(i int) int {
	if i == 0 {
		panic("root has no parent")
	}

	return (i - 1) / 2
}

// siblingIndex returns the index of the sibling of the node at the given index.
func siblingIndex(i int) int {
	if i == 0 {
		panic("root has no sibling")
	}

	if i%2 == 0 {
		return i - 1
	}

	return i + 1
}

// sortBytes returns the given byte slices sorted in ascending order.
func sortBytes(buf ...[]byte) [][]byte {
	sort.Slice(buf, func(i, j int) bool {
		return bytes.Compare(buf[i], buf[j]) < 0
	})

	return buf
}

// concatBytes returns the concatenation of the given byte slices.
func concatBytes(buf ...[]byte) []byte {
	var resp []byte
	for _, b := range buf {
		resp = append(resp, b...)
	}

	return resp
}

// hashPair returns the 32 byte keccak256 hash of the sorted concatenation of the given byte arrays.
func hashPair(a [32]byte, b [32]byte) [32]byte {
	return hash(concatBytes(sortBytes(a[:], b[:])...))
}

// hash returns the 32 byte keccak256 hash of the given byte slice.
func hash(buf []byte) [32]byte {
	resp := crypto.Keccak256(buf)

	return [32]byte(resp)
}
