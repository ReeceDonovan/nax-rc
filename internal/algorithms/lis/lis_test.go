package lis

import (
	"fmt"
	"testing"
)

type TestCase struct {
	array []int
}

type TestFunc struct {
	name string
	f    func([]int) []int
}

/*
Test Case 1:
"array": [5, 7, -24, 12, 10, 2, 3, 12, 5, 6, 35]

Test Case 2:
"array": [-1]

Test Case 3:
"array": [-1, 2]

Test Case 4:
"array": [-1, 2, 1, 2]

Test Case 5:
"array": [1, 5, -1, 10]

Test Case 6:
"array": [1, 5, -1, 0, 6, 2, 4]

Test Case 7:
"array": [3, 4, -1]

Test Case 8:
"array": [29, 2, 32, 12, 30, 31]

Test Case 9:
"array": [10, 22, 9, 33, 21, 61, 41, 60, 80]

Test Case 10:
"array": [100, 1, 2, 3, 4, 101]
*/

var testCases = []TestCase{
	{[]int{5, 7, -24, 12, 10, 2, 3, 12, 5, 6, 35}},
	{[]int{-1}},
	{[]int{-1, 2}},
	{[]int{-1, 2, 1, 2}},
	{[]int{1, 5, -1, 10}},
	{[]int{1, 5, -1, 0, 6, 2, 4}},
	{[]int{3, 4, -1}},
	{[]int{29, 2, 32, 12, 30, 31}},
	{[]int{10, 22, 9, 33, 21, 61, 41, 60, 80}},
	{[]int{100, 1, 2, 3, 4, 101}},
}

var testFunctions = []TestFunc{
	{"Slow", longestIncreasingSubsequenceSlow},
	{"Fast", longestIncreasingSubsequenceFast},
}

func BenchmarkLongestIncreasingSubsequence(b *testing.B) {
	for _, impl := range testFunctions {
		for i, tc := range testCases {
			b.Run(fmt.Sprintf("%s_%d", impl.name, i+1), func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					impl.f(tc.array)
				}
			})
		}
	}
}
