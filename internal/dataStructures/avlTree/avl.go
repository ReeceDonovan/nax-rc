package avlTree

import (
	"errors"

	"github.com/ReeceDonovan/nax-rc/internal/types"
)

// TreeNode represents a node in the AVL tree, containing a revision.
type TreeNode struct {
	Revision types.Revision
	Height   int
	Left     *TreeNode
	Right    *TreeNode
}

// Insert inserts a revision into the AVL tree.
func (node *TreeNode) Insert(revision types.Revision) *TreeNode {
	if node == nil {
		return &TreeNode{Revision: revision, Height: 1}
	}
	if revision.ID() < node.Revision.ID() {
		node.Left = node.Left.Insert(revision)
	} else {
		node.Right = node.Right.Insert(revision)
	}
	return node.balance()
}

// getHeight returns the height of the given node.
func (node *TreeNode) getHeight() int {
	if node == nil {
		return 0
	}
	return node.Height
}

// updateHeight updates the height of the given node based on its children.
func (node *TreeNode) updateHeight() {
	node.Height = max(node.Left.getHeight(), node.Right.getHeight()) + 1
}

// getBalanceFactor returns the balance factor of the given node.
func (node *TreeNode) getBalanceFactor() int {
	if node == nil {
		return 0
	}
	return node.Left.getHeight() - node.Right.getHeight()
}

// balance checks if the node is unbalanced and performs rotations to balance it.
func (node *TreeNode) balance() *TreeNode {
	if node == nil {
		return nil
	}
	node.updateHeight()

	balanceFactor := node.getBalanceFactor()

	if balanceFactor > 1 {
		if node.Left.getBalanceFactor() < 0 {
			node.Left = node.Left.rotateLeft()
		}
		return node.rotateRight()
	} else if balanceFactor < -1 {
		if node.Right.getBalanceFactor() > 0 {
			node.Right = node.Right.rotateRight()
		}
		return node.rotateLeft()
	}
	return node
}

// rotateLeft performs a left rotation on the given node.
func (node *TreeNode) rotateLeft() *TreeNode {
	newRoot := node.Right
	node.Right = newRoot.Left
	newRoot.Left = node
	node.updateHeight()
	newRoot.updateHeight()
	return newRoot
}

// rotateRight performs a right rotation on the given node.
func (node *TreeNode) rotateRight() *TreeNode {
	newRoot := node.Left
	node.Left = newRoot.Right
	newRoot.Right = node
	node.updateHeight()
	newRoot.updateHeight()
	return newRoot
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Contains checks if a revision with the given ID exists in the AVL tree.
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

// Remove removes a revision with the given ID from the AVL tree.
func (node *TreeNode) Remove(id int) (*TreeNode, error) {
	if node == nil {
		return nil, errors.New("cannot remove from an empty tree")
	}
	node.remove(id, nil)
	return node, nil
}

// remove is a helper function for Remove. It removes a revision with the given ID from the AVL tree.
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

// getMinRevision is a helper function to find the minimum revision in the tree.
func (node *TreeNode) getMinRevision() types.Revision {
	currentNode := node
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode.Revision
}
