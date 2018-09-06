package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type byID []Record

func (s byID) Len() int           { return len(s) }
func (s byID) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byID) Less(i, j int) bool { return s[i].ID < s[j].ID }

func Build(records []Record) (*Node, error) {
	if len(records) <= 0 {
		return nil, nil
	}

	sort.Sort(byID(records))

	nodes := make([]*Node, len(records))

	for i, r := range records {
		nodes[i] = &Node{ID: r.ID}
		switch {
		case i == 0 && (r.ID != 0 || r.Parent != 0):
			return nil, errors.New("Invalid root record")
		case i == 0:
			continue
		case i != r.ID || r.ID <= r.Parent:
			return nil, errors.New("Invalid record")
		}

		if parent := nodes[r.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[i])
		}
	}
	return nodes[0], nil
}
