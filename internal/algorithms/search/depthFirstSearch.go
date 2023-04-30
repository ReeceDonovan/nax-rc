package search

import (
	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/graph"
)

// DepthFirstSearch performs a depth-first search on the given Directed Acyclic Graph (DAG),
// starting at the node with the given revision ID, and applies the provided visit function
// to each visited node.
func DepthFirstSearch(dag *graph.DAG, startRevisionID int, visitFunc func(*graph.DAGNode)) {
	startNode := dag.GetNode(startRevisionID)
	if startNode == nil {
		return
	}

	visitedNodes := make(map[int]bool)
	depthFirstTraversal(startNode, visitedNodes, visitFunc)
}

func depthFirstTraversal(node *graph.DAGNode, visitedNodes map[int]bool, visitFunc func(*graph.DAGNode)) {
	if visitedNodes[node.Revision.ID()] {
		return
	}

	visitedNodes[node.Revision.ID()] = true
	visitFunc(node)

	for _, childNode := range node.Children {
		depthFirstTraversal(childNode, visitedNodes, visitFunc)
	}
}
