package flatten

func Flatten(nested interface{}) []interface{} {
	flatt := []interface{}{} // initialie a new array of interface{}

	for _, element := range nested.([]interface{}) { // iterate over nested array
		switch t := element.(type) {
		case int: // if element is an int we append it to flatt
			flatt = append(flatt, t)
		case []interface{}: // if element is another []interface{} the call flatten again
			flatt = append(flatt, Flatten(t)...)
		} // if element is any other type we do nothing with the element
	}

	return flatt
}

/* Benchmark
=== RUN   BenchmarkFlatten
BenchmarkFlatten
BenchmarkFlatten-8        233034              4883 ns/op            3400 B/op         77 allocs/op
PASS
ok      flatten 1.196s
*/
