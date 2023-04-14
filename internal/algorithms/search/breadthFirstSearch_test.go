package search

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkBreadthFirstSearch(b *testing.B) {
	for _, dataSize := range DATA_SIZE {
		b.Run(fmt.Sprintf("BreadthFirstSearch_%d", dataSize.dataSize), func(b *testing.B) {
			dataSize := dataSize.dataSize
			// g, revisions := createBigGraph(dataSize)
			for i := 0; i < b.N; i++ {
				g, revisions := createBigGraph(dataSize)
				// select a random target revision
				targetRevisionID := rand.Intn(26) + 65
				var targetRevision = revisions[fmt.Sprintf("%c", targetRevisionID)]
				if targetRevision == nil {
					b.Errorf("expected target revision to not be nil, %v", targetRevisionID)
				}

				result := BreadthFirstSearch(g, revisions["A"], targetRevision.ID())

				if result == nil {
					b.Errorf("BreadthFirstSearch(%v)", targetRevision.ID())
				}

				// b.Logf("Target ID: %v, Result ID: %v", targetRevision.ID(), result.ID())
			}
		})
	}
}
