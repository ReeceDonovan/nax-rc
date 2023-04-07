package linkedList

import (
	"testing"
)

// Helper function to create a list for testing
func createTestList() *LinkedList {
	list := &LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	return list
}

func TestLinkedList_IsEmpty(t *testing.T) {
	list := &LinkedList{}
	if !list.IsEmpty() {
		t.Error("Expected list to be empty")
	}

	list.Append(1)
	if list.IsEmpty() {
		t.Error("Expected list to not be empty")
	}
}

func TestLinkedList_Append(t *testing.T) {
	list := createTestList()
	expected := "1 -> 2 -> 3 -> nil"

	if list.String() != expected {
		t.Errorf("Expected list to be %s, got %s", expected, list.String())
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	list := createTestList()
	list.Prepend(0)
	expected := "0 -> 1 -> 2 -> 3 -> nil"

	if list.String() != expected {
		t.Errorf("Expected list to be %s, got %s", expected, list.String())
	}
}

func TestLinkedList_Remove(t *testing.T) {
	list := createTestList()
	list.Remove(list.Head.Next) // Remove '2'
	expected := "1 -> 3 -> nil"

	if list.String() != expected {
		t.Errorf("Expected list to be %s, got %s", expected, list.String())
	}
}

func BenchmarkLinkedList_Append(b *testing.B) {
	list := &LinkedList{}
	for i := 0; i < b.N; i++ {
		list.Append(i)
	}
}

func BenchmarkLinkedList_Prepend(b *testing.B) {
	list := &LinkedList{}
	for i := 0; i < b.N; i++ {
		list.Prepend(i)
	}
}

func BenchmarkLinkedList_Remove(b *testing.B) {
	list := createTestList()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.Remove(list.Head)
		list.Append(i)
	}
}
