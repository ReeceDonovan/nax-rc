package graph

import (
	"bytes"
	"fmt"
	"sort"
)

// Vertex is a simple interface that represents a vertex in a graph.
type Vertex interface{}

// LabelledVertex is an interface that represents a vertex with a label.
type LabelledVertex interface {
	Vertex
	VertexLabel() string
}

type Graph struct {
	// The set of vertices in the graph
	vertices GenericSet
	// The set of edges in the graph
	edges GenericSet
	// For each vertex, the set of edges that have it as their destination
	incomingEdges map[interface{}]GenericSet
	// For each vertex, the set of edges that have it as their source
	outgoingEdges map[interface{}]GenericSet
}

func (graph *Graph) Vertices() []Vertex {
	vertices := make([]Vertex, 0, len(graph.vertices))
	for _, elem := range graph.vertices {
		vertices = append(vertices, elem.(Vertex))
	}
	return vertices
}

func (graph *Graph) Edges() []Edge {
	edges := make([]Edge, 0, len(graph.edges))
	edgeList := graph.edges.List()
	for _, elem := range edgeList {
		edges = append(edges, elem.(Edge))
	}
	return edges
}

func (graph *Graph) OutgoingEdges(vertex Vertex) []Edge {
	var outEdges []Edge
	sourceVertexHashcode := GetElementHashcode(vertex)
	for _, elem := range graph.Edges() {
		if GetElementHashcode(elem.SourceVertex()) == sourceVertexHashcode {
			outEdges = append(outEdges, elem)
		}
	}
	return outEdges
}

func (graph *Graph) IncomingEdges(vertex Vertex) []Edge {
	var inEdges []Edge
	targetVertexHashcode := GetElementHashcode(vertex)
	for _, elem := range graph.Edges() {
		if GetElementHashcode(elem.TargetVertex()) == targetVertexHashcode {
			inEdges = append(inEdges, elem)
		}
	}
	return inEdges
}

func (graph *Graph) AddVertex(vertex Vertex) Vertex {
	graph.init()
	graph.vertices.Add(vertex)
	return vertex
}

func (graph *Graph) RemoveVertex(vertex Vertex) Vertex {
	graph.vertices.Delete(vertex)
	for _, targetVertex := range graph.outgoingEdgesRaw(vertex) {
		edge := StandardEdge(vertex, targetVertex)
		graph.RemoveEdge(edge)
	}
	for _, sourceVertex := range graph.incomingEdgesRaw(vertex) {
		edge := StandardEdge(sourceVertex, vertex)
		graph.RemoveEdge(edge)
	}
	return nil
}

func (graph *Graph) AddEdge(edge Edge) {
	graph.init()

	sourceVertex := edge.SourceVertex()
	sourceVertexHashcode := GetElementHashcode(sourceVertex)
	targetVertex := edge.TargetVertex()
	targetVertexHashcode := GetElementHashcode(targetVertex)

	// Ensure that the edge is not already in the graph
	if set, ok := graph.outgoingEdges[sourceVertexHashcode]; ok && set.Contains(targetVertex) {
		return
	}

	// Add the edge to the graph
	graph.edges.Add(edge)

	// Add the edge to the outgoing edges of the source vertex
	set, ok := graph.outgoingEdges[sourceVertexHashcode]
	if !ok {
		set = make(GenericSet)
		graph.outgoingEdges[sourceVertexHashcode] = set
	}
	set.Add(targetVertex)

	// Add the edge to the incoming edges of the target vertex
	set, ok = graph.incomingEdges[targetVertexHashcode]
	if !ok {
		set = make(GenericSet)
		graph.incomingEdges[targetVertexHashcode] = set
	}
	set.Add(sourceVertex)
}

func (graph *Graph) RemoveEdge(edge Edge) {
	graph.init()
	graph.edges.Delete(edge)
	if set, ok := graph.outgoingEdges[GetElementHashcode(edge.SourceVertex())]; ok {
		set.Delete(edge.TargetVertex())
	}
	if set, ok := graph.incomingEdges[GetElementHashcode(edge.TargetVertex())]; ok {
		set.Delete(edge.SourceVertex())
	}
}

// =========================================
// Helper functions
// =========================================

func (graph *Graph) init() {
	if graph.vertices == nil {
		graph.vertices = make(GenericSet)
	}
	if graph.edges == nil {
		graph.edges = make(GenericSet)
	}
	if graph.incomingEdges == nil {
		graph.incomingEdges = make(map[interface{}]GenericSet)
	}
	if graph.outgoingEdges == nil {
		graph.outgoingEdges = make(map[interface{}]GenericSet)
	}
}

func (graph *Graph) incomingEdgesRaw(vertex Vertex) GenericSet {
	graph.init()
	return graph.incomingEdges[GetElementHashcode(vertex)]
}

func (graph *Graph) outgoingEdgesRaw(vertex Vertex) GenericSet {
	graph.init()
	return graph.outgoingEdges[GetElementHashcode(vertex)]
}

func (graph *Graph) String() string {
	var buffer bytes.Buffer

	vertices := graph.Vertices()
	labels := make([]string, len(vertices))
	labelMap := make(map[string]Vertex, len(vertices))
	for _, vertex := range vertices {
		label := VertexLabel(vertex)
		labels = append(labels, label)
		labelMap[label] = vertex
	}
	sort.Strings(labels)

	for _, label := range labels {
		vertex := labelMap[label]
		targetVertices := graph.outgoingEdges[GetElementHashcode(vertex)]

		buffer.WriteString(fmt.Sprintf("%s\n", label))

		dependantLabels := make([]string, targetVertices.Length())
		for _, targetVertex := range targetVertices {
			dependantLabels = append(dependantLabels, VertexLabel(targetVertex))
		}
		sort.Strings(dependantLabels)

		for _, dependantLabel := range dependantLabels {
			buffer.WriteString(fmt.Sprintf("\t%s\n", dependantLabel))
		}
	}

	return buffer.String()
}

func VertexLabel(rawVertex Vertex) string {
	switch vertex := rawVertex.(type) {
	case LabelledVertex:
		return vertex.VertexLabel()
	case fmt.Stringer:
		return vertex.String()
	default:
		return fmt.Sprintf("%v", vertex)
	}
}
