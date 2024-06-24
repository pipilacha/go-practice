package tree

import (
	"errors"
	"sort"
)

// Exploring a community solution

func Build2(records []Record) (*Node, error) {

	nodes := map[int]*Node{}

	// lets sort by the record id first.
	// a record like [{2,0}, {0,0}, {1,0}, {3, 1}] will be sorted => [{0,0},{1,0},{2,0},{3, 1}]
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("not in sequence or bad parent")
		}
		/*
			i= 0 {0,0}
				nodes[0] = node{ID: 0}
			i=1 {1,2}
				nodes[1] = node{ID: 1}
				nodes[0].Children = nodes[0].Children + nodes[1] -> [] + nodes[1] => Node{ID: 0 Children:[node{ID: 1}]
			i=2 {2,0}
				nodes[2] = node{ID: 2}
				nodes[0].Children = nodes[0].Children + nodes[2] -> [node{ID: 1}] + nodes[2] => Node{ID:0, Children:[Node{ID: 1 Children:[]}, Node{ID: 2 Children:[]}}
			i=3 {3,0}
				nodes[3] = node{ID: 3}
				nodes[0].Children = nodes[0].Children + nodes[3] -> [node{ID: 1}, node{ID: 2}] + nodes[3] => Node{ID:0, Children:[Node{ID: 1 Children:[]}, Node{ID: 2 Children:[]}, Node{ID: 3 Children:[]}}
		*/
		nodes[r.ID] = &Node{ID: r.ID}

		if r.ID != 0 {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
		}
	}

	return nodes[0], nil
}

/*
BenchmarkTwoTree-8                    51          26442448 ns/op         9908331 B/op     197703 allocs/op
BenchmarkTenTree-8                   409           5619406 ns/op         1007842 B/op      20042 allocs/op
BenchmarkShallowTree-8                42          29445291 ns/op          988027 B/op      10038 allocs/op

BenchmarkTwoTree2-8                   70          17236277 ns/op         8242714 B/op     133317 allocs/op
BenchmarkTenTree2-8                  409           5414321 ns/op         1244390 B/op      15216 allocs/op
BenchmarkShallowTree2-8              748           4184031 ns/op         1306914 B/op      10235 allocs/op
PASS
ok      tree    13.521s
*/
