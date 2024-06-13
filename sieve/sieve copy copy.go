package sieve

func Sieve3(limit int) []int {

	numbers := map[int]bool{}

	sieved := make([]int, limit/2)
	pidx := 0

	for i := 2; i <= limit; i++ {

		if !numbers[i] {

			for j := i; j*i <= limit; j++ {
				numbers[j*i] = true
			}

			numbers[i] = false
			sieved[pidx] = i
			pidx++
			//sieved = append(sieved, i)
		}
	}

	return sieved[:pidx]
}

/*
=== RUN   BenchmarkSieve
BenchmarkSieve
BenchmarkSieve-8            9686            107172 ns/op           57887 B/op         79 allocs/op
PASS
ok      sieve   1.056s
*/
