package djikstra

// PathVertex is a Vertex on the path with a copy of his Parent
type PathVertex struct {
	Parent, Actual Vertexer
}

// Path represents calculated shortest path between to nodes on the graph
type Path struct {
	Vertexes      []*PathVertex
	TotalDistance float64
}
