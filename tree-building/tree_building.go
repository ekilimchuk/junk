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

	for r, rec := range records {
		nodes[r] = &Node{ID: rec.ID}
		switch {
		case r == 0 && (rec.ID != 0 || rec.Parent != 0):
			return nil, errors.New("Invalid root record")
		case r == 0:
			continue
		case r != rec.ID || rec.ID <= rec.Parent:
			return nil, errors.New("Invalid record")
		}

		if parent := nodes[rec.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[r])
		}
	}
	return nodes[0], nil
}
