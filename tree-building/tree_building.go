package tree

import (
	"errors"
	"sort"
)

type (
	Node struct {
		ID       int
		Children []*Node
	}
	Record struct {
		ID, Parent int
	}

	ByID []Record
)

func (b ByID) Len() int {
	return len(b)
}

func (b ByID) Less(i, j int) bool {
	return b[i].ID < b[j].ID
}

func (b ByID) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	root := new(Node)
	sort.Sort(ByID(records))

	for i := range records {
		if i < len(records)-1 {
			if records[i].ID != records[i+1].ID-1 {
				return nil, errors.New("not continues")
			}
		}
	}

	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, errors.New("root has parent")
	}
	records = records[1:]

	for _, v := range records {
		if err := root.add(v); err != nil {
			return nil, err
		}
	}

	return root, nil
}

func (n *Node) add(r Record) error {
	if r.ID == n.ID {
		return errors.New("child has parent id")
	}
	if r.ID == r.Parent {
		return errors.New("child is its own parrent")
	}
	if r.ID < r.Parent {
		return errors.New("cycle indirectly")
	}
	if r.Parent == n.ID {
		newN := &Node{ID: r.ID}
		if len(n.Children) == 0 {
			n.Children = []*Node{newN}

		} else {
			for _, v := range n.Children {
				if v.ID == newN.ID {
					return errors.New("duplicated node")
				}
			}
			n.Children = append(n.Children, newN)
		}
		return nil
	}
	for _, v := range n.Children {
		if err := v.add(r); err != nil {
			return err
		}
	}
	return nil
}
