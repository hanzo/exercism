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
	var root *Node
	nodeLookup := make(map[int]*Node)
	// sort records by ID in ascending order
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for _, r := range records {
		if r.ID < 0 || r.ID >= len(records) {
			return nil, fmt.Errorf("invalid ID: %d", r.ID)
		}
		if r.ID == r.Parent {
			if root != nil {
				return nil, fmt.Errorf("multiple root nodes found. first ID: %d, second ID: %d", root.ID, r.ID)
			}
			root = &Node{ID: r.ID}
			nodeLookup[r.ID] = root
			continue
		}
		if _, found := nodeLookup[r.ID]; found {
			return nil, fmt.Errorf("duplicate node found with ID %d", r.ID)
		}
		// every node's ID is greater than its parent's ID, so we've already stored
		// a reference to the parent by the time we visit any child node
		parent, found := nodeLookup[r.Parent]
		if !found {
			return nil, fmt.Errorf("failed to find parent with ID %d", r.Parent)
		}
		childNode := &Node{ID: r.ID}
		nodeLookup[r.ID] = childNode
		parent.Children = append(parent.Children, childNode)
	}
	return root, nil
}
