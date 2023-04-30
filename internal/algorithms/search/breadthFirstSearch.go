package search

import (
	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/graph"
)

// BreadthFirstSearch performs a breadth-first search on the given Directed Acyclic Graph (DAG),
// starting from the node with the specified revision ID. The visitFunc function is called
// for each visited node during the search.
func BreadthFirstSearch(dag *graph.DAG, startRevisionID int, visitFunc func(node *graph.DAGNode)) {
	startNode := dag.GetNode(startRevisionID)
	if startNode == nil {
		return
	}

	visitedNodes := make(map[int]bool)
	nodeQueue := []*graph.DAGNode{startNode}

	for len(nodeQueue) > 0 {
		currentNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		if visitedNodes[currentNode.Revision.ID()] {
			continue
		}

		visitFunc(currentNode)
		visitedNodes[currentNode.Revision.ID()] = true

		for _, childNode := range currentNode.Children {
			if !visitedNodes[childNode.Revision.ID()] {
				nodeQueue = append(nodeQueue, childNode)
			}
		}
	}
}
