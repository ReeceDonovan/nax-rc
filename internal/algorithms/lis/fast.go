package lis

// O(n log n) time | O(n) space
func longestIncreasingSubsequenceFast(array []int) []int {
	// Initialize slices to store predecessor indices and LIS indices.
	predecessors := make([]int, len(array))
	indices := make([]int, len(array)+1)

	// Set default values for each element in the array.
	for i := range array {
		predecessors[i] = -1
		indices[i] = -1
	}

	length := 0
	for i, num := range array {
		// Perform binary search to find the new length.
		newLength := binarySearch(1, length, indices, array, num)
		// Update the predecessors and indices arrays.
		predecessors[i] = indices[newLength-1]
		indices[newLength] = i
		// Update the LIS length.
		length = max(length, newLength)
	}

	// Build the LIS using the array, predecessors, and the maximum length index.
	return buildLisSequence(array, predecessors, indices[length])
}
