package djikstra

import "gonum.org/v1/gonum/mat"

// Vertex describes position in multidimensional space
type Vertex struct {
	connections []int
	position    []float64
	key         int
}

// NewVertex creates instance of Vertex implementing vertexer interface
func NewVertex(key int, position []float64, connections []int) *Vertex {
	return &Vertex{connections, position, key}
}

func (v Vertex) GetKey() int {
	return v.key
}

func (v Vertex) GetPosition() []float64 {
	return v.position
}

func (v Vertex) GetDistance(np Positioner) float64 {
	p := np.GetPosition()
	vec1 := mat.NewVecDense(len(v.position), v.position)
	vec2 := mat.NewVecDense(len(p), p)
	rv := new(mat.VecDense)
	rv.SubVec(vec1.TVec(), vec2.TVec())
	return rv.Norm(2)
}

func (v Vertex) GetConnections() []int {
	return v.connections
}
