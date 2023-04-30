package search

import (
	"errors"

	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/binaryTree"
	"github.com/ReeceDonovan/nax-rc/internal/types"
)

// BinarySearch takes a binary tree node and a revision ID, and returns the revision with the given ID if found.
func BinarySearch(tree *binaryTree.TreeNode, id int) (types.Revision, error) {
	currentNode := tree
	for currentNode != nil {
		if id < currentNode.Revision.ID() {
			currentNode = currentNode.Left
		} else if id > currentNode.Revision.ID() {
			currentNode = currentNode.Right
		} else {
			return currentNode.Revision, nil
		}
	}
	return nil, errors.New("revision not found")
}
