package graph

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ReeceDonovan/nax-rc/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestDAG_AddNode(t *testing.T) {
	dag := NewDAG()
	revision := types.NewBlankRevision(1)
	node, err := dag.AddNode(revision)

	assert.NoError(t, err)
	assert.Equal(t, revision, node.Revision)
	assert.True(t, dag.NodeExists(1))
}

func TestDAG_RemoveNode(t *testing.T) {
	dag := NewDAG()
	revision := types.NewBlankRevision(1)
	dag.AddNode(revision)
	err := dag.RemoveNode(1)

	assert.NoError(t, err)
	assert.False(t, dag.NodeExists(1))
}

func TestDAG_AddEdge(t *testing.T) {
	dag := NewDAG()
	revision1 := types.NewBlankRevision(1)
	revision2 := types.NewBlankRevision(2)
	dag.AddNode(revision1)
	dag.AddNode(revision2)

	err := dag.AddEdge(1, 2)

	assert.NoError(t, err)
	parentNode := dag.GetNode(1)
	childNode := dag.GetNode(2)

	assert.Contains(t, parentNode.Children, childNode)
	assert.Contains(t, childNode.Parents, parentNode)
}

func TestDAG_RemoveEdge(t *testing.T) {
	dag := NewDAG()
	revision1 := types.NewBlankRevision(1)
	revision2 := types.NewBlankRevision(2)
	dag.AddNode(revision1)
	dag.AddNode(revision2)
	dag.AddEdge(1, 2)

	err := dag.RemoveEdge(1, 2)

	assert.NoError(t, err)
	parentNode := dag.GetNode(1)
	childNode := dag.GetNode(2)

	assert.NotContains(t, parentNode.Children, childNode)
	assert.NotContains(t, childNode.Parents, parentNode)
}

func generateRandomRevision() types.Revision {
	revisionID := rand.Intn(1000)
	dataSize := rand.Intn(100)
	return types.NewRandomRevision(revisionID, dataSize)
}

func BenchmarkDAG_AddNode(b *testing.B) {
	dag := NewDAG()

	for i := 0; i < b.N; i++ {
		revision := generateRandomRevision()
		dag.AddNode(revision)
	}
}

func BenchmarkDAG_RemoveNode(b *testing.B) {
	dag := NewDAG()

	for i := 0; i < b.N; i++ {
		revision := generateRandomRevision()
		dag.AddNode(revision)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		revisionID := rand.Intn(b.N)
		dag.RemoveNode(revisionID)
	}
}

func BenchmarkDAG_AddEdge(b *testing.B) {
	dag := NewDAG()

	for i := 0; i < b.N; i++ {
		revision := generateRandomRevision()
		dag.AddNode(revision)
	}

	b.ResetTimer()

	for i := 0; i < b.N-1; i++ {
		dag.AddEdge(i, i+1)
	}
}

func BenchmarkDAG_RemoveEdge(b *testing.B) {
	dag := NewDAG()

	for i := 0; i < b.N; i++ {
		revision := generateRandomRevision()
		dag.AddNode(revision)
	}

	for i := 0; i < b.N-1; i++ {
		dag.AddEdge(i, i+1)
	}

	b.ResetTimer()

	for i := 0; i < b.N-1; i++ {
		dag.RemoveEdge(i, i+1)
	}
}

const (
	smallNumNodes  = 200
	mediumNumNodes = 2000
	largeNumNodes  = 20000
	largeDataSize  = 10000 // 10 Kb
	mediumDataSize = 1000  // 1 Kb
	smallDataSize  = 100   // 100 bytes
)

// Create benchmarks for constructing a DAG of revisions with different sizes
func BenchmarkDAG_RevisionConstruction(b *testing.B) {
	for _, numNodes := range []int{smallNumNodes, mediumNumNodes, largeNumNodes} {
		for _, dataSize := range []int{smallDataSize, mediumDataSize, largeDataSize} {
			b.Run(fmt.Sprintf("numRevisions=%d,dataSize=%d", numNodes, dataSize), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					// Create a DAG
					dag := NewDAG()
					// Add nodes with revisions
					for j := 0; j < numNodes; j++ {
						revision := types.NewRandomRevision(j, dataSize)
						dag.AddNode(revision)
					}

					// Add edges
					for j := 0; j < numNodes-1; j++ {
						dag.AddEdge(j, j+1)
					}
				}
			})
		}
	}
}
