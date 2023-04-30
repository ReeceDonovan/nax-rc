package types

import (
	"bytes"
	"testing"
)

func TestNewRevision(t *testing.T) {
	id := 1
	data := []byte("test data")
	rev := NewRevision(id, data)

	if rev.ID() != id {
		t.Errorf("expected ID to be %d, got %d", id, rev.ID())
	}

	if !bytes.Equal(rev.Data(), data) {
		t.Errorf("expected Data to be %v, got %v", data, rev.Data())
	}
}

func TestNewBlankRevision(t *testing.T) {
	id := 2
	rev := NewBlankRevision(id)

	if rev.ID() != id {
		t.Errorf("expected ID to be %d, got %d", id, rev.ID())
	}

	if rev.Data() != nil {
		t.Errorf("expected Data to be nil, got %v", rev.Data())
	}
}

func TestNewRandomRevision(t *testing.T) {
	id := 3
	dataSize := 10
	rev := NewRandomRevision(id, dataSize)

	if rev.ID() != id {
		t.Errorf("expected ID to be %d, got %d", id, rev.ID())
	}

	if len(rev.Data()) != dataSize {
		t.Errorf("expected Data to have length %d, got %d", dataSize, len(rev.Data()))
	}
}

func BenchmarkNewRevision(b *testing.B) {
	id := 1
	data := []byte("test data")

	for i := 0; i < b.N; i++ {
		NewRevision(id, data)
	}
}

func BenchmarkNewBlankRevision(b *testing.B) {
	id := 2

	for i := 0; i < b.N; i++ {
		NewBlankRevision(id)
	}
}

func BenchmarkNewRandomRevision(b *testing.B) {
	id := 3
	dataSize := 10

	for i := 0; i < b.N; i++ {
		NewRandomRevision(id, dataSize)
	}
}

func BenchmarkGenerateRandomBytes(b *testing.B) {
	dataSize := 100

	for i := 0; i < b.N; i++ {
		_, err := generateRandomBytes(dataSize)
		if err != nil {
			b.Fatal(err)
		}
	}
}
