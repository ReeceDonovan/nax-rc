package lis

// binarySearch is a helper function to find the index at which the number should be inserted.
func binarySearch(startIdx, endIdx int, indices []int, array []int, num int) int {
	if startIdx > endIdx {
		return startIdx
	}
	midIdx := (startIdx + endIdx) / 2
	if array[indices[midIdx]] < num {
		startIdx = midIdx + 1
	} else {
		endIdx = midIdx - 1
	}
	return binarySearch(startIdx, endIdx, indices, array, num)
}

// buildLisSequence constructs the LIS using the input array, predecessors, and a starting index.
func buildLisSequence(array, predecessors []int, idx int) []int {
	out := []int{}
	// Iterate through the predecessors array, adding elements to the output slice.
	for idx != -1 {
		out = append(out, array[idx])
		idx = predecessors[idx]
	}

	// Reverse the output slice to obtain the LIS.
	reverseInts(out)
	return out
}

// max returns the maximum of the first argument and the rest of the arguments.
func max(arg int, rest ...int) int {
	for _, num := range rest {
		if num > arg {
			arg = num
		}
	}
	return arg
}

// reverseInts reverses an int slice in place.
func reverseInts(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}
