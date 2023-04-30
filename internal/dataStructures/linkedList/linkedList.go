package linkedList

import (
	"github.com/ReeceDonovan/nax-rc/internal/types"
)

// DLLNode represents a doubly linked list node containing a revision.
type DLLNode struct {
	Revision types.Revision
	Next     *DLLNode
	Prev     *DLLNode
}

// DoublyLinkedList represents a doubly linked list data structure.
type DoublyLinkedList struct {
	Head *DLLNode
	Tail *DLLNode
}

// NewDoublyLinkedList initializes and returns an empty doubly linked list.
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// AssignHead sets the given node as the head of the doubly linked list.
func (dll *DoublyLinkedList) AssignHead(node *DLLNode) {
	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		return
	}
	dll.InsertPrior(dll.Head, node)
}

// AssignTail sets the given node as the tail of the doubly linked list.
func (dll *DoublyLinkedList) AssignTail(node *DLLNode) {
	if dll.Tail == nil {
		dll.AssignHead(node)
		return
	}
	dll.InsertSubsequent(dll.Tail, node)
}

// InsertPrior inserts a new node before the specified node in the doubly linked list.
func (dll *DoublyLinkedList) InsertPrior(node, nodeToInsert *DLLNode) {
	if nodeToInsert == dll.Head && nodeToInsert == dll.Tail {
		return
	}
	dll.Remove(nodeToInsert)
	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node
	if node.Prev == nil {
		dll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}
	node.Prev = nodeToInsert
}

// InsertSubsequent inserts a new node after the specified node in the doubly linked list.
func (dll *DoublyLinkedList) InsertSubsequent(node, nodeToInsert *DLLNode) {
	if nodeToInsert == dll.Head && nodeToInsert == dll.Tail {
		return
	}
	dll.Remove(nodeToInsert)
	nodeToInsert.Prev = node
	nodeToInsert.Next = node.Next
	if node.Next == nil {
		dll.Tail = nodeToInsert
	} else {
		node.Next.Prev = nodeToInsert
	}
	node.Next = nodeToInsert
}

// InsertAtPosition inserts a new node at the specified position in the doubly linked list.
func (dll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *DLLNode) {
	if position == 1 {
		dll.AssignHead(nodeToInsert)
		return
	}
	currentNode := dll.Head
	currentPosition := 1
	for currentNode != nil && currentPosition != position {
		currentNode = currentNode.Next
		currentPosition++
	}
	if currentNode != nil {
		dll.InsertPrior(currentNode, nodeToInsert)
	} else {
		dll.AssignTail(nodeToInsert)
	}
}

// RemoveNodesWithID removes all nodes with the specified revision ID from the doubly linked list.
func (dll *DoublyLinkedList) RemoveNodesWithID(id int) {
	currentNode := dll.Head
	for currentNode != nil {
		nodeToRemove := currentNode
		currentNode = currentNode.Next
		if nodeToRemove.Revision.ID() == id {
			dll.Remove(nodeToRemove)
		}
	}
}

// Remove removes the specified node from the doubly linked list.
func (dll *DoublyLinkedList) Remove(node *DLLNode) {
	if node == dll.Head {
		dll.Head = dll.Head.Next
	}
	if node == dll.Tail {
		dll.Tail = dll.Tail.Prev
	}
	dll.removeNodeBindings(node)
}

// ContainsNodeWithID returns true if the doubly linked list contains a node with the specified revision ID.
func (dll *DoublyLinkedList) ContainsNodeWithID(id int) bool {
	currentNode := dll.Head
	for currentNode != nil && currentNode.Revision.ID() != id {
		currentNode = currentNode.Next
	}
	return currentNode != nil
}

// removeNodeBindings removes the bindings of the specified node from the doubly linked list.
func (dll *DoublyLinkedList) removeNodeBindings(node *DLLNode) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}
