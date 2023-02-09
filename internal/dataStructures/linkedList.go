package datastructures

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Prev *Node
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

// Create an empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Display the linked list
func (ll *LinkedList) Show() { // O(n) time | O(1) space
	node := ll.Head
	for node != nil {
		fmt.Printf("%v ", node.Data)
		node = node.Next
	}
	fmt.Println()
}

// Display the linked list in reverse
func (ll *LinkedList) ShowReverse() { // O(n) time | O(1) space
	node := ll.Tail
	for node != nil {
		fmt.Printf("%v ", node.Data)
		node = node.Prev
	}
	fmt.Println()
}

// Set the head of the linked list
func (ll *LinkedList) SetHead(node *Node) { // O(1) time | O(1) space
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.InsertBefore(ll.Head, node)
}

// Set the tail of the linked list
func (ll *LinkedList) SetTail(node *Node) { // O(1) time | O(1) space
	if ll.Tail == nil {
		ll.SetHead(node)
		return
	}
	ll.InsertAfter(ll.Tail, node)
}

// Insert a node before another node in the linked list
func (ll *LinkedList) InsertBefore(node, nodeToInsert *Node) { // O(1) time | O(1) space
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}
	ll.Remove(nodeToInsert)
	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node
	if node.Prev == nil {
		ll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}
	node.Prev = nodeToInsert
}

// Insert a node after another node in the linked list
func (ll *LinkedList) InsertAfter(node, nodeToInsert *Node) { // O(1) time | O(1) space
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}
	ll.Remove(nodeToInsert)
	nodeToInsert.Prev = node
	nodeToInsert.Next = node.Next
	if node.Next == nil {
		ll.Tail = nodeToInsert
	} else {
		node.Next.Prev = nodeToInsert
	}
	node.Next = nodeToInsert
}

// Insert a node at a position in the linked list
func (ll *LinkedList) InsertAtPosition(position int, nodeToInsert *Node) { // O(p) time | O(1) space
	if position == 1 {
		ll.SetHead(nodeToInsert)
		return
	}
	node := ll.Head
	currentPosition := 1
	for node != nil && currentPosition != position {
		node = node.Next
		currentPosition++
	}
	if node != nil {
		ll.InsertBefore(node, nodeToInsert)
	} else {
		ll.SetTail(nodeToInsert)
	}
}

// Check if a node exists in the linked list
func (ll *LinkedList) ContainsNodeWithData(data interface{}) bool { // O(n) time | O(1) space
	node := ll.Head
	for node != nil && node.Data != data {
		node = node.Next
	}
	return node != nil
}

// Remove a node with specific data from the linked list
func (ll *LinkedList) RemoveNodesWithData(data interface{}) { // O(n) time | O(1) space
	node := ll.Head
	for node != nil {
		nodeToRemove := node
		node = node.Next
		if nodeToRemove.Data == data {
			ll.Remove(nodeToRemove)
		}
	}
}

// Remove a node from the linked list
func (ll *LinkedList) Remove(node *Node) { // O(1) time | O(1) space
	if node == ll.Head {
		ll.Head = ll.Head.Next
	}
	if node == ll.Tail {
		ll.Tail = ll.Tail.Prev
	}
	ll.removeNodeBindings(node)
}

// Remove all nodes from the linked list
func (ll *LinkedList) RemoveAllNodes() { // O(n) time | O(1) space
	node := ll.Head
	for node != nil {
		next := node.Next
		ll.Remove(node)
		node = next
	}
}

// Remove all bindings from a node
func (ll *LinkedList) removeNodeBindings(node *Node) { // O(1) time | O(1) space
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}
