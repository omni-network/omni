package merkle_test

import (
	"math/rand"
	"testing"

	"github.com/omni-network/omni/lib/merkle"

	"github.com/ethereum/go-ethereum/crypto"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

// TestLeaveProvable tests that a leaf can be proven.
func TestLeaveProvable(t *testing.T) {
	t.Parallel()

	// Create random leaves
	var leaves [][32]byte
	fuzz.New().NilChance(0).NumElements(1, 256).Fuzz(&leaves)

	// Make tree
	tree, err := merkle.MakeTree(leaves)
	require.NoError(t, err)

	// Pick random leaf
	leafIndex := rand.Intn(len(leaves))
	treeIndex := merkle.LeafToTreeIndex(tree, leafIndex)

	// Get the proof
	proof, err := merkle.GetProof(tree, treeIndex)
	require.NoError(t, err)

	leaf := leaves[leafIndex]
	root := merkle.ProcessProof(leaf, proof)
	require.Equal(t, tree[0], root)
}

// TestLeavesProvable tests that multiple leaves can be proven.
func TestLeavesProvable(t *testing.T) {
	t.Parallel()

	// Create random leaves
	var leaves [][32]byte
	fuzz.New().NilChance(0).NumElements(1, 256).Fuzz(&leaves)

	// Make tree
	tree, err := merkle.MakeTree(leaves)
	require.NoError(t, err)

	// Pick random leaves
	leafIndices := randomIndicesRange(leaves)
	treeIndices := make([]int, len(leafIndices))
	for i, leafIndex := range leafIndices {
		treeIndices[i] = merkle.LeafToTreeIndex(tree, leafIndex)
	}

	// Get the multi proof
	multi, err := merkle.GetMultiProof(tree, treeIndices...)
	require.NoError(t, err)

	// Check that the proof contains the leaves
	require.Equal(t, len(leafIndices), len(multi.Leaves))
	for _, i := range leafIndices {
		require.Contains(t, multi.Leaves, leaves[i])
	}

	// Check that the proof is valid
	root, err := merkle.ProcessMultiProof(multi)
	require.NoError(t, err)
	require.Equal(t, tree[0], root)
}

func TestEvenTree(t *testing.T) {
	t.Parallel()

	// Invalid tree with an even number of nodes.
	tree := [][32]byte{
		crypto.Keccak256Hash([]byte("node1")),
		crypto.Keccak256Hash([]byte("node2")),
	}
	treeIndices := []int{1}

	_, err := merkle.GetMultiProof(tree, treeIndices...)
	require.ErrorContains(t, err, "invalid even tree")
}

// randomIndicesRange returns a random range of indices of the provided slice.
func randomIndicesRange(slice [][32]byte) []int {
	start := rand.Intn(len(slice))
	count := rand.Intn(len(slice) - start)
	if count == 0 {
		count = 1
	}

	indices := make([]int, count)
	for i := range indices {
		indices[i] = start + i
	}

	return indices
}
