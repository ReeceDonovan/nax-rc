package datastructures

import (
	"fmt"
	"strconv"
	"strings"
)

// Rudimentary implementation of a version control system.

// Revision is a single revision of a file.
type Revision struct {
	// The revision id.
	ID int
	// The file contents.
	Data []byte
	// TODO: Consider adding a timestamp.
}

// Revlog is a log of revisions to a file.
type Revlog interface {
	// Add a new revision to the log.
	Add(rev Revision) error
	// Get the revision with the given id.
	Get(rId int) (Revision, error)
	// Remove the revision with the given id.
	Remove(rId int) error
	// Get the number of revisions in the log.
	Count() int
}

// Create a new revision from a given log line.
func ParseRevisionFromLog(line string) (Revision, error) {

	// Format: `<epoch_time_in_microseconds> <revision_id> <content_added> <file_content>`

	// Create variables to hold the parts of the line.
	var revisionID int
	var fileContent string

	// Split the line and assign the parts to the variables.
	parts := strings.Split(line, " ")
	// Check that the line has the correct number of parts.
	if len(parts) != 4 {
		return Revision{}, fmt.Errorf("invalid line format")
	}

	// Assign the parts to the variables.
	revisionID, err := strconv.Atoi(parts[1])
	if err != nil {
		return Revision{}, err
	}
	fileContent = parts[3]

	// TODO: Utilize the content added and epoch time.
	// var epochTime int64
	// var contentAdded string
	// epochTime, err := strconv.ParseInt(parts[0], 10, 64)
	// if err != nil {
	// 	return Revision{}, err
	// }
	// contentAdded = parts[2]

	// Create a new revision.
	revision := Revision{
		ID:   revisionID,
		Data: []byte(fileContent),
	}

	return revision, nil
}
