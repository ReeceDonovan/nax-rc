package diff

import (
	"fmt"
	"testing"

	"github.com/ReeceDonovan/nax-rc/pkg/ioutil"
)

// func TestMyersDiff(t *testing.T) {
// 	// 	a := "kitten"
// 	// 	b := "sitting"

// 	src := []string{"k", "i", "t", "t", "e", "n"}
// 	dst := []string{"s", "i", "t", "t", "i", "n", "g"}

// 	// src := []string{"kitten", "sitting"}
// 	// dst := []string{"sitting", "kitten"}

// 	// src := strings.Split("kitten", "")
// 	// dst := strings.Split("sitting", "")

// 	// src := []string{"1681286115299102 10 C CHJBABAFAC", "1681286115299489 11 G CHJBABAFACG"}
// 	// dst := []string{"1681286115334532 13 G JGJAFFBGHAEAG", "1681286115334895 14 H JGJAFFBGHAEAGH", "1681286115348568 41 G JGJAFFBGHAEAGHEIHEFBAJFHGJDFHEIJCGAJFDBCG"}

// 	operations := MyersDiff(src, dst)

// 	t.Log(operations)

// 	// stringifiedOperations := stringifyDiff(src, dst, operations)

// 	// t.Log(stringifiedOperations)
// }

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

// Create helper functions for reading in the file and parsing the lines
func getFilesFromDirectory(dirPath string) ([]string, error) {
	// Get the files in the directory
	files, err := ioutil.GetFilesInDirectory(dirPath)
	if err != nil {
		return nil, err
	}

	// Check that the directory is not empty
	if len(files) == 0 {
		return nil, fmt.Errorf("directory is empty")
	}

	return files, nil
}

// BenchmarkMyersDiff benchmarks the MyersDiff function.
func BenchmarkMyersDiffFromFile(b *testing.B) {
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
		edits := MyersDiff(source, destination)

		// Log the number of edits
		b.Logf("number of edits: %d", len(edits))
	}
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
