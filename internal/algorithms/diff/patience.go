package diff

import (
	"container/heap"
	"sort"
)

// PatienceDiff computes the diff between two slices of strings (source and destination)
func PatienceDiff(source, destination []string) []string {
	lcs := patienceLCS(source, destination)
	editScript := generateEditScript(source, destination, lcs)
	return generateDiffOutput(source, destination, editScript)
}

// IndexPair represents a pair of indices in the source and destination slices
type IndexPair struct {
	sourceIndex      int
	destinationIndex int
}

// patienceLCS computes the Longest Common Subsequence (LCS) using the patience sorting algorithm
func patienceLCS(source, destination []string) []IndexPair {
	var piles []IntHeap
	sourceIndices := make(map[string][]int)
	for i, value := range source {
		sourceIndices[value] = append(sourceIndices[value], i)
	}

	for i, value := range destination {
		if sourceIndexList, found := sourceIndices[value]; found && len(sourceIndexList) > 0 {
			sourceIndex := sourceIndices[value][0]
			sourceIndices[value] = sourceIndices[value][1:]

			inserted := false
			for _, pile := range piles {
				if top := pile.Top(); top != nil && top.sourceIndex < sourceIndex {
					heap.Push(&pile, IndexPair{sourceIndex, i})
					inserted = true
					break
				}
			}

			if !inserted {
				var newPile IntHeap
				heap.Push(&newPile, IndexPair{sourceIndex, i})
				piles = append(piles, newPile)
			}
		}
	}

	var lcs []IndexPair
	for _, pile := range piles {
		lcs = append(lcs, *pile.Top())
	}

	sort.Slice(lcs, func(i, j int) bool {
		return lcs[i].sourceIndex < lcs[j].sourceIndex
	})

	return lcs
}

// generateEditScript generates the edit script from the LCS
func generateEditScript(source, destination []string, lcs []IndexPair) []EditOperation {
	var editScript []EditOperation

	sourceIndex, destinationIndex := 0, 0
	for _, pair := range lcs {
		for sourceIndex < pair.sourceIndex {
			editScript = append(editScript, DELETE)
			sourceIndex++
		}

		for destinationIndex < pair.destinationIndex {
			editScript = append(editScript, INSERT)
			destinationIndex++
		}

		editScript = append(editScript, MATCH)
		sourceIndex++
		destinationIndex++
	}

	for sourceIndex < len(source) {
		editScript = append(editScript, DELETE)
		sourceIndex++
	}

	for destinationIndex < len(destination) {
		editScript = append(editScript, INSERT)
		destinationIndex++
	}

	return editScript
}

// IntHeap is a min-heap of IndexPair elements
type IntHeap []IndexPair

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].destinationIndex < h[j].destinationIndex }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(IndexPair))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Top returns the top element of the IntHeap
func (h *IntHeap) Top() *IndexPair {
	if len(*h) == 0 {
		return nil
	}
	return &((*h)[0])
}
