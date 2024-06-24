package tree

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

func traverseNesting(id int, mapp map[int][]int) *Node {
	/*
		tarverse to the childs of the node recursively.
		if the child does not has childs we added directly to he children array
		we create the node and assing the children
		then return the node
	*/
	children := []*Node{}
	for _, c := range mapp[id] {
		if len(mapp[c]) > 0 {
			children = append(children, traverseNesting(c, mapp))
		} else {
			children = append(children, &Node{ID: c, Children: nil})
		}
	}
	node := Node{ID: id, Children: children}
	return &node
}

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	mapping := map[int][]int{} // we will use a map to keep track of parents and their childs

	root_found := false

	for _, record := range records {
		if record.ID == 0 && record.Parent != 0 {
			return nil, fmt.Errorf("root node cannot have parent: %v", record)
		}

		if record.ID < record.Parent {
			return nil, fmt.Errorf("higher id parent of lower id for node: %v", record)
		}

		if record.ID == record.Parent && record.ID != 0 {
			return nil, fmt.Errorf("only the root node can have same parent and id: %v", record)
		}

		if record.ID == 0 && record.Parent == 0 {
			if root_found {
				return nil, fmt.Errorf("duplicate root node: %v", record)
			} else {
				root_found = true
			}
		}

		for _, v := range mapping[record.Parent] {
			if v == record.ID {
				return nil, fmt.Errorf("duplicate node: %v", record)
			}
		}

		if record.ID > len(records)-1 {
			return nil, fmt.Errorf("non-continuous node: %v", record)
		}

		// We are going to add record to map by adding parent and appending id to parent childs
		if record.ID != 0 {
			mapping[record.Parent] = append(mapping[record.Parent], record.ID)
		}
	}

	if !root_found {
		return nil, errors.New("no root node")
	}

	parents := []int{} // We are gonna make an array with all the parents and sort it
	for k := range mapping {
		sort.Ints(mapping[k]) // We are also going to sort the childs of each parent
		parents = append(parents, k)
	}
	sort.Ints(parents)

	children := []*Node{} // We are gonna traverse trough the childs of root recursively
	for _, parent := range mapping[0] {
		children = append(children, traverseNesting(parent, mapping))
	}

	root := Node{ID: 0, Children: children}

	return &root, nil
}

/*
BenchmarkTwoTree-8                    50          24157042 ns/op         9908996 B/op     197706 allocs/op
BenchmarkTenTree-8                   318           4892310 ns/op         1007645 B/op      20042 allocs/op
BenchmarkShallowTree-8                85          14789961 ns/op          988025 B/op      10038 allocs/op
PASS
ok      tree    5.213s
*/
