package graph

// Edge is an interface that represents a directed edge connecting two vertices in a graph.
type Edge interface {
	// SourceVertex returns the vertex at the start of the edge.
	SourceVertex() Vertex
	// TargetVertex returns the vertex at the end of the edge.
	TargetVertex() Vertex
	// HashableElement is an interface that represents any type that can be hashed.
	HashableElement
}

// StandardEdge creates a new standard edge connecting two vertices.
func StandardEdge(source, target Vertex) Edge {
	return &standardEdge{Source: source, Target: target}
}

// standardEdge is a simple implementation of the Edge interface.
type standardEdge struct {
	Source, Target Vertex
}

// Hashcode returns a unique hashcode for the edge.
func (edge *standardEdge) Hashcode() interface{} {
	return [...]interface{}{edge.Source, edge.Target}
}

// SourceVertex returns the source vertex of the edge.
func (edge *standardEdge) SourceVertex() Vertex {
	return edge.Source
}

// TargetVertex returns the target vertex of the edge.
func (edge *standardEdge) TargetVertex() Vertex {
	return edge.Target
}

type Vertex interface{}
