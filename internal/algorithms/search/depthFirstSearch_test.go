package search

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/graph"
	"github.com/ReeceDonovan/nax-rc/internal/types"
)

func createTestDAG() *graph.DAG {
	dag := graph.NewDAG()

	// Add nodes
	for i := 1; i <= 5; i++ {
		dag.AddNode(types.NewBlankRevision(i))
	}

	// Add edges
	dag.AddEdge(1, 2)
	dag.AddEdge(1, 3)
	dag.AddEdge(2, 4)
	dag.AddEdge(3, 4)
	dag.AddEdge(4, 5)

	return dag
}

func TestDepthFirstSearch(t *testing.T) {
	dag := createTestDAG()

	visited := make([]int, 0)

	DepthFirstSearch(dag, 1, func(node *graph.DAGNode) {
		visited = append(visited, node.Revision.ID())
	})

	expected := []int{1, 2, 4, 5, 3}

	for i, id := range visited {
		if id != expected[i] {
			t.Errorf("Unexpected order in DepthFirstSearch: got %v, want %v", visited, expected)
			break
		}
	}
}

func BenchmarkDepthFirstSearch_RevisionSearch(b *testing.B) {
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
					DepthFirstSearch(dag, rand.Intn(numRevisions), visitor)
				}
			})
		}
	}
}
