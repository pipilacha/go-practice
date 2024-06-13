package sieve

func Sieve(limit int) (sieved []int) {

	numbers := map[int]bool{}

	for i := 2; i <= limit; i++ {

		if !numbers[i] {

			for j := i; j*i <= limit; j++ {
				numbers[j*i] = true
			}

			numbers[i] = false
			sieved = append(sieved, i)
		}
	}

	return
}

/*
=== RUN   BenchmarkSieve
BenchmarkSieve
BenchmarkSieve-8            9404            111243 ns/op           57962 B/op         92 allocs/op
PASS
ok      sieve   1.064s
*/
