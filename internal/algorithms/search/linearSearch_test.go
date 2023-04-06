package search

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/linkedlist"
	"github.com/ReeceDonovan/nax-rc/internal/dataStructures/revlog"
)

/*
These tests and benchmarks will involve generating a mock linked list that acts as a basic version control system.
The mocked linked list will contain `numRevisions` number of revisions (nodes). Each node will have an ID that is its revision number, and a Data field that is a `dataSize` length byte slice of random data.
We want to test the performance of the linear search algorithm on a linked list of varying sizes, and with varying data sizes.

Create a table of values for the number of revisions and data size, and run benchmarks for each combination.
*/

var NUMBER_OF_REVISIONS = []struct {
	numRevisions int
}{
	{20},
	{200},
	{2000},
	{20000},
}

var DATA_SIZE = []struct {
	dataSize int
}{
	{100},
	{1000},
	{10000},
	{100000},
	{1000000},
}

func BenchmarkLinearSearch(b *testing.B) {
	// var numRevisions = NUMBER_OF_REVISIONS[2].numRevisions
	for _, numRevisions := range NUMBER_OF_REVISIONS {
		for _, dataSize := range DATA_SIZE {
			b.Run(fmt.Sprintf("LinearSearch_%d_%d", numRevisions.numRevisions, dataSize.dataSize), func(b *testing.B) {
				numRevisions := numRevisions.numRevisions
				dataSize := dataSize.dataSize
				// // Create a linked list of revisions
				// list := linkedlist.LinkedList{}
				// // TODO: Consider moving this into the main b.N loop, so that each iteration of the benchmark is testing a different set of revisions
				// for i := 0; i < numRevisions; i++ {
				// 	// Create a revision with random data of size `dataSize`
				// 	revision := revlog.NewRandomRevision(i, dataSize)
				// 	// Add the revision to the linked list
				// 	list.Append(revision)
				// }

				// b.ResetTimer()

				// Run the benchmark
				for i := 0; i < b.N; i++ {
					// Create a linked list of revisions
					list := linkedlist.LinkedList{}
					// TODO: Consider moving this out of the main b.N loop
					for i := 0; i < numRevisions; i++ {
						// Create a revision with random data of size `dataSize`
						revision := revlog.NewRandomRevision(i, dataSize)
						// Add the revision to the linked list
						list.Append(revision)
					}
					// Select a random revision ID to search for in the linked list
					var targetRevision = rand.Intn(int(float64(numRevisions)*0.3)) + int(float64(numRevisions)*0.55)
					// b.Logf("Searching for revision %d", targetRevision)
					// Search for the revision in the linked list
					result := LinearSearch(list, targetRevision)
					// Check that the result is not nil
					if result == nil {
						b.Errorf("Expected to find revision %d, but got nil", targetRevision)
					}
					// b.Logf("Result: %v", result.Data.(revlog.Revision).Data())
				}

			})
		}
	}
}
