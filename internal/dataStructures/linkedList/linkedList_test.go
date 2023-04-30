package linkedList

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestNewDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList()

	assert.Nil(t, list.Head)
	assert.Nil(t, list.Tail)
}

func TestDoublyLinkedList_AssignHead(t *testing.T) {
	list := NewDoublyLinkedList()
	node := &DLLNode{Revision: types.NewBlankRevision(1)}

	list.AssignHead(node)

	assert.Equal(t, node, list.Head)
	assert.Equal(t, node, list.Tail)
}

func TestDoublyLinkedList_AssignTail(t *testing.T) {
	list := NewDoublyLinkedList()
	node := &DLLNode{Revision: types.NewBlankRevision(1)}

	list.AssignTail(node)

	assert.Equal(t, node, list.Head)
	assert.Equal(t, node, list.Tail)
}

func TestDoublyLinkedList_InsertPrior(t *testing.T) {
	list := NewDoublyLinkedList()
	node1 := &DLLNode{Revision: types.NewBlankRevision(1)}
	node2 := &DLLNode{Revision: types.NewBlankRevision(2)}

	list.AssignHead(node1)
	list.InsertPrior(node1, node2)

	assert.Equal(t, node2, list.Head)
	assert.Equal(t, node1, list.Tail)
	assert.Equal(t, node1, node2.Next)
	assert.Equal(t, node2, node1.Prev)
}

func TestDoublyLinkedList_InsertSubsequent(t *testing.T) {
	list := NewDoublyLinkedList()
	node1 := &DLLNode{Revision: types.NewBlankRevision(1)}
	node2 := &DLLNode{Revision: types.NewBlankRevision(2)}

	list.AssignHead(node1)
	list.InsertSubsequent(node1, node2)

	assert.Equal(t, node1, list.Head)
	assert.Equal(t, node2, list.Tail)
	assert.Equal(t, node2, node1.Next)
	assert.Equal(t, node1, node2.Prev)
}

func TestDoublyLinkedList_InsertAtPosition(t *testing.T) {
	list := NewDoublyLinkedList()
	node1 := &DLLNode{Revision: types.NewBlankRevision(1)}
	node2 := &DLLNode{Revision: types.NewBlankRevision(2)}
	node3 := &DLLNode{Revision: types.NewBlankRevision(3)}

	list.AssignHead(node1)
	list.AssignTail(node3)
	list.InsertAtPosition(2, node2)

	assert.Equal(t, node1, list.Head)
	assert.Equal(t, node3, list.Tail)
	assert.Equal(t, node2, node1.Next)
	assert.Equal(t, node1, node2.Prev)
	assert.Equal(t, node3, node2.Next)
	assert.Equal(t, node2, node3.Prev)
}

func TestDoublyLinkedList_RemoveNodesWithID(t *testing.T) {
	list := NewDoublyLinkedList()
	node1 := &DLLNode{Revision: types.NewBlankRevision(1)}
	node2 := &DLLNode{Revision: types.NewBlankRevision(2)}
	node3 := &DLLNode{Revision: types.NewBlankRevision(2)}

	list.AssignHead(node1)
	list.AssignTail(node2)
	list.InsertSubsequent(node2, node3)
	list.RemoveNodesWithID(2)

	assert.Equal(t, node1, list.Head)
	assert.Equal(t, node1, list.Tail)
	assert.Nil(t, node1.Prev)
	assert.Nil(t, node1.Next)
}

func TestDoublyLinkedList_ContainsNodeWithID(t *testing.T) {
	list := NewDoublyLinkedList()
	node1 := &DLLNode{Revision: types.NewBlankRevision(1)}
	node2 := &DLLNode{Revision: types.NewBlankRevision(2)}
	node3 := &DLLNode{Revision: types.NewBlankRevision(3)}

	list.AssignHead(node1)
	list.AssignTail(node2)
	list.InsertSubsequent(node2, node3)

	assert.True(t, list.ContainsNodeWithID(1))
	assert.True(t, list.ContainsNodeWithID(2))
	assert.True(t, list.ContainsNodeWithID(3))
	assert.False(t, list.ContainsNodeWithID(4))
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	list := NewDoublyLinkedList()
	node1 := &DLLNode{Revision: types.NewBlankRevision(1)}
	node2 := &DLLNode{Revision: types.NewBlankRevision(2)}
	node3 := &DLLNode{Revision: types.NewBlankRevision(3)}

	list.AssignHead(node1)
	list.AssignTail(node2)
	list.InsertSubsequent(node2, node3)
	list.Remove(node2)

	assert.Equal(t, node1, list.Head)
	assert.Equal(t, node3, list.Tail)
	assert.Equal(t, node3, node1.Next)
	assert.Equal(t, node1, node3.Prev)
}

func TestDoublyLinkedList_removeNodeBindings(t *testing.T) {
	list := NewDoublyLinkedList()
	node1 := &DLLNode{Revision: types.NewBlankRevision(1)}
	node2 := &DLLNode{Revision: types.NewBlankRevision(2)}
	node3 := &DLLNode{Revision: types.NewBlankRevision(3)}

	list.AssignHead(node1)
	list.AssignTail(node2)
	list.InsertSubsequent(node2, node3)
	list.removeNodeBindings(node2)

	assert.Equal(t, node1, list.Head)
	assert.Equal(t, node3, list.Tail)
	assert.Equal(t, node3, node1.Next)
	assert.Equal(t, node1, node3.Prev)
	assert.Nil(t, node2.Prev)
	assert.Nil(t, node2.Next)
}

func generateRandomNode() *DLLNode {
	revisionID := rand.Intn(1000)
	dataSize := rand.Intn(100)
	revision := types.NewRandomRevision(revisionID, dataSize)
	return &DLLNode{Revision: revision}
}

func BenchmarkDoublyLinkedList_InsertPrior(b *testing.B) {
	list := NewDoublyLinkedList()
	node := generateRandomNode()
	list.AssignHead(node)

	for i := 0; i < b.N; i++ {
		newNode := generateRandomNode()
		list.InsertPrior(node, newNode)
	}
}

func BenchmarkDoublyLinkedList_InsertSubsequent(b *testing.B) {
	list := NewDoublyLinkedList()
	node := generateRandomNode()
	list.AssignHead(node)

	for i := 0; i < b.N; i++ {
		newNode := generateRandomNode()
		list.InsertSubsequent(node, newNode)
	}
}

func BenchmarkDoublyLinkedList_InsertAtPosition(b *testing.B) {
	list := NewDoublyLinkedList()
	node := generateRandomNode()
	list.AssignHead(node)

	for i := 0; i < b.N; i++ {
		newNode := generateRandomNode()
		position := rand.Intn(i+2) + 1
		list.InsertAtPosition(position, newNode)
	}
}

func BenchmarkDoublyLinkedList_RemoveNodesWithID(b *testing.B) {
	list := NewDoublyLinkedList()
	node := generateRandomNode()
	list.AssignHead(node)

	for i := 0; i < b.N; i++ {
		newNode := generateRandomNode()
		list.InsertSubsequent(node, newNode)
		node = newNode
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		revisionID := rand.Intn(b.N)
		list.RemoveNodesWithID(revisionID)
	}
}

func BenchmarkDoublyLinkedList_ContainsNodeWithID(b *testing.B) {
	list := NewDoublyLinkedList()
	node := generateRandomNode()
	list.AssignHead(node)

	for i := 0; i < b.N; i++ {
		newNode := generateRandomNode()
		list.InsertSubsequent(node, newNode)
		node = newNode
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		revisionID := rand.Intn(b.N)
		list.ContainsNodeWithID(revisionID)
	}
}

const (
	smallNumRevisions  = 200
	mediumNumRevisions = 2000
	largeNumRevisions  = 20000
	largeDataSize      = 10000 // 10 Kb
	mediumDataSize     = 1000  // 1 Kb
	smallDataSize      = 100   // 100 bytes
)

// Create a benchmarks for constructing a list of revisions with different sizes
func BenchmarkDoublyLinkedList_RevisionConstruction(b *testing.B) {
	for _, numRevisions := range []int{smallNumRevisions, mediumNumRevisions, largeNumRevisions} {
		for _, dataSize := range []int{smallDataSize, mediumDataSize, largeDataSize} {
			b.Run(fmt.Sprintf("numRevisions=%d,dataSize=%d", numRevisions, dataSize), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					list := NewDoublyLinkedList()
					for j := 0; j < numRevisions; j++ {
						revision := types.NewRandomRevision(j, dataSize)
						list.InsertAtPosition(j+1, &DLLNode{Revision: revision})
					}
				}
			})
		}
	}
}
