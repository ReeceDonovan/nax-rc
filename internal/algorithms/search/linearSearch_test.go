package search

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/linkedList"
	"github.com/ReeceDonovan/nax-rc/internal/types"
	"github.com/stretchr/testify/assert"
)

func createTestList() *linkedList.DoublyLinkedList {
	list := linkedList.NewDoublyLinkedList()
	for i := 0; i < 10; i++ {
		revision := types.NewBlankRevision(i)
		list.InsertAtPosition(i+1, &linkedList.DLLNode{Revision: revision})
	}
	return list
}

func TestLinearSearch(t *testing.T) {
	list := createTestList()

	for i := 0; i < 10; i++ {
		node := LinearSearch(list, i)
		assert.NotNil(t, node, "Node with ID %d not found", i)
		assert.Equal(t, i, node.Revision.ID())
	}

	node := LinearSearch(list, -1)
	assert.Nil(t, node, "Node with ID -1 should not be found")
}

func BenchmarkLinearSearch(b *testing.B) {
	list := createTestList()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		revisionID := rand.Intn(10)
		LinearSearch(list, revisionID)
	}
}
func createListWithRevisions(numRevisions, dataSize int) *linkedList.DoublyLinkedList {
	list := linkedList.NewDoublyLinkedList()
	for i := 0; i < numRevisions; i++ {
		revision := types.NewRandomRevision(i, dataSize)
		list.InsertAtPosition(i+1, &linkedList.DLLNode{Revision: revision})
	}
	return list
}

func BenchmarkLinearSearch_RevisionSearch(b *testing.B) {
	for _, numRevisions := range []int{smallNumRevisions, mediumNumRevisions, largeNumRevisions} {
		for _, dataSize := range []int{smallDataSize, mediumDataSize, largeDataSize} {
			b.Run(fmt.Sprintf("numRevisions=%d,dataSize=%d", numRevisions, dataSize), func(b *testing.B) {
				list := createListWithRevisions(numRevisions, dataSize)
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					revisionID := rand.Intn(numRevisions)
					LinearSearch(list, revisionID)
				}
			})
		}
	}
}
