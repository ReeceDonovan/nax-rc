package graph

import (
	"errors"

	"github.com/ReeceDonovan/nax-rc/internal/types"
)

// DAGNode represents a node in the directed acyclic graph (DAG), containing a revision.
type DAGNode struct {
	Revision types.Revision
	Parents  []*DAGNode
	Children []*DAGNode
}

// DAG represents a directed acyclic graph (DAG) data structure.
type DAG struct {
	nodes map[int]*DAGNode
}

// NewDAG initializes and returns an empty directed acyclic graph.
func NewDAG() *DAG {
	return &DAG{nodes: make(map[int]*DAGNode)}
}

// AddNode adds a new node with the specified revision to the DAG.
func (dag *DAG) AddNode(revision types.Revision) (*DAGNode, error) {
	if _, exists := dag.nodes[revision.ID()]; exists {
		return nil, errors.New("node with the given revision ID already exists in the DAG")
	}
	node := &DAGNode{Revision: revision}
	dag.nodes[revision.ID()] = node
	return node, nil
}

// RemoveNode removes a node with the specified revision ID from the DAG.
func (dag *DAG) RemoveNode(revisionID int) error {
	node, exists := dag.nodes[revisionID]
	if !exists {
		return errors.New("node with the given revision ID does not exist in the DAG")
	}

	for _, parent := range node.Parents {
		for i, child := range parent.Children {
			if child == node {
				parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
				break
			}
		}
	}

	for _, child := range node.Children {
		for i, parent := range child.Parents {
			if parent == node {
				child.Parents = append(child.Parents[:i], child.Parents[i+1:]...)
				break
			}
		}
	}

	delete(dag.nodes, revisionID)
	return nil
}

// AddEdge adds a directed edge between two nodes with the specified revision IDs in the DAG.
func (dag *DAG) AddEdge(parentID, childID int) error {
	parent, exists := dag.nodes[parentID]
	if !exists {
		return errors.New("parent node with the given revision ID does not exist in the DAG")
	}

	child, exists := dag.nodes[childID]
	if !exists {
		return errors.New("child node with the given revision ID does not exist in the DAG")
	}

	parent.Children = append(parent.Children, child)
	child.Parents = append(child.Parents, parent)
	return nil
}

// RemoveEdge removes a directed edge between two nodes with the specified revision IDs in the DAG.
func (dag *DAG) RemoveEdge(parentID, childID int) error {
	parent, exists := dag.nodes[parentID]
	if !exists {
		return errors.New("parent node with the given revision ID does not exist in the DAG")
	}

	child, exists := dag.nodes[childID]
	if !exists {
		return errors.New("child node with the given revision ID does not exist in the DAG")
	}

	for i, childNode := range parent.Children {
		if childNode == child {
			parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			break
		}
	}

	for i, parentNode := range child.Parents {
		if parentNode == parent {
			child.Parents = append(child.Parents[:i], child.Parents[i+1:]...)
			break
		}
	}

	return nil
}

// NodeExists checks if a node with the given revision ID exists in the DAG.
func (dag *DAG) NodeExists(revisionID int) bool {
	_, exists := dag.nodes[revisionID]
	return exists
}

// GetNode retrieves a node with the given revision ID from the DAG, returning nil if it doesn't exist.
func (dag *DAG) GetNode(revisionID int) *DAGNode {
	return dag.nodes[revisionID]
}
