package resistorcolortrio

import "fmt"

var colorcodes map[string]int = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	exponent := 1

	for i := 0; i < colorcodes[colors[2]]; i++ {
		exponent *= 10
	}

	prefix := ""

	amount := (colorcodes[colors[0]]*10 + colorcodes[colors[1]]) * exponent

	switch {
	case amount >= 1000000000:
		prefix = "giga"
		amount /= 1000000000
	case amount >= 1000000:
		prefix = "mega"
		amount /= 1000000
	case amount >= 1000:
		prefix = "kilo"
		amount /= 1000
	}

	return fmt.Sprintf("%d %sohms", amount, prefix)
}

/* Benchmark
=== RUN   BenchmarkLabel
BenchmarkLabel
BenchmarkLabel-8          506622              2248 ns/op             248 B/op         19 allocs/op
PASS
ok      resistorcolortrio       1.168s

Using func with switch instead of map seems to similar
=== RUN   BenchmarkLabel
BenchmarkLabel
BenchmarkLabel-8          641270              1756 ns/op             240 B/op         18 allocs/op
PASS
ok      resistorcolortrio       1.149s
*/
