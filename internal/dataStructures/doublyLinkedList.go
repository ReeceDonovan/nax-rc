package datastructures

import "fmt"

type DoublyLinkedListNode struct {
	Data Revision
	Prev *DoublyLinkedListNode
	Next *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	Head *DoublyLinkedListNode
	Tail *DoublyLinkedListNode
}

// Create an empty doubly linked list.
func NewDoublyLinkedList() Revlog {
	return &DoublyLinkedList{}
}

// Append a new node to the end of the list.
func (dll *DoublyLinkedList) Add(rev Revision) error {
	node := &DoublyLinkedListNode{Data: rev}
	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		return nil
	}
	dll.Tail.Next = node
	node.Prev = dll.Tail
	dll.Tail = node
	return nil
}

// Get the revision with the given revision id.
func (dll *DoublyLinkedList) Get(rId int) (Revision, error) {
	node := dll.Head
	for node != nil {
		if node.Data.ID == rId {
			return node.Data, nil
		}
		node = node.Next
	}
	return Revision{}, fmt.Errorf("Revision with id %d not found", rId)
}

// Remove the revision with the given revision id.
func (dll *DoublyLinkedList) Remove(rId int) error {
	node := dll.Head
	for node != nil {
		if node.Data.ID == rId {
			if node == dll.Head {
				dll.Head = dll.Head.Next
			}
			if node == dll.Tail {
				dll.Tail = dll.Tail.Prev
			}
			dll.removeBindings(node)
			return nil
		}
		node = node.Next
	}
	return fmt.Errorf("Revision with id %d not found", rId)
}

// Remove all bindings from a node.
func (dll *DoublyLinkedList) removeBindings(node *DoublyLinkedListNode) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}

// Count the number of revisions in the log.
func (dll *DoublyLinkedList) Count() int {
	count := 0
	node := dll.Head
	for node != nil {
		count++
		node = node.Next
	}
	return count
}
