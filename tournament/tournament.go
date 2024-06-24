package tournament

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

func Tally(reader io.Reader, writer io.Writer) error {

	all, err := io.ReadAll(reader) // read all from reader

	if err != nil {
		return err
	}

	str := string(all) // parse to string

	matches := []string{} // will use matches array to store each match data
	for _, v := range strings.Split(str, "\n") {
		if len(v) > 1 && v[0] != '#' { // ignore break lines and comments
			matches = append(matches, v)
		}
	}

	type stats map[string]int // stats will help us keep track of each stat for each team

	tally := map[string]stats{} // tally will keep track of the scoreboard

	for _, entry := range matches {

		match := strings.Split(entry, ";") // split each match into [team1 team2 outcome]

		if len(match) != 3 || match[0] == match[1] {
			return fmt.Errorf("invalid entry. %s", entry)
		}

		for _, v := range match[0:2] { // if team is not in scoreboard then we add it
			_, ok := tally[v]
			if !ok {
				tally[v] = stats{"MP": 0, "W": 0, "D": 0, "L": 0, "P": 0}
			}
		}

		// increase scoreboard according to outcome
		tally[match[0]]["MP"] += 1
		tally[match[1]]["MP"] += 1

		switch match[2] {
		case "win":
			tally[match[0]]["W"] += 1
			tally[match[0]]["P"] += 3

			tally[match[1]]["L"] += 1
		case "loss":
			tally[match[0]]["L"] += 1

			tally[match[1]]["W"] += 1
			tally[match[1]]["P"] += 3
		case "draw":
			tally[match[0]]["D"] += 1
			tally[match[0]]["P"] += 1

			tally[match[1]]["D"] += 1
			tally[match[1]]["P"] += 1
		default:
			return fmt.Errorf("invalid entry. %s", entry)
		}
	}

	ordered := map[int][]string{}
	for k, v := range tally { // score board is printed ordered according to points. Let's use a map to group teams according to points
		names := append(ordered[v["P"]], k) // since teams can score the same points, we will sort their name in alphabetical order
		sort.Strings(names)
		ordered[v["P"]] = names
	}

	points := []int{}
	for k := range ordered { // let's create an array to store the points and sort it
		points = append(points, k)
	}
	sort.Ints(points)

	writer.Write([]byte(`Team                           | MP |  W |  D |  L |  P`))
	for i := len(points) - 1; i >= 0; i-- { // we will print points in reverse, highest first
		teams := ordered[points[i]]
		for _, team := range teams { // print the teams that scored those points
			writer.Write([]byte(fmt.Sprintf("\n%-31s|  %d |  %d |  %d |  %d |  %d", team, tally[team]["MP"], tally[team]["W"], tally[team]["D"], tally[team]["L"], tally[team]["P"])))
		}
	}
	writer.Write([]byte("\n"))

	return nil
}

/* Benchmark
=== RUN   BenchmarkTally
BenchmarkTally
BenchmarkTally-8           31803             37745 ns/op           20648 B/op        232 allocs/op
PASS
ok      tournament      1.588s
*/
