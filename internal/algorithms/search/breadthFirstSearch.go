package search

import (
	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/graph"
	"github.com/ReeceDonovan/nax-rc/internal/types"
)

func BreadthFirstSearch(g graph.Graph, startVertex graph.Vertex, id int) types.Revision {
	visited := make(map[graph.Vertex]bool)
	queue := []graph.Vertex{startVertex}

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]

		currentVertexHashcode := graph.GetElementHashcode(currentVertex)
		if !visited[currentVertexHashcode] {
			visited[currentVertexHashcode] = true
			currentVertexData := currentVertex.(types.Revision)
			if currentVertexData.ID() == id {
				return currentVertexData
			}
			outgoingEdges := g.OutgoingEdges(currentVertex)
			for _, edge := range outgoingEdges {
				queue = append(queue, edge.TargetVertex())
			}
		}
	}
	return nil
}
