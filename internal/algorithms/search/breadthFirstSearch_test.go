package search

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/graph"
	"github.com/ReeceDonovan/nax-rc/internal/types"
)

func TestBreadthFirstSearch(t *testing.T) {
	dag := graph.NewDAG()

	nodes := make([]*graph.DAGNode, 10)
	for i := 0; i < 10; i++ {
		nodes[i], _ = dag.AddNode(types.NewBlankRevision(i))
	}

	for i := 0; i < 9; i++ {
		_ = dag.AddEdge(i, i+1)
	}

	visitedNodes := make([]int, 0)
	visitor := func(node *graph.DAGNode) {
		visitedNodes = append(visitedNodes, node.Revision.ID())
	}

	BreadthFirstSearch(dag, 0, visitor)

	expectedOrder := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, nodeID := range visitedNodes {
		if nodeID != expectedOrder[i] {
			t.Errorf("Unexpected order of visited nodes, expected %v, got %v", expectedOrder, visitedNodes)
			break
		}
	}
}

func BenchmarkBreadthFirstSearch_RevisionSearch(b *testing.B) {
	for _, numRevisions := range []int{smallNumRevisions, mediumNumRevisions, largeNumRevisions} {
		for _, dataSize := range []int{smallDataSize, mediumDataSize, largeDataSize} {
			b.Run(fmt.Sprintf("numRevisions=%d,dataSize=%d", numRevisions, dataSize), func(b *testing.B) {
				dag := graph.NewDAG()

				nodes := make([]*graph.DAGNode, numRevisions)
				for i := 0; i < numRevisions; i++ {
					nodes[i], _ = dag.AddNode(types.NewRandomRevision(i, dataSize))
				}

				for i := 0; i < numRevisions-1; i++ {
					_ = dag.AddEdge(i, i+1)
				}

				visitor := func(node *graph.DAGNode) {}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					BreadthFirstSearch(dag, rand.Intn(numRevisions), visitor)
				}
			})
		}
	}
}
