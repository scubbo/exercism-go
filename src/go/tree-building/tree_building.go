package tree

// This file has intentionally not been "cleaned up" because this was the first exercise that I
// really struggled with, and I want to be able to look back on it when I understand Go (and
// especially pointers) to see how far I've come.

import (
	"errors"
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

// This was my first attempt at this challenge, which failed because the children of nodes that
// had already been created were not updated.
// It seems like pointers are the way to complete this challenge? (that doesn't make sense to my
// understanding of pointers - I thought that they let you make changes to variables across
// function boundaries, but we're not calling any other functions here, so why don't the variables
// update?). Need to ask an experienced Gopher why this doesn't work.
func BuildBroken(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	nodeMap := make(map[int]Node)
	for _, r := range records {
		fmt.Printf("Operating on record with id %v, parent %v\n", r.ID, r.Parent)
		var nodeForRecord Node
		if fetchedNode, exists := nodeMap[r.ID]; !exists {
			fmt.Printf("Node did not already exist, creating it\n")
			nodeForRecord = Node{
				ID: r.ID,
				// Children initialized to zero-value, which should be an empty slice - right?
			}
			nodeMap[r.ID] = nodeForRecord
		} else {
			fmt.Printf("Node already existed\n")
			nodeForRecord = fetchedNode
		}
		parentId := r.Parent
		if r.ID == 0 && parentId == 0 {
			continue
		}
		// Map should be initialized to the zero-values, right? So we shouldn't need to do a
		// "if nodeMap[parentId] doesn't exist, create it" (I bet this fails...)
		//
		// https://stackoverflow.com/a/69006398/1040915 states that you can't do
		// `nodeMap[parentId].Children = append(nodeMap[parentId].Children, &nodeForRecord)`,
		// though doesn't explain why
		//
		if parentNode, exists := nodeMap[parentId]; !exists {
			fmt.Printf("Node for parent %v did not exist - creating it\n", parentId)
			nodeMap[parentId] = Node{
				ID:       parentId,
				Children: []*Node{&nodeForRecord},
			}
		} else {
			fmt.Printf("Node for parent %v already existed (and looks like %v) - appending the child node\n", parentId, parentNode)
			parentNode.Children = append(parentNode.Children, &nodeForRecord)
			fmt.Printf("After appending, the parentNode look like %v\n", parentNode)
			nodeMap[parentId] = parentNode
			fmt.Printf("After reassigning back into the map, parentNode-in-map looks like %v\n", nodeMap[parentId])
		}
		fmt.Printf("Node of parent %v now has children %v\n", nodeMap[parentId], nodeMap[parentId].Children)
	}

	// Now we need to sort the Children
	for _, v := range nodeMap {
		sort.Slice(v.Children, func(i, j int) bool {
			return v.Children[i].ID < v.Children[j].ID
		})
	}

	rootNode := nodeMap[0]
	return &rootNode, nil
}

// This fails `go test -run TestMakeTreeFailure/no_root_node`. The input has no root node, and so
// should fail - but my code creates a root node because the root node (id 0) is defined as a parent
// of record 1.
//
// Thing is, I'm not sure how I could resolve this. If I _don't_ create a Node when encountering a
// reference to a parent that hasn't yet been processed, then how can I create the Children of
// said parent?
// I suppose I could sort the records before processing them (then I could guarantee that any
// records I process _should_ have already had their parent created), but that seems inefficient.
func BuildSecondAttempt(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	nodeMap := make(map[int]*Node)
	for _, r := range records {
		fmt.Printf("Operating on record with id %v, parent %v\n", r.ID, r.Parent)
		var nodeForRecord *Node
		if fetchedNode, exists := nodeMap[r.ID]; !exists {
			fmt.Printf("Node did not already exist, creating it\n")
			nodeForRecord = &Node{
				ID: r.ID,
				// Children initialized to zero-value, which should be an empty slice - right?
			}
			nodeMap[r.ID] = nodeForRecord
		} else {
			fmt.Printf("Node already existed\n")
			nodeForRecord = fetchedNode
		}
		parentId := r.Parent
		if r.ID == 0 && parentId == 0 {
			continue
		}
		if r.ID == 0 && parentId != 0 {
			return nil, errors.New("Root node cannot have parent")
		}
		// https://stackoverflow.com/a/69006398/1040915 states that you can't do
		// `nodeMap[parentId].Children = append(nodeMap[parentId].Children, &nodeForRecord)`,
		// though doesn't explain why
		//
		if parentNode, exists := nodeMap[parentId]; !exists {
			fmt.Printf("Node for parent %v did not exist - creating it\n", parentId)
			nodeMap[parentId] = &Node{
				ID:       parentId,
				Children: []*Node{nodeForRecord},
			}
		} else {
			fmt.Printf("Node for parent %v already existed (and looks like %v) - appending the child node\n", parentId, parentNode)
			parentNode.Children = append(parentNode.Children, nodeForRecord)
			fmt.Printf("After appending, the parentNode look like %v\n", parentNode)
			//nodeMap[parentId] = parentNode
			//fmt.Printf("After reassigning back into the map, parentNode-in-map looks like %v\n", nodeMap[parentId])
		}
		fmt.Printf("Node of parent %v now has children %v\n", nodeMap[parentId], nodeMap[parentId].Children)
	}

	// Now we need to sort the Children
	for _, v := range nodeMap {
		sort.Slice(v.Children, func(i, j int) bool {
			return v.Children[i].ID < v.Children[j].ID
		})
	}

	rootNode := nodeMap[0]
	return rootNode, nil
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	if records[0].ID != 0 {
		return nil, errors.New("No root node")
	}

	// Now that the records are sorted, we can guarantee that the parent of any record we process
	// will have already had its parent created

	nodeMap := make(map[int]*Node)

	for idx, r := range records {
		if r.ID != idx {
			return nil, errors.New("non-continuous ids")
		}
		fmt.Printf("Operating on record with id %v, parent %v\n", r.ID, r.Parent)
		node := &Node{ID: r.ID}
		if _, exists := nodeMap[r.ID]; exists {
			return nil, fmt.Errorf("Duplicate record with id %v", r.ID)
		}
		nodeMap[r.ID] = node

		parentId := r.Parent
		if r.ID == 0 && parentId == 0 {
			continue
		}
		if r.ID == 0 && parentId != 0 {
			return nil, errors.New("Root node cannot have parent")
		}
		if parentId >= r.ID {
			return nil, errors.New("ID is lower than parentId")
		}
		parentNode, exists := nodeMap[parentId]
		if !exists {
			return nil, fmt.Errorf("No record for id %v which is parent of id %v", parentId, r.ID)
		}
		parentNode.Children = append(parentNode.Children, node)
	}
	return nodeMap[0], nil
}
