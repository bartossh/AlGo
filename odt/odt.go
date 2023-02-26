// Optimal Dependencies Traverse package
//
// Calculates traverse between dependant nodes
// to be most optimal based on the cost between
// dependant nodes.

package odt

import (
	"errors"
	"fmt"
	"math"
)

// Identifier allows to identify unique node.
type Identifier interface {
	Id() string
}

// Evaluator allows to evaluate cost between parent and child.
type Evaluator interface {
	Cost(parent, child Identifier) int
}

// Node is a unit of a graph.
// Node should be unique entity connection to children nodes
// and having having state such as visited, begin and finnish.
type Node struct {
	id              string
	visited         bool
	begin           bool
	terminal        bool
	accumulatedCost int
	connections     map[*Connection]struct{}
}

// NewNode creates a new node.
func NewNode(id string, v, b, t bool) Node {
	return Node{
		id:          id,
		visited:     v,
		begin:       b,
		terminal:    t,
		connections: make(map[*Connection]struct{}),
	}
}

// AddConnection adds a new Connection to the node.
// If new connection is added then returns true if connection already exists
// returns false.
func (n *Node) AddConnection(c *Connection) bool {
	if _, ok := n.connections[c]; ok {
		return false
	}
	n.connections[c] = struct{}{}
	return true
}

// Id returns node identifier.
func (n Node) Id() string {
	return n.id
}

// Connection connects Node to another Node in one direction.
type Connection struct {
	child     *Node
	evaluator Evaluator
}

// NewConnection creates new Connection.
func NewConnection(c *Node, e Evaluator) Connection {
	return Connection{c, e}
}

// Vertex is a node that is a part of a Path and
// contains calculated cost
// of getting to the underlining Node.
type Vertex struct {
	node *Node
	cost int
}

// Path contains consecutive Vertex Nodes.
type Path = []Vertex

// Graph contains all the Nodes with Connections
type Graph struct {
	inner map[*Node]struct{}
}

// NewGraph creates new Graph validating given nodes
// for repetition, beginning and termination nodes.
func NewGraph(nodes []*Node) (Graph, error) {
	var terminal bool
	var begin bool
	inner := make(map[*Node]struct{}, len(nodes))
	for _, n := range nodes {
		if _, ok := inner[n]; ok {
			return Graph{}, fmt.Errorf("repeated node: %v", n.id)
		}
		switch {
		case n.terminal:
			terminal = true
		case n.begin:
			begin = true
			fallthrough
		default:
			if len(n.connections) == 0 {
				return Graph{}, errors.New("begin or intermediate node needs at least one connection")
			}
		}
		inner[n] = struct{}{}
	}

	if !begin {
		return Graph{}, errors.New("no beginning node found in the given nodes slice")
	}
	if !terminal {
		return Graph{}, errors.New("no terminate node found in the given nodes slice")
	}

	g := Graph{inner}
	g.Reset()

	return g, nil
}

// Reset resets the graph
func (g *Graph) Reset() {
	for n := range g.inner {
		n.accumulatedCost = math.MaxInt64
	}
}

// CalculateResultPath calculates path from start to finnish Node.
func (g *Graph) CalculateResultPath(begin, terminal *Node) (Path, error) {
	if begin == terminal {
		return nil, errors.New("begin node is a terminal node, path has a zero length")
	}
	if _, ok := g.inner[begin]; !ok {
		return nil, errors.New("start node is not in the graph")
	}
	if _, ok := g.inner[terminal]; !ok {
		return nil, errors.New("terminal node is not in the graph")
	}
	if !begin.begin {
		return nil, errors.New("start node should be set as begin in the graph")
	}
	if !terminal.terminal {
		return nil, errors.New("finnish node should be set as finnish in the graph")
	}

	return g.calculatePath(begin, terminal)
}

func (g *Graph) calculatePath(begin, terminal *Node) (Path, error) {
	res := make(Path, 0)

	begin.accumulatedCost = 0
	act := begin
GraphLoop:
	for {
		for connection := range act.connections {
			if !connection.child.visited {
				dist := connection.evaluator.Cost(*act, *connection.child)
				dist = dist + act.accumulatedCost
				if dist < connection.child.accumulatedCost {
					connection.child.accumulatedCost = dist
				}
			}
		}
		act.visited = true
		minDist := math.MaxInt64
		for n := range g.inner {
			if !n.visited && n.accumulatedCost < minDist {
				minDist = n.accumulatedCost
				act = n
			}
		}
		if act.id == terminal.id {
			act.visited = true
			break GraphLoop
		}
		if minDist == math.MaxInt64 {
			return res, fmt.Errorf("there is no connection between nodes %v %v", begin.id, terminal.id)
		}
	}

	act = begin
	res = append(res, Vertex{act, 0})

PathLoop:
	for {
		if act.id == terminal.id {
			break PathLoop
		}
		minDist := math.MaxInt64
		for connection := range act.connections {
			if connection.child.visited && connection.child.accumulatedCost < minDist {
				minDist = connection.child.accumulatedCost
				act = connection.child
			}
		}
		pn := Vertex{
			node: act,
			cost: act.accumulatedCost,
		}
		res = append(res, pn)
	}

	g.Reset()
	return res, nil
}
