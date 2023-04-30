package types

import (
	"crypto/rand"
	"log"
)

// Revision represents a single revision of a file in a Version Control System (VCS).
type Revision interface {
	ID() int
	Data() []byte
}

// NewRevision creates a new revision with the given ID and data.
func NewRevision(revisionID int, revisionData []byte) Revision {
	return &standardRevision{revisionID, revisionData}
}

// NewBlankRevision creates a new revision with the given ID and no data.
func NewBlankRevision(revisionID int) Revision {
	return NewRevision(revisionID, nil)
}

// NewRandomRevision creates a new revision with the given ID and random data of the specified size.
func NewRandomRevision(revisionID int, dataSize int) Revision {
	data, err := generateRandomBytes(dataSize)
	if err != nil {
		log.Fatalf("error generating random bytes: %v", err)
	}
	return NewRevision(revisionID, data)
}

// standardRevision is a concrete implementation of the Revision interface.
type standardRevision struct {
	id   int
	data []byte
}

// ID returns the ID of the revision.
func (rev *standardRevision) ID() int {
	return rev.id
}

// Data returns the data of the revision.
func (rev *standardRevision) Data() []byte {
	return rev.data
}

// generateRandomBytes generates a slice of random bytes of the specified length.
func generateRandomBytes(length int) ([]byte, error) {
	buf := make([]byte, length)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
