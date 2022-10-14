package djikstra

import (
	"errors"
	"fmt"
	"math"
)

type node struct {
	neighbours map[*node]struct{}
	vertex     Vertexer
	total      float64
	visited    bool
}

func newStartNode(vertex Vertexer) *node {
	neighbours := make(map[*node]struct{})
	return &node{vertex: vertex, neighbours: neighbours}
}

func (n *node) getNeighboursDistance(nn *node) float64 {
	_, ok := n.neighbours[nn]
	if !ok {
		return math.MaxFloat64
	}
	return n.vertex.GetDistance(nn.vertex)
}

type graph struct {
	nodes map[int]*node
}

// NewGraph creates new graph for dijkstra the shortest path calculation
func NewGraph(vertexes []Vertexer) *graph {
	return &graph{nodes: createUnvisited(vertexes)}
}

// CalculateResultGraph provides Path of the shortest path calculation and error if path has no solution
func (g *graph) CalculateResultGraph(st, fn Vertexer) (Path, error) {
	stn, fnn := g.findStartFinnishNode(st, fn)
	if stn == nil {
		return Path{}, errors.New("cannot find start node")
	}
	if fnn == nil {
		return Path{}, errors.New("cannot find finish node")
	}
	return g.calcResultGraph(stn, fnn)
}

func (g *graph) calcResultGraph(st, fn *node) (Path, error) {
	if st.getNeighboursDistance(fn) == 0 {
		return Path{}, nil
	}
	st.total = 0
	res := Path{}

	act := st
Loop:
	for {
		for n := range act.neighbours {
			n.neighbours[act] = struct{}{} // assure that it is possible to get back to this neighbour
			if !n.visited {
				dist := act.getNeighboursDistance(n)
				dist = dist + act.total
				if dist < n.total {
					n.total = dist
				}
			}
		}
		g.nodes[act.vertex.GetKey()].visited = true
		minDist := math.MaxFloat64
		for _, n := range g.nodes {
			if !n.visited && n.total < minDist {
				minDist = n.total
				act = n
			}
		}
		if act.vertex.GetKey() == fn.vertex.GetKey() {
			g.nodes[act.vertex.GetKey()].visited = true
			res.TotalDistance = act.total
			break Loop
		}
		if minDist == math.MaxFloat64 {
			return res, fmt.Errorf("there is no connection between nodes of key %#v and %#v", st.vertex.GetKey(), fn.vertex.GetKey())
		}
	}
	act = fn
	var parent *node
	for {
		if act.vertex.GetKey() == st.vertex.GetKey() {
			break
		}
		minDist := math.MaxFloat64
		parent = act
		for n := range act.neighbours {
			if n.visited && n.total < minDist {
				minDist = n.total
				act = n
			}
		}
		pn := &PathVertex{
			Parent: parent.vertex,
			Actual: act.vertex,
		}
		res.Vertexes = append(res.Vertexes, pn)
	}
	for _, n := range g.nodes {
		n.visited = false
	}

	return res, nil
}

func createUnvisited(vrx []Vertexer) map[int]*node {
	kv := make(map[int]Vertexer)
	unvisited := make(map[int]*node)

	for _, v := range vrx {
		kv[v.GetKey()] = v
		nd := newStartNode(v)
		nd.total = math.MaxFloat64
		unvisited[v.GetKey()] = nd
	}

	for k, v := range kv {
		for _, nk := range v.GetConnections() {
			nd := unvisited[k]
			nd.neighbours[unvisited[nk]] = struct{}{}
			unvisited[k] = nd
		}
	}
	return unvisited
}

func (g *graph) findStartFinnishNode(a, b Vertexer) (*node, *node) {
	na := g.nodes[a.GetKey()]
	nb := g.nodes[b.GetKey()]
	return na, nb
}
