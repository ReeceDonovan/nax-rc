package search

import (
	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/graph"
	"github.com/ReeceDonovan/nax-rc/internal/types"
)

func DepthFirstSearch(g graph.Graph, startVertex graph.Vertex, id int) types.Revision {
	visited := make(map[graph.Vertex]bool)
	stack := []graph.Vertex{startVertex}

	for len(stack) > 0 {
		currentVertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currentVertexHashcode := graph.GetElementHashcode(currentVertex)
		if !visited[currentVertexHashcode] {
			visited[currentVertexHashcode] = true
			currentVertexData := currentVertex.(types.Revision)
			if currentVertexData.ID() == id {
				return currentVertexData
			}
			outgoingEdges := g.OutgoingEdges(currentVertex)
			for _, edge := range outgoingEdges {
				stack = append(stack, edge.TargetVertex())
			}
		}
	}
	return nil
}
