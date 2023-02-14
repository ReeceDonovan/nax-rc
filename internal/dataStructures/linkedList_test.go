package datastructures

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ReeceDonovan/nax-rc/pkg/ioutil"
)

// Create helper functions for reading in the file and parsing the lines
func getLinesFromFile(file string) ([]string, error) {
	// Get the content of the file
	content, err := ioutil.GetFileContent(file)
	if err != nil {
		return nil, err
	}

	// Split the content into lines
	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")

	// Remove the last line if it is empty
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	// Check that the file is not empty
	if len(lines) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	return lines, nil
}

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

func TestCreatingLinkedListFromRevlog(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList()

	// Set directory path
	dirPath := "../../test_data"

	// Get the files in the directory
	files, err := getFilesFromDirectory(dirPath)
	if err != nil {
		t.Errorf("error getting files from directory: %v", err)
	}

	// Check that the directory is not empty
	if len(files) == 0 {
		t.Errorf("directory is empty")
	}

	// Use the first file in the directory for testing
	file := files[0]

	// Get the lines from the file
	lines, err := getLinesFromFile(file)
	if err != nil {
		t.Errorf("error getting lines from file: %v", err)
	}

	// Log the number of lines in the file
	t.Logf("number of lines in file: %v", len(lines))

	// Set the first line as the head of the linked list
	ll.SetHead(&Node{
		Data: lines[0],
	})

	// Iterate over the remaining lines and insert the revision data into the linked list as a node
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			t.Errorf("line is empty")
		}

		// Log the progress of the test
		t.Logf("inserting node %v of %v", i, len(lines))
		fmt.Printf("inserting node %v of %v", i, len(lines))

		// Create a new node with the revision data
		node := &Node{
			Data: lines[i],
		}

		// Insert the node into the linked list
		ll.InsertAfter(ll.Tail, node)

		// Log the progress of the test
		t.Logf("inserted node %v of %v", i, len(lines))
	}

	// Iterate over the linked list and compare the data to the file
	node := ll.Head
	for i := 0; i < len(lines); i++ {
		// Log the progress of the test
		t.Logf("comparing node %v of %v", i, len(lines))

		if node == nil {
			t.Errorf("node is nil")
		}

		if node.Data == nil {
			t.Errorf("data is nil")
		}

		if node.Data == "" {
			t.Errorf("data is empty")
		}

		if node.Data != lines[i] {
			t.Errorf("data does not match")
		}

		if node.Next == nil && i != len(lines)-1 {
			t.Errorf("next node is nil but there are more lines")
		}

		node = node.Next

		// Log the progress of the test
		t.Logf("compared node %v of %v", i, len(lines))
	}
}

// Benchmark the linked list data structure for basic version control operations.
func BenchmarkCreatingLinkedListFromRevlog(b *testing.B) {
	// Create a new linked list
	ll := NewLinkedList()

	// Set directory path
	dirPath := "../../test_data"

	// Get the files in the directory
	files, err := getFilesFromDirectory(dirPath)
	if err != nil {
		b.Errorf("error getting files from directory: %v", err)
	}

	// Check that the directory is not empty
	if len(files) == 0 {
		b.Errorf("directory is empty")
	}

	// Use the first file in the directory for testing
	file := files[0]

	// Get the lines from the file
	lines, err := getLinesFromFile(file)
	if err != nil {
		b.Errorf("error getting lines from file: %v", err)
	}

	// Set the first line as the head of the linked list
	ll.SetHead(&Node{
		Data: lines[0],
	})

	// Iterate over the remaining lines and insert the revision data into the linked list as a node
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			b.Errorf("line is empty")
		}

		// Create a new node with the revision data
		node := &Node{
			Data: lines[i],
		}

		// Insert the node into the linked list
		ll.InsertAfter(ll.Tail, node)
	}
}
