package diff

// Longest Common Subsequence (LCS) algorithm implementation.

func LongestCommonSubsequence(str1, str2 string) []byte {
	// Create the matrix of size (len(str2)+1) x (len(str1)+1)
	matrix := make([][]int, len(str2)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(str1)+1)
	}

	// Fill the matrix
	for i := 1; i < len(str2)+1; i++ {
		for j := 1; j < len(str1)+1; j++ {
			if str2[i-1] == str1[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				matrix[i][j] = max(matrix[i-1][j], matrix[i][j-1])
			}
		}
	}

	return buildSequence(matrix, str1)
}

func buildSequence(matrix [][]int, str1 string) []byte {
	// Build the sequence
	sequence := make([]byte, 0)
	i := len(matrix) - 1
	j := len(matrix[0]) - 1
	for i != 0 && j != 0 {
		if matrix[i][j] == matrix[i-1][j] {
			i -= 1
		} else if matrix[i][j] == matrix[i][j-1] {
			j -= 1
		} else {
			sequence = append(sequence, str1[j-1])
			i -= 1
			j -= 1
		}
	}

	return reverse(sequence)
}

func reverse(sequence []byte) []byte {
	for i, j := 0, len(sequence)-1; i < j; i, j = i+1, j-1 {
		sequence[i], sequence[j] = sequence[j], sequence[i]
	}
	return sequence
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
