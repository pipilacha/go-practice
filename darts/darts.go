package darts

import (
	"math"
)

func Score(x, y float64) (points int) {
	//r := math.Sqrt(x*x + y*y)
	r := math.Hypot(x, y)

	switch {
	case r > 10:
		points = 0
	case r > 5:
		points = 1
	case r > 1:
		points = 5
	case r <= 1:
		points = 10
	}

	return points
}
