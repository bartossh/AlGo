package rectangles

import (
	"sort"
)

type (
	cornerNode struct {
		x, y        int
		right, down []*cornerNode
	}
)

type ByXY []*cornerNode

func (l ByXY) Len() int {
	return len(l)
}

func (l ByXY) Less(i, j int) bool {
	return l[i].x < l[j].x && l[i].y < l[j].y
}

func (l ByXY) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func Count(strRect []string) int {
	if len(strRect) == 0 {
		return 0
	}
	if len(strRect[0]) == 0 {
		return 0
	}
	var nodes ByXY = make([]*cornerNode, 0)
	for i, vi := range strRect {
		lineNodes := make([]*cornerNode, 0)
		for j, vj := range vi {
			if vj != '+' && vj != '-' {
				nodes = append(nodes, lineNodes...)
				lineNodes = make([]*cornerNode, 0)
				continue
			}
			if vj == '+' {
				n := &cornerNode{x: i, y: j}
				for _, nd := range lineNodes {
					if nd.right == nil {
						nd.right = make([]*cornerNode, 0, 1)
					}
					nd.right = append(nd.right, n)
				}
				lineNodes = append(lineNodes, n)
			}
		}
		nodes = append(nodes, lineNodes...)
	}
	for j := range strRect[0] {
		lineNodes := make([]*cornerNode, 0)
		for i := range strRect {
			if strRect[i][j] != '+' && strRect[i][j] != '|' {
				findAndAdd(nodes, lineNodes)
				lineNodes = make([]*cornerNode, 0)
				continue
			}
			if strRect[i][j] == '+' {
				n := &cornerNode{x: i, y: j}
				for _, nd := range lineNodes {
					if nd.down == nil {
						nd.down = make([]*cornerNode, 0, 1)
					}
					nd.down = append(nd.down, n)
				}
				lineNodes = append(lineNodes, n)
			}
		}
		findAndAdd(nodes, lineNodes)
	}
	sort.Sort(nodes)

	count := 0
	for _, n := range nodes {
		count += calculateSquares(n, nodes)
	}

	return count
}

func findAndAdd(nodes []*cornerNode, lineNodes []*cornerNode) {
	for _, ln := range lineNodes {
		for _, nd := range nodes {
			if ln.x == nd.x && ln.y == nd.y {
				if nd.down == nil {
					nd.down = make([]*cornerNode, 0, 1)
				}
				nd.down = append(nd.down, ln.down...)
			}
		}
	}
}

func calculateSquares(n *cornerNode, nodes []*cornerNode) int {
	count := 0
	if len(n.right) > 0 {
		for _, rn := range n.right {
			if len(rn.down) > 0 {
				for _, dn := range rn.down {
					if findLeftNode(dn, n, nodes) {
						count++
					}
					count += calculateSquares(dn, nodes)
				}
			}
		}
	}
	return count
}

func findLeftNode(n, parent *cornerNode, nodes []*cornerNode) bool {
	for _, ln := range nodes {
		if len(ln.right) > 0 {
			for _, rn := range ln.right {
				if rn.x == n.x && rn.y == n.y {
					if len(parent.down) > 0 {
						for _, dn := range parent.down {
							if dn.x == ln.x && dn.y == ln.y {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}
