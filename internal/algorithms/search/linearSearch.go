package search

import (
	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/linkedList"
)

// LinearSearch takes a doubly linked list and a revision ID, and returns the node containing the revision with the given ID or nil if not found.
func LinearSearch(list *linkedList.DoublyLinkedList, revisionID int) *linkedList.Node {
	currentNode := list.Head
	for currentNode != nil {
		if currentNode.Revision.ID() == revisionID {
			return currentNode
		}
		currentNode = currentNode.Next
	}
	return nil
}
