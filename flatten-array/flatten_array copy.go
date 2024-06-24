package flatten

var flatt = []interface{}{}

func flatten(nested interface{}) {
	for _, element := range nested.([]interface{}) {
		switch t := element.(type) {
		case int:
			flatt = append(flatt, t)
		case []interface{}:
			flatten(t)
		}
	}
}

func Flatten2(nested interface{}) []interface{} {

	flatten(nested)

	return flatt
}

/* Benchmark
=== RUN   BenchmarkFlatten2
BenchmarkFlatten2
BenchmarkFlatten2-8       318109              3854 ns/op            4567 B/op         27 allocs/op
PASS
ok      flatten 1.368s
*/
