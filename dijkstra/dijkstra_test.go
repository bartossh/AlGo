package djikstra

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
)

func TestDistanceCalculation(t *testing.T) {
	// given
	cases := []struct {
		title    string
		points   [2][]float64
		distance float64
	}{
		{
			title:    "2D distance calculation 1",
			points:   [2][]float64{{1, 1}, {2, 2}},
			distance: 1.4142135623730951,
		},
		{
			title:    "2D distance calculation 2",
			points:   [2][]float64{{0, 0}, {1, 1}},
			distance: 1.4142135623730951,
		},
		{
			title:    "2D distance calculation 3",
			points:   [2][]float64{{0, 1}, {0, 2}},
			distance: 1,
		},
		{
			title:    "2D distance calculation 4",
			points:   [2][]float64{{1, 0}, {0, 2}},
			distance: 2.23606797749979,
		},
		{
			title:    "2D distance calculation 5",
			points:   [2][]float64{{1, 0}, {1, 0}},
			distance: 0,
		},
		{
			title:    "2D distance calculation 6",
			points:   [2][]float64{{1, 0}, {0, 1}},
			distance: 1.4142135623730951,
		},
		{
			title:    "2D distance calculation 7",
			points:   [2][]float64{{3, 0}, {0, 4}},
			distance: 5,
		},
		{
			title:    "2D distance calculation 8",
			points:   [2][]float64{{0, 8}, {15, 0}},
			distance: 17,
		},
		{
			title:    "3D distance calculation 1",
			points:   [2][]float64{{0, 0, 0}, {1, 0, 1}},
			distance: 1.4142135623730951,
		},
		{
			title:    "3D distance calculation 2",
			points:   [2][]float64{{0, 0, 0}, {1, 1, 1}},
			distance: 1.7320508075688772,
		},
		{
			title:    "3D distance calculation 3",
			points:   [2][]float64{{0, 0, 0}, {0, 0, 1}},
			distance: 1,
		},
		{
			title:    "3D distance calculation 4",
			points:   [2][]float64{{0, 0, 0}, {10, 10, 10}},
			distance: 17.32050807568877,
		},
		{
			title:    "3D distance calculation 5",
			points:   [2][]float64{{2, 3, 6}, {0, 0, 0}},
			distance: 7,
		},
		{
			title:    "4D distance calculation 1",
			points:   [2][]float64{{2, 3, 6, 0}, {0, 0, 0, 0}},
			distance: 7,
		},
		{
			title:    "4D distance calculation 2",
			points:   [2][]float64{{3, 4, 0, 0}, {0, 0, 0, 0}},
			distance: 5,
		},
		{
			title:    "4D distance calculation 3",
			points:   [2][]float64{{1, 1, 1, 1}, {2, 2, 2, 2}},
			distance: 2,
		},
	}

	asr := assert.New(t)

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// when
			p1 := NewVertex(1, c.points[0], []int{})
			p2 := NewVertex(2, c.points[1], []int{})

			// then
			dist := p1.GetDistance(p2)
			asr.Equal(c.distance, dist, fmt.Sprintf("Expected %v, got %v", c.distance, dist))
		})
	}
}

func TestNodesCreation(t *testing.T) {
	type connection struct {
		pos  []float64
		conn []int
	}

	// given
	cases := []struct {
		title string
		graph []connection
	}{
		{
			title: "small graph",
			graph: []connection{
				{
					pos:  []float64{0, 0, 0},
					conn: []int{1},
				},
				{
					pos:  []float64{0, 1, 0},
					conn: []int{0, 2},
				},
				{
					pos:  []float64{0, 0, 1},
					conn: []int{1, 3},
				},
				{
					pos:  []float64{1, 1, 0},
					conn: []int{2},
				},
			},
		},
		{
			title: "big graph",
			graph: []connection{
				{
					pos:  []float64{0, 0, 0},
					conn: []int{1},
				},
				{
					pos:  []float64{0, 1, 0},
					conn: []int{0, 2},
				},
				{
					pos:  []float64{0, 0, 1},
					conn: []int{1, 3},
				},
				{
					pos:  []float64{1, 1, 0},
					conn: []int{2, 4},
				},
				{
					pos:  []float64{2, 0, 0},
					conn: []int{3, 5},
				},
				{
					pos:  []float64{0, 2, 0},
					conn: []int{4, 6},
				},
				{
					pos:  []float64{0, 2, 1},
					conn: []int{5, 7},
				},
				{
					pos:  []float64{1, 2, 0},
					conn: []int{6, 8},
				},
				{
					pos:  []float64{3, 0, 0},
					conn: []int{7, 9},
				},
				{
					pos:  []float64{0, 3, 0},
					conn: []int{8, 10},
				},
				{
					pos:  []float64{0, 3, 1},
					conn: []int{9, 11},
				},
				{
					pos:  []float64{1, 3, 0},
					conn: []int{10},
				},
			},
		},
		{
			title: "complex graph",
			graph: []connection{
				{
					pos:  []float64{0, 0, 0},
					conn: []int{1, 10},
				},
				{
					pos:  []float64{0, 1, 0},
					conn: []int{0, 2, 11},
				},
				{
					pos:  []float64{0, 0, 1},
					conn: []int{1, 3, 5},
				},
				{
					pos:  []float64{1, 1, 0},
					conn: []int{2, 4, 7},
				},
				{
					pos:  []float64{2, 0, 0},
					conn: []int{3, 5, 9},
				},
				{
					pos:  []float64{0, 2, 0},
					conn: []int{4, 6, 5, 8},
				},
				{
					pos:  []float64{0, 2, 1},
					conn: []int{5, 7, 1, 11, 9, 4},
				},
				{
					pos:  []float64{1, 2, 0},
					conn: []int{6, 8, 9, 10},
				},
				{
					pos:  []float64{3, 0, 0},
					conn: []int{7, 9, 8, 6, 4, 2},
				},
				{
					pos:  []float64{0, 3, 0},
					conn: []int{8, 10, 9, 7, 6, 5, 4, 3, 2},
				},
				{
					pos:  []float64{0, 3, 1},
					conn: []int{9, 11, 0},
				},
				{
					pos:  []float64{1, 3, 0},
					conn: []int{10},
				},
			},
		},
	}

	asr := assert.New(t)
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// when
			nodes := make([]*node, 0, len(c.graph))
			for i, g := range c.graph {
				ps := NewVertex(i, g.pos, []int{})
				n := newStartNode(ps)
				nodes = append(nodes, n)
			}
			for i, g := range c.graph {
				for _, idx := range g.conn {
					nodes[i].neighbours[nodes[idx]] = struct{}{}
				}
			}
			// then
			for i, n := range nodes {
				poss := make([]int, 0)
				for nn := range n.neighbours {
					poss = append(poss, nn.vertex.GetKey())
				}
				sort.Ints(poss)
				conn := c.graph[i].conn
				sort.Ints(conn)
				asr.Equal(conn, poss, "should have the same number od nodes as neighbours", conn, poss)
			}
		})
	}
}

func TestResultGraph(t *testing.T) {
	cases := []struct {
		title    string
		vertexes []Vertexer
		result   Path
		st, fn   Vertexer
	}{
		{
			title: "small simple one way graph",
			vertexes: []Vertexer{
				NewVertex(5, []float64{5, 0}, []int{4}),
				NewVertex(0, []float64{0, 0}, []int{1}),
				NewVertex(1, []float64{1, 0}, []int{0, 2}),
				NewVertex(2, []float64{2, 0}, []int{1, 3}),
				NewVertex(4, []float64{4, 0}, []int{3, 5}),
				NewVertex(3, []float64{3, 0}, []int{2, 4}),
			},
			result: Path{TotalDistance: 3},
			st:     NewVertex(0, []float64{0, 0}, []int{1}),
			fn:     NewVertex(3, []float64{3, 0}, []int{2, 4}),
		},
		{
			title: "small circular graph",
			vertexes: []Vertexer{
				NewVertex(5, []float64{0, 1}, []int{4, 0}),
				NewVertex(0, []float64{0, 0}, []int{1, 5}),
				NewVertex(1, []float64{1, 0}, []int{0, 2}),
				NewVertex(2, []float64{2, 0}, []int{1, 3}),
				NewVertex(4, []float64{4, 0}, []int{3, 5}),
				NewVertex(3, []float64{3, 0}, []int{2, 4}),
			},
			result: Path{TotalDistance: 1},
			st:     NewVertex(0, []float64{0, 0}, []int{1, 5}),
			fn:     NewVertex(5, []float64{0, 1}, []int{4, 0}),
		},
		{
			title: "large graph",
			vertexes: []Vertexer{
				NewVertex(0, []float64{0, 0}, []int{1, 5, 11}),
				NewVertex(1, []float64{10, 0}, []int{0, 2}),
				NewVertex(2, []float64{20, 0}, []int{1, 3}),
				NewVertex(3, []float64{30, 0}, []int{2, 4}),
				NewVertex(4, []float64{40, 0}, []int{3, 5}),
				NewVertex(5, []float64{5, 0}, []int{4, 6, 11, 0}),
				NewVertex(6, []float64{60, 0}, []int{5, 7}),
				NewVertex(7, []float64{70, 0}, []int{6, 8}),
				NewVertex(8, []float64{80, 0}, []int{7, 9}),
				NewVertex(9, []float64{90, 0}, []int{7, 10, 12, 14}),
				NewVertex(10, []float64{5, 5}, []int{9, 14, 11}),
				NewVertex(11, []float64{0, 5}, []int{0, 5, 12, 10}),
				NewVertex(12, []float64{0, 20}, []int{11, 9}),
				NewVertex(13, []float64{0, 30}, []int{12, 14}),
				NewVertex(14, []float64{0, 40}, []int{13, 9, 10}),
			},
			result: Path{TotalDistance: 10},
			st:     NewVertex(0, []float64{0, 0}, []int{1, 5, 11}),
			fn:     NewVertex(10, []float64{5, 5}, []int{9, 14, 11}),
		},
	}

	asr := assert.New(t)

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			g := NewGraph(c.vertexes)
			result, err := g.CalculateResultGraph(c.st, c.fn)
			asr.Equal(nil, err, "error should be nil")
			asr.Equal(c.result.TotalDistance, result.TotalDistance, "total distance isn't correct")
		})
	}
}

var vertexes = []Vertexer{
	NewVertex(0, []float64{0, 0}, []int{1, 5, 11}),
	NewVertex(1, []float64{10, 0}, []int{0, 2}),
	NewVertex(2, []float64{20, 0}, []int{1, 3}),
	NewVertex(3, []float64{30, 0}, []int{2, 4}),
	NewVertex(4, []float64{40, 0}, []int{3, 5}),
	NewVertex(5, []float64{5, 0}, []int{4, 6, 11, 0}),
	NewVertex(6, []float64{60, 0}, []int{5, 7}),
	NewVertex(7, []float64{70, 0}, []int{6, 8}),
	NewVertex(8, []float64{80, 0}, []int{7, 9}),
	NewVertex(9, []float64{90, 0}, []int{7, 10, 12, 14}),
	NewVertex(10, []float64{5, 5}, []int{9, 14, 11}),
	NewVertex(11, []float64{0, 5}, []int{0, 5, 12, 10}),
	NewVertex(12, []float64{0, 20}, []int{11, 9}),
	NewVertex(13, []float64{0, 30}, []int{12, 14}),
	NewVertex(14, []float64{0, 40}, []int{13, 9, 10}),
}

var benchcases = []struct {
	st, fn Vertexer
	dst    float64
}{
	{
		st:  NewVertex(0, []float64{0, 0}, []int{1, 5, 11}),
		fn:  NewVertex(10, []float64{5, 5}, []int{9, 14, 11}),
		dst: 10,
	},
	{
		st:  NewVertex(0, []float64{0, 0}, []int{1, 5, 11}),
		fn:  NewVertex(14, []float64{0, 40}, []int{13, 9, 10}),
		dst: 45.35533905932738,
	},
	{
		st:  NewVertex(0, []float64{0, 0}, []int{1, 5, 11}),
		fn:  NewVertex(7, []float64{70, 0}, []int{6, 8}),
		dst: 70,
	},
	{
		st:  NewVertex(0, []float64{0, 0}, []int{1, 5, 11}),
		fn:  NewVertex(13, []float64{0, 30}, []int{12, 14}),
		dst: 55.35533905932738,
	},
	{
		st:  NewVertex(0, []float64{0, 0}, []int{1, 5, 11}),
		fn:  NewVertex(11, []float64{0, 5}, []int{0, 5, 12, 10}),
		dst: 5,
	},
	{
		st:  NewVertex(0, []float64{0, 0}, []int{1, 5, 11}),
		fn:  NewVertex(12, []float64{0, 20}, []int{11, 9}),
		dst: 20,
	},
	{
		st:  NewVertex(0, []float64{0, 0}, []int{1, 5, 11}),
		fn:  NewVertex(6, []float64{60, 0}, []int{5, 7}),
		dst: 60,
	},
}

func BenchmarkGraphSmallPredefined(b *testing.B) {
	asr := assert.New(b)
	g := NewGraph(vertexes)
	for _, c := range benchcases {
		b.Run(fmt.Sprintf("from %v to %v", c.st.GetKey(), c.fn.GetKey()), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, err := g.CalculateResultGraph(c.st, c.fn)
				asr.Equal(err, nil, "error should be nil")
			}
		})
	}
}

func performanceTestVertexesCreator(numOfNodes, connPerNode int) []Vertexer {
	vertexes := make([]Vertexer, 0, numOfNodes)

	for i := 0; i < numOfNodes; i++ {
		connections := make([]int, 0, connPerNode+2)
		if i < numOfNodes-1 {
			connections = append(connections, i+1)
		}
		if i > 0 {
			connections = append(connections, i-1)
			for j := connPerNode; j > 0; j-- {
				r := rand.Intn(numOfNodes - 2)
				connections = append(connections, r+1)
			}
		}
		y := float64(rand.Intn(numOfNodes - 1))
		z := float64(rand.Intn(numOfNodes - 1))
		n := NewVertex(i, []float64{float64(i), y, z}, connections)
		vertexes = append(vertexes, n)
	}
	return vertexes
}

func BenchmarkGraphHugeGenerated(b *testing.B) {
	cases := []struct {
		numOfNodes, connPerNode int
	}{
		{1000, 10},
		{1000, 100},
		{10_000, 10},
		{10_000, 100},
		{10_000, 1000},
		{10_000, 5_000},
		{25_000, 1000},
	}
	asr := assert.New(b)
	for _, c := range cases {
		vertexes := performanceTestVertexesCreator(c.numOfNodes, c.connPerNode)
		b.Run(fmt.Sprintf("number of nodes %v, number off connections per node %v", c.numOfNodes, c.connPerNode), func(b *testing.B) {
			g := NewGraph(vertexes)
			for n := 0; n < b.N; n++ {
				_, err := g.CalculateResultGraph(vertexes[0], vertexes[len(vertexes)-1])
				asr.Equal(err, nil, "error should be nil")
			}
		})
	}
}
