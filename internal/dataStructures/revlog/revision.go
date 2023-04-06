package revlog

import (
	"crypto/rand"
	"log"
)

type Revision interface {
	ID() int
	Data() []byte
}

func NewRevision(id int, data []byte) Revision {
	return &standardRevision{id, data}
}

func NewBlankRevision(id int) Revision {
	return &standardRevision{id, nil}
}

func NewRandomRevision(id int, size int) Revision {
	return &standardRevision{id, generateRandomBytes(size)}
}

type standardRevision struct {
	id   int
	data []byte
}

func (rev *standardRevision) ID() int {
	return rev.id
}

func (rev *standardRevision) Data() []byte {
	return rev.data
}

func generateRandomBytes(n int) []byte {
	buf := make([]byte, n)
	_, err := rand.Read(buf)
	if err != nil {
		log.Fatalf("error generating random bytes: %v", err)
	}
	return buf
}
