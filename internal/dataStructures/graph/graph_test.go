package graph

import (
	"fmt"
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/types"
)

// Benchmark adding a very large number of vertices and edges to a graph.
func BenchmarkGraphAdd(b *testing.B) {
	var numRevisions = 10
	for i := 0; i < b.N; i++ {
		var graph Graph
		for i := 0; i < numRevisions; i++ {
			graph.AddVertex(i)
		}
		for i := 0; i < numRevisions; i++ {
			graph.AddEdge(StandardEdge(i, i+1))
		}
		b.Logf("Vertices: %d", graph.vertices.Length())
		b.Logf("Edges: %d", graph.edges.Length())
		b.Logf("Graph: %s", graph.String())
	}
}

// Benchmark removing a very large number of vertices and edges from a graph.
func BenchmarkGraphRemove(b *testing.B) {
	var numRevisions = 1000
	for i := 0; i < b.N; i++ {
		var graph Graph
		for i := 0; i < numRevisions; i++ {
			graph.AddVertex(i)
		}
		for i := 0; i < numRevisions; i++ {
			graph.AddEdge(StandardEdge(i, i+1))
		}

		b.Logf("Vertices: %d", graph.vertices.Length())

		for i := 0; i < numRevisions; i++ {
			graph.RemoveVertex(i)
		}

		b.Logf("Remaining vertices: %d", graph.vertices.Length())
	}
}

func createGenericRevFromID(id int) types.Revision {
	return types.NewBlankRevision(id)
}

func BenchmarkGraphRev(b *testing.B) {
	var numRevisions = 1000
	for i := 0; i < b.N; i++ {
		var graph Graph
		for i := 0; i < numRevisions; i++ {
			graph.AddVertex(createGenericRevFromID(i))
		}
		for i := 0; i < numRevisions; i++ {
			graph.AddEdge(StandardEdge(createGenericRevFromID(i), createGenericRevFromID(i+1)))
		}
		b.Logf("Vertices: %d", graph.vertices.Length())
		b.Logf("Edges: %d", graph.edges.Length())
		b.Logf("Graph:\n%s", graph.String())
	}
}

const (
	smallNumRevisions  = 20
	mediumNumRevisions = 200
	largeNumRevisions  = 2000
	largeDataSize      = 1000000 // 1 Mb
	mediumDataSize     = 1000    // 1 Kb
	smallDataSize      = 100     // 100 bytes
)

// Create a benchmarks for constructing a graph of revisions with different sizes
func BenchmarkGraph_NewGraph(b *testing.B) {
	for _, numRevisions := range []int{smallNumRevisions, mediumNumRevisions, largeNumRevisions} {
		for _, dataSize := range []int{smallDataSize, mediumDataSize, largeDataSize} {
			b.Run(fmt.Sprintf("numRevisions=%d,dataSize=%d", numRevisions, dataSize), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					var graph Graph
					for j := 0; j < numRevisions; j++ {
						revision := types.NewRandomRevision(j, dataSize)
						graph.AddVertex(revision)
					}
				}
			})
		}
	}
}
