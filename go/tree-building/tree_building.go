package tree

import (
	"fmt"
	"sort"
)

// Define a function Build(records []Record) (*Node, error)
// where Record is a struct containing int fields ID and Parent
// and Node is a struct containing int field ID and []*Node field Children.

type Node struct {
	ID       int
	Children []*Node
}

type Record struct {
	ID     int
	Parent int
}

// Build returns the root node of a tree representation of the given list of records.
func Build(records []Record) (*Node, error) {
	var root *Node
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
			continue
		}
		if root == nil {
			return nil, fmt.Errorf("non-root record found before root: ID %d", r.ID)
		}
		// find Node with ID == r.Parent
		parent := DFS(root, r.Parent)
		if parent == nil {
			return nil, fmt.Errorf("failed to find parent with ID %d", r.Parent)
		}
		for _, child := range parent.Children {
			if child.ID == r.ID {
				return nil, fmt.Errorf("duplicate node found with ID %d and parent ID %d", r.ID, r.Parent)
			}
		}
		parent.Children = append(parent.Children, &Node{ID: r.ID})
	}
	return root, nil
}

// DFS performs a depth-first search to return the node with the given ID.
func DFS(node *Node, id int) *Node {
	if node.ID == id {
		return node
	}
	for _, c := range node.Children {
		ret := DFS(c, id)
		if ret != nil {
			return ret
		}
	}
	return nil
}
