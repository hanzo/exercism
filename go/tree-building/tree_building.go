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

// the constraints of the question imply that the root node ID must be zero
const rootNodeID = 0

// Build returns the root node of a tree representation of the given list of records.
func Build(records []Record) (*Node, error) {
	nodeLookup := make(map[int]*Node)
	// sort records by ID in ascending order
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i, r := range records {
		parent, foundParent := nodeLookup[r.Parent]
		if r.ID != i || (r.ID == rootNodeID && r.ID != r.Parent) || (r.ID > rootNodeID && !foundParent) {
			return nil, fmt.Errorf("invalid record with ID %d and parent ID %d", r.ID, r.Parent)
		}
		node := &Node{ID: r.ID}
		nodeLookup[r.ID] = node
		if r.ID > rootNodeID {
			parent.Children = append(parent.Children, node)
		}
	}
	return nodeLookup[0], nil
}
