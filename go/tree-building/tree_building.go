package tree

import (
	"fmt"
	"sort"
)

// Record represents a database object.
type Record struct {
	ID     int
	Parent int
}

// Node represents a member of a tree.
type Node struct {
	ID       int
	Children []*Node
}

// Build returns the root node of a tree representation of the given list of records.
func Build(records []Record) (*Node, error) {
	nodes := make(map[int]*Node)
	// sort records by ID in ascending order
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, fmt.Errorf("invalid record with ID %d and parent ID %d", r.ID, r.Parent)
		}
		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID > 0 {
			p := nodes[r.Parent]
			p.Children = append(p.Children, nodes[r.ID])
		}
	}
	return nodes[0], nil
}
