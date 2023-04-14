package lis

// O(n^2) time | O(n) space
func longestIncreasingSubsequenceSlow(array []int) []int {
	// Initialize slices to store predecessor indices and LIS lengths for subproblems.
	predecessors := make([]int, len(array))
	lengths := make([]int, len(array))

	// Set default values for each element in the array.
	for i := range array {
		predecessors[i] = -1
		lengths[i] = 1
	}

	// Iterate through the input array to compute LIS lengths.
	for i := range array {
		currentNum := array[i]
		for j := 0; j < i; j++ {
			otherNum := array[j]
			// If the other number is smaller and the updated LIS length is greater or equal,
			// update the LIS length and predecessor index.
			if otherNum < currentNum && lengths[j]+1 >= lengths[i] {
				lengths[i] = lengths[j] + 1
				predecessors[i] = j
			}
		}
	}

	// Find the index with the maximum LIS length.
	maxLengthIdx := 0
	for i := range array {
		if lengths[i] > lengths[maxLengthIdx] {
			maxLengthIdx = i
		}
	}

	// Build the LIS using the array, predecessors, and the maximum length index.
	return buildLisSequence(array, predecessors, maxLengthIdx)
}
