package linkedList

import (
	"github.com/ReeceDonovan/nax-rc/internal/types"
)

// Node represents a doubly linked list node containing a revision.
type Node struct {
	Revision types.Revision
	Next     *Node
	Prev     *Node
}

// DoublyLinkedList represents a doubly linked list data structure.
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

// NewDoublyLinkedList initializes and returns an empty doubly linked list.
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// SetHead sets the given node as the head of the doubly linked list.
func (list *DoublyLinkedList) SetHead(node *Node) {
	if list.Head == nil {
		list.Head = node
		list.Tail = node
		return
	}
	list.InsertBefore(list.Head, node)
}

// SetTail sets the given node as the tail of the doubly linked list.
func (list *DoublyLinkedList) SetTail(node *Node) {
	if list.Tail == nil {
		list.SetHead(node)
		return
	}
	list.InsertAfter(list.Tail, node)
}

// InsertBefore inserts a new node before the specified node in the doubly linked list.
func (list *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	if nodeToInsert == list.Head && nodeToInsert == list.Tail {
		return
	}
	list.Remove(nodeToInsert)
	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node
	if node.Prev == nil {
		list.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}
	node.Prev = nodeToInsert
}

// InsertAfter inserts a new node after the specified node in the doubly linked list.
func (list *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	if nodeToInsert == list.Head && nodeToInsert == list.Tail {
		return
	}
	list.Remove(nodeToInsert)
	nodeToInsert.Prev = node
	nodeToInsert.Next = node.Next
	if node.Next == nil {
		list.Tail = nodeToInsert
	} else {
		node.Next.Prev = nodeToInsert
	}
	node.Next = nodeToInsert
}

// InsertAtPosition inserts a new node at the specified position in the doubly linked list.
func (list *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	if position == 1 {
		list.SetHead(nodeToInsert)
		return
	}
	currentNode := list.Head
	currentPosition := 1
	for currentNode != nil && currentPosition != position {
		currentNode = currentNode.Next
		currentPosition++
	}
	if currentNode != nil {
		list.InsertBefore(currentNode, nodeToInsert)
	} else {
		list.SetTail(nodeToInsert)
	}
}

// RemoveNodesWithID removes all nodes with the specified revision ID from the doubly linked list.
func (list *DoublyLinkedList) RemoveNodesWithID(id int) {
	currentNode := list.Head
	for currentNode != nil {
		nodeToRemove := currentNode
		currentNode = currentNode.Next
		if nodeToRemove.Revision.ID() == id {
			list.Remove(nodeToRemove)
		}
	}
}

// Remove removes the specified node from the doubly linked list.
func (list *DoublyLinkedList) Remove(node *Node) {
	if node == list.Head {
		list.Head = list.Head.Next
	}
	if node == list.Tail {
		list.Tail = list.Tail.Prev
	}
	list.removeNodeBindings(node)
}

// ContainsNodeWithID returns true if the doubly linked list contains a node with the specified revision ID.
func (list *DoublyLinkedList) ContainsNodeWithID(id int) bool {
	currentNode := list.Head
	for currentNode != nil && currentNode.Revision.ID() != id {
		currentNode = currentNode.Next
	}
	return currentNode != nil
}

// removeNodeBindings removes the bindings of the specified node from the doubly linked list.
func (list *DoublyLinkedList) removeNodeBindings(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}
