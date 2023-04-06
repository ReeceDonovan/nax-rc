package graph

import (
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/revlog"
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

func createGenericRevFromID(id int) *revlog.Revision {
	return &revlog.Revision{ID: id, Data: nil}
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
