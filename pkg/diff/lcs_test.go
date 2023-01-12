package diff

// Longest Common Subsequence (LCS) algorithm test and benchmark

import (
	"bytes"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	str1 := "ZXVVYZW"
	str2 := "XKYKZPW"
	expected := []byte("XYZW")
	if !bytes.Equal(LongestCommonSubsequence(str1, str2), expected) {
		t.Errorf("Expected %v, got %v", expected, LongestCommonSubsequence(str1, str2))
	}
}
