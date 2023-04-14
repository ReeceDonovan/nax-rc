package lcs

// O(nm*min(n,m)) time | O(nm*min(n,m)) space
func longestCommonSubsequenceSlow(str1, str2 string) []byte {
	// Initialize a 3D byte slice to store LCS solutions for subproblems.
	subproblemSolutions := make([][][]byte, len(str2)+1)
	for i := range subproblemSolutions {
		subproblemSolutions[i] = make([][]byte, len(str1)+1)
	}

	// Iterate through the characters of both input strings.
	for i := 1; i < len(subproblemSolutions); i++ {
		for j := 1; j < len(subproblemSolutions[i]); j++ {
			// If the current characters match, add the character to the LCS of the previous subproblem.
			if str2[i-1] == str1[j-1] {
				tmp := make([]byte, len(subproblemSolutions[i-1][j-1]))
				copy(tmp, subproblemSolutions[i-1][j-1])
				subproblemSolutions[i][j] = append(tmp, str2[i-1])
			} else {
				// If the characters don't match, take the longer LCS from the previous subproblems.
				if len(subproblemSolutions[i-1][j]) < len(subproblemSolutions[i][j-1]) {
					subproblemSolutions[i][j] = subproblemSolutions[i][j-1]
				} else {
					subproblemSolutions[i][j] = subproblemSolutions[i-1][j]
				}
			}
		}
	}

	// Return the final LCS solution.
	return subproblemSolutions[len(str2)][len(str1)]
}
