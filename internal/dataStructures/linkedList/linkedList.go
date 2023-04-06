package linkedlist

import (
	"bytes"
	"fmt"
)

type LinkedListNode struct {
	Data interface{}
	Next *LinkedListNode
	Prev *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
	Tail *LinkedListNode
}

func (list *LinkedList) IsEmpty() bool {
	return list.Head == nil
}

func (list *LinkedList) Append(data interface{}) {
	node := &LinkedListNode{Data: data}
	if list.IsEmpty() {
		list.Head = node
		list.Tail = node
		return
	}
	list.Tail.Next = node
	node.Prev = list.Tail
	list.Tail = node
}

func (list *LinkedList) Prepend(data interface{}) {
	node := &LinkedListNode{Data: data}
	if list.IsEmpty() {
		list.Head = node
		list.Tail = node
		return
	}
	list.Head.Prev = node
	node.Next = list.Head
	list.Head = node
}

func (list *LinkedList) Remove(node *LinkedListNode) {
	if node == list.Head {
		list.Head = list.Head.Next
	}
	if node == list.Tail {
		list.Tail = list.Tail.Prev
	}
	list.removeBindings(node)
}

func (list *LinkedList) removeBindings(node *LinkedListNode) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}

func (list *LinkedList) String() string {
	var buffer bytes.Buffer
	current := list.Head
	for current != nil {
		buffer.WriteString(fmt.Sprintf("%v -> ", current.Data))
		current = current.Next
	}
	buffer.WriteString("nil")
	return buffer.String()
}
