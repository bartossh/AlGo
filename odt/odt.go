// Optimal Dependencies Traverse package
//
// Calculates traverse between dependant nodes
// to be most optimal based on the cost between
// dependant nodes.

package odt

// Identifier allows to identify unique node.
type Identifier interface {
	Id() string
}

// Evaluator allows to evaluate cost between parent and child.
type Evaluator interface {
	Cost(parent, child Identifier) int
}

type Connector interface {
	Child() *Node
}

// Node is a unit of a graph.
// Node should be unique entity connection to children nodes
// and having having state such as visited, begin and finnish.
type Node struct {
	id          string
	visited     bool
	begin       bool
	finnish     bool
	connections map[*Connection]struct{}
}

// NewNode creates a new node.
func NewNode(id string, v, b, f bool) Node {
	return Node{
		id:          id,
		visited:     v,
		begin:       b,
		finnish:     f,
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
	id        string
	evaluator Evaluator
	child     *Node
}

// NewConnection creates new Connection.
func NewConnection(id string, e Evaluator, c *Node) Connection {
	return Connection{id, e, c}
}

// Id returns node identifier.
func (c *Connection) Id() string {
	return c.id
}

// Child return connection child,
// which is the Node that Connection connects to.
func (c *Connection) Child() *Node {
	return c.child
}

// Vertex is a node that is a part of a Path and
// contains calculated accumulated cost
// of getting to underlining Node.
type Vertex struct {
	node *Node
	cost int
}

// Path contains consecutive Vertex Nodes
// that has accumulated cost.
type Path = []Vertex

// Graph contains all the Nodes with Connections
type Graph struct {
	inner map[*Node]struct{}
}

func (g *Graph) CalculateResultGraph(start, finnish *Node) (Path, error) {
	return nil, nil
}
