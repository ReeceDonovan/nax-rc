package lcs

import (
	"fmt"
	"testing"
)

// Create a struct to hold the two input strings for each test case.
type TestCase struct {
	str1 string
	str2 string
}

type TestFunc struct {
	name string
	f    func(string, string) []byte
}

// Create a slice of test cases.
var testCases = []TestCase{
	{
		"ZXVVYZW",
		"XKYKZPW",
	},
	{
		"",
		"",
	},
	{
		"",
		"ABCDEFG",
	},
	{
		"ABCDEFG",
		"",
	},
	{
		"ABCDEFG",
		"ABCDEFG",
	},
	{
		"ABCDEFG",
		"APPLES",
	},
	{
		"8111111111111111142",
		"222222222822222222222222222222433333333332",
	},
	{
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"CCCDDEGDHAGKGLWAJWKJAWGKGWJAKLGGWAFWLFFWAGJWKAG",
	},
	{
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"CCCDDEGDHAGKGLWAJWKJAWGKGWJAKLGGWAFWLFFWAGJWKAGTUV",
	},
}

var testFunctions = []TestFunc{
	{"Slow", longestCommonSubsequenceSlow},
	{"Fast", longestCommonSubsequenceFast},
}

func BenchmarkLongestCommonSubsequence(b *testing.B) {
	for _, impl := range testFunctions {
		for i, tc := range testCases {
			b.Run(fmt.Sprintf("%s_%d", impl.name, i+1), func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					impl.f(tc.str1, tc.str2)
				}
			})
		}
	}
}
