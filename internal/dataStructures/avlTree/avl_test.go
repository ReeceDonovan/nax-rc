package avlTree

import (
	"fmt"
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/types"
)

const (
	smallNumRevisions  = 20
	mediumNumRevisions = 200
	largeNumRevisions  = 2000
	largeDataSize      = 1000000 // 1 Mb
	mediumDataSize     = 1000    // 1 Kb
	smallDataSize      = 100     // 100 bytes
)

// Create benchmarks for inserting revisions into a AVL with different sizes
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
