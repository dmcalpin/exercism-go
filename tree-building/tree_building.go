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

// SortRecord implements the Sort interface
type SortRecord []Record

func (r SortRecord) Len() int           { return len(r) }
func (r SortRecord) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r SortRecord) Less(i, j int) bool { return r[i].ID < r[j].ID }

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
	sort.Sort(SortRecord(records))

	rootRecord := records[0]
	// root record must be ID:0, Parent:0
	// becase Root and Parent are ints, not
	// setting either of these fields defaults
	// them to 0
	if rootRecord.Parent != 0 {
		return nil, errors.New("Root nood cannot have a Parent")
	}
	if rootRecord.ID != 0 {
		return nil, errors.New("Root node must have ID: 0")
	}

	// nodeMap will keep allow us to access nodes
	// by their ID so we don't have to iterate
	// over the existing tree structure to find
	// the parent
	nodeMap := make(map[int]*Node)

	// insert our root node before iterating.
	// root node is the only node that has the
	// same ID and Parent {0,0}
	rootNode := Node{rootRecord.ID, nil}
	nodeMap[0] = &rootNode

	// loop over all subsequent records
	for i := 1; i < len(records); i++ {
		record := records[i]
		node := Node{record.ID, nil}

		// do some error checking
		switch {
		case nodeMap[record.ID] != nil:
			// duplicate check
			return nil, errors.New(
				"Node already exists, duplicates not allowed")
		case record.ID != i:
			// continuity check
			return nil, errors.New(
				"Node ID's are not sequential")
		case record.ID == record.Parent:
			// cycle directly check
			return nil, errors.New(
				"Node ID is same as Node Parent")
		case nodeMap[record.Parent] == nil:
			// existing parent check
			return nil, errors.New(
				"Node's Parent does not exist")
		}

		// if all is good, add the node to our map
		nodeMap[record.ID] = &node
		// and add it to the children of the
		// parent. this is the part which
		// actually creates our tree structure
		nodeMap[record.Parent].Children = append(nodeMap[record.Parent].Children, &node)
	}
	return &rootNode, nil
}
