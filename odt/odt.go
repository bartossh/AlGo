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

// Node is a unit of a graph.
// Node should be unique entity connection to children nodes
// and having having state such as visited, begin and finnish.
type Node struct {
	id       string
	visited  bool
	begin    bool
	finnish  bool
	children []*Node
}

// NewNode creates a new node.
func NewNode(id string, v, b, f bool, children []*Node) Node {
	return Node{
		id:       id,
		visited:  v,
		begin:    b,
		finnish:  f,
		children: children,
	}
}

// Id returns node identifier.
func (n Node) Id() string {
	return n.id
}
