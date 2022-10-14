package djikstra

// Locator locates Vertex on the graph by unique key
type Locator interface {
	// GetKey returns unique Vertex key.
	// Each Vertex has to have unique key.
	GetKey() int
}

// Positioner describes Vertex position in multidimensional space
type Positioner interface {
	// GetPosition returns Vertex position in multidimensional space,
	// for example []int{1.1, 0.5, 2.3} is Vertex position in 3D space
	GetPosition() []float64
}

// Ruler describes distance between vertexes
type Ruler interface {
	// GetDistance returns calculated distance between vertexes based on provided Positioner implementation
	GetDistance(p Positioner) float64
}

// Connector describes Vertex connections
type Connector interface {
	// GetConnections returns slice of keys of all vertexes that given Vertex has connection to,
	// it is possible that connection is one way only
	GetConnections() []int
}

// Vertexer allows to get all Vertex related information needed to graph the shortest path
type Vertexer interface {
	Connector
	Ruler
	Positioner
	Locator
}
