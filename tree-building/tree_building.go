package tree

import "errors"

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
	numRecords := len(records)
	if numRecords == 0 {
		return nil, nil
	}

	// nodeLocs is a sorted slice of nodes
	// this is faster than sort.Sort but we
	// take a hit on memory
	nodeLocs := make([]*Node, numRecords)
	for _, record := range records {
		// error checking
		if record.ID >= numRecords {
			return nil, errors.New("Out of bounds")
		}
		if nodeLocs[record.ID] != nil {
			return nil, errors.New("Duplicate Node")
		}

		// insert the node at it's proper location
		nodeLocs[record.ID] = &Node{record.ID, nil}
	}

	for _, record := range records {
		parentNode := nodeLocs[record.Parent]

		// error checking
		if parentNode.ID > record.ID {
			return nil, errors.New("Bad parent ID")
		}
		if record.ID == record.Parent {
			if record.ID == 0 {
				continue
			}
			return nil, errors.New("Bad parent ID")
		}

		// if we already have children in the slice
		// find out where to insert the new Node
		// while maintaining order. This is faster
		// than calling sort.Sort after simply
		// appending the child
		if len(parentNode.Children) > 0 {
			var i = 0
			for _, child := range parentNode.Children {
				if record.ID < child.ID {
					break
				}
				i++
			}

			// increases the length of the slice
			parentNode.Children = append(parentNode.Children, nil)
			// shift "half" of the slice over by one
			copy(parentNode.Children[i+1:], parentNode.Children[i:])
			// insert the new node in the "middle"
			parentNode.Children[i] = nodeLocs[record.ID]
		} else {
			// if this is the first child just create the
			// node slice
			parentNode.Children = []*Node{nodeLocs[record.ID]}
		}

	}

	// return the root node
	return nodeLocs[0], nil
}
