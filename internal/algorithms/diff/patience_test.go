package diff

import (
	"strings"
	"testing"
)

func TestPatienceDiff(t *testing.T) {
	originalLines := []string{
		"package main",
		"",
		"import (",
		"\t\"fmt\"",
		")",
		"",
		"func main() {",
		"\tprintHelloWorld()",
		"}",
		"",
		"func printHelloWorld() {",
		"\tfmt.Println(\"Hello, world!\")",
		"}",
	}

	modifiedLines := []string{
		"package main",
		"",
		"import (",
		"\t\"fmt\"",
		"\t\"time\"",
		")",
		"",
		"func main() {",
		"\tprintGreeting()",
		"\tprintCurrentTime()",
		"}",
		"",
		"func printGreeting() {",
		"\tfmt.Println(\"Hello, everyone!\")",
		"}",
		"",
		"func printCurrentTime() {",
		"\tcurrentTime := time.Now()",
		"\tfmt.Println(\"Current time:\", currentTime)",
		"}",
	}

	operations := PatienceDiff(originalLines, modifiedLines)

	t.Logf("Operations: %v", operations)
}

func BenchmarkPatienceDiff(b *testing.B) {
	originalLines := []string{
		"package main",
		"",
		"import (",
		"\t\"fmt\"",
		")",
		"",
		"func main() {",
		"\tprintHelloWorld()",
		"}",
		"",
		"func printHelloWorld() {",
		"\tfmt.Println(\"Hello, world!\")",
		"}",
	}

	modifiedLines := []string{
		"package main",
		"",
		"import (",
		"\t\"fmt\"",
		"\t\"time\"",
		")",
		"",
		"func main() {",
		"\tprintGreeting()",
		"\tprintCurrentTime()",
		"}",
		"",
		"func printGreeting() {",
		"\tfmt.Println(\"Hello, everyone!\")",
		"}",
		"",
		"func printCurrentTime() {",
		"\tcurrentTime := time.Now()",
		"\tfmt.Println(\"Current time:\", currentTime)",
		"}",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// operations := PatienceDiff(originalLines, modifiedLines)
		// // Log the number of edits
		// b.Logf("number of edits: %d", len(operations))
		PatienceDiff(originalLines, modifiedLines)
	}
}

func BenchmarkPatienceDiffSmall(b *testing.B) {
	sourceText := "The quick brown fox jumps over the lazy dog"
	destinationText := "The quick red fox jumped over the lazy dog"

	source := strings.Split(sourceText, " ")
	destination := strings.Split(destinationText, " ")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PatienceDiff(source, destination)
	}
}

func BenchmarkPatienceDiffLarge(b *testing.B) {
	sourceText := strings.Repeat("A ", 500) + strings.Repeat("B ", 500) + strings.Repeat("C ", 500)
	destinationText := strings.Repeat("B ", 500) + strings.Repeat("A ", 500) + strings.Repeat("D ", 500)

	source := strings.Split(sourceText, " ")
	destination := strings.Split(destinationText, " ")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PatienceDiff(source, destination)
	}
}

func BenchmarkPatienceDiffExtreme(b *testing.B) {
	sourceText := strings.Repeat("A ", 1000)
	destinationText := strings.Repeat("B ", 1000)

	source := strings.Split(sourceText, " ")
	destination := strings.Split(destinationText, " ")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PatienceDiff(source, destination)
	}
}
