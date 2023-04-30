package diff

import (
	"strings"
	"testing"
)

func TestMyersDiff(t *testing.T) {
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

	operations := MyersDiff(originalLines, modifiedLines)

	t.Logf("Operations: %v", operations)
}

func BenchmarkMyersDiff(b *testing.B) {
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
		// edits := MyersDiff(originalLines, modifiedLines)
		// // Log the number of edits
		// b.Logf("number of edits: %d", len(edits))
		MyersDiff(originalLines, modifiedLines)
	}
}

func BenchmarkMyersDiffSmall(b *testing.B) {
	sourceText := "The quick brown fox jumps over the lazy dog"
	destinationText := "The quick red fox jumped over the lazy dog"

	source := strings.Split(sourceText, " ")
	destination := strings.Split(destinationText, " ")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MyersDiff(source, destination)
	}
}

func BenchmarkMyersDiffLarge(b *testing.B) {
	sourceText := strings.Repeat("A ", 500) + strings.Repeat("B ", 500) + strings.Repeat("C ", 500)
	destinationText := strings.Repeat("B ", 500) + strings.Repeat("A ", 500) + strings.Repeat("D ", 500)

	source := strings.Split(sourceText, " ")
	destination := strings.Split(destinationText, " ")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MyersDiff(source, destination)
	}
}

func BenchmarkMyersDiffExtreme(b *testing.B) {
	sourceText := strings.Repeat("A ", 1000)
	destinationText := strings.Repeat("B ", 1000)

	source := strings.Split(sourceText, " ")
	destination := strings.Split(destinationText, " ")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MyersDiff(source, destination)
	}
}
