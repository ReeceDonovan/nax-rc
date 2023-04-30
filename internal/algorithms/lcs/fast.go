package lcs

// O(nm) time | O(nm) space
func longestCommonSubsequenceFast(str1, str2 string) []byte {
	// Initialize a 2D int slice to store lengths of LCS for subproblems.
	lengths := make([][]int, len(str2)+1)
	for i := range lengths {
		lengths[i] = make([]int, len(str1)+1)
	}

	// Iterate through the characters of both input strings.
	for i := 1; i < len(str2)+1; i++ {
		for j := 1; j < len(str1)+1; j++ {
			// If the current characters match, increment the length of the LCS of the previous subproblem.
			if str2[i-1] == str1[j-1] {
				lengths[i][j] = lengths[i-1][j-1] + 1
			} else {
				// If the characters don't match, take the maximum LCS length from the previous subproblems.
				lengths[i][j] = max(lengths[i-1][j], lengths[i][j-1])
			}
		}
	}

	// Build the LCS sequence from the lengths matrix.
	return buildLcsSequence(lengths, str1)
}

// buildLcsSequence constructs the LCS sequence using the lengths matrix and the first input string.
func buildLcsSequence(lengths [][]int, str1 string) []byte {
	sequence := make([]byte, 0)
	i, j := len(lengths)-1, len(lengths[0])-1
	for i != 0 && j != 0 {
		if lengths[i][j] == lengths[i-1][j] {
			i--
		} else if lengths[i][j] == lengths[i][j-1] {
			j--
		} else {
			sequence = append(sequence, str1[j-1])
			i--
			j--
		}
	}

	// Reverse the sequence to obtain the LCS.
	return reverseBytes(sequence)
}

// reverseBytes reverses a byte slice in place.
func reverseBytes(sequence []byte) []byte {
	for i, j := 0, len(sequence)-1; i < j; i, j = i+1, j-1 {
		sequence[i], sequence[j] = sequence[j], sequence[i]
	}
	return sequence
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
