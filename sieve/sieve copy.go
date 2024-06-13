package sieve

import (
	"sort"
)

func Sieve2(limit int) (sieved []int) {
	numbers := map[int]bool{}

	for i := 2; i <= limit; i++ {
		if !numbers[i] {
			for j := i; j*i <= limit; j++ {
				numbers[j*i] = true
			}
			numbers[i] = false
		}
	}

	for k, v := range numbers {
		if !v {
			sieved = append(sieved, k)
		}
	}

	sort.Ints(sieved)

	return
}

/*
=== RUN   BenchmarkSieve
BenchmarkSieve
BenchmarkSieve-8            7654            132101 ns/op           58101 B/op         98 allocs/op
PASS
ok      sieve   1.031s
*/
