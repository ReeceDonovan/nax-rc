package search

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/binaryTree"
	"github.com/ReeceDonovan/nax-rc/internal/types"
	"github.com/stretchr/testify/assert"
)

func createTreeWithRevisions(numRevisions, dataSize int) *binaryTree.TreeNode {
	rootRevision := types.NewRandomRevision(0, dataSize)
	root := &binaryTree.TreeNode{Revision: rootRevision}

	for i := 1; i < numRevisions; i++ {
		revision := types.NewRandomRevision(i, dataSize)
		root.Insert(revision)
	}
	return root
}

func TestBinarySearch(t *testing.T) {
	tree := createTreeWithRevisions(100, 10)

	foundRevision, err := BinarySearch(tree, 42)
	assert.NoError(t, err)
	assert.NotNil(t, foundRevision)
	assert.Equal(t, 42, foundRevision.ID())

	notFoundRevision, err := BinarySearch(tree, 999)
	assert.Error(t, err)
	assert.Equal(t, errors.New("revision not found"), err)
	assert.Nil(t, notFoundRevision)
}

func BenchmarkBinarySearch_RevisionSearch(b *testing.B) {
	for _, numRevisions := range []int{smallNumRevisions, mediumNumRevisions, largeNumRevisions} {
		for _, dataSize := range []int{smallDataSize, mediumDataSize, largeDataSize} {
			b.Run(fmt.Sprintf("numRevisions=%d,dataSize=%d", numRevisions, dataSize), func(b *testing.B) {
				tree := createTreeWithRevisions(numRevisions, dataSize)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					revisionID := rand.Intn(numRevisions)
					_, _ = BinarySearch(tree, revisionID)
				}
			})
		}
	}
}
