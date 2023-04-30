package binaryTree

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestTreeNode_Insert(t *testing.T) {
	root := &TreeNode{Revision: types.NewBlankRevision(10)}
	root.Insert(types.NewBlankRevision(5))
	root.Insert(types.NewBlankRevision(15))

	assert.Equal(t, 5, root.Left.Revision.ID())
	assert.Equal(t, 15, root.Right.Revision.ID())
}

func TestTreeNode_Contains(t *testing.T) {
	root := &TreeNode{Revision: types.NewBlankRevision(10)}
	root.Insert(types.NewBlankRevision(5))
	root.Insert(types.NewBlankRevision(15))

	assert.True(t, root.Contains(5))
	assert.True(t, root.Contains(10))
	assert.True(t, root.Contains(15))
	assert.False(t, root.Contains(20))
}

func TestTreeNode_Remove(t *testing.T) {
	root := &TreeNode{Revision: types.NewBlankRevision(10)}
	root.Insert(types.NewBlankRevision(5))
	root.Insert(types.NewBlankRevision(15))

	root, _ = root.Remove(15)

	assert.False(t, root.Contains(15))
}

func generateRandomRevision() types.Revision {
	revisionID := rand.Intn(1000)
	dataSize := rand.Intn(100)
	return types.NewRandomRevision(revisionID, dataSize)
}

func BenchmarkTreeNode_Insert(b *testing.B) {
	root := &TreeNode{Revision: generateRandomRevision()}

	for i := 0; i < b.N; i++ {
		revision := generateRandomRevision()
		root.Insert(revision)
	}
}

func BenchmarkTreeNode_Contains(b *testing.B) {
	root := &TreeNode{Revision: generateRandomRevision()}

	for i := 0; i < b.N; i++ {
		revision := generateRandomRevision()
		root.Insert(revision)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		revisionID := rand.Intn(b.N)
		root.Contains(revisionID)
	}
}

func BenchmarkTreeNode_Remove(b *testing.B) {
	root := &TreeNode{Revision: generateRandomRevision()}

	for i := 0; i < b.N; i++ {
		revision := generateRandomRevision()
		root.Insert(revision)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		revisionID := rand.Intn(b.N)
		root, _ = root.Remove(revisionID)
	}
}

const (
	smallNumRevisions  = 200
	mediumNumRevisions = 2000
	largeNumRevisions  = 20000
	largeDataSize      = 10000 // 10 Kb
	mediumDataSize     = 1000  // 1 Kb
	smallDataSize      = 100   // 100 bytes
)

// Create benchmarks for constructing a BST of revisions with different sizes
func BenchmarkTreeNode_RevisionConstruction(b *testing.B) {
	for _, numRevisions := range []int{smallNumRevisions, mediumNumRevisions, largeNumRevisions} {
		for _, dataSize := range []int{smallDataSize, mediumDataSize, largeDataSize} {
			b.Run(fmt.Sprintf("numRevisions=%d,dataSize=%d", numRevisions, dataSize), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					// Create a root node
					rootRevision := types.NewRandomRevision(0, dataSize)
					root := &TreeNode{Revision: rootRevision}

					// Insert revisions
					for j := 1; j < numRevisions; j++ {
						revision := types.NewRandomRevision(j, dataSize)
						root.Insert(revision)
					}
				}
			})
		}
	}
}
