package datastructures

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ReeceDonovan/nax-rc/pkg/ioutil"
)

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

func testRevlogAPI(revlog Revlog, t *testing.T) {
	if revlog.Count() != 0 {
		t.Errorf("expected count to be 0")
	}
	for i := 1; i <= 100; i++ {
		revlog.Add(Revision{ID: i, Data: []byte(fmt.Sprintf("Revision %d", i))})
	}
	if revlog.Count() != 100 {
		t.Errorf("expected count to be 100")
	}
	for i := 1; i <= 50; i++ {
		revlog.Remove(i)
	}
	if revlog.Count() != 50 {
		t.Errorf("expected count to be 50")
	}
}

func TestDoublyLinkedList(t *testing.T) {
	var revlog Revlog = NewDoublyLinkedList()
	testRevlogAPI(revlog, t)
}

// ---------------------------------------------

// Test creating a revlog from a log file
func testRevlogFromFile(revlog Revlog, t *testing.T) {
	// Set directory path
	dirPath := "../../test_data"

	// Get the files in the directory
	files, err := getFilesFromDirectory(dirPath)
	if err != nil {
		t.Errorf("error getting files from directory: %v", err)
	}

	// TODO: Fix this to use all files in the directory
	// Use the first file in the directory for testing
	file := files[0]

	// Get the lines from the file
	lines, err := ioutil.GetFileContentLines(file)
	if err != nil {
		t.Errorf("error getting lines from file: %v", err)
	}

	// Remove the last line if it is empty
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	// Check that the file is not empty
	if len(lines) == 0 {
		t.Errorf("file is empty")
	}

	// Log the number of lines in the file
	t.Logf("number of lines in file: %d", len(lines))

	// Iterate through the lines and add them to the revlog
	for i, line := range lines {
		// Check that the line is not empty
		if line == "" {
			t.Errorf("line is empty")
		}

		// Log the progress
		t.Logf("adding item %d to revlog", i)

		// Parse the revision from the line
		revision, err := ParseRevisionFromLog(line)
		if err != nil {
			t.Errorf("error parsing revision from line: %v", err)
		}

		// Add the revision to the revlog
		err = revlog.Add(revision)
		if err != nil {
			t.Errorf("error adding revision to revlog: %v", err)
		}
	}

	// Check that the revlog has the correct number of revisions
	if revlog.Count() != len(lines) {
		t.Errorf("expected revlog to have %d revisions, got %d", len(lines), revlog.Count())
	}
}

func TestDoublyLinkedListRevlog(t *testing.T) {
	var revlog Revlog = NewDoublyLinkedList()
	testRevlogFromFile(revlog, t)
}

// ---------------------------------------------

// Benchmark finding a specific revision in a revlog using the revision id
func benchmarkFindRevision(revlog Revlog, rId int, b *testing.B) {
	r, err := revlog.Get(rId)
	if err != nil {
		b.Errorf("error finding revision: %v", err)
	}
	b.Logf("found revision: %v", r)
}

// func BenchmarkDoublyLinkedListFindRevision(b *testing.B) {
// 	// b.N = 500
// 	for n := 0; n < b.N; n++ {
// 		// set num revisions to a random number between 100000 and 250000
// 		var numRevisions = rand.Intn(250000) + 250000
// 		// set target revision to a random number between 45% and 85% of numRevisions
// 		var targetRevision = rand.Intn(int(float64(numRevisions)*0.4)) + int(float64(numRevisions)*0.45)

// 		// b.Logf("number of revisions: %d", numRevisions)
// 		// b.Logf("target revision: %d", targetRevision)

// 		var revlog Revlog = NewDoublyLinkedList()
// 		for i := 0; i <= numRevisions; i++ {
// 			revlog.Add(Revision{ID: i, Data: []byte(fmt.Sprintf("Revision %d", i))})
// 		}

// 		b.ResetTimer()
// 		benchmarkFindRevision(revlog, targetRevision, b)
// 	}
// }

func BenchmarkDoublyLinkedListFindRevision(b *testing.B) {
	var numRevisions = 500000
	var revlog Revlog = NewDoublyLinkedList()
	for i := 0; i <= numRevisions; i++ {
		revlog.Add(Revision{ID: i, Data: []byte(fmt.Sprintf("Revision %d", i))})
	}
	for n := 0; n < b.N; n++ {
		// set target revision to a random number between 55% and 85% of numRevisions
		var targetRevision = rand.Intn(int(float64(numRevisions)*0.3)) + int(float64(numRevisions)*0.55)

		// b.Logf("number of revisions: %d", numRevisions)
		b.Logf("target revision: %d", targetRevision)

		// b.ResetTimer()
		benchmarkFindRevision(revlog, targetRevision, b)
	}
}
