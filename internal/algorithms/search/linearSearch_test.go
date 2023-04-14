package search

import (
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/linkedList"
	"github.com/ReeceDonovan/nax-rc/internal/types"
)

const (
	smallNumRevisions  = 20
	mediumNumRevisions = 200
	largeNumRevisions  = 2000
	largeDataSize      = 1000000 // 1 Mb
	mediumDataSize     = 1000    // 1 Kb
	smallDataSize      = 100     // 100 bytes
)

func BenchmarkLinearSearch_SmallListSmallData(b *testing.B) {
	list := createList(smallNumRevisions, smallDataSize)
	targetID := smallNumRevisions - 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(list, targetID)
	}
}

func BenchmarkLinearSearch_SmallListMediumData(b *testing.B) {
	list := createList(smallNumRevisions, mediumDataSize)
	targetID := smallNumRevisions - 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(list, targetID)
	}
}

func BenchmarkLinearSearch_SmallListLargeData(b *testing.B) {
	list := createList(smallNumRevisions, largeDataSize)
	targetID := smallNumRevisions - 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(list, targetID)
	}
}

func BenchmarkLinearSearch_MediumListSmallData(b *testing.B) {
	list := createList(mediumNumRevisions, smallDataSize)
	targetID := mediumNumRevisions - 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(list, targetID)
	}
}

func BenchmarkLinearSearch_MediumListMediumData(b *testing.B) {
	list := createList(mediumNumRevisions, mediumDataSize)
	targetID := mediumNumRevisions - 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(list, targetID)
	}
}

func BenchmarkLinearSearch_MediumListLargeData(b *testing.B) {
	list := createList(mediumNumRevisions, largeDataSize)
	targetID := mediumNumRevisions - 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(list, targetID)
	}
}

func BenchmarkLinearSearch_LargeListSmallData(b *testing.B) {
	list := createList(largeNumRevisions, smallDataSize)
	targetID := largeNumRevisions - 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(list, targetID)
	}
}

func BenchmarkLinearSearch_LargeListMediumData(b *testing.B) {
	list := createList(largeNumRevisions, mediumDataSize)
	targetID := largeNumRevisions - 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(list, targetID)
	}
}

func BenchmarkLinearSearch_LargeListLargeData(b *testing.B) {
	list := createList(largeNumRevisions, largeDataSize)
	targetID := largeNumRevisions - 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(list, targetID)
	}
}

func createList(numRevisions, dataSize int) *linkedList.DoublyLinkedList {
	list := linkedList.NewDoublyLinkedList()
	for i := 0; i < numRevisions; i++ {
		revision := types.NewRandomRevision(i, dataSize)
		node := &linkedList.Node{Revision: revision}
		list.SetTail(node)
	}
	return list
}
