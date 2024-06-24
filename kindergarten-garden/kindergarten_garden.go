package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// Define the Garden type here.
type Garden map[string][]string

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

var plants = map[byte]string{
	71: "grass",
	67: "clover",
	82: "radishes",
	86: "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {

	rows := strings.Split(diagram, "\n")

	if len(rows) != 3 {
		return nil, fmt.Errorf("wrong diagram format. expected two rows of plants")
	}

	if len(rows[1]) != len(rows[2]) {
		return nil, fmt.Errorf("mismatch between number of plants in each row. cups in row 1: %d, cups in row 2: %d", len(rows[1]), len(rows[2]))
	}

	if len(rows[1])%2 != 0 {
		return nil, fmt.Errorf("rows cannot have odd number of cups: %s", diagram)
	}

	if len(rows[1])/2 != len(children) {
		return nil, fmt.Errorf("mismatch between number of cups and number of childrens: %d cups, %d children", len(rows[1]), len(children))
	}

	names := make([]string, len(children))

	copy(names, children)

	sort.Strings(names)

	garden := Garden{}

	j := 0
	for i := 0; i < len(rows[1]); i += 2 {

		p1, p2, p3, p4 := plants[rows[1][i]], plants[rows[1][i+1]], plants[rows[2][i]], plants[rows[2][i+1]]

		if p1 == "" || p2 == "" || p3 == "" || p4 == "" {
			return nil, fmt.Errorf("unknown plant encoding: %s%s%s%s", p1, p2, p3, p4)
		}

		if len(garden[names[j]]) != 0 {
			return nil, fmt.Errorf("duplicate children name: %s", names[j])
		}

		garden[names[j]] = append(garden[names[j]], p1, p2, p3, p4)
		j++
	}

	return &garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (*g)[child]
	return plants, ok
}
