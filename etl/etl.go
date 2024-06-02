package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	points := map[string]int{}

	for point, letters := range in {
		for _, letter := range letters {
			points[strings.ToLower(letter)] = point
		}
	}
	return points
}
