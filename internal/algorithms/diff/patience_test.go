package diff

import (
	"testing"

	"github.com/ReeceDonovan/nax-rc/pkg/ioutil"
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

func BenchmarkPatienceDiffFromFile(b *testing.B) {
	dirPath := "../../../test_data"
	files, err := getFilesFromDirectory(dirPath)
	if err != nil {
		b.Errorf("error getting files from directory: %v", err)
	}

	if len(files) < 2 {
		b.Errorf("expected at least 2 files in directory")
	}

	source, err := ioutil.GetFileContentLines(files[0])
	if err != nil {
		b.Errorf("error getting file content: %v", err)
	}

	destination, err := ioutil.GetFileContentLines(files[1])
	if err != nil {
		b.Errorf("error getting file content: %v", err)
	}

	// Remove the last line if it is empty
	if len(source) > 0 && source[len(source)-1] == "" {
		source = source[:len(source)-1]
	}

	if len(destination) > 0 && destination[len(destination)-1] == "" {
		destination = destination[:len(destination)-1]
	}

	// Log number of lines in each file
	b.Logf("source file has %d lines", len(source))
	b.Logf("destination file has %d lines", len(destination))

	// Run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		operations := PatienceDiff(source, destination)

		// Log the number of edits
		b.Logf("number of edits: %d", len(operations))
	}
}
