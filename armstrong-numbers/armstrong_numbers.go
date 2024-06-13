package armstrong

import (
	"strconv"
)

func pow(x, y int) (pow int) {
	pow = x
	for i := 2; i <= y; i++ {
		pow *= x
	}
	return
}

func IsNumber(n int) bool {
	n_str := strconv.Itoa(n)

	sum, n_digits := 0, len(n_str)

	for _, r := range n_str {
		//v, _ := strconv.Atoi(string(r))
		sum += pow(int(r-'0'), n_digits)
	}

	return n == sum
}

/*
Benchmark using math.Pow()
=== RUN   BenchmarkIsNumber
BenchmarkIsNumber
BenchmarkIsNumber-8      1161601              1040 ns/op              30 B/op          6 allocs/op
PASS
ok      armstrong       2.208s


Benchmark using own function to calculate pow
=== RUN   BenchmarkIsNumber
BenchmarkIsNumber
BenchmarkIsNumber-8      2246650               489.2 ns/op            30 B/op          6 allocs/op
PASS
ok      armstrong       1.645s
*/
