package search

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/graph"
	"github.com/ReeceDonovan/nax-rc/internal/types"
)

// var DATA_SIZE = []struct {
// 	dataSize int
// }{
// 	{100},
// 	{1000},
// 	{10000},
// 	{100000},
// 	{1000000},
// }

var DATA_SIZE = []struct {
	dataSize int
}{
	{100},
	{1000},
	{10000},
	{100000},
	{1000000},
}

func createBigGraph(dataSize int) (graph.Graph, map[string]types.Revision) {
	var g graph.Graph

	// Create a map of letters A-Z to use as vertex labels, each letter will correspond to a random revision
	var revisions = make(map[string]types.Revision)
	for i := 65; i < 91; i++ {
		revision := types.NewRandomRevision(i, dataSize)
		revisions[fmt.Sprintf("%c", i)] = revision
		g.AddVertex(revision)
	}

	// 	graph2.AddEdge(StandardEdge("A", "B"))
	// 	graph2.AddEdge(StandardEdge("A", "C"))
	// 	graph2.AddEdge(StandardEdge("A", "D"))
	// 	graph2.AddEdge(StandardEdge("A", "E"))
	// 	graph2.AddEdge(StandardEdge("A", "F"))

	// 	graph2.AddEdge(StandardEdge("B", "G"))
	// 	graph2.AddEdge(StandardEdge("B", "H"))
	// 	graph2.AddEdge(StandardEdge("B", "I"))

	// 	graph2.AddEdge(StandardEdge("C", "J"))

	// 	graph2.AddEdge(StandardEdge("D", "K"))
	// 	graph2.AddEdge(StandardEdge("D", "L"))

	// 	graph2.AddEdge(StandardEdge("F", "M"))
	// 	graph2.AddEdge(StandardEdge("F", "N"))

	// 	graph2.AddEdge(StandardEdge("H", "O"))
	// 	graph2.AddEdge(StandardEdge("H", "P"))
	// 	graph2.AddEdge(StandardEdge("H", "Q"))
	// 	graph2.AddEdge(StandardEdge("H", "R"))

	// 	graph2.AddEdge(StandardEdge("P", "T"))
	// 	graph2.AddEdge(StandardEdge("P", "U"))

	// 	graph2.AddEdge(StandardEdge("R", "V"))

	// 	graph2.AddEdge(StandardEdge("V", "W"))
	// 	graph2.AddEdge(StandardEdge("V", "X"))
	// 	graph2.AddEdge(StandardEdge("V", "Y"))

	// 	graph2.AddEdge(StandardEdge("X", "Z"))

	g.AddEdge(graph.StandardEdge(revisions["A"], revisions["B"]))
	g.AddEdge(graph.StandardEdge(revisions["A"], revisions["C"]))
	g.AddEdge(graph.StandardEdge(revisions["A"], revisions["D"]))
	g.AddEdge(graph.StandardEdge(revisions["A"], revisions["E"]))
	g.AddEdge(graph.StandardEdge(revisions["A"], revisions["F"]))
	g.AddEdge(graph.StandardEdge(revisions["B"], revisions["G"]))
	g.AddEdge(graph.StandardEdge(revisions["B"], revisions["H"]))
	g.AddEdge(graph.StandardEdge(revisions["B"], revisions["I"]))
	g.AddEdge(graph.StandardEdge(revisions["C"], revisions["J"]))
	g.AddEdge(graph.StandardEdge(revisions["D"], revisions["K"]))
	g.AddEdge(graph.StandardEdge(revisions["D"], revisions["L"]))
	g.AddEdge(graph.StandardEdge(revisions["F"], revisions["M"]))
	g.AddEdge(graph.StandardEdge(revisions["F"], revisions["N"]))
	g.AddEdge(graph.StandardEdge(revisions["H"], revisions["O"]))
	g.AddEdge(graph.StandardEdge(revisions["H"], revisions["P"]))
	g.AddEdge(graph.StandardEdge(revisions["H"], revisions["Q"]))
	g.AddEdge(graph.StandardEdge(revisions["H"], revisions["R"]))
	g.AddEdge(graph.StandardEdge(revisions["K"], revisions["S"]))
	g.AddEdge(graph.StandardEdge(revisions["P"], revisions["T"]))
	g.AddEdge(graph.StandardEdge(revisions["P"], revisions["U"]))
	g.AddEdge(graph.StandardEdge(revisions["R"], revisions["V"]))
	g.AddEdge(graph.StandardEdge(revisions["V"], revisions["W"]))
	g.AddEdge(graph.StandardEdge(revisions["V"], revisions["X"]))
	g.AddEdge(graph.StandardEdge(revisions["V"], revisions["Y"]))
	g.AddEdge(graph.StandardEdge(revisions["X"], revisions["Z"]))
	return g, revisions
}

func BenchmarkDepthFirstSearch(b *testing.B) {
	for _, dataSize := range DATA_SIZE {
		b.Run(fmt.Sprintf("DepthFirstSearch_%d", dataSize.dataSize), func(b *testing.B) {
			dataSize := dataSize.dataSize
			// g, revisions := createBigGraph(dataSize)
			for i := 0; i < b.N; i++ {
				g, revisions := createBigGraph(dataSize)
				// select a random target revision
				targetRevisionID := rand.Intn(26) + 65
				var targetRevision = revisions[fmt.Sprintf("%c", targetRevisionID)]
				if targetRevision == nil {
					b.Errorf("expected target revision to not be nil, %v", targetRevisionID)
				}

				result := DepthFirstSearch(g, revisions["A"], targetRevision.ID())

				if result == nil {
					b.Errorf("DepthFirstSearch(%v)", targetRevision.ID())
				}

				// b.Logf("Target ID: %v, Result ID: %v", targetRevision.ID(), result.ID())
			}
		})
	}
}
