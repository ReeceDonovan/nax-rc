package binaryTree

import (
	"errors"

	"github.com/ReeceDonovan/nax-rc/internal/types"
)

// TreeNode represents a node in the binary search tree, containing a revision.
type TreeNode struct {
	Revision types.Revision
	Left     *TreeNode
	Right    *TreeNode
}

// Insert inserts a revision into the binary search tree.
func (node *TreeNode) Insert(revision types.Revision) *TreeNode {
	currentNode := node
	for {
		if revision.ID() < currentNode.Revision.ID() {
			if currentNode.Left == nil {
				currentNode.Left = &TreeNode{Revision: revision}
				break
			} else {
				currentNode = currentNode.Left
			}
		} else {
			if currentNode.Right == nil {
				currentNode.Right = &TreeNode{Revision: revision}
				break
			} else {
				currentNode = currentNode.Right
			}
		}
	}
	return node
}

// Contains checks if a revision with the given ID exists in the binary search tree.
func (node *TreeNode) Contains(id int) bool {
	currentNode := node
	for currentNode != nil {
		if id < currentNode.Revision.ID() {
			currentNode = currentNode.Left
		} else if id > currentNode.Revision.ID() {
			currentNode = currentNode.Right
		} else {
			return true
		}
	}
	return false
}

// Remove removes a revision with the given ID from the binary search tree.
func (node *TreeNode) Remove(id int) (*TreeNode, error) {
	if node == nil {
		return nil, errors.New("cannot remove from an empty tree")
	}
	node.remove(id, nil)
	return node, nil
}

// remove is a helper function for Remove. It removes a revision with the given ID from the binary search tree.
func (node *TreeNode) remove(id int, parentNode *TreeNode) {
	currentNode := node
	for currentNode != nil {
		if id < currentNode.Revision.ID() {
			parentNode = currentNode
			currentNode = currentNode.Left
		} else if id > currentNode.Revision.ID() {
			parentNode = currentNode
			currentNode = currentNode.Right
		} else {
			if currentNode.Left != nil && currentNode.Right != nil {
				currentNode.Revision = currentNode.Right.getMinRevision()
				currentNode.Right.remove(currentNode.Revision.ID(), currentNode)
			} else if parentNode == nil {
				if currentNode.Left != nil {
					currentNode.Revision = currentNode.Left.Revision
					currentNode.Right = currentNode.Left.Right
					currentNode.Left = currentNode.Left.Left
				} else if currentNode.Right != nil {
					currentNode.Revision = currentNode.Right.Revision
					currentNode.Left = currentNode.Right.Left
					currentNode.Right = currentNode.Right.Right
				} else {
					// This is a single-node tree; do nothing.
				}
			} else if parentNode.Left == currentNode {
				if currentNode.Left != nil {
					parentNode.Left = currentNode.Left
				} else {
					parentNode.Left = currentNode.Right
				}
			} else if parentNode.Right == currentNode {
				if currentNode.Left != nil {
					parentNode.Right = currentNode.Left
				} else {
					parentNode.Right = currentNode.Right
				}
			}
			break
		}
	}
}

// getMinRevision is a helper function to find the minimum revision in the binary search tree.
func (node *TreeNode) getMinRevision() types.Revision {
	currentNode := node
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode.Revision
}
