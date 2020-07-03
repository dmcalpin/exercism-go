package tree

import (
	"errors"
	"sort"
)

// Record is a database record
type Record struct {
	ID     int // Unique ID
	Parent int // Parent Record ID
}

// Node is a node in the tree
type Node struct {
	ID       int     // Unique ID
	Children []*Node // Nodes which belong to this one
}

// Build creates a tree from a slice of Records
// The root node is returned, along with a
// possible error
func Build(records []Record) (*Node, error) {
	// return nil instead of empty node if
	// there are no records
	if len(records) == 0 {
		return nil, nil
	}
	// sorting makes it easier to work with
	// the data
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	// root record must be ID:0, Parent:0
	// becase Root and Parent are ints, not
	// setting either of these fields defaults
	// them to 0
	if records[0] != (Record{}) {
		return nil, errors.New("Root must be 0,0")
	}

	// nodeMap will keep allow us to access nodes
	// by their ID so we don't have to iterate
	// over the existing tree structure to find
	// the parent
	nodeMap := make(map[int]*Node)

	// insert our root node before iterating.
	// root node is the only node that has the
	// same ID and Parent {0,0}
	nodeMap[0] = &(Node{records[0].ID, nil})

	// loop over all subsequent records
	for i := 1; i < len(records); i++ {
		record := records[i]
		node := Node{record.ID, nil}

		// do some error checking
		switch {
		case record.ID != i:
			// continuity check
			return nil, errors.New(
				"Node ID's are not sequential")
		case record.Parent >= i:
			// existing parent check
			return nil, errors.New(
				"Node's Parent does not exist")
		}

		// if all is good, add the node to our map
		nodeMap[record.ID] = &node
		// and add it to the children of the
		// parent. this is the part which
		// actually creates our tree structure
		nodeMap[record.Parent].Children =
			append(nodeMap[record.Parent].Children, &node)
	}
	return nodeMap[0], nil
}
