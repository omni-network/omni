package merkle

import "github.com/omni-network/omni/lib/errors"

// These functions are also ported from OpenZeppelin's library, but they
// are not used by omni's production code, so they are part of the
// tests to decrease prod code surface.

// GetProof returns a merkle proof for the given leaf index.
func GetProof(tree [][32]byte, index int) ([][32]byte, error) {
	if err := checkLeafNode(tree, index); err != nil {
		return nil, err
	}

	var proof [][32]byte
	for index > 0 {
		proof = append(proof, tree[siblingIndex(index)])
		index = parentIndex(index)
	}

	return proof, nil
}

// ProcessProof returns the root hash of the merkle tree given the leaf hash and the proof.
func ProcessProof(leaf [32]byte, proof [][32]byte) [32]byte {
	node := leaf
	for _, p := range proof {
		node = hashPair(node, p)
	}

	return node
}

// ProcessMultiProof returns the root hash of the tree given a multi proof.
func ProcessMultiProof(multi MultiProof) ([32]byte, error) {
	if err := verifyMultiProof(multi); err != nil {
		return [32]byte{}, err
	}

	// Copy leaves and proof.
	stack := make([][32]byte, len(multi.Leaves))
	copy(stack, multi.Leaves)
	proof := make([][32]byte, len(multi.Proof))
	copy(proof, multi.Proof)

	for _, flag := range multi.ProofFlags {
		// Pop from the beginning of the stack.
		a := stack[0]
		stack = stack[1:]

		// Either pop from the stack or the proof, depending on the flag.
		var b [32]byte
		if flag {
			b = stack[0]
			stack = stack[1:]
		} else {
			b = proof[0]
			proof = proof[1:]
		}

		stack = append(stack, hashPair(a, b)) //nolint:makezero // Appending to non-zero initialized slice is ok
	}

	// Either the stack or the proof should have one element left.
	if len(stack)+len(proof) != 1 {
		return [32]byte{}, errors.New("broken invariant")
	}

	if len(stack) > 0 {
		return stack[0], nil
	}

	return proof[0], nil
}

// LeafToTreeIndex returns the index of the leaf in the tree given the original index in the leaves slice.
func LeafToTreeIndex(tree [][32]byte, leafIndex int) int {
	return len(tree) - 1 - leafIndex
}

// verifyMultiProof returns an error if the given multi proof is invalid.
func verifyMultiProof(multi MultiProof) error {
	var falseFlags int
	for _, flag := range multi.ProofFlags {
		if !flag {
			falseFlags++
		}
	}
	if len(multi.Proof) != falseFlags {
		return errors.New("false proof flags don't match proof")
	}

	if len(multi.Leaves)+len(multi.Proof) != len(multi.ProofFlags)+1 {
		return errors.New("proof flags don't match leaves and proof")
	}

	if len(multi.Leaves) == 0 {
		return errors.New("no leaves provided")
	}

	return nil
}
