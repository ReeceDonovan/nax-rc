package dataStructures

import (
	"fmt"
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/binaryTree"
	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/graph"
	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/linkedList"
	"github.com/ReeceDonovan/nax-rc/internal/types"
)

const (
	smallNumRevisions  = 200
	mediumNumRevisions = 2000
	largeNumRevisions  = 20000
	largeDataSize      = 10000 // 10 Kb
	mediumDataSize     = 1000  // 1 Kb
	smallDataSize      = 100   // 100 bytes
)

// Create a benchmarks for constructing a list of revisions with different sizes
func BenchmarkDoublyLinkedList_RevisionConstruction(b *testing.B) {
	for _, numRevisions := range []int{smallNumRevisions, mediumNumRevisions, largeNumRevisions} {
		for _, dataSize := range []int{smallDataSize, mediumDataSize, largeDataSize} {
			b.Run(fmt.Sprintf("numRevisions=%d,dataSize=%d", numRevisions, dataSize), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					list := linkedList.NewDoublyLinkedList()
					for j := 0; j < numRevisions; j++ {
						revision := types.NewRandomRevision(j, dataSize)
						list.InsertAtPosition(j+1, &linkedList.DLLNode{Revision: revision})
					}
				}
			})
		}
	}
}

// Create benchmarks for constructing a BST of revisions with different sizes
func BenchmarkTreeNode_RevisionConstruction(b *testing.B) {
	for _, numRevisions := range []int{smallNumRevisions, mediumNumRevisions, largeNumRevisions} {
		for _, dataSize := range []int{smallDataSize, mediumDataSize, largeDataSize} {
			b.Run(fmt.Sprintf("numRevisions=%d,dataSize=%d", numRevisions, dataSize), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					// Create a root node
					rootRevision := types.NewRandomRevision(0, dataSize)
					root := &binaryTree.TreeNode{Revision: rootRevision}

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

// Create benchmarks for constructing a DAG of revisions with different sizes
func BenchmarkDAG_RevisionConstruction(b *testing.B) {
	for _, numNodes := range []int{smallNumRevisions, mediumNumRevisions, largeNumRevisions} {
		for _, dataSize := range []int{smallDataSize, mediumDataSize, largeDataSize} {
			b.Run(fmt.Sprintf("numRevisions=%d,dataSize=%d", numNodes, dataSize), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					// Create a DAG
					dag := graph.NewDAG()
					// Add nodes with revisions
					for j := 0; j < numNodes; j++ {
						revision := types.NewRandomRevision(j, dataSize)
						dag.AddNode(revision)
					}

					// Add edges
					for j := 0; j < numNodes-1; j++ {
						dag.AddEdge(j, j+1)
					}
				}
			})
		}
	}
}

// ------------------------------------------

func BenchmarkDataStructures(b *testing.B) {
	b.Run("LinkedList", func(b *testing.B) {
		BenchmarkDoublyLinkedList_RevisionConstruction(b)
	})

	b.Run("Graph", func(b *testing.B) {
		BenchmarkDAG_RevisionConstruction(b)
	})

	b.Run("BinaryTree", func(b *testing.B) {
		BenchmarkTreeNode_RevisionConstruction(b)
	})
}

// ------------------------------------------
